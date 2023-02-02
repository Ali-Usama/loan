package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"loan/x/loan/types"
)

func (k msgServer) ApproveLoan(goCtx context.Context, msg *types.MsgApproveLoan) (*types.MsgApproveLoanResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	loan, found := k.GetLoan(ctx, msg.Id) // retrieves a loan object
	if !found {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "key %d doesn't exist", msg.Id)
	}
	if loan.State != "requested" {
		return nil, sdkerrors.Wrapf(types.ErrWrongLoanState, "%v", loan.State)
	}

	lender, _ := sdk.AccAddressFromBech32(msg.Creator)
	borrower, _ := sdk.AccAddressFromBech32(loan.Borrower)
	amount, err := sdk.ParseCoinsNormalized(loan.Amount)
	if err != nil {
		return nil, sdkerrors.Wrap(types.ErrWrongLoanState, "Cannot parse coins into loan amount")
	}

	err = k.bankKeeper.SendCoins(ctx, lender, borrower, amount)
	if err != nil {
		return nil, err
	}

	loan.Lender = msg.Creator
	loan.State = "approved"
	k.SetLoan(ctx, loan)

	return &types.MsgApproveLoanResponse{}, nil
}
