package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateItem{}

func NewMsgCreateItem(creator string, name string, productType string, amount int32, price int32, discounted bool) *MsgCreateItem {
	return &MsgCreateItem{
		Creator:     creator,
		Name:        name,
		ProductType: productType,
		Amount:      amount,
		Price:       price,
		Discounted:  discounted,
	}
}

func (msg *MsgCreateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateItem{}

func NewMsgUpdateItem(creator string, id uint64, name string, productType string, amount int32, price int32, discounted bool) *MsgUpdateItem {
	return &MsgUpdateItem{
		Id:          id,
		Creator:     creator,
		Name:        name,
		ProductType: productType,
		Amount:      amount,
		Price:       price,
		Discounted:  discounted,
	}
}

func (msg *MsgUpdateItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteItem{}

func NewMsgDeleteItem(creator string, id uint64) *MsgDeleteItem {
	return &MsgDeleteItem{
		Id:      id,
		Creator: creator,
	}
}

func (msg *MsgDeleteItem) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
