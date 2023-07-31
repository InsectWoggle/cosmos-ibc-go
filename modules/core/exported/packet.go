package exported

// PacketI defines the standard interface for IBC packets
type PacketI interface {
	GetSequence() uint64
	GetTimeoutHeight() Height
	GetTimeoutTimestamp() uint64
	GetSourcePort() string
	GetSourceChannel() string
	GetDestPort() string
	GetDestChannel() string
	GetData() []byte
	ValidateBasic() error
}

// Acknowledgement defines the interface used to return acknowledgements in the OnRecvPacket callback.
// The Acknowledgement interface is used by core IBC to ensure partial state changes are not committed
// when packet receives have not properly succeeded (typically resulting in an error acknowledgement being returned).
// The interface also allows core IBC to obtain the acknowledgement bytes whose encoding is determined by each IBC application or middleware.
// Each custom acknowledgement type must implement this interface.
type Acknowledgement interface {
	// Success determines if the IBC application state should be persisted when handling `RecvPacket`.
	// During `OnRecvPacket` IBC application callback execution, all state changes are held in a cache store and committed if:
	// - the acknowledgement.Success() returns true
	// - a nil acknowledgement is returned (asynchronous acknowledgements)
	//
	// Note 1: IBC application callback events are always persisted so long as `RecvPacket` succeeds without error.
	//
	// Note 2: The return value should account for the success of the underlying IBC application or middleware. Thus the  `acknowledgement.Success` is representative of the entire IBC stack's success when receiving a packet. The individual success of each acknowledgement associated with an IBC application or middleware must be determined by obtaining the actual acknowledgement type after decoding the acknowledgement bytes.
	//
	// See https://github.com/cosmos/ibc-go/blob/v7.0.0/docs/ibc/apps.md for further explanations.
	Success() bool
	Acknowledgement() []byte
}

// PacketData defines an optional interface which application's packet data may implement.
type PacketData interface {
	// GetPacketSender returns the sender address of the packet.
	// If the packet sender is unknown or undefined, an empty string should be returned.
	GetPacketSender(sourcePortID string) string
}
