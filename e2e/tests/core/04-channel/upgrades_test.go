//go:build !test_e2e

package channel

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	test "github.com/strangelove-ventures/interchaintest/v8/testutil"
	testifysuite "github.com/stretchr/testify/suite"

	sdkmath "cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"

	"github.com/cosmos/ibc-go/e2e/testsuite"
	"github.com/cosmos/ibc-go/e2e/testvalues"
	feetypes "github.com/cosmos/ibc-go/v8/modules/apps/29-fee/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
)

func TestChannelTestSuite(t *testing.T) {
	testifysuite.Run(t, new(ChannelTestSuite))
}

type ChannelTestSuite struct {
	testsuite.E2ETestSuite
}

// TestChannelUpgrade_WithFeeMiddleware_Succeeds tests upgrading a transfer channel to wire up fee middleware
func (s *ChannelTestSuite) TestChannelUpgrade_WithFeeMiddleware_Succeeds() {
	t := s.T()
	ctx := context.TODO()

	relayer, channelA := s.SetupChainsRelayerAndChannel(ctx, s.TransferChannelOptions())
	channelB := channelA.Counterparty
	chainA, chainB := s.GetChains()

	chainADenom := chainA.Config().Denom
	chainBDenom := chainB.Config().Denom
	chainAIBCToken := testsuite.GetIBCToken(chainBDenom, channelA.PortID, channelA.ChannelID)
	chainBIBCToken := testsuite.GetIBCToken(chainADenom, channelB.PortID, channelB.ChannelID)

	chainAWallet := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	chainAAddress := chainAWallet.FormattedAddress()

	chainBWallet := s.CreateUserOnChainB(ctx, testvalues.StartingTokenAmount)
	chainBAddress := chainBWallet.FormattedAddress()

	chainAwalletAmount := ibc.WalletAmount{
		Address: chainBWallet.FormattedAddress(), // destination address
		Denom:   chainADenom,
		Amount:  sdkmath.NewInt(testvalues.IBCTransferAmount),
	}

	var (
		transferTxResp                           ibc.Tx
		payPacketFeeTxResp                       sdk.TxResponse
		chainARelayerWallet, chainBRelayerWallet ibc.Wallet
		testFee                                  = testvalues.DefaultFee(chainADenom)
	)

	s.Require().NoError(test.WaitForBlocks(ctx, 1, chainA, chainB), "failed to wait for blocks")

	// trying to create some inflight packets, although they might get relayed before the upgrade starts
	t.Run("create inflight transfer packets between chain A and chain B", func(t *testing.T) {
		transferTxResp, err := chainA.SendIBCTransfer(ctx, channelA.ChannelID, chainAWallet.KeyName(), chainAwalletAmount, ibc.TransferOptions{})
		s.Require().NoError(err)
		s.Require().NoError(transferTxResp.Validate(), "chain-a ibc transfer tx is invalid")

		chainBwalletAmount := ibc.WalletAmount{
			Address: chainAWallet.FormattedAddress(), // destination address
			Denom:   chainBDenom,
			Amount:  sdkmath.NewInt(testvalues.IBCTransferAmount),
		}

		transferTxResp, err = chainB.SendIBCTransfer(ctx, channelB.ChannelID, chainBWallet.KeyName(), chainBwalletAmount, ibc.TransferOptions{})
		s.Require().NoError(err)
		s.Require().NoError(transferTxResp.Validate(), "chain-a ibc transfer tx is invalid")
	})

	t.Run("execute gov proposal to initiate channel upgrade", func(t *testing.T) {
		chA, err := s.QueryChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)

		s.initiateChannelUpgrade(ctx, chainA, chainAWallet, channelA.PortID, channelA.ChannelID, s.createUpgradeFields(chA))
	})

	t.Run("start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	s.Require().NoError(test.WaitForBlocks(ctx, 10, chainA, chainB), "failed to wait for blocks")

	t.Run("packets are relayed between chain A and chain B", func(t *testing.T) {
		// packet from chain A to chain B
		s.AssertPacketRelayed(ctx, chainA, channelA.PortID, channelA.ChannelID, 1)
		actualBalance, err := s.QueryBalance(ctx, chainB, chainBAddress, chainBIBCToken.IBCDenom())
		s.Require().NoError(err)
		expected := testvalues.IBCTransferAmount
		s.Require().Equal(expected, actualBalance.Int64())

		// packet from chain B to chain A
		s.AssertPacketRelayed(ctx, chainB, channelB.PortID, channelB.ChannelID, 1)
		actualBalance, err = s.QueryBalance(ctx, chainA, chainAAddress, chainAIBCToken.IBCDenom())
		s.Require().NoError(err)
		expected = testvalues.IBCTransferAmount
		s.Require().Equal(expected, actualBalance.Int64())
	})

	t.Run("verify channel A upgraded and is fee enabled", func(t *testing.T) {
		channel, err := s.QueryChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)

		// check the channel version include the fee version
		version, err := feetypes.MetadataFromVersion(channel.Version)
		s.Require().NoError(err)
		s.Require().Equal(feetypes.Version, version.FeeVersion, "the channel version did not include ics29")

		// extra check
		feeEnabled, err := s.QueryFeeEnabledChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)
		s.Require().Equal(true, feeEnabled)
	})

	t.Run("verify channel B upgraded and is fee enabled", func(t *testing.T) {
		channel, err := s.QueryChannel(ctx, chainB, channelB.PortID, channelB.ChannelID)
		s.Require().NoError(err)

		// check the channel version include the fee version
		version, err := feetypes.MetadataFromVersion(channel.Version)
		s.Require().NoError(err)
		s.Require().Equal(feetypes.Version, version.FeeVersion, "the channel version did not include ics29")

		// extra check
		feeEnabled, err := s.QueryFeeEnabledChannel(ctx, chainB, channelB.PortID, channelB.ChannelID)
		s.Require().NoError(err)
		s.Require().Equal(true, feeEnabled)
	})

	t.Run("prune packet acknowledgements", func(t *testing.T) {
		// there should be one ack for the packet that we sent before the upgrade
		acks, err := s.QueryPacketAcknowledgements(ctx, chainA, channelA.PortID, channelA.ChannelID, []uint64{})
		s.Require().NoError(err)
		s.Require().Len(acks, 1)
		s.Require().Equal(uint64(1), acks[0].Sequence)

		pruneAcksTxResponse := s.PruneAcknowledgements(ctx, chainA, chainAWallet, channelA.PortID, channelA.ChannelID, uint64(1))
		s.AssertTxSuccess(pruneAcksTxResponse)

		// after pruning there should not be any acks
		acks, err = s.QueryPacketAcknowledgements(ctx, chainA, channelA.PortID, channelA.ChannelID, []uint64{})
		s.Require().NoError(err)
		s.Require().Empty(acks)
	})

	t.Run("stop relayer", func(t *testing.T) {
		s.StopRelayer(ctx, relayer)
	})

	t.Run("recover relayer wallets", func(t *testing.T) {
		err := s.RecoverRelayerWallets(ctx, relayer)
		s.Require().NoError(err)

		chainARelayerWallet, chainBRelayerWallet, err = s.GetRelayerWallets(relayer)
		s.Require().NoError(err)
	})

	s.Require().NoError(test.WaitForBlocks(ctx, 1, chainA, chainB), "failed to wait for blocks")

	t.Run("register and verify counterparty payee", func(t *testing.T) {
		_, chainBRelayerUser := s.GetRelayerUsers(ctx)
		resp := s.RegisterCounterPartyPayee(ctx, chainB, chainBRelayerUser, channelA.Counterparty.PortID, channelA.Counterparty.ChannelID, chainBRelayerWallet.FormattedAddress(), chainARelayerWallet.FormattedAddress())
		s.AssertTxSuccess(resp)

		address, err := s.QueryCounterPartyPayee(ctx, chainB, chainBRelayerWallet.FormattedAddress(), channelA.Counterparty.ChannelID)
		s.Require().NoError(err)
		s.Require().Equal(chainARelayerWallet.FormattedAddress(), address)
	})

	var err error
	t.Run("send IBC transfer from chain A to chain B", func(t *testing.T) {
		transferTxResp, err = chainA.SendIBCTransfer(ctx, channelA.ChannelID, chainAWallet.KeyName(), chainAwalletAmount, ibc.TransferOptions{})
		s.Require().NoError(err)
		s.Require().NoError(transferTxResp.Validate(), "chain-a ibc transfer tx is invalid")
	})

	t.Run("incentivize packet", func(t *testing.T) {
		t.Run("pay packet fee", func(t *testing.T) {
			// before adding fees for the packet, there should not be incentivized packets
			packets, err := s.QueryIncentivizedPacketsForChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
			s.Require().NoError(err)
			s.Require().Empty(packets)

			packetID := channeltypes.NewPacketID(channelA.PortID, channelA.ChannelID, transferTxResp.Packet.Sequence)
			packetFee := feetypes.NewPacketFee(testFee, chainAWallet.FormattedAddress(), nil)

			payPacketFeeTxResp = s.PayPacketFeeAsync(ctx, chainA, chainAWallet, packetID, packetFee)
			s.AssertTxSuccess(payPacketFeeTxResp)

			// after adding fees for the packet, there should be incentivized packets
			packets, err = s.QueryIncentivizedPacketsForChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
			s.Require().NoError(err)
			s.Require().Len(packets, 1)
			actualFee := packets[0].PacketFees[0].Fee

			s.Require().True(actualFee.RecvFee.Equal(testFee.RecvFee))
			s.Require().True(actualFee.AckFee.Equal(testFee.AckFee))
			s.Require().True(actualFee.TimeoutFee.Equal(testFee.TimeoutFee))
		})

		t.Run("balance should be lowered by sum of recv ack and timeout", func(t *testing.T) {
			// The balance should be lowered by the sum of the recv, ack and timeout fees.
			actualBalance, err := s.GetChainANativeBalance(ctx, chainAWallet)
			s.Require().NoError(err)

			expected := testvalues.StartingTokenAmount - 2*chainAwalletAmount.Amount.Int64() - testFee.Total().AmountOf(chainADenom).Int64()
			s.Require().Equal(expected, actualBalance)
		})
	})

	t.Run("start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	t.Run("packets are relayed", func(t *testing.T) {
		packets, err := s.QueryIncentivizedPacketsForChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)
		s.Require().Empty(packets)
	})
}

// TestChannelUpgrade_WithFeeMiddleware_FailsWithTimeoutOnAck tests upgrading a transfer channel to wire up fee middleware but fails on ACK because of timeout
func (s *ChannelTestSuite) TestChannelUpgrade_WithFeeMiddleware_FailsWithTimeoutOnAck() {
	t := s.T()
	ctx := context.TODO()

	relayer, channelA := s.SetupChainsRelayerAndChannel(ctx, s.TransferChannelOptions())
	channelB := channelA.Counterparty
	chainA, chainB := s.GetChains()

	chainAWallet := s.CreateUserOnChainA(ctx, testvalues.StartingTokenAmount)
	chainBWallet := s.CreateUserOnChainB(ctx, testvalues.StartingTokenAmount)

	s.Require().NoError(test.WaitForBlocks(ctx, 1, chainA, chainB), "failed to wait for blocks")

	t.Run("execute gov proposal to set upgrade timeout", func(t *testing.T) {
		s.setUpgradeTimeoutParam(ctx, chainB, chainBWallet)
	})

	t.Run("execute gov proposal to initiate channel upgrade", func(t *testing.T) {
		chA, err := s.QueryChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)

		s.initiateChannelUpgrade(ctx, chainA, chainAWallet, channelA.PortID, channelA.ChannelID, s.createUpgradeFields(chA))
	})

	t.Run("start relayer", func(t *testing.T) {
		s.StartRelayer(relayer)
	})

	s.Require().NoError(test.WaitForBlocks(ctx, 10, chainA, chainB), "failed to wait for blocks")

	t.Run("verify channel A did not upgrade", func(t *testing.T) {
		channel, err := s.QueryChannel(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)

		s.Require().Equal(channeltypes.OPEN, channel.State, "the channel state is not OPEN")
		s.Require().Equal(transfertypes.Version, channel.Version, "the channel version is not ics20-1")

		errorReceipt, err := s.QueryUpgradeError(ctx, chainA, channelA.PortID, channelA.ChannelID)
		s.Require().NoError(err)
		s.Require().Equal(uint64(1), errorReceipt.Sequence)
		s.Require().Contains(errorReceipt.Message, "restored channel to pre-upgrade state")
	})

	t.Run("verify channel B did not upgrade", func(t *testing.T) {
		channel, err := s.QueryChannel(ctx, chainB, channelB.PortID, channelB.ChannelID)
		s.Require().NoError(err)

		s.Require().Equal(channeltypes.OPEN, channel.State, "the channel state is not OPEN")
		s.Require().Equal(transfertypes.Version, channel.Version, "the channel version is not ics20-1")

		errorReceipt, err := s.QueryUpgradeError(ctx, chainB, channelB.PortID, channelB.ChannelID)
		s.Require().NoError(err)
		s.Require().Equal(uint64(1), errorReceipt.Sequence)
		s.Require().Contains(errorReceipt.Message, "restored channel to pre-upgrade state")
	})
}

// createUpgradeFields created the upgrade fields for channel
func (s *ChannelTestSuite) createUpgradeFields(channel channeltypes.Channel) channeltypes.UpgradeFields {
	versionMetadata := feetypes.Metadata{
		FeeVersion: feetypes.Version,
		AppVersion: transfertypes.Version,
	}
	versionBytes, err := feetypes.ModuleCdc.MarshalJSON(&versionMetadata)
	s.Require().NoError(err)

	return channeltypes.NewUpgradeFields(channel.Ordering, channel.ConnectionHops, string(versionBytes))
}

// setUpgradeTimeoutParam creates and submits a governance proposal to execute the message to update 04-channel params with a timeout of 1s
func (s *ChannelTestSuite) setUpgradeTimeoutParam(ctx context.Context, chain ibc.Chain, wallet ibc.Wallet) {
	const timeoutDelta = 1000000000 // use 1 second as relative timeout to force upgrade timeout on the counterparty
	govModuleAddress, err := s.QueryModuleAccountAddress(ctx, govtypes.ModuleName, chain)
	s.Require().NoError(err)
	s.Require().NotNil(govModuleAddress)

	upgradeTimeout := channeltypes.NewTimeout(channeltypes.DefaultTimeout.Height, timeoutDelta)
	msg := channeltypes.NewMsgUpdateChannelParams(govModuleAddress.String(), channeltypes.NewParams(upgradeTimeout))
	s.ExecuteAndPassGovV1Proposal(ctx, msg, chain, wallet)
}

// initiateChannelUpgrade creates and submits a governance proposal to execute the message to initiate a channel upgrade
func (s *ChannelTestSuite) initiateChannelUpgrade(ctx context.Context, chain ibc.Chain, wallet ibc.Wallet, portID, channelID string, upgradeFields channeltypes.UpgradeFields) {
	govModuleAddress, err := s.QueryModuleAccountAddress(ctx, govtypes.ModuleName, chain)
	s.Require().NoError(err)
	s.Require().NotNil(govModuleAddress)

	msg := channeltypes.NewMsgChannelUpgradeInit(portID, channelID, upgradeFields, govModuleAddress.String())
	s.ExecuteAndPassGovV1Proposal(ctx, msg, chain, wallet)
}
