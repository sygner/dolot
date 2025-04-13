import { AddTransactionRequest, BooleanResult, getTransactionsByUserIdAndPaginationRequest, getTransactionsByWalletIdAndUserIdAndPaginationRequest, GetTransactionsByWalletIdRequest, PreTransactionDetail, Transaction, TransactionId, Transactions, UserId, WalletId } from "../../proto/api/service_pb";
import { TransactionService } from "../services/transaction"
import * as grpc from '@grpc/grpc-js';
import { CustomError } from "../types/error";
import { getUnixTimestamp } from "../utils/time";
import { AddTransactionRequest as AddTransactionRequestModel } from "../models/transaction";

export class TransactionHandler {
    private transactionService: TransactionService;

    constructor(ts: TransactionService) {
        this.transactionService = ts;
        this.checkTransactionExists = this.checkTransactionExists.bind(this);
        this.getTransactionByTxId = this.getTransactionByTxId.bind(this);
        this.getTransactionsByWalletId = this.getTransactionsByWalletId.bind(this);
        this.AddTransaction = this.AddTransaction.bind(this);
        this.getTransactionsByUserId = this.getTransactionsByUserId.bind(this);
        this.getPreTransactionDetail = this.getPreTransactionDetail.bind(this);
        this.getTransactionsByWalletIdAndUserIdAndPagination = this.getTransactionsByWalletIdAndUserIdAndPagination.bind(this);
        this.getTransactionsByUserIdAndPagination = this.getTransactionsByUserIdAndPagination.bind(this);
    }

    public async checkTransactionExists(
        call: grpc.ServerUnaryCall<TransactionId, BooleanResult>,
        callback: grpc.sendUnaryData<BooleanResult>
    ) {
        try {
            const tx_id = call.request.getTxId();
            const exists = await this.transactionService.checkTransactionExists(tx_id);
            callback(null, new BooleanResult().setResult(exists));
        } catch (error) {
            console.error('Error in checkTransactionExists:', error);

            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }



    public async getTransactionByTxId(
        call: grpc.ServerUnaryCall<TransactionId, Transaction>,
        callback: grpc.sendUnaryData<Transaction>
    ) {
        try {
            const tx_id = call.request.getTxId();
            const result = await this.transactionService.getTransactionByTxId(tx_id);
            const transaction = new Transaction();
            transaction.setTxId(result.tx_id!);
            transaction.setCurrencyId(result.currency_id!);
            transaction.setCurrencyName(result.currency_name!);
            transaction.setAmount(result.amount!);
            transaction.setFromAddress(result.from_address!);
            transaction.setToAddress(result.to_address!);
            transaction.setFromWalletId(result.from_wallet_id!);
            transaction.setFromPublicKey(result.from_public_key!);
            transaction.setTransactionAt(getUnixTimestamp(result.transaction_at!).toString());
            callback(null, transaction);
        } catch (error) {
            console.error('Error in getTransactionByTxId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async getTransactionsByWalletId(
        call: grpc.ServerUnaryCall<GetTransactionsByWalletIdRequest, Transactions>,
        callback: grpc.sendUnaryData<Transactions>
    ) {
        try {
            const walletId = call.request.getWalletId();
            const userId = call.request.getUserId();
            const result = await this.transactionService.getTransactionsByWalletIdAndUserId(walletId, userId);
            const transactions = new Transactions();
            result.forEach((transaction) => {
                const t = new Transaction();
                t.setTxId(transaction.tx_id!);
                t.setCurrencyId(transaction.currency_id!);
                t.setCurrencyName(transaction.currency_name!);
                t.setAmount(transaction.amount!);
                t.setFromAddress(transaction.from_address!);
                t.setToAddress(transaction.to_address!);
                t.setFromWalletId(transaction.from_wallet_id!);
                t.setFromPublicKey(transaction.from_public_key!);
                t.setTransactionAt(getUnixTimestamp(transaction.transaction_at!).toString());
                transactions.addTransactions(t);
            });

            callback(null, transactions);
        } catch (error) {
            console.error('Error in getTransactionsByWalletId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async getTransactionsByUserId(
        call: grpc.ServerUnaryCall<UserId, Transactions>,
        callback: grpc.sendUnaryData<Transactions>
    ) {
        try {
            const userId = call.request.getUserId();
            const result = await this.transactionService.getTransacitonsByUserId(userId);
            const transactions = new Transactions();
            result.forEach((transaction) => {
                const t = new Transaction();
                t.setTxId(transaction.tx_id!);
                t.setCurrencyId(transaction.currency_id!);
                t.setCurrencyName(transaction.currency_name!);
                t.setAmount(transaction.amount!);
                t.setFromAddress(transaction.from_address!);
                t.setToAddress(transaction.to_address!);
                t.setFromWalletId(transaction.from_wallet_id!);
                t.setFromPublicKey(transaction.from_public_key!);
                t.setTransactionAt(getUnixTimestamp(transaction.transaction_at!).toString());
                transactions.addTransactions(t);
            });

            callback(null, transactions);
        } catch (error) {
            console.error('Error in getTransactionsByWalletId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async getPreTransactionDetail(
        call: grpc.ServerUnaryCall<AddTransactionRequest, PreTransactionDetail>,
        callback: grpc.sendUnaryData<PreTransactionDetail>
    ) {
        try {
            const request: AddTransactionRequestModel =
            {
                amount: call.request.getAmount(),
                from_wallet_id: call.request.getFromWalletId(),
                from_wallet_user_id: call.request.getFromWalletUserId(),
                to_wallet_address: call.request.getToWalletAddress()
            }
            const result = await this.transactionService.preTransaction(request);
            let detail = new PreTransactionDetail()

            const amount = result.amount.get('uluna')?.amount!.toNumber()!;
            const gas_limit = result.gas_limit;
            detail.setAmount(amount);
            detail.setGasLimit(gas_limit);

            callback(null, detail);
        } catch (error) {
            console.error('Error in AddTransaction:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async getTransactionsByWalletIdAndUserIdAndPagination(
        call: grpc.ServerUnaryCall<getTransactionsByWalletIdAndUserIdAndPaginationRequest, Transactions>,
        callback: grpc.sendUnaryData<Transactions>
    ) {
        try {
            const walletId = call.request.getWalletId();
            const userId = call.request.getUserId();
            const limit = call.request.getPagination()?.getLimit() || 10;
            const offset = call.request.getPagination()?.getOffset() || 0;
            const total = call.request.getPagination()?.getTotal() || false;
            const result = await this.transactionService.getTransactionsByWalletIdAndUserIdAndPagination(walletId, userId, limit, offset,total);
            console.log(result.count!)
            const transactions = new Transactions();
            result.transactions.forEach((transaction) => {
                const t = new Transaction();
                t.setTxId(transaction.tx_id!);
                t.setCurrencyId(transaction.currency_id!);
                t.setCurrencyName(transaction.currency_name!);
                t.setAmount(transaction.amount!);
                t.setFromAddress(transaction.from_address!);
                t.setToAddress(transaction.to_address!);
                t.setFromWalletId(transaction.from_wallet_id!);
                t.setFromPublicKey(transaction.from_public_key!);
                t.setTransactionAt(getUnixTimestamp(transaction.transaction_at!).toString());
                transactions.addTransactions(t);
            });
            if (total){
                transactions.setTotal(result.count!);
            }
            callback(null, transactions);

        } catch (error) {
            console.error('Error in AddTransaction:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async getTransactionsByUserIdAndPagination(
        call: grpc.ServerUnaryCall<getTransactionsByUserIdAndPaginationRequest, Transactions>,
        callback: grpc.sendUnaryData<Transactions>
    ) {
        try {
            const userId = call.request.getUserId();
            const limit = call.request.getPagination()?.getLimit() || 10;
            const offset = call.request.getPagination()?.getOffset() || 0;
            const total = call.request.getPagination()?.getTotal() || false;
            const result = await this.transactionService.getTransactionsByUserIdAndPagination(userId, limit, offset,total);
            const transactions = new Transactions();
            result.transactions.forEach((transaction) => {
                const t = new Transaction();
                t.setTxId(transaction.tx_id!);
                t.setCurrencyId(transaction.currency_id!);
                t.setCurrencyName(transaction.currency_name!);
                t.setAmount(transaction.amount!);
                t.setFromAddress(transaction.from_address!);
                t.setToAddress(transaction.to_address!);
                t.setFromWalletId(transaction.from_wallet_id!);
                t.setFromPublicKey(transaction.from_public_key!);
                t.setTransactionAt(getUnixTimestamp(transaction.transaction_at!).toString());
                transactions.addTransactions(t);
            });
            if (total){
                transactions.setTotal(result.count!);
            }
            callback(null, transactions);

        } catch (error) {
            console.error('Error in AddTransaction:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    public async AddTransaction(
        call: grpc.ServerUnaryCall<AddTransactionRequest, Transaction>,
        callback: grpc.sendUnaryData<Transaction>
    ) {
        try {
            const request: AddTransactionRequestModel =
            {
                amount: call.request.getAmount(),
                from_wallet_id: call.request.getFromWalletId(),
                from_wallet_user_id: call.request.getFromWalletUserId(),
                to_wallet_address: call.request.getToWalletAddress()
            }
            const result = await this.transactionService.createTransaction(request);
            const transaction = new Transaction();
            transaction.setTxId(result.tx_id!);
            transaction.setCurrencyId(result.currency_id!);
            transaction.setCurrencyName(result.currency_name!);
            transaction.setAmount(result.amount!);
            transaction.setFromAddress(result.from_address!);
            transaction.setToAddress(result.to_address!);
            transaction.setFromWalletId(result.from_wallet_id!);
            transaction.setFromPublicKey(result.from_public_key!);
            transaction.setTransactionAt(getUnixTimestamp(result.transaction_at!).toString());
            console.log("Ready to return grpc")
            callback(null, transaction);

        } catch (error) {
            console.error('Error in AddTransaction:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

}