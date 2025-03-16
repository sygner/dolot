export interface Wallet{
    id?:number,
    sid?:string,
    user_id?:number,
    coin_id?:number,
    balance?:number,
    address?:string,
    public_key?:string,
    private_key?:string,
    mnemonic?:string,
    created_at?:Date,
    updated_at?:Date
}

