package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerros "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) CancelLoan(goCtx context.Context, msg *types.MsgCancelLoan) (*types.MsgCancelLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id)
	if !found {
		return nil, sdkerros.Wrapf(sdkerros.ErrKeyNotFound, "key %d not found", msg.Id)
	}
	if loan.Borrower != msg.Creator {
		return nil, sdkerros.Wrap(sdkerros.ErrUnauthorized, "Cannot cancel: Not the borrower")
	}
	if loan.State != "requested" {
		return nil, sdkerros.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	collateral, _ := sdk.ParseCoinsNormalized(loan.Collateral)
	err := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, borrower, collateral)
	if err != nil {
		return nil, err
	}
	loan.State = "cancelled"
	k.SetLoan(ctx, loan)

	return &types.MsgCancelLoanResponse{}, nil
}
