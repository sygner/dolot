import { PoolClient, QueryResult } from "pg";

export class Repository {
    protected _db: PoolClient;

    constructor(db: PoolClient) {
        this._db = db;
    }

    protected async executeQuery(query: string, values?: any[]): Promise<QueryResult<any>> {
        const rows = await this._db.query(query, values || []);
        return rows;
    }
}
