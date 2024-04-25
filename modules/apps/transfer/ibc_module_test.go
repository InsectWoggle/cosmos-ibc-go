package transfer_test

import (
	"math"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	sdk "github.com/cosmos/cosmos-sdk/types"

	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer"
	"github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	multidenom "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types/v3"
	connectiontypes "github.com/cosmos/ibc-go/v8/modules/core/03-connection/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	host "github.com/cosmos/ibc-go/v8/modules/core/24-host"
	ibctesting "github.com/cosmos/ibc-go/v8/testing"
)

func (suite *TransferTestSuite) TestOnChanOpenInit() {
	var (
		channel      *channeltypes.Channel
		path         *ibctesting.Path
		chanCap      *capabilitytypes.Capability
		counterparty channeltypes.Counterparty
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			// connection hops is not used in the transfer application callback,
			// it is already validated in the core OnChanUpgradeInit.
			"success: invalid connection hops", func() {
				path.EndpointA.ConnectionID = "invalid-connection-id"
			}, true,
		},
		{
			"empty version string", func() {
				channel.Version = ""
			}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid version", func() {
				channel.Version = "version" //nolint:goconst
			}, false,
		},
		{
			"capability already claimed", func() {
				err := suite.chainA.GetSimApp().ScopedTransferKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
				suite.Require().NoError(err)
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset
			path = ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.SetupConnections()
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty = channeltypes.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.INIT,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}

			var err error
			chanCap, err = suite.chainA.App.GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			transferModule := transfer.NewIBCModule(suite.chainA.GetSimApp().TransferKeeper, types.Version)

			tc.malleate() // explicitly change fields in channel and testChannel

			version, err := transferModule.OnChanOpenInit(suite.chainA.GetContext(), channel.Ordering, channel.ConnectionHops,
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, counterparty, channel.Version,
			)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal(version, "")
			}
		})
	}
}

func (suite *TransferTestSuite) TestOnChanOpenTry() {
	var (
		channel             *channeltypes.Channel
		chanCap             *capabilitytypes.Capability
		path                *ibctesting.Path
		counterparty        channeltypes.Counterparty
		counterpartyVersion string
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"max channels reached", func() {
				path.EndpointA.ChannelID = channeltypes.FormatChannelIdentifier(math.MaxUint32 + 1)
			}, false,
		},
		{
			"capability already claimed", func() {
				err := suite.chainA.GetSimApp().ScopedTransferKeeper.ClaimCapability(suite.chainA.GetContext(), chanCap, host.ChannelCapabilityPath(path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID))
				suite.Require().NoError(err)
			}, false,
		},
		{
			"invalid order - ORDERED", func() {
				channel.Ordering = channeltypes.ORDERED
			}, false,
		},
		{
			"invalid port ID", func() {
				path.EndpointA.ChannelConfig.PortID = ibctesting.MockPort
			}, false,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path = ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.SetupConnections()
			path.EndpointA.ChannelID = ibctesting.FirstChannelID

			counterparty = channeltypes.NewCounterparty(path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID)
			channel = &channeltypes.Channel{
				State:          channeltypes.TRYOPEN,
				Ordering:       channeltypes.UNORDERED,
				Counterparty:   counterparty,
				ConnectionHops: []string{path.EndpointA.ConnectionID},
				Version:        types.Version,
			}
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			chanCap, err = suite.chainA.App.GetScopedIBCKeeper().NewCapability(suite.chainA.GetContext(), host.ChannelCapabilityPath(ibctesting.TransferPort, path.EndpointA.ChannelID))
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			version, err := cbs.OnChanOpenTry(suite.chainA.GetContext(), channel.Ordering, channel.ConnectionHops,
				path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, chanCap, channel.Counterparty, counterpartyVersion,
			)

			if tc.expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
				suite.Require().Equal("", version)
			}
		})
	}
}

func (suite *TransferTestSuite) TestOnChanOpenAck() {
	var counterpartyVersion string

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success", func() {}, true,
		},
		{
			"invalid counterparty version", func() {
				counterpartyVersion = "version"
			}, false,
		},
	}

	for _, tc := range testCases {
		tc := tc

		suite.Run(tc.name, func() {
			suite.SetupTest() // reset

			path := ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.SetupConnections()
			path.EndpointA.ChannelID = ibctesting.FirstChannelID
			counterpartyVersion = types.Version

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), ibctesting.TransferPort)
			suite.Require().NoError(err)

			cbs, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			tc.malleate() // explicitly change fields in channel and testChannel

			err = cbs.OnChanOpenAck(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointA.Counterparty.ChannelID, counterpartyVersion)

			if tc.expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
			}
		})
	}
}

func (suite *TransferTestSuite) TestOnChanUpgradeInit() {
	var path *ibctesting.Path

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {}, // successful happy path for a standalone transfer app is swapping out the underlying connection
			nil,
		},
		{
			"invalid upgrade connection",
			func() {
				path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{"connection-100"}
				path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{"connection-100"}
			},
			connectiontypes.ErrConnectionNotFound,
		},
		{
			"invalid upgrade ordering",
			func() {
				path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.Ordering = channeltypes.ORDERED
				path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.Ordering = channeltypes.ORDERED
			},
			channeltypes.ErrInvalidChannelOrdering,
		},
		{
			"invalid upgrade version",
			func() {
				path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.Version = ibctesting.InvalidID
				path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.Version = ibctesting.InvalidID
			},
			types.ErrInvalidVersion,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			path = ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.Setup()

			// configure the channel upgrade to modify the underlying connection
			upgradePath := ibctesting.NewPath(suite.chainA, suite.chainB)
			upgradePath.SetupConnections()

			path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointA.ConnectionID}
			path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointB.ConnectionID}

			tc.malleate()

			err := path.EndpointA.ChanUpgradeInit()

			expPass := tc.expError == nil
			if expPass {
				suite.Require().NoError(err)
				upgrade := path.EndpointA.GetChannelUpgrade()
				suite.Require().Equal(upgradePath.EndpointA.ConnectionID, upgrade.Fields.ConnectionHops[0])
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expError.Error())
			}
		})
	}
}

func (suite *TransferTestSuite) TestOnChanUpgradeTry() {
	var (
		counterpartyUpgrade channeltypes.Upgrade
		path                *ibctesting.Path
	)

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {}, // successful happy path for a standalone transfer app is swapping out the underlying connection
			nil,
		},
		{
			"invalid upgrade ordering",
			func() {
				counterpartyUpgrade.Fields.Ordering = channeltypes.ORDERED
			},
			channeltypes.ErrInvalidChannelOrdering,
		},
		{
			"invalid upgrade version",
			func() {
				counterpartyUpgrade.Fields.Version = ibctesting.InvalidID
			},
			types.ErrInvalidVersion,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			path = ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.Setup()

			// configure the channel upgrade to modify the underlying connection
			upgradePath := ibctesting.NewPath(suite.chainA, suite.chainB)
			upgradePath.SetupConnections()

			path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointA.ConnectionID}
			path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointB.ConnectionID}

			err := path.EndpointA.ChanUpgradeInit()
			suite.Require().NoError(err)

			counterpartyUpgrade = path.EndpointA.GetChannelUpgrade()

			tc.malleate()

			module, _, err := suite.chainB.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainB.GetContext(), types.PortID)
			suite.Require().NoError(err)

			app, ok := suite.chainB.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			cbs, ok := app.(porttypes.UpgradableModule)
			suite.Require().True(ok)

			version, err := cbs.OnChanUpgradeTry(
				suite.chainB.GetContext(), path.EndpointB.ChannelConfig.PortID, path.EndpointB.ChannelID,
				counterpartyUpgrade.Fields.Ordering, counterpartyUpgrade.Fields.ConnectionHops, counterpartyUpgrade.Fields.Version,
			)

			expPass := tc.expError == nil
			if expPass {
				suite.Require().NoError(err)
				suite.Require().Equal(types.Version, version)
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expError.Error())
			}
		})
	}
}

func (suite *TransferTestSuite) TestOnChanUpgradeAck() {
	var path *ibctesting.Path

	testCases := []struct {
		name     string
		malleate func()
		expError error
	}{
		{
			"success",
			func() {}, // successful happy path for a standalone transfer app is swapping out the underlying connection
			nil,
		},
		{
			"invalid upgrade version",
			func() {
				path.EndpointB.ChannelConfig.Version = ibctesting.InvalidID
			},
			types.ErrInvalidVersion,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			suite.SetupTest()

			path = ibctesting.NewTransferPath(suite.chainA, suite.chainB)
			path.Setup()

			// configure the channel upgrade to modify the underlying connection
			upgradePath := ibctesting.NewPath(suite.chainA, suite.chainB)
			upgradePath.SetupConnections()

			path.EndpointA.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointA.ConnectionID}
			path.EndpointB.ChannelConfig.ProposedUpgrade.Fields.ConnectionHops = []string{upgradePath.EndpointB.ConnectionID}

			err := path.EndpointA.ChanUpgradeInit()
			suite.Require().NoError(err)

			err = path.EndpointB.ChanUpgradeTry()
			suite.Require().NoError(err)

			tc.malleate()

			module, _, err := suite.chainA.App.GetIBCKeeper().PortKeeper.LookupModuleByPort(suite.chainA.GetContext(), types.PortID)
			suite.Require().NoError(err)

			app, ok := suite.chainA.App.GetIBCKeeper().Router.GetRoute(module)
			suite.Require().True(ok)

			cbs, ok := app.(porttypes.UpgradableModule)
			suite.Require().True(ok)

			err = cbs.OnChanUpgradeAck(suite.chainA.GetContext(), path.EndpointA.ChannelConfig.PortID, path.EndpointA.ChannelID, path.EndpointB.ChannelConfig.Version)

			expPass := tc.expError == nil
			if expPass {
				suite.Require().NoError(err)
			} else {
				suite.Require().Error(err)
				suite.Require().Contains(err.Error(), tc.expError.Error())
			}
		})
	}
}

func (suite *TransferTestSuite) TestPacketDataUnmarshalerInterface() {
	var (
		sender   = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()
		receiver = sdk.AccAddress(secp256k1.GenPrivKey().PubKey().Address()).String()

		data              []byte
		initialPacketData interface{}
	)

	testCases := []struct {
		name     string
		malleate func()
		expPass  bool
	}{
		{
			"success: valid packet data single denom -> multidenom conversion with memo",
			func() {
				initialPacketData = types.FungibleTokenPacketData{
					Denom:    ibctesting.TestCoin.Denom,
					Amount:   ibctesting.TestCoin.Amount.String(),
					Sender:   sender,
					Receiver: receiver,
					Memo:     "some memo",
				}

				expPacketData := multidenom.FungibleTokenPacketData{
					Tokens: []*multidenom.Token{
						{
							Denom:  ibctesting.TestCoin.Denom,
							Amount: ibctesting.TestCoin.Amount.String(),
							Trace:  []string{""},
						},
					},
					Sender:   sender,
					Receiver: receiver,
					Memo:     "some memo",
				}
				data = expPacketData.GetBytes()
			},
			true,
		},
		{
			"success: valid packet data single denom -> multidenom conversion without memo",
			func() {
				initialPacketData = types.FungibleTokenPacketData{
					Denom:    ibctesting.TestCoin.Denom,
					Amount:   ibctesting.TestCoin.Amount.String(),
					Sender:   sender,
					Receiver: receiver,
					Memo:     "",
				}
				expPacketData := multidenom.FungibleTokenPacketData{
					Tokens: []*multidenom.Token{
						{
							Denom:  ibctesting.TestCoin.Denom,
							Amount: ibctesting.TestCoin.Amount.String(),
							Trace:  []string{""},
						},
					},
					Sender:   sender,
					Receiver: receiver,
					Memo:     "",
				}
				data = expPacketData.GetBytes()
			},
			true,
		},
		{
			"success: valid packet data single denom with trace -> multidenom conversion with trace",
			func() {
				initialPacketData = types.FungibleTokenPacketData{
					Denom:    "transfer/channel-0/atom",
					Amount:   ibctesting.TestCoin.Amount.String(),
					Sender:   sender,
					Receiver: receiver,
					Memo:     "",
				}

				expPacketData := multidenom.FungibleTokenPacketData{
					Tokens: []*multidenom.Token{
						{
							Denom:  "atom",
							Amount: ibctesting.TestCoin.Amount.String(),
							Trace:  []string{"transfer/channel-0"},
						},
					},
					Sender:   sender,
					Receiver: receiver,
					Memo:     "",
				}
				data = expPacketData.GetBytes()
			},
			true,
		},
		{
			"success: valid packet data multidenom with memo",
			func() {
				initialPacketData = multidenom.FungibleTokenPacketData{
					Tokens: []*multidenom.Token{
						{
							Denom:  ibctesting.TestCoin.Denom,
							Amount: ibctesting.TestCoin.Amount.String(),
							Trace:  []string{""},
						},
					},
					Sender:   sender,
					Receiver: receiver,
					Memo:     "some memo",
				}

				data = initialPacketData.(multidenom.FungibleTokenPacketData).GetBytes()
			},
			true,
		},
		{
			"success: valid packet data multidenom without memo",
			func() {
				initialPacketData = multidenom.FungibleTokenPacketData{
					Tokens: []*multidenom.Token{
						{
							Denom:  ibctesting.TestCoin.Denom,
							Amount: ibctesting.TestCoin.Amount.String(),
							Trace:  []string{""},
						},
					},
					Sender:   sender,
					Receiver: receiver,
					Memo:     "",
				}

				data = initialPacketData.(multidenom.FungibleTokenPacketData).GetBytes()
			},
			true,
		},
		{
			"failure: invalid packet data",
			func() {
				data = []byte("invalid packet data")
			},
			false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		suite.Run(tc.name, func() {
			tc.malleate()

			packetData, err := transfer.IBCModule{}.UnmarshalPacketData(data)

			if tc.expPass {
				suite.Require().NoError(err)
				if _, ok := initialPacketData.(types.FungibleTokenPacketData); ok {
					// Note: testing of the denom trace parsing/conversion should be done as part of testing internal conversion functions
					suite.Require().Equal(initialPacketData.(types.FungibleTokenPacketData).Amount, packetData.(multidenom.FungibleTokenPacketData).Tokens[0].Amount)
					suite.Require().Equal(initialPacketData.(types.FungibleTokenPacketData).Sender, packetData.(multidenom.FungibleTokenPacketData).Sender)
					suite.Require().Equal(initialPacketData.(types.FungibleTokenPacketData).Receiver, packetData.(multidenom.FungibleTokenPacketData).Receiver)
					suite.Require().Equal(initialPacketData.(types.FungibleTokenPacketData).Memo, packetData.(multidenom.FungibleTokenPacketData).Memo)
				} else {
					suite.Require().Equal(initialPacketData.(multidenom.FungibleTokenPacketData), packetData.(multidenom.FungibleTokenPacketData))
				}
			} else {
				suite.Require().Error(err)
				suite.Require().Nil(packetData)
			}
		})
	}
}
