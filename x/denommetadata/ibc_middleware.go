package denommetadata

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	errortypes "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitytypes "github.com/cosmos/ibc-go/modules/capability/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v8/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	"github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/dymensionxyz/gerr-cosmos/gerrc"

	"github.com/dymensionxyz/sdk-utils/utils/uevent"
	"github.com/dymensionxyz/sdk-utils/utils/uibc"

	"github.com/dymensionxyz/dymension/v3/x/denommetadata/types"
)

var _ porttypes.IBCModule = &IBCModule{}

// IBCModule implements the ICS26 callbacks for the transfer middleware
type IBCModule struct {
	porttypes.IBCModule
	keeper        types.DenomMetadataKeeper
	rollappKeeper types.RollappKeeper
}

// NewIBCModule creates a new IBCModule given the keepers and underlying application
func NewIBCModule(
	app porttypes.IBCModule,
	keeper types.DenomMetadataKeeper,
	rollappKeeper types.RollappKeeper,
) IBCModule {
	return IBCModule{
		IBCModule:     app,
		keeper:        keeper,
		rollappKeeper: rollappKeeper,
	}
}

// OnRecvPacket registers the denom metadata if it does not exist.
// It will intercept an incoming packet and check if the denom metadata exists.
// If it does not, it will register the denom metadata.
// The handler will expect a 'denom_metadata' object in the memo field of the packet.
// If the memo is not an object, or does not contain the metadata, it moves on to the next handler.
func (im IBCModule) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	transferData, err := im.rollappKeeper.GetValidTransfer(ctx, packet.Data, packet.DestinationPort, packet.DestinationChannel)
	if err != nil {
		return uevent.NewErrorAcknowledgement(ctx, err)
	}

	rollapp, packetData := transferData.Rollapp, transferData.FungibleTokenPacketData

	if packetData.Memo == "" {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	// at this point it's safe to assume that we are not handling a native token of the rollapp
	denomTrace := uibc.GetForeignDenomTrace(packet.GetDestChannel(), packetData.Denom)
	ibcDenom := denomTrace.IBCDenom()

	dm := types.ParsePacketMetadata(packetData.Memo)
	if dm == nil {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	if err = dm.Validate(); err != nil {
		return uevent.NewErrorAcknowledgement(ctx, err)
	}

	if dm.Base != packetData.Denom {
		return uevent.NewErrorAcknowledgement(ctx, gerrc.ErrInvalidArgument)
	}

	// if denom metadata was found in the memo, it means we should have the rollapp record
	if rollapp == nil {
		return uevent.NewErrorAcknowledgement(ctx, gerrc.ErrNotFound)
	}

	if im.keeper.HasDenomMetadata(ctx, ibcDenom) {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	// adjust the denom metadata with the IBC denom
	dm.Base = ibcDenom
	dm.DenomUnits[0].Denom = dm.Base
	if err = im.keeper.CreateDenomMetadata(ctx, *dm); err != nil {
		// TODO: remove? already checked above
		if errorsmod.IsOf(err, gerrc.ErrAlreadyExists) {
			return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
		}
		return uevent.NewErrorAcknowledgement(ctx, err)
	}

	return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
}

// OnAcknowledgementPacket adds the token metadata to the rollapp if it doesn't exist
// It marks the completion of the denom metadata registration process on the rollapp
func (im IBCModule) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var ack channeltypes.Acknowledgement
	if err := transfertypes.ModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return errorsmod.Wrapf(errortypes.ErrJSONUnmarshal, "unmarshal ICS-20 transfer packet acknowledgement: %v", err)
	}

	if !ack.Success() {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	transferData, err := im.rollappKeeper.GetValidTransfer(ctx, packet.Data, packet.GetSourcePort(), packet.GetSourceChannel())
	if err != nil {
		return errorsmod.Wrapf(errortypes.ErrInvalidRequest, "get valid transfer data: %s", err.Error())
	}

	rollapp, packetData := transferData.Rollapp, transferData.FungibleTokenPacketData

	if packetData.Memo == "" {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	dm := types.ParsePacketMetadata(packetData.Memo)
	if dm == nil {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	if err = dm.Validate(); err != nil {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	// if denom metadata was found in the memo, it means we should have the rollapp record
	if rollapp == nil {
		return gerrc.ErrNotFound
	}

	// TODO: simplify: can do with just Set*
	has, err := im.rollappKeeper.HasRegisteredDenom(ctx, rollapp.RollappId, dm.Base)
	if err != nil {
		return errorsmod.Wrapf(errortypes.ErrKeyNotFound, "check if rollapp has registered denom: %s", err.Error())
	}
	if !has {
		// add the new token denom base to the list of rollapp's registered denoms
		if err = im.rollappKeeper.SetRegisteredDenom(ctx, rollapp.RollappId, dm.Base); err != nil {
			return errorsmod.Wrapf(errortypes.ErrKeyNotFound, "set registered denom: %s", err.Error())
		}
	}

	return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// ICS4Wrapper intercepts outgoing IBC packets and adds token metadata to the memo if the rollapp doesn't have it.
// This is a solution for adding token metadata to fungible tokens transferred over IBC,
// targeted at rollapps that don't have the token metadata for the token being transferred.
// More info here: https://www.notion.so/dymension/ADR-x-IBC-Denom-Metadata-Transfer-From-Hub-to-Rollapp-d3791f524ac849a9a3eb44d17968a30b
type ICS4Wrapper struct {
	porttypes.ICS4Wrapper

	rollappKeeper types.RollappKeeper
	bankKeeper    types.BankKeeper
}

// NewICS4Wrapper creates a new ICS4Wrapper
func NewICS4Wrapper(
	ics porttypes.ICS4Wrapper,
	rollappKeeper types.RollappKeeper,
	bankKeeper types.BankKeeper,
) *ICS4Wrapper {
	return &ICS4Wrapper{
		ICS4Wrapper:   ics,
		rollappKeeper: rollappKeeper,
		bankKeeper:    bankKeeper,
	}
}

// SendPacket wraps IBC ChannelKeeper's SendPacket function
func (m *ICS4Wrapper) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	sourcePort string,
	sourceChannel string,
	timeoutHeight clienttypes.Height,
	timeoutTimestamp uint64,
	data []byte,
) (sequence uint64, err error) {
	packet := new(transfertypes.FungibleTokenPacketData)
	if err = transfertypes.ModuleCdc.UnmarshalJSON(data, packet); err != nil {
		return 0, errorsmod.Wrapf(errortypes.ErrJSONUnmarshal, "unmarshal ICS-20 transfer packet data: %s", err.Error())
	}

	if types.MemoHasPacketMetadata(packet.Memo) {
		return 0, types.ErrMemoDenomMetadataAlreadyExists
	}

	transferData, err := m.rollappKeeper.GetValidTransfer(ctx, data, sourcePort, sourceChannel)
	if err != nil {
		return 0, errorsmod.Wrapf(errortypes.ErrInvalidRequest, "get valid transfer data: %s", err.Error())
	}

	rollapp := transferData.Rollapp
	// TODO: currently we check if receiving chain is a rollapp, consider that other chains also might want this feature
	// meaning, find a better way to check if the receiving chain supports this middleware
	if rollapp == nil {
		return m.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	}

	if transfertypes.ReceiverChainIsSource(sourcePort, sourceChannel, packet.Denom) {
		return m.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	}

	// Check if the rollapp already contains the denom metadata by matching the base of the denom metadata.
	// At the first match, we assume that the rollapp already contains the metadata.
	// It would be technically possible to have a race condition where the denom metadata is added to the rollapp
	// from another packet before this packet is acknowledged.
	// The value of `packet.Denom` here can be one of two things:
	// 		1. Base denom (e.g. "adym") for the native token of the hub, and
	// 		2. IBC trace (e.g. "transfer/channel-1/arax") for a third party token.
	// We need to handle both cases:
	// 		1. We use the value of `packet.Denom` as the baseDenom
	//		2. We parse the IBC denom trace into IBC denom hash and prepend it with "ibc/" to get the baseDenom
	baseDenom := transfertypes.ParseDenomTrace(packet.Denom).IBCDenom() // TODO: rename base denom to ibc denom https://github.com/dymensionxyz/dymension/issues/1650

	has, err := m.rollappKeeper.HasRegisteredDenom(ctx, rollapp.RollappId, baseDenom)
	if err != nil {
		return 0, errorsmod.Wrapf(errortypes.ErrKeyNotFound, "check if rollapp has registered denom: %s", err.Error()) /// TODO: no .Error()
	}
	if has {
		return m.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	}

	// get the denom metadata from the bank keeper, if it doesn't exist, move on to the next middleware in the chain
	denomMetadata, ok := m.bankKeeper.GetDenomMetaData(ctx, baseDenom)
	if !ok {
		return m.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
	}

	packet.Memo, err = types.AddDenomMetadataToMemo(packet.Memo, denomMetadata)
	if err != nil {
		return 0, errorsmod.Wrapf(gerrc.ErrInvalidArgument, "add denom metadata to memo: %s", err.Error()) /// TODO: no .Error()
	}

	data, err = transfertypes.ModuleCdc.MarshalJSON(packet)
	if err != nil {
		return 0, errorsmod.Wrapf(errortypes.ErrJSONMarshal, "marshal ICS-20 transfer packet data: %s", err.Error()) /// TODO: no .Error()
	}

	return m.ICS4Wrapper.SendPacket(ctx, chanCap, sourcePort, sourceChannel, timeoutHeight, timeoutTimestamp, data)
}
