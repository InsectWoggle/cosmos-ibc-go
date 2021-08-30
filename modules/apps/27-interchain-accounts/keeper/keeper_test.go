package keeper_test

import (
	"testing"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/suite"

	"github.com/cosmos/ibc-go/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/modules/core/04-channel/types"
	ibctesting "github.com/cosmos/ibc-go/testing"
)

type KeeperTestSuite struct {
	suite.Suite

	coordinator *ibctesting.Coordinator

	// testing chains used for convenience and readability
	chainA *ibctesting.TestChain
	chainB *ibctesting.TestChain
	chainC *ibctesting.TestChain
}

func (suite *KeeperTestSuite) SetupTest() {
	suite.coordinator = ibctesting.NewCoordinator(suite.T(), 3)
	suite.chainA = suite.coordinator.GetChain(ibctesting.GetChainID(0))
	suite.chainB = suite.coordinator.GetChain(ibctesting.GetChainID(1))
	suite.chainC = suite.coordinator.GetChain(ibctesting.GetChainID(2))
}

func NewICAPath(chainA, chainB *ibctesting.TestChain) *ibctesting.Path {
	path := ibctesting.NewPath(chainA, chainB)
	path.EndpointA.ChannelConfig.PortID = types.PortID
	path.EndpointB.ChannelConfig.PortID = types.PortID
	path.EndpointA.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointB.ChannelConfig.Order = channeltypes.ORDERED
	path.EndpointA.ChannelConfig.Version = types.Version
	path.EndpointB.ChannelConfig.Version = types.Version

	return path
}

// SetupICAPath invokes the InterchainAccounts entrypoint and subsequent channel handshake handlers
func SetupICAPath(path *ibctesting.Path, owner string) error {
	if err := InitInterchainAccount(path.EndpointA, owner); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenTry(); err != nil {
		return err
	}

	if err := path.EndpointA.ChanOpenAck(); err != nil {
		return err
	}

	if err := path.EndpointB.ChanOpenConfirm(); err != nil {
		return err
	}

	return nil
}

// InitInterchainAccount is a helper function for starting the channel handshake
// TODO: parse identifiers from events
func InitInterchainAccount(endpoint *ibctesting.Endpoint, owner string) error {
	portID, err := types.GeneratePortID(owner, endpoint.ConnectionID, endpoint.Counterparty.ConnectionID)
	if err != nil {
		return err
	}
	channelSequence := endpoint.Chain.App.GetIBCKeeper().ChannelKeeper.GetNextChannelSequence(endpoint.Chain.GetContext())

	if err := endpoint.Chain.GetSimApp().ICAKeeper.InitInterchainAccount(endpoint.Chain.GetContext(), endpoint.ConnectionID, endpoint.Counterparty.ConnectionID, owner); err != nil {
		return err
	}

	// commit state changes for proof verification
	endpoint.Chain.App.Commit()
	endpoint.Chain.NextBlock()

	// update port/channel ids
	endpoint.ChannelID = channeltypes.FormatChannelIdentifier(channelSequence)
	endpoint.ChannelConfig.PortID = portID
	return nil
}

func TestKeeperTestSuite(t *testing.T) {
	suite.Run(t, new(KeeperTestSuite))
}

func (suite *KeeperTestSuite) TestIsBound() {
	isBound := suite.chainA.GetSimApp().ICAKeeper.IsBound(suite.chainA.GetContext(), types.PortID)
	suite.Require().True(isBound)
}

func (suite *KeeperTestSuite) TestGetPort() {
	port := suite.chainA.GetSimApp().ICAKeeper.GetPort(suite.chainA.GetContext())
	suite.Require().Equal(types.PortID, port)
}

func (suite *KeeperTestSuite) TestGetInterchainAccountAddress() {
	suite.SetupTest()
	path := NewICAPath(suite.chainA, suite.chainB)
	suite.coordinator.SetupConnections(path)

	err := SetupICAPath(path, "testing")
	suite.Require().NoError(err)

	counterpartyPortID := path.EndpointA.ChannelConfig.PortID
	expectedAddr := authtypes.NewBaseAccountWithAddress(types.GenerateAddress(counterpartyPortID)).GetAddress()

	retrievedAddr, found := suite.chainB.GetSimApp().ICAKeeper.GetInterchainAccountAddress(suite.chainB.GetContext(), counterpartyPortID)
	suite.Require().True(found)
	suite.Require().Equal(expectedAddr.String(), retrievedAddr)

	retrievedAddr, found = suite.chainA.GetSimApp().ICAKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), "invalid port")
	suite.Require().False(found)
	suite.Require().Empty(retrievedAddr)
}

func (suite *KeeperTestSuite) TestSetInterchainAccountAddress() {
	expectedAddr, portID := "address", "port"
	suite.chainA.GetSimApp().ICAKeeper.SetInterchainAccountAddress(suite.chainA.GetContext(), portID, expectedAddr)

	retrievedAddr, found := suite.chainA.GetSimApp().ICAKeeper.GetInterchainAccountAddress(suite.chainA.GetContext(), portID)
	suite.Require().True(found)
	suite.Require().Equal(expectedAddr, retrievedAddr)
}
