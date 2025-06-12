import { LCDClient, MsgSend, MnemonicKey, Coins, MsgExecuteContract, Fee, SignMode } from '@terra-money/feather.js';
import { AddTransactionRequest, Transaction } from "../models/transaction";
import { TransactionRepository } from "../repository/transaction.repository";
import { CustomError } from "../types/error";
import { WalletRepository } from "../repository/wallet.repository";
import { GetLunaBalanceByAddress, GetLunaBalanceByCoinIdAndUserId, MAX_BALANCE_RETRIES, terra, chainID } from "./contant"
import { randomAlphabet, randomAlphanumeric } from '../utils/rand';
import { Wallet } from '../models/wallet';

export class TransactionService {
    private transactionRepository: TransactionRepository;
    private walletRepository: WalletRepository;

    constructor(tp: TransactionRepository, wp: WalletRepository) {
        this.transactionRepository = tp;
        this.walletRepository = wp;
    }

    public async checkTransactionExists(tx_id: string): Promise<boolean> {
        return await this.transactionRepository.checkTransactionExists(tx_id);
    }

    public async getTransactionByTxId(tx_id: string) {
        const transaction = await this.transactionRepository.getTransactionByTxId(tx_id);

        if (!transaction)
            throw new CustomError('Transaction not found', 404);

        return transaction;
    }

    public async getTransacitonsByUserId(userId: number) {
        const transactions = await this.transactionRepository.getTransactionsByUserId(userId);

        if (!transactions)
            throw new CustomError('Transaction not found', 404);

        return transactions;
    }

    public async getTransactionsByWalletId(walletId: number) {
        const transactions = await this.transactionRepository.getTransactionsByWalletId(walletId);

        if (!transactions)
            throw new CustomError('Transaction not found', 404);

        return transactions;
    }

    public async getTransactionsByWalletIdAndUserId(walletId: number, userId: number) {
        const transactions = await this.transactionRepository.getTransactionsByWalletIdAndUserId(walletId, userId);

        if (!transactions)
            throw new CustomError('Transaction not found', 404);

        return transactions;
    }

    public async createTransaction(transactionDetails: AddTransactionRequest) {
        const walletResult = await this.walletRepository.getWalletByWalletId(transactionDetails.from_wallet_id!);
        if (!walletResult)
            throw new CustomError('Wallet not found', 404)
        console.log("Transfer ", walletResult, transactionDetails)

        if (walletResult.user_id != 0) {
            if (walletResult.user_id !== transactionDetails.from_wallet_user_id)
                throw new CustomError('User not authorized', 401);
        }
        if (!walletResult)
            throw new CustomError('Wallet not found', 404);

        const latestBalance = await GetLunaBalanceByCoinIdAndUserId(this.walletRepository, walletResult.coin_id!, walletResult.user_id!)
        if (latestBalance <= transactionDetails.amount!) {
            throw new CustomError('Insufficient balance', 400);
        }


        // const send = new MsgExecuteContract(
        //     walletResult.address!,
        //     transactionDetails.to_wallet_address!,
        //     {
        //         transfer: {
        //             recipient: transactionDetails.to_wallet_address!,
        //             amount: transactionDetails.amount!.toString(),  // Ensure amount is string
        //         },
        //     },
        // );

        let currencyName: string;
        try {
            switch (walletResult.coin_id) {
                case 1:
                    await this.sendLuncTransaction(walletResult, transactionDetails)
                    currencyName = "uluna"
                    break;
                default:
                    throw new CustomError('Coin not supported', 400);
            }
        } catch (error) {
            throw error
        }
        const transaction: Transaction = {
            tx_id: randomAlphabet(60),
            amount: transactionDetails.amount!,
            currency_id: walletResult.coin_id!,
            from_wallet_id: walletResult.id!,
            from_address: walletResult.address!,
            from_public_key: walletResult.public_key!,
            to_address: transactionDetails.to_wallet_address!,
            currency_name: "lunc",
            transaction_at: new Date(),
        };
        console.log(transaction)
        const response = await this.transactionRepository.createTransaction(transaction)
        return response;

    }

    public async updateTransaction(transaction: Transaction) {
        return await this.transactionRepository.updateTransaction(transaction);
    }

    public async deleteTransactionByTxId(tx_id: string) {
        return await this.transactionRepository.deleteTransactionByTxId(tx_id);
    }

    public async deleteTransactionsByWalletId(walletId: number) {
        return await this.transactionRepository.deleteTransactionsByWalletId(walletId);
    }

    public async getTransactionsByWalletIdAndUserIdAndPagination(walletId: number, userId: number, limit: number, offset: number, total: boolean) {
        const transactions = await this.transactionRepository.getTransactionsByWalletIdAndUserIdAndPagination(walletId, userId, limit, offset);

        if (!transactions)
            throw new CustomError('Transaction not found', 404);

        if (total) {
            const count = await this.transactionRepository.getTransactionsByWalletIdAndUserIdAndPaginationCount(walletId, userId);
            return { transactions, count };
        }

        return { transactions };
    }

    public async getTransactionsByUserIdAndPagination(userId: number, limit: number, offset: number, total: boolean) {
        const transactions = await this.transactionRepository.getTransactionsByUserIdAndPagination(userId, limit, offset);

        if (!transactions)
            throw new CustomError('Transaction not found', 404);

        if (total) {
            const count = await this.transactionRepository.getTransactionsByUserIdAndPaginationCount(userId);
            return { transactions, count };
        }


        return { transactions };
    }

    public async preTransaction(transactionDetails: AddTransactionRequest): Promise<Fee> {
        const walletResult = await this.walletRepository.getWalletByWalletId(transactionDetails.from_wallet_id!);
        if (!walletResult)
            throw new CustomError('Wallet not found', 404)
        console.log("Transfer ", walletResult, transactionDetails)

        if (walletResult.user_id != 0) {
            if (walletResult.user_id !== transactionDetails.from_wallet_user_id)
                throw new CustomError('User not authorized', 401);
        }
        if (!walletResult)
            throw new CustomError('Wallet not found', 404);

        const latestBalance = await GetLunaBalanceByCoinIdAndUserId(this.walletRepository, walletResult.coin_id!, walletResult.user_id!)
        if (latestBalance <= transactionDetails.amount!) {
            throw new CustomError('Insufficient balance', 400);
        }
        const mk = new MnemonicKey({
            mnemonic: walletResult.mnemonic!,
        });
      const accountInfo = await terra.auth.accountInfo(mk.accAddress('terra'));
        const send = new MsgSend(
            mk.accAddress('terra'),
            transactionDetails.to_wallet_address!,
            { uluna: transactionDetails.amount! }
        );


const estimatedFee = await terra.tx.estimateFee(
    [{ sequenceNumber: accountInfo.getSequenceNumber(), publicKey: mk.publicKey }],
    {
        chainID: chainID,
        msgs: [send],
    }
);

        const fee = new Fee(estimatedFee.gas_limit, { uluna: estimatedFee.amount.get('uluna')?.amount! })
        return fee
    }
    protected async sendLuncTransaction(walletResult: Wallet, transactionDetails: AddTransactionRequest) {

        const mk = new MnemonicKey({
            mnemonic: walletResult.mnemonic!,
        });

        const wallet = terra.wallet(mk);
        let accountInfo = await terra.auth.accountInfo(mk.accAddress('terra'));
        let sequenceNumber = accountInfo.getSequenceNumber();
        console.log("#sequenceNumber1", sequenceNumber)
        const send = new MsgSend(
            mk.accAddress('terra'),
            transactionDetails.to_wallet_address!,
            { uluna: transactionDetails.amount! }
        );

        let tx;
        try {
            console.log("#1");


            const estimatedFee = await terra.tx.estimateFee(
                [{ sequenceNumber: sequenceNumber, publicKey: mk.publicKey }],
                {
                    chainID: chainID,
                    msgs: [send],

                }
            );
            // const estimatedFee = new Fee(91414, { uluna: "2000" }); // Adjust values as needed

            console.log("#2");

            console.log("Estimated Fee:", estimatedFee.toData());

            const amountInUluna = Number(transactionDetails.amount); // Convert to number
            const minFee = 100;  // Minimum fee in uluna
            const percentageFee = amountInUluna * 0.002; // Example: 0.2% of the transaction
            const finalFeeAmount = Math.max(minFee, Math.round(percentageFee)); // Ensure minimum fee

            console.log(`Final Fee: ${finalFeeAmount} uluna`);
            console.log("#3");
            // await Delay(200);
            tx = await wallet.createAndSignTx({
                msgs: [send],
                memo: randomAlphanumeric(32),
                chainID: chainID,
                sequence: sequenceNumber,
                fee: new Fee(estimatedFee.gas_limit, { uluna: estimatedFee.amount.get('uluna')?.amount! }),
            });

            console.log("#4");

            if (!tx) throw new CustomError("Transaction creation failed", 500);

        } catch (error) {
            console.error("Transaction Error:", error);
            throw new CustomError(`Transaction failed: ${error}`, 500);
        }
        try {
            console.log("#5")
            const result = await terra.tx.broadcastAsync(tx, chainID);
            console.log(`TX Full: `, tx, result);
            console.log(`TX hash: ${result.txhash} From ${walletResult.address} To ${transactionDetails.to_wallet_address} Amount ${transactionDetails.amount}`);
        } catch (error) {
            throw new CustomError(`Transaction failed to broadcast ${error}`, 500);
        }

        try {
            await GetLunaBalanceByAddress(this.walletRepository, transactionDetails.to_wallet_address!);
        } catch (error) {
            console.warn("Error fetching Luna balance by address. Ignoring error:", error);
        }
        try {
            await GetLunaBalanceByCoinIdAndUserId(this.walletRepository, walletResult.coin_id!, walletResult.user_id!);

        } catch (error) {
            console.warn("Error fetching Luna balance by Coin ID and User ID. Ignoring error:", error);
        }
        console.log("__()##@@##()__")
        return tx;
    }
}

