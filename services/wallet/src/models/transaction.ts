export interface Transaction {
    tx_id?: string;
    currency_id?: number;
    currency_name?: string;
    from_address?: string;
    to_address?: string;
    from_wallet_id?: number;
    from_public_key?: string;
    amount?: number;  // This can be used for double/float
    transaction_at?: Date;
}

export interface AddTransactionRequest {
    amount?:number;
    from_wallet_id?:number;
    from_wallet_user_id?:number;
    to_wallet_address?:string;
}