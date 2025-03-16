import { Repository } from './repository';
import { Coin } from "../models/coin";

export class CoinRepository extends Repository {
    public async getAllCoins(): Promise<Array<Coin>> {
        const query = `SELECT * FROM coins`;
        const result = await this.executeQuery(query);
        return result.rows;
    }

    public async getCoinById(id: number): Promise<Coin> {
        const query = `SELECT * FROM coins WHERE id = $1`;
        const values = [id];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }

    public async getCoinBySymbol(symbol: string): Promise<Coin> {
        const query = `SELECT * FROM coins WHERE currency_symbol = $1`;
        const values = [symbol];
        const result = await this.executeQuery(query, values);
        return result.rows[0];
    }
}
