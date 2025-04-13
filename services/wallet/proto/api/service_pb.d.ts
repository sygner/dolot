// package: wallet
// file: service.proto

/* tslint:disable */
/* eslint-disable */

import * as jspb from "google-protobuf";

export class Empty extends jspb.Message { 

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Empty.AsObject;
    static toObject(includeInstance: boolean, msg: Empty): Empty.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Empty, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Empty;
    static deserializeBinaryFromReader(message: Empty, reader: jspb.BinaryReader): Empty;
}

export namespace Empty {
    export type AsObject = {
    }
}

export class Pagination extends jspb.Message { 
    getOffset(): number;
    setOffset(value: number): Pagination;
    getLimit(): number;
    setLimit(value: number): Pagination;
    getTotal(): boolean;
    setTotal(value: boolean): Pagination;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Pagination.AsObject;
    static toObject(includeInstance: boolean, msg: Pagination): Pagination.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Pagination, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Pagination;
    static deserializeBinaryFromReader(message: Pagination, reader: jspb.BinaryReader): Pagination;
}

export namespace Pagination {
    export type AsObject = {
        offset: number,
        limit: number,
        total: boolean,
    }
}

export class getTransactionsByWalletIdAndUserIdAndPaginationRequest extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): getTransactionsByWalletIdAndUserIdAndPaginationRequest;
    getWalletId(): number;
    setWalletId(value: number): getTransactionsByWalletIdAndUserIdAndPaginationRequest;

    hasPagination(): boolean;
    clearPagination(): void;
    getPagination(): Pagination | undefined;
    setPagination(value?: Pagination): getTransactionsByWalletIdAndUserIdAndPaginationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): getTransactionsByWalletIdAndUserIdAndPaginationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: getTransactionsByWalletIdAndUserIdAndPaginationRequest): getTransactionsByWalletIdAndUserIdAndPaginationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: getTransactionsByWalletIdAndUserIdAndPaginationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): getTransactionsByWalletIdAndUserIdAndPaginationRequest;
    static deserializeBinaryFromReader(message: getTransactionsByWalletIdAndUserIdAndPaginationRequest, reader: jspb.BinaryReader): getTransactionsByWalletIdAndUserIdAndPaginationRequest;
}

export namespace getTransactionsByWalletIdAndUserIdAndPaginationRequest {
    export type AsObject = {
        userId: number,
        walletId: number,
        pagination?: Pagination.AsObject,
    }
}

export class getTransactionsByUserIdAndPaginationRequest extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): getTransactionsByUserIdAndPaginationRequest;

    hasPagination(): boolean;
    clearPagination(): void;
    getPagination(): Pagination | undefined;
    setPagination(value?: Pagination): getTransactionsByUserIdAndPaginationRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): getTransactionsByUserIdAndPaginationRequest.AsObject;
    static toObject(includeInstance: boolean, msg: getTransactionsByUserIdAndPaginationRequest): getTransactionsByUserIdAndPaginationRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: getTransactionsByUserIdAndPaginationRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): getTransactionsByUserIdAndPaginationRequest;
    static deserializeBinaryFromReader(message: getTransactionsByUserIdAndPaginationRequest, reader: jspb.BinaryReader): getTransactionsByUserIdAndPaginationRequest;
}

export namespace getTransactionsByUserIdAndPaginationRequest {
    export type AsObject = {
        userId: number,
        pagination?: Pagination.AsObject,
    }
}

export class BooleanResult extends jspb.Message { 
    getResult(): boolean;
    setResult(value: boolean): BooleanResult;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): BooleanResult.AsObject;
    static toObject(includeInstance: boolean, msg: BooleanResult): BooleanResult.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: BooleanResult, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): BooleanResult;
    static deserializeBinaryFromReader(message: BooleanResult, reader: jspb.BinaryReader): BooleanResult;
}

export namespace BooleanResult {
    export type AsObject = {
        result: boolean,
    }
}

export class GetWalletsByUserIdsAndCoinIdRequest extends jspb.Message { 
    clearUserIdsList(): void;
    getUserIdsList(): Array<number>;
    setUserIdsList(value: Array<number>): GetWalletsByUserIdsAndCoinIdRequest;
    addUserIds(value: number, index?: number): number;
    getCoinId(): number;
    setCoinId(value: number): GetWalletsByUserIdsAndCoinIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetWalletsByUserIdsAndCoinIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetWalletsByUserIdsAndCoinIdRequest): GetWalletsByUserIdsAndCoinIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetWalletsByUserIdsAndCoinIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetWalletsByUserIdsAndCoinIdRequest;
    static deserializeBinaryFromReader(message: GetWalletsByUserIdsAndCoinIdRequest, reader: jspb.BinaryReader): GetWalletsByUserIdsAndCoinIdRequest;
}

export namespace GetWalletsByUserIdsAndCoinIdRequest {
    export type AsObject = {
        userIdsList: Array<number>,
        coinId: number,
    }
}

export class GetTransactionsByWalletIdRequest extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): GetTransactionsByWalletIdRequest;
    getWalletId(): number;
    setWalletId(value: number): GetTransactionsByWalletIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetTransactionsByWalletIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetTransactionsByWalletIdRequest): GetTransactionsByWalletIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetTransactionsByWalletIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetTransactionsByWalletIdRequest;
    static deserializeBinaryFromReader(message: GetTransactionsByWalletIdRequest, reader: jspb.BinaryReader): GetTransactionsByWalletIdRequest;
}

export namespace GetTransactionsByWalletIdRequest {
    export type AsObject = {
        userId: number,
        walletId: number,
    }
}

export class Coin extends jspb.Message { 
    getCoinId(): number;
    setCoinId(value: number): Coin;
    getCurrencyName(): string;
    setCurrencyName(value: string): Coin;
    getCurrencySymbol(): string;
    setCurrencySymbol(value: string): Coin;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Coin.AsObject;
    static toObject(includeInstance: boolean, msg: Coin): Coin.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Coin, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Coin;
    static deserializeBinaryFromReader(message: Coin, reader: jspb.BinaryReader): Coin;
}

export namespace Coin {
    export type AsObject = {
        coinId: number,
        currencyName: string,
        currencySymbol: string,
    }
}

export class Coins extends jspb.Message { 
    clearCoinsList(): void;
    getCoinsList(): Array<Coin>;
    setCoinsList(value: Array<Coin>): Coins;
    addCoins(value?: Coin, index?: number): Coin;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Coins.AsObject;
    static toObject(includeInstance: boolean, msg: Coins): Coins.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Coins, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Coins;
    static deserializeBinaryFromReader(message: Coins, reader: jspb.BinaryReader): Coins;
}

export namespace Coins {
    export type AsObject = {
        coinsList: Array<Coin.AsObject>,
    }
}

export class Sid extends jspb.Message { 
    getSid(): string;
    setSid(value: string): Sid;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Sid.AsObject;
    static toObject(includeInstance: boolean, msg: Sid): Sid.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Sid, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Sid;
    static deserializeBinaryFromReader(message: Sid, reader: jspb.BinaryReader): Sid;
}

export namespace Sid {
    export type AsObject = {
        sid: string,
    }
}

export class CreateWalletRequest extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): CreateWalletRequest;
    getCoinId(): number;
    setCoinId(value: number): CreateWalletRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CreateWalletRequest.AsObject;
    static toObject(includeInstance: boolean, msg: CreateWalletRequest): CreateWalletRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CreateWalletRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CreateWalletRequest;
    static deserializeBinaryFromReader(message: CreateWalletRequest, reader: jspb.BinaryReader): CreateWalletRequest;
}

export namespace CreateWalletRequest {
    export type AsObject = {
        userId: number,
        coinId: number,
    }
}

export class AddTransactionRequest extends jspb.Message { 
    getAmount(): number;
    setAmount(value: number): AddTransactionRequest;
    getFromWalletId(): number;
    setFromWalletId(value: number): AddTransactionRequest;
    getFromWalletUserId(): number;
    setFromWalletUserId(value: number): AddTransactionRequest;
    getToWalletAddress(): string;
    setToWalletAddress(value: string): AddTransactionRequest;
    getCoinId(): number;
    setCoinId(value: number): AddTransactionRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): AddTransactionRequest.AsObject;
    static toObject(includeInstance: boolean, msg: AddTransactionRequest): AddTransactionRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: AddTransactionRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): AddTransactionRequest;
    static deserializeBinaryFromReader(message: AddTransactionRequest, reader: jspb.BinaryReader): AddTransactionRequest;
}

export namespace AddTransactionRequest {
    export type AsObject = {
        amount: number,
        fromWalletId: number,
        fromWalletUserId: number,
        toWalletAddress: string,
        coinId: number,
    }
}

export class GetWalletByCoinIdAndUserIdRequest extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): GetWalletByCoinIdAndUserIdRequest;
    getCoinId(): number;
    setCoinId(value: number): GetWalletByCoinIdAndUserIdRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): GetWalletByCoinIdAndUserIdRequest.AsObject;
    static toObject(includeInstance: boolean, msg: GetWalletByCoinIdAndUserIdRequest): GetWalletByCoinIdAndUserIdRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: GetWalletByCoinIdAndUserIdRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): GetWalletByCoinIdAndUserIdRequest;
    static deserializeBinaryFromReader(message: GetWalletByCoinIdAndUserIdRequest, reader: jspb.BinaryReader): GetWalletByCoinIdAndUserIdRequest;
}

export namespace GetWalletByCoinIdAndUserIdRequest {
    export type AsObject = {
        userId: number,
        coinId: number,
    }
}

export class CoinId extends jspb.Message { 
    getCoinId(): number;
    setCoinId(value: number): CoinId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinId.AsObject;
    static toObject(includeInstance: boolean, msg: CoinId): CoinId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinId;
    static deserializeBinaryFromReader(message: CoinId, reader: jspb.BinaryReader): CoinId;
}

export namespace CoinId {
    export type AsObject = {
        coinId: number,
    }
}

export class CoinSymbol extends jspb.Message { 
    getSymbol(): string;
    setSymbol(value: string): CoinSymbol;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): CoinSymbol.AsObject;
    static toObject(includeInstance: boolean, msg: CoinSymbol): CoinSymbol.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: CoinSymbol, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): CoinSymbol;
    static deserializeBinaryFromReader(message: CoinSymbol, reader: jspb.BinaryReader): CoinSymbol;
}

export namespace CoinSymbol {
    export type AsObject = {
        symbol: string,
    }
}

export class UserId extends jspb.Message { 
    getUserId(): number;
    setUserId(value: number): UserId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UserId.AsObject;
    static toObject(includeInstance: boolean, msg: UserId): UserId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UserId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UserId;
    static deserializeBinaryFromReader(message: UserId, reader: jspb.BinaryReader): UserId;
}

export namespace UserId {
    export type AsObject = {
        userId: number,
    }
}

export class TransactionId extends jspb.Message { 
    getTxId(): string;
    setTxId(value: string): TransactionId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): TransactionId.AsObject;
    static toObject(includeInstance: boolean, msg: TransactionId): TransactionId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: TransactionId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): TransactionId;
    static deserializeBinaryFromReader(message: TransactionId, reader: jspb.BinaryReader): TransactionId;
}

export namespace TransactionId {
    export type AsObject = {
        txId: string,
    }
}

export class Address extends jspb.Message { 
    getAddress(): string;
    setAddress(value: string): Address;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Address.AsObject;
    static toObject(includeInstance: boolean, msg: Address): Address.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Address, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Address;
    static deserializeBinaryFromReader(message: Address, reader: jspb.BinaryReader): Address;
}

export namespace Address {
    export type AsObject = {
        address: string,
    }
}

export class WalletId extends jspb.Message { 
    getId(): number;
    setId(value: number): WalletId;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): WalletId.AsObject;
    static toObject(includeInstance: boolean, msg: WalletId): WalletId.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: WalletId, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): WalletId;
    static deserializeBinaryFromReader(message: WalletId, reader: jspb.BinaryReader): WalletId;
}

export namespace WalletId {
    export type AsObject = {
        id: number,
    }
}

export class Balance extends jspb.Message { 
    getBalance(): number;
    setBalance(value: number): Balance;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Balance.AsObject;
    static toObject(includeInstance: boolean, msg: Balance): Balance.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Balance, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Balance;
    static deserializeBinaryFromReader(message: Balance, reader: jspb.BinaryReader): Balance;
}

export namespace Balance {
    export type AsObject = {
        balance: number,
    }
}

export class UpdateBalanceRequest extends jspb.Message { 
    getWalletId(): string;
    setWalletId(value: string): UpdateBalanceRequest;
    getBalance(): number;
    setBalance(value: number): UpdateBalanceRequest;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): UpdateBalanceRequest.AsObject;
    static toObject(includeInstance: boolean, msg: UpdateBalanceRequest): UpdateBalanceRequest.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: UpdateBalanceRequest, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): UpdateBalanceRequest;
    static deserializeBinaryFromReader(message: UpdateBalanceRequest, reader: jspb.BinaryReader): UpdateBalanceRequest;
}

export namespace UpdateBalanceRequest {
    export type AsObject = {
        walletId: string,
        balance: number,
    }
}

export class Wallet extends jspb.Message { 
    getId(): number;
    setId(value: number): Wallet;
    getSid(): string;
    setSid(value: string): Wallet;
    getUserId(): number;
    setUserId(value: number): Wallet;
    getCoinId(): number;
    setCoinId(value: number): Wallet;
    getBalance(): number;
    setBalance(value: number): Wallet;
    getAddress(): string;
    setAddress(value: string): Wallet;
    getPublicKey(): string;
    setPublicKey(value: string): Wallet;
    getPrivateKey(): string;
    setPrivateKey(value: string): Wallet;
    getMnemonic(): string;
    setMnemonic(value: string): Wallet;
    getCreatedAt(): string;
    setCreatedAt(value: string): Wallet;
    getUpdatedAt(): string;
    setUpdatedAt(value: string): Wallet;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Wallet.AsObject;
    static toObject(includeInstance: boolean, msg: Wallet): Wallet.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Wallet, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Wallet;
    static deserializeBinaryFromReader(message: Wallet, reader: jspb.BinaryReader): Wallet;
}

export namespace Wallet {
    export type AsObject = {
        id: number,
        sid: string,
        userId: number,
        coinId: number,
        balance: number,
        address: string,
        publicKey: string,
        privateKey: string,
        mnemonic: string,
        createdAt: string,
        updatedAt: string,
    }
}

export class Wallets extends jspb.Message { 
    clearWalletsList(): void;
    getWalletsList(): Array<Wallet>;
    setWalletsList(value: Array<Wallet>): Wallets;
    addWallets(value?: Wallet, index?: number): Wallet;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Wallets.AsObject;
    static toObject(includeInstance: boolean, msg: Wallets): Wallets.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Wallets, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Wallets;
    static deserializeBinaryFromReader(message: Wallets, reader: jspb.BinaryReader): Wallets;
}

export namespace Wallets {
    export type AsObject = {
        walletsList: Array<Wallet.AsObject>,
    }
}

export class Transaction extends jspb.Message { 
    getTxId(): string;
    setTxId(value: string): Transaction;
    getCurrencyId(): number;
    setCurrencyId(value: number): Transaction;
    getCurrencyName(): string;
    setCurrencyName(value: string): Transaction;
    getFromAddress(): string;
    setFromAddress(value: string): Transaction;
    getToAddress(): string;
    setToAddress(value: string): Transaction;
    getFromWalletId(): number;
    setFromWalletId(value: number): Transaction;
    getFromPublicKey(): string;
    setFromPublicKey(value: string): Transaction;
    getAmount(): number;
    setAmount(value: number): Transaction;
    getTransactionAt(): string;
    setTransactionAt(value: string): Transaction;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Transaction.AsObject;
    static toObject(includeInstance: boolean, msg: Transaction): Transaction.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Transaction, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Transaction;
    static deserializeBinaryFromReader(message: Transaction, reader: jspb.BinaryReader): Transaction;
}

export namespace Transaction {
    export type AsObject = {
        txId: string,
        currencyId: number,
        currencyName: string,
        fromAddress: string,
        toAddress: string,
        fromWalletId: number,
        fromPublicKey: string,
        amount: number,
        transactionAt: string,
    }
}

export class Transactions extends jspb.Message { 
    clearTransactionsList(): void;
    getTransactionsList(): Array<Transaction>;
    setTransactionsList(value: Array<Transaction>): Transactions;
    addTransactions(value?: Transaction, index?: number): Transaction;

    hasTotal(): boolean;
    clearTotal(): void;
    getTotal(): number | undefined;
    setTotal(value: number): Transactions;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): Transactions.AsObject;
    static toObject(includeInstance: boolean, msg: Transactions): Transactions.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: Transactions, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): Transactions;
    static deserializeBinaryFromReader(message: Transactions, reader: jspb.BinaryReader): Transactions;
}

export namespace Transactions {
    export type AsObject = {
        transactionsList: Array<Transaction.AsObject>,
        total?: number,
    }
}

export class PreTransactionDetail extends jspb.Message { 
    getAmount(): number;
    setAmount(value: number): PreTransactionDetail;
    getGasLimit(): number;
    setGasLimit(value: number): PreTransactionDetail;

    serializeBinary(): Uint8Array;
    toObject(includeInstance?: boolean): PreTransactionDetail.AsObject;
    static toObject(includeInstance: boolean, msg: PreTransactionDetail): PreTransactionDetail.AsObject;
    static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
    static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
    static serializeBinaryToWriter(message: PreTransactionDetail, writer: jspb.BinaryWriter): void;
    static deserializeBinary(bytes: Uint8Array): PreTransactionDetail;
    static deserializeBinaryFromReader(message: PreTransactionDetail, reader: jspb.BinaryReader): PreTransactionDetail;
}

export namespace PreTransactionDetail {
    export type AsObject = {
        amount: number,
        gasLimit: number,
    }
}
