package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgCreateSporechainEntry = "create_sporechain_entry"

var _ sdk.Msg = &MsgCreateSporechainEntry{}

func NewMsgCreateSporechainEntry(creator string, userAddress string, contentHash string, time string) *MsgCreateSporechainEntry {
  return &MsgCreateSporechainEntry{
		Creator: creator,
    UserAddress: userAddress,
    ContentHash: contentHash,
    Time: time,
	}
}

func (msg *MsgCreateSporechainEntry) Route() string {
  return RouterKey
}

func (msg *MsgCreateSporechainEntry) Type() string {
  return TypeMsgCreateSporechainEntry
}

func (msg *MsgCreateSporechainEntry) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgCreateSporechainEntry) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateSporechainEntry) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

