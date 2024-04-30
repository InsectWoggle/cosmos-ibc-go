package types

import (
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	icatypes "github.com/cosmos/ibc-go/v7/modules/apps/27-interchain-accounts/types"
	channeltypes "github.com/cosmos/ibc-go/v7/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v7/modules/core/24-host"
)

var _ sdk.Msg = &MsgRegisterInterchainAccount{}

// NewMsgRegisterInterchainAccountWithOrdering creates a new instance of MsgRegisterInterchainAccount.
func NewMsgRegisterInterchainAccountWithOrdering(connectionID, owner, version string, ordering channeltypes.Order) *MsgRegisterInterchainAccount {
	return &MsgRegisterInterchainAccount{
		ConnectionId: connectionID,
		Owner:        owner,
		Version:      version,
		Ordering:     ordering,
	}
}

// NewMsgRegisterInterchainAccount creates a new instance of MsgRegisterInterchainAccount.
// It uses channeltypes.ORDERED as the default ordering. Breakage in v9.0.0 will allow the ordering to be provided
// directly. Use NewMsgRegisterInterchainAccountWithOrder to provide the ordering in previous versions.
func NewMsgRegisterInterchainAccount(connectionID, owner, version string) *MsgRegisterInterchainAccount {
	return &MsgRegisterInterchainAccount{
		ConnectionId: connectionID,
		Owner:        owner,
		Version:      version,
		Ordering:     channeltypes.ORDERED,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgRegisterInterchainAccount) ValidateBasic() error {
	if err := host.ConnectionIdentifierValidator(msg.ConnectionId); err != nil {
		return sdkerrors.Wrap(err, "invalid connection ID")
	}

	if strings.TrimSpace(msg.Owner) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgRegisterInterchainAccount) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}

// NewMsgSendTx creates a new instance of MsgSendTx
func NewMsgSendTx(owner, connectionID string, relativeTimeoutTimestamp uint64, packetData icatypes.InterchainAccountPacketData) *MsgSendTx {
	return &MsgSendTx{
		ConnectionId:    connectionID,
		Owner:           owner,
		RelativeTimeout: relativeTimeoutTimestamp,
		PacketData:      packetData,
	}
}

// ValidateBasic implements sdk.Msg
func (msg MsgSendTx) ValidateBasic() error {
	if err := host.ConnectionIdentifierValidator(msg.ConnectionId); err != nil {
		return sdkerrors.Wrap(err, "invalid connection ID")
	}

	if strings.TrimSpace(msg.Owner) == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "owner address cannot be empty")
	}

	if err := msg.PacketData.ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "invalid interchain account packet data")
	}

	if msg.RelativeTimeout == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "relative timeout cannot be zero")
	}

	return nil
}

// GetSigners implements sdk.Msg
func (msg MsgSendTx) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Owner)
	if err != nil {
		panic(err)
	}

	return []sdk.AccAddress{accAddr}
}
