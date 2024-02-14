package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateTransaction{}

func NewMsgCreateTransaction(creator string, amount string, remarks string) *MsgCreateTransaction {
	return &MsgCreateTransaction{
		Creator: creator,
		Amount:  amount,
		Remarks: remarks,
	}
}

func (msg *MsgCreateTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateTransaction{}

func NewMsgUpdateTransaction(creator string, id uint64, amount string, remarks string) *MsgUpdateTransaction {
	return &MsgUpdateTransaction{
		Id:      id,
		Creator: creator,
		Amount:  amount,
		Remarks: remarks,
	}
}

func (msg *MsgUpdateTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteTransaction{}

func NewMsgDeleteTransaction(creator string, id uint64) *MsgDeleteTransaction {
	return &MsgDeleteTransaction{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteTransaction) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
