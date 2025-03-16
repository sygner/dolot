import { Pool, PoolClient } from "pg";
import { exit } from "process";
import * as fs from "fs";

export class Database {
    private _address: string;
    private _port: number;
    private _username: string;
    private _password: string;
    private _database: string;
    protected _pool!: PoolClient;

    constructor(address: string, port: number = 5432, username: string, password: string, database: string) {
        this._address = address;
        this._port = port;
        this._username = username;
        this._password = password;
        this._database = database;
    }

    public get address(): string {
        return this._address;
    }

    public get port(): number {
        return this._port;
    }

    public get username(): string {
        return this._username;
    }

    public get password(): string {
        return this._password;
    }

    public get database(): string {
        return this._database;
    }

    public async connect(): Promise<PoolClient> {
        console.log(`\x1b[33m[DB]\x1b[0m Connecting ${this._address}:${this._port}`);
    
        const pool = new Pool({
            host: this._address,
            port: this._port,
            password: this._password,
            user: this._username,
            database: this._database,
            idleTimeoutMillis: 30000,
        });
    
        try {
            const client = await pool.connect();
            console.log(`\x1b[32m[DB]\x1b[0m Connected to database: ${this._address}:${this._port}`);
            this._pool = client; // Assign to instance variable if needed
            return client; // Always return the client
        } catch (err) {
            console.error(`\x1b[31m[DB]\x1b[0m Error connecting to database: ${err}`);
            throw new Error("Failed to connect to the database."); // Throw an error for the caller to handle
        }
    }
    
    public async run_migration(migration_path:string) {
        console.log(`\x1b[33m[DB]\x1b[0m Running migration`);

        try {
            const migrationSQL = fs.readFileSync(migration_path, "utf8");

            await this._pool.query(migrationSQL);

            console.log(`\x1b[32m[DB]\x1b[0m Migration completed successfully`);
        } catch (err) {
            console.error(`\x1b[31m[DB]\x1b[0m Error running migration: ${err}`);
        }
    }
}
