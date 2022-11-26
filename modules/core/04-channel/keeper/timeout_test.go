package keeper_test

import (
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"

	clienttypes "github.com/cosmos/ibc-go/v6/modules/core/02-client/types"
	connectiontypes "github.com/cosmos/ibc-go/v6/modules/core/03-connection/types"
	"github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v6/modules/core/24-host"
	"github.com/cosmos/ibc-go/v6/modules/core/exported"
	ibctesting "github.com/cosmos/ibc-go/v6/testing"
)

// TestTimeoutPacket test the TimeoutPacket call on chainA by ensuring the timeout has passed
// on chainB, but that no ack has been written yet. Test cases expected to reach proof
// verification must specify which proof to use using the ordered bool.
func (suite *KeeperTestSuite) TestTimeoutPacket() {
	var (
		path        *ibctesting.Path
		packet      types.Packet
		nextSeqRecv uint64
		ordered     bool
		expError    *sdkerrors.Error
		hasEvents   bool
	)

	testCases := []testCase{
		{"success: ORDERED", func() {
			ordered = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()
		}, true},
		{"success: UNORDERED", func() {
			ordered = false
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()
		}, true},
		{"packet already timed out: ORDERED", func() {
			expError = types.ErrNoOpMsg
			ordered = true
			hasEvents = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			err = path.EndpointA.TimeoutPacket(packet)
			suite.Require().NoError(err)
		}, false},
		{"packet already timed out: UNORDERED", func() {
			expError = types.ErrNoOpMsg
			ordered = false
			hasEvents = true
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			err = path.EndpointA.TimeoutPacket(packet)
			suite.Require().NoError(err)
		}, false},
		{"channel not found", func() {
			expError = types.ErrChannelNotFound
			// use wrong channel naming
			suite.coordinator.Setup(path)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, ibctesting.InvalidID, ibctesting.InvalidID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"channel not open", func() {
			expError = types.ErrInvalidChannelState
			suite.coordinator.Setup(path)

			timeoutHeight := path.EndpointA.GetClientState().GetLatestHeight().Increment().(clienttypes.Height)

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			err = path.EndpointA.SetChannelClosed()
			suite.Require().NoError(err)
		}, false},
		{"packet destination port ≠ channel counterparty port", func() {
			expError = types.ErrInvalidPacket
			suite.coordinator.Setup(path)
			// use wrong port for dest
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, ibctesting.InvalidID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"packet destination channel ID ≠ channel counterparty channel ID", func() {
			expError = types.ErrInvalidPacket
			suite.coordinator.Setup(path)
			// use wrong channel for dest
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, ibctesting.InvalidID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"connection not found", func() {
			expError = connectiontypes.ErrConnectionNotFound
			// pass channel check
			suite.chainA.App.GetIBCKeeper().ChannelKeeper.SetChannel(
				suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID,
				types.NewChannel(types.OPEN, types.ORDERED, types.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID), []string{connIDA}, path.EndpointA.ChannelConfig.Version),
			)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"timeout", func() {
			expError = types.ErrPacketTimeout
			suite.coordinator.Setup(path)
			sequence, err := path.EndpointA.SendPacket(defaultTimeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
			path.EndpointA.UpdateClient()
		}, false},
		{"packet already received ", func() {
			expError = types.ErrPacketReceived
			ordered = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			nextSeqRecv = 2
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(defaultTimeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, timeoutTimestamp)
			path.EndpointA.UpdateClient()
		}, false},
		{"packet hasn't been sent", func() {
			expError = types.ErrNoOpMsg
			ordered = true
			hasEvents = true
			path.SetChannelOrdered()

			suite.coordinator.Setup(path)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, uint64(suite.chainB.GetContext().BlockTime().UnixNano()))
			path.EndpointA.UpdateClient()
		}, false},
		{"next seq receive verification failed", func() {
			// skip error check, error occurs in light-clients

			// set ordered to false resulting in wrong proof provided
			ordered = false

			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			path.EndpointA.UpdateClient()
		}, false},
		{"packet ack verification failed", func() {
			// skip error check, error occurs in light-clients

			// set ordered to true resulting in wrong proof provided
			ordered = true

			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			path.EndpointA.UpdateClient()
		}, false},
	}

	for i, tc := range testCases {
		tc := tc
		suite.Run(fmt.Sprintf("Case %s, %d/%d tests", tc.msg, i, len(testCases)), func() {
			var (
				proof       []byte
				proofHeight exported.Height
			)

			suite.SetupTest() // reset
			expError = nil    // must be expliticly changed by failed cases
			hasEvents = false // must be explicitly changed by functions emitting events
			nextSeqRecv = 1   // must be explicitly changed
			path = ibctesting.NewPath(suite.chainA, suite.chainB)

			tc.malleate()

			orderedPacketKey := host.NextSequenceRecvKey(packet.GetDestPort(), packet.GetDestChannel())
			unorderedPacketKey := host.PacketReceiptKey(packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence())

			if path.EndpointB.ConnectionID != "" {
				if ordered {
					proof, proofHeight = path.EndpointB.QueryProof(orderedPacketKey)
				} else {
					proof, proofHeight = path.EndpointB.QueryProof(unorderedPacketKey)
				}
			}

			ctx := suite.chainA.GetContext()
			err := suite.chainA.App.GetIBCKeeper().ChannelKeeper.TimeoutPacket(ctx, packet, proof, proofHeight, nextSeqRecv)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				// only check if expError is set, since not all error codes can be known
				if expError != nil {
					suite.Require().True(errors.Is(err, expError))
				}
			}

			// Verify events
			events := ctx.EventManager().Events()
			expEvents := map[string]map[string]string{
				"timeout_packet": {
					"packet_timeout_height":    packet.GetTimeoutHeight().String(),
					"packet_timeout_timestamp": fmt.Sprintf("%d", packet.GetTimeoutTimestamp()),
					"packet_sequence":          fmt.Sprintf("%d", packet.GetSequence()),
					"packet_src_port":          path.EndpointA.ChannelConfig.PortID,
					"packet_src_channel":       path.EndpointA.ChannelID,
					"packet_dst_port":          path.EndpointB.ChannelConfig.PortID,
					"packet_dst_channel":       path.EndpointB.ChannelID,
					"packet_channel_ordering":  path.EndpointA.ChannelConfig.Order.String(),
				},
				"message": {
					"module": "ibc_channel",
				},
			}

			if hasEvents {
				ibctesting.AssertEvents(&suite.Suite, expEvents, events)
			} else {
				suite.Require().Len(events, 0)
			}
		})
	}
}

// TestTimeoutExectued verifies that packet commitments are deleted on chainA after the
// channel capabilities are verified.
func (suite *KeeperTestSuite) TestTimeoutExecuted() {
	var (
		path      *ibctesting.Path
		packet    types.Packet
		chanCap   *capabilitytypes.Capability
		hasEvents bool
	)

	testCases := []testCase{
		{"success ORDERED", func() {
			hasEvents = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, true},
		{"channel not found", func() {
			// use wrong channel naming
			suite.coordinator.Setup(path)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, ibctesting.InvalidID, ibctesting.InvalidID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"incorrect capability ORDERED", func() {
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			chanCap = capabilitytypes.NewCapability(100)
		}, false},
	}

	for i, tc := range testCases {
		tc := tc
		suite.Run(fmt.Sprintf("Case %s, %d/%d tests", tc.msg, i, len(testCases)), func() {
			hasEvents = false
			suite.SetupTest() // reset
			path = ibctesting.NewPath(suite.chainA, suite.chainB)

			tc.malleate()

			ctx := suite.chainA.GetContext()
			err := suite.chainA.App.GetIBCKeeper().ChannelKeeper.TimeoutExecuted(ctx, chanCap, packet)
			pc := suite.chainA.App.GetIBCKeeper().ChannelKeeper.GetPacketCommitment(ctx, packet.GetSourcePort(), packet.GetSourceChannel(), packet.GetSequence())

			if tc.expPass {
				suite.NoError(err)
				suite.Nil(pc)
			} else {
				suite.Error(err)
			}

			// Verify events
			events := ctx.EventManager().Events()
			expEvents := map[string]map[string]string{
				"timeout_packet": {
					"packet_timeout_height":    packet.GetTimeoutHeight().String(),
					"packet_timeout_timestamp": fmt.Sprintf("%d", packet.GetTimeoutTimestamp()),
					"packet_sequence":          fmt.Sprintf("%d", packet.GetSequence()),
					"packet_src_port":          path.EndpointA.ChannelConfig.PortID,
					"packet_src_channel":       path.EndpointA.ChannelID,
					"packet_dst_port":          path.EndpointB.ChannelConfig.PortID,
					"packet_dst_channel":       path.EndpointB.ChannelID,
					"packet_channel_ordering":  path.EndpointA.ChannelConfig.Order.String(),
				},
				"message": {
					"module": "ibc_channel",
				},
				"channel_close": {
					"port_id":                 path.EndpointA.ChannelConfig.PortID,
					"channel_id":              path.EndpointA.ChannelID,
					"counterparty_port_id":    path.EndpointB.ChannelConfig.PortID,
					"counterparty_channel_id": path.EndpointB.ChannelID,
					"packet_channel_ordering": path.EndpointA.ChannelConfig.Order.String(),
					"connection_id":           path.EndpointA.ConnectionID,
				},
			}

			if hasEvents {
				ibctesting.AssertEvents(&suite.Suite, expEvents, events)
			} else {
				suite.Require().Len(events, 0)
			}
		})
	}
}

// TestTimeoutOnClose tests the call TimeoutOnClose on chainA by closing the corresponding
// channel on chainB after the packet commitment has been created.
func (suite *KeeperTestSuite) TestTimeoutOnClose() {
	var (
		path        *ibctesting.Path
		packet      types.Packet
		chanCap     *capabilitytypes.Capability
		nextSeqRecv uint64
		ordered     bool
		hasEvents   bool
	)

	testCases := []testCase{
		{"success: ORDERED", func() {
			ordered = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, true},
		{"success: UNORDERED", func() {
			ordered = false
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, true},
		{"channel not found", func() {
			// use wrong channel naming
			suite.coordinator.Setup(path)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, ibctesting.InvalidID, ibctesting.InvalidID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
		}, false},
		{"packet dest port ≠ channel counterparty port", func() {
			suite.coordinator.Setup(path)
			// use wrong port for dest
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, ibctesting.InvalidID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"packet dest channel ID ≠ channel counterparty channel ID", func() {
			suite.coordinator.Setup(path)
			// use wrong channel for dest
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, ibctesting.InvalidID, defaultTimeoutHeight, disabledTimeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"connection not found", func() {
			// pass channel check
			suite.chainA.App.GetIBCKeeper().ChannelKeeper.SetChannel(
				suite.chainA.GetContext(),
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID,
				types.NewChannel(types.OPEN, types.ORDERED, types.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID), []string{connIDA}, path.EndpointA.ChannelConfig.Version),
			)
			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, defaultTimeoutHeight, disabledTimeoutTimestamp)

			// create chancap
			suite.chainA.CreateChannelCapability(suite.chainA.GetSimApp().ScopedIBCMockKeeper, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"packet hasn't been sent ORDERED", func() {
			hasEvents = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			packet = types.NewPacket(ibctesting.MockPacketData, 1, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.GetSelfHeight(suite.chainB.GetContext()), uint64(suite.chainB.GetContext().BlockTime().UnixNano()))
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"packet already received ORDERED", func() {
			path.SetChannelOrdered()
			nextSeqRecv = 2
			ordered = true
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"channel verification failed ORDERED", func() {
			ordered = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, timeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"next seq receive verification failed ORDERED", func() {
			// set ordered to false providing the wrong proof for ORDERED case
			ordered = false
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			path.EndpointA.UpdateClient()
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.GetSelfHeight(suite.chainB.GetContext()), uint64(suite.chainB.GetContext().BlockTime().UnixNano()))
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"packet ack verification failed", func() {
			// set ordered to true providing the wrong proof for UNORDERED case
			ordered = true
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, disabledTimeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			path.EndpointA.UpdateClient()
			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, timeoutHeight, disabledTimeoutTimestamp)
			chanCap = suite.chainA.GetChannelCapability(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID)
		}, false},
		{"channel capability not found ORDERED", func() {
			ordered = true
			path.SetChannelOrdered()
			suite.coordinator.Setup(path)

			timeoutHeight := clienttypes.GetSelfHeight(suite.chainB.GetContext())
			timeoutTimestamp := uint64(suite.chainB.GetContext().BlockTime().UnixNano())

			sequence, err := path.EndpointA.SendPacket(timeoutHeight, timeoutTimestamp, ibctesting.MockPacketData)
			suite.Require().NoError(err)
			path.EndpointB.SetChannelClosed()
			// need to update chainA's client representing chainB to prove missing ack
			path.EndpointA.UpdateClient()

			packet = types.NewPacket(ibctesting.MockPacketData, sequence, path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID, clienttypes.GetSelfHeight(suite.chainB.GetContext()), uint64(suite.chainB.GetContext().BlockTime().UnixNano()))
			chanCap = capabilitytypes.NewCapability(100)
		}, false},
	}

	for i, tc := range testCases {
		tc := tc
		suite.Run(fmt.Sprintf("Case %s, %d/%d tests", tc.msg, i, len(testCases)), func() {
			var proof []byte

			suite.SetupTest() // reset
			nextSeqRecv = 1   // must be explicitly changed
			hasEvents = false
			path = ibctesting.NewPath(suite.chainA, suite.chainB)

			tc.malleate()

			channelKey := host.ChannelKey(packet.GetDestPort(), packet.GetDestChannel())
			unorderedPacketKey := host.PacketReceiptKey(packet.GetDestPort(), packet.GetDestChannel(), packet.GetSequence())
			orderedPacketKey := host.NextSequenceRecvKey(packet.GetDestPort(), packet.GetDestChannel())

			proofClosed, proofHeight := suite.chainB.QueryProof(channelKey)

			if ordered {
				proof, _ = suite.chainB.QueryProof(orderedPacketKey)
			} else {
				proof, _ = suite.chainB.QueryProof(unorderedPacketKey)
			}

			ctx := suite.chainA.GetContext()
			err := suite.chainA.App.GetIBCKeeper().ChannelKeeper.TimeoutOnClose(ctx, chanCap, packet, proof, proofClosed, proofHeight, nextSeqRecv)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}

			// Verify events
			events := ctx.EventManager().Events()
			expEvents := map[string]map[string]string{
				"timeout_packet": {
					"packet_timeout_height":    packet.GetTimeoutHeight().String(),
					"packet_timeout_timestamp": fmt.Sprintf("%d", packet.GetTimeoutTimestamp()),
					"packet_sequence":          fmt.Sprintf("%d", packet.GetSequence()),
					"packet_src_port":          path.EndpointA.ChannelConfig.PortID,
					"packet_src_channel":       path.EndpointA.ChannelID,
					"packet_dst_port":          path.EndpointB.ChannelConfig.PortID,
					"packet_dst_channel":       path.EndpointB.ChannelID,
					"packet_channel_ordering":  path.EndpointA.ChannelConfig.Order.String(),
				},
				"message": {
					"module": "ibc_channel",
				},
			}

			if hasEvents {
				ibctesting.AssertEvents(&suite.Suite, expEvents, events)
			} else {
				suite.Require().Len(events, 0)
			}
		})
	}
}
