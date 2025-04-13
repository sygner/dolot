import { Repository } from "./repository";
import { Wallet } from "../models/wallet";

export class WalletRepository extends Repository {
    public async checkWalletExists(userId: number, coinId: number): Promise<boolean> {
        const query = `SELECT * FROM wallets WHERE user_id = $1 AND coin_id = $2`;
        const values = [userId, coinId];
        const result = await this.executeQuery(query, values);
        return result?.rowCount ? result.rowCount > 0 : false;
    }
    
    public async checkWalletExistsByUserIdAndSymbol(userId: number, symbol: string): Promise<boolean> {
        const query = `SELECT * FROM wallets w JOIN coins c ON w.coin_id = c.id WHERE w.user_id = $1 AND c.symbol = $2`;
        const values = [userId, symbol];
        const result = await this.executeQuery(query, values);
        return result?.rowCount ? result.rowCount > 0 : false;
    }
    
    public async getWalletsByUserId(userId: number): Promise<Array<Wallet>> {
        const query = `SELECT * FROM wallets WHERE user_id = $1`;
        const values = [userId];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async getWalletByAddress(address: string): Promise<Wallet> {
        const query = `SELECT * FROM wallets WHERE address = $1`;
        const values = [address];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async getWalletByPublicKey(publicKey: string): Promise<Wallet> {
        const query = `SELECT * FROM wallets WHERE public_key = $1`;
        const values = [publicKey];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }
    
    public async getWalletByWalletId(walletId: number): Promise<Wallet> {
        const query = `SELECT * FROM wallets WHERE id = $1`;
        const values = [walletId];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async getWalletBySid(sid: string): Promise<Wallet> {
        const query = `SELECT * FROM wallets WHERE sid = $1`;
        const values = [sid];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async createWallet(wallet: Wallet): Promise<Wallet> {
        const query = `
            INSERT INTO wallets (sid, user_id, coin_id, balance, address, public_key, private_key, mnemonic, created_at, updated_at)
            VALUES ($1, $2, $3, $4, $5, $6, $7, $8, NOW(), NOW())
            RETURNING *`;
        const values = [
            wallet.sid,
            wallet.user_id,
            wallet.coin_id,
            wallet.balance,
            wallet.address,
            wallet.public_key,
            wallet.private_key,
            wallet.mnemonic,
        ];
        console.log(query,values)
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async updateWallet(wallet: Wallet): Promise<Wallet> {
        const query = `
            UPDATE wallets SET balance = $1, updated_at = NOW()
            WHERE id = $2
            RETURNING *`;
        const values = [wallet.balance, wallet.id];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async updateBalance(walletId: number, balance: number): Promise<Wallet> {
        const query = `
            UPDATE wallets SET balance = $1, updated_at = NOW()
            WHERE id = $2
            RETURNING *`;
        const values = [balance, walletId];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async deleteWalletByWalletId(walletId: number): Promise<boolean> {
        const query = `DELETE FROM wallets WHERE id = $1`;
        const values = [walletId];
        await this.executeQuery(query, values);
        return true;
    }

    public async deleteWalletsByUserId(userId: number): Promise<boolean> {
        const query = `DELETE FROM wallets WHERE user_id = $1`;
        const values = [userId];
        await this.executeQuery(query, values);
        return true;
    }

    public async deleteWalletByAddress(address: string): Promise<boolean> {
        const query = `DELETE FROM wallets WHERE address = $1`;
        const values = [address];
        await this.executeQuery(query, values);
        return true;
    }

    public async getWalletByCoinIdAndUserId(coinId:number,userId:number):Promise<Wallet>{
        const query = `SELECT * FROM wallets WHERE user_id = $1 AND coin_id = $2`;
        const values = [userId,coinId]
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async getWalletsByUserIdsAndCoinId(userIds: Array<number>, coinId: number): Promise<Array<Wallet>> {
        const query = `SELECT * FROM wallets WHERE user_id = ANY($1) AND coin_id = $2`;
        const values = [userIds, coinId];
        const result = await this.executeQuery(query, values);
        return result.rows;
    }

    public async makeTheMainAccountIndexToZero():Promise<boolean>{
        const query = `UPDATE wallets SET id = '0' WHERE  user_id = 0`
        await this.executeQuery(query, [])
        return true
    }
    
}
