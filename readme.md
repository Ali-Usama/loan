# A Sovereign Rollup using Rollkit
**loan** is a blockchain built using Cosmos SDK and Tendermint and created with [Ignite CLI](https://ignite.com/cli).


A loan is a financial transaction in which one party, the borrower, receives a certain amount of assets, such as money or digital tokens, and agrees to pay back the loan amount plus a fee to the lender by a predetermined deadline. To secure the loan, the borrower provides collateral, which may be seized by the lender if the borrower fails to pay back the loan as agreed.

## Properties
A loan has several properties that define its terms and conditions.

- The `id` is a unique identifier that is used to identify the loan on a blockchain.

- The `amount` is the amount of assets that are being lent to the borrower.

- The `fee` is the cost that the borrower must pay to the lender for the loan.

- The `collateral` is the asset or assets that the borrower provides to the lender as security for the loan.

- The `deadline` is the date by which the borrower must pay back the loan. If the borrower fails to pay back the loan by the deadline, the lender may choose to liquidate the loan and seize the collateral.

- The `state` of a loan describes the current status of the loan and can take on several values, such as `requested`, `approved`, `paid`, `cancelled`, or `liquidated`


In a loan transaction, there are two parties involved: the borrower and the lender. The borrower is the party that requests the loan and agrees to pay back the loan amount plus a fee to the lender by a predetermined deadline. The lender is the party that approves the loan request and provides the borrower with the loan amount.


## How to run

First, run the `local-celestia-devnet` by running the following command:

```
docker run --platform linux/amd64 -p 26650:26657 -p 26659:26659 ghcr.io/celestiaorg/local-celestia-devnet:main
```

Clone the repo: 
```
git clone https://github.com/Ali-Usama/loan.git
cd loan
bash init-local.sh
```

Keep this script running, while in another tab, test the App's functionality with following commands:
```
loand tx bank send alice bob 42069stake --keyring-backend test                /// Test transaction
loand tx loan request-loan 1000token 100token 1000foocoin 500 --from alice     /// For requesting a loan
loand tx loan request-loan 1000token 100token 1000foocoin 500 --from alice     /// For loan repayment
loand tx loan approve-loan 0 --from bob                                        /// Approve loan
loand q loan list-loan                                                         /// Query loan
```