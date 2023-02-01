import { GeneratedType } from "@cosmjs/proto-signing";
import { MsgRequestLoan } from "./types/loan/loan/tx";
import { MsgLiquidateLoan } from "./types/loan/loan/tx";
import { MsgApproveLoan } from "./types/loan/loan/tx";
import { MsgRepayLoan } from "./types/loan/loan/tx";
import { MsgCancelLoan } from "./types/loan/loan/tx";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/loan.loan.MsgRequestLoan", MsgRequestLoan],
    ["/loan.loan.MsgLiquidateLoan", MsgLiquidateLoan],
    ["/loan.loan.MsgApproveLoan", MsgApproveLoan],
    ["/loan.loan.MsgRepayLoan", MsgRepayLoan],
    ["/loan.loan.MsgCancelLoan", MsgCancelLoan],
    
];

export { msgTypes }