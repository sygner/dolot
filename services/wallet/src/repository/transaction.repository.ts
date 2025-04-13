import { Repository } from "./repository";
import { Transaction } from "../models/transaction"; // Assuming there's a Transaction model similar to Wallet

export class TransactionRepository extends Repository {

    public async checkTransactionExists(tx_id: string): Promise<boolean> {
        const query = `SELECT * FROM transactions WHERE tx_id = $1`;
        const values = [tx_id];
        const result = await this.executeQuery(query, values);
        return result?.rowCount ? result.rowCount > 0 : false;
    }

    public async getTransactionByTxId(tx_id: string): Promise<Transaction> {
        const query = `SELECT * FROM transactions WHERE tx_id = $1`;
        const values = [tx_id];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async getTransactionsByWalletId(walletId: number): Promise<Array<Transaction>> {
        const query = `SELECT * FROM transactions WHERE from_wallet_id = $1 ORDER BY transaction_at DESC`;
        const values = [walletId];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async getTransactionsByWalletIdAndUserId(walletId: number, userId: number): Promise<Array<Transaction>> {
        console.log(walletId,userId)
        const query = `SELECT * FROM transactions WHERE from_wallet_id = $1 AND EXISTS (SELECT 1 FROM wallets WHERE user_id = $2) ORDER BY transaction_at DESC`;
        const values = [walletId, userId];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async getTransactionsByWalletIdAndUserIdAndPagination(walletId: number, userId: number, limit: number, offset: number): Promise<Array<Transaction>> {
        const query = `SELECT * FROM transactions WHERE from_wallet_id = $1 AND EXISTS (SELECT 1 FROM wallets WHERE user_id = $2) ORDER BY transaction_at DESC LIMIT $3 OFFSET $4`;
        const values = [walletId, userId, limit, offset];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }
    
    public async getTransactionsByWalletIdAndUserIdAndPaginationCount(walletId: number, userId: number): Promise<number> {
        const query = `SELECT COUNT(*) FROM transactions WHERE from_wallet_id = $1 AND EXISTS (SELECT 1 FROM wallets WHERE user_id = $2)`;
        const values = [walletId, userId];
        const result = await this.executeQuery(query, values);

        return parseInt(result.rows[0].count);
    }

    public async getTransactionsByUserId(userId: number): Promise<Array<Transaction>> {
        const query = `SELECT * FROM transactions WHERE from_wallet_id IN (SELECT id FROM wallets WHERE user_id = $1) ORDER BY transaction_at DESC`;
        const values = [userId];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async getTransactionsByUserIdAndPagination(userId: number, limit: number, offset: number): Promise<Array<Transaction>> {
        const query = `SELECT * FROM transactions WHERE from_wallet_id IN (SELECT id FROM wallets WHERE user_id = $1) ORDER BY transaction_at DESC LIMIT $2 OFFSET $3`;
        const values = [userId, limit, offset];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async getTransactionsByUserIdAndPaginationCount(userId: number): Promise<number> {
        const query = `SELECT COUNT(*) FROM transactions WHERE from_wallet_id IN (SELECT id FROM wallets WHERE user_id = $1)`;
        const values = [userId];
        const result = await this.executeQuery(query, values);
        return parseInt(result.rows[0].count);
    }

    public async createTransaction(transaction: Transaction): Promise<Transaction> {
        const query = `
            INSERT INTO transactions (tx_id, currency_id, currency_name, from_address, to_address, from_wallet_id, from_public_key, amount, transaction_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW())
            RETURNING *`;
        const values = [
            transaction.tx_id,
            transaction.currency_id,
            transaction.currency_name,
            transaction.from_address,
            transaction.to_address,
            transaction.from_wallet_id,
            transaction.from_public_key,
            transaction.amount,
        ];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async updateTransaction(transaction: Transaction): Promise<Transaction> {
        const query = `
            UPDATE transactions 
            SET currency_id = $1, currency_name = $2, from_address = $3, to_address = $4, from_wallet_id = $5, from_public_key = $6, amount = $7, transaction_at = $8
            WHERE tx_id = $9
            RETURNING *`;
        const values = [
            transaction.currency_id,
            transaction.currency_name,
            transaction.from_address,
            transaction.to_address,
            transaction.from_wallet_id,
            transaction.from_public_key,
            transaction.amount,
            transaction.transaction_at,
            transaction.tx_id
        ];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async deleteTransactionByTxId(tx_id: string): Promise<boolean> {
        const query = `DELETE FROM transactions WHERE tx_id = $1`;
        const values = [tx_id];
        await this.executeQuery(query, values);
        return true;
    }

    public async deleteTransactionsByWalletId(walletId: number): Promise<boolean> {
        const query = `DELETE FROM transactions WHERE from_wallet_id = $1`;
        const values = [walletId];
        await this.executeQuery(query, values);
        return true;
    }
}
