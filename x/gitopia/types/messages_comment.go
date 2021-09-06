package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateComment{}

func NewMsgCreateComment(creator string, parentId uint64, body string, attachments []string, diffHunk string, path string, system bool, authorAssociation string, commentType string) *MsgCreateComment {
	return &MsgCreateComment{
		Creator:           creator,
		ParentId:          parentId,
		Body:              body,
		Attachments:       attachments,
		DiffHunk:          diffHunk,
		Path:              path,
		System:            system,
		AuthorAssociation: authorAssociation,
		CommentType:       commentType,
	}
}

func (msg *MsgCreateComment) Route() string {
	return RouterKey
}

func (msg *MsgCreateComment) Type() string {
	return "CreateComment"
}

func (msg *MsgCreateComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgCreateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgCreateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	_, exists := Comment_Type_value[msg.CommentType]
	if !exists {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid commentType (%s)", msg.CommentType)
	}
	return nil
}

var _ sdk.Msg = &MsgUpdateComment{}

func NewMsgUpdateComment(creator string, id uint64, body string, attachments []string) *MsgUpdateComment {
	return &MsgUpdateComment{
		Id:          id,
		Creator:     creator,
		Body:        body,
		Attachments: attachments,
	}
}

func (msg *MsgUpdateComment) Route() string {
	return RouterKey
}

func (msg *MsgUpdateComment) Type() string {
	return "UpdateComment"
}

func (msg *MsgUpdateComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgUpdateComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgUpdateComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}

var _ sdk.Msg = &MsgDeleteComment{}

func NewMsgDeleteComment(creator string, id uint64) *MsgDeleteComment {
	return &MsgDeleteComment{
		Id:      id,
		Creator: creator,
	}
}
func (msg *MsgDeleteComment) Route() string {
	return RouterKey
}

func (msg *MsgDeleteComment) Type() string {
	return "DeleteComment"
}

func (msg *MsgDeleteComment) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgDeleteComment) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgDeleteComment) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
