// package: wallet
// file: service.proto

/* tslint:disable */
/* eslint-disable */

import * as grpc from "@grpc/grpc-js";
import * as service_pb from "./service_pb";

interface IWalletServiceService extends grpc.ServiceDefinition<grpc.UntypedServiceImplementation> {
    getAllCoins: IWalletServiceService_IGetAllCoins;
    getCoinById: IWalletServiceService_IGetCoinById;
    getCoinBySymbol: IWalletServiceService_IGetCoinBySymbol;
    createWallet: IWalletServiceService_ICreateWallet;
    getBalanceByCoinIdAndUserId: IWalletServiceService_IGetBalanceByCoinIdAndUserId;
    getWalletsByUserId: IWalletServiceService_IGetWalletsByUserId;
    getWalletByAddress: IWalletServiceService_IGetWalletByAddress;
    getWalletBySid: IWalletServiceService_IGetWalletBySid;
    getWalletByWalletId: IWalletServiceService_IGetWalletByWalletId;
    updateBalance: IWalletServiceService_IUpdateBalance;
    deleteWalletByWalletId: IWalletServiceService_IDeleteWalletByWalletId;
    deleteWalletsByUserId: IWalletServiceService_IDeleteWalletsByUserId;
    checkTransactionExists: IWalletServiceService_ICheckTransactionExists;
    getTransactionByTxId: IWalletServiceService_IGetTransactionByTxId;
    getTransactionsByWalletId: IWalletServiceService_IGetTransactionsByWalletId;
    addTransaction: IWalletServiceService_IAddTransaction;
    getTransactionsByUserId: IWalletServiceService_IGetTransactionsByUserId;
    getWalletsByUserIdsAndCoinId: IWalletServiceService_IGetWalletsByUserIdsAndCoinId;
    getPreTransactionDetail: IWalletServiceService_IGetPreTransactionDetail;
    getTransactionsByWalletIdAndUserIdAndPagination: IWalletServiceService_IgetTransactionsByWalletIdAndUserIdAndPagination;
    getTransactionsByUserIdAndPagination: IWalletServiceService_IgetTransactionsByUserIdAndPagination;
}

interface IWalletServiceService_IGetAllCoins extends grpc.MethodDefinition<service_pb.Empty, service_pb.Coins> {
    path: "/wallet.WalletService/GetAllCoins";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.Empty>;
    requestDeserialize: grpc.deserialize<service_pb.Empty>;
    responseSerialize: grpc.serialize<service_pb.Coins>;
    responseDeserialize: grpc.deserialize<service_pb.Coins>;
}
interface IWalletServiceService_IGetCoinById extends grpc.MethodDefinition<service_pb.CoinId, service_pb.Coin> {
    path: "/wallet.WalletService/GetCoinById";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.CoinId>;
    requestDeserialize: grpc.deserialize<service_pb.CoinId>;
    responseSerialize: grpc.serialize<service_pb.Coin>;
    responseDeserialize: grpc.deserialize<service_pb.Coin>;
}
interface IWalletServiceService_IGetCoinBySymbol extends grpc.MethodDefinition<service_pb.CoinSymbol, service_pb.Coin> {
    path: "/wallet.WalletService/GetCoinBySymbol";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.CoinSymbol>;
    requestDeserialize: grpc.deserialize<service_pb.CoinSymbol>;
    responseSerialize: grpc.serialize<service_pb.Coin>;
    responseDeserialize: grpc.deserialize<service_pb.Coin>;
}
interface IWalletServiceService_ICreateWallet extends grpc.MethodDefinition<service_pb.CreateWalletRequest, service_pb.Wallet> {
    path: "/wallet.WalletService/CreateWallet";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.CreateWalletRequest>;
    requestDeserialize: grpc.deserialize<service_pb.CreateWalletRequest>;
    responseSerialize: grpc.serialize<service_pb.Wallet>;
    responseDeserialize: grpc.deserialize<service_pb.Wallet>;
}
interface IWalletServiceService_IGetBalanceByCoinIdAndUserId extends grpc.MethodDefinition<service_pb.GetWalletByCoinIdAndUserIdRequest, service_pb.Balance> {
    path: "/wallet.WalletService/GetBalanceByCoinIdAndUserId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.GetWalletByCoinIdAndUserIdRequest>;
    requestDeserialize: grpc.deserialize<service_pb.GetWalletByCoinIdAndUserIdRequest>;
    responseSerialize: grpc.serialize<service_pb.Balance>;
    responseDeserialize: grpc.deserialize<service_pb.Balance>;
}
interface IWalletServiceService_IGetWalletsByUserId extends grpc.MethodDefinition<service_pb.UserId, service_pb.Wallets> {
    path: "/wallet.WalletService/GetWalletsByUserId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.UserId>;
    requestDeserialize: grpc.deserialize<service_pb.UserId>;
    responseSerialize: grpc.serialize<service_pb.Wallets>;
    responseDeserialize: grpc.deserialize<service_pb.Wallets>;
}
interface IWalletServiceService_IGetWalletByAddress extends grpc.MethodDefinition<service_pb.Address, service_pb.Wallet> {
    path: "/wallet.WalletService/GetWalletByAddress";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.Address>;
    requestDeserialize: grpc.deserialize<service_pb.Address>;
    responseSerialize: grpc.serialize<service_pb.Wallet>;
    responseDeserialize: grpc.deserialize<service_pb.Wallet>;
}
interface IWalletServiceService_IGetWalletBySid extends grpc.MethodDefinition<service_pb.Sid, service_pb.Wallet> {
    path: "/wallet.WalletService/GetWalletBySid";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.Sid>;
    requestDeserialize: grpc.deserialize<service_pb.Sid>;
    responseSerialize: grpc.serialize<service_pb.Wallet>;
    responseDeserialize: grpc.deserialize<service_pb.Wallet>;
}
interface IWalletServiceService_IGetWalletByWalletId extends grpc.MethodDefinition<service_pb.WalletId, service_pb.Wallet> {
    path: "/wallet.WalletService/GetWalletByWalletId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.WalletId>;
    requestDeserialize: grpc.deserialize<service_pb.WalletId>;
    responseSerialize: grpc.serialize<service_pb.Wallet>;
    responseDeserialize: grpc.deserialize<service_pb.Wallet>;
}
interface IWalletServiceService_IUpdateBalance extends grpc.MethodDefinition<service_pb.WalletId, service_pb.Wallet> {
    path: "/wallet.WalletService/UpdateBalance";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.WalletId>;
    requestDeserialize: grpc.deserialize<service_pb.WalletId>;
    responseSerialize: grpc.serialize<service_pb.Wallet>;
    responseDeserialize: grpc.deserialize<service_pb.Wallet>;
}
interface IWalletServiceService_IDeleteWalletByWalletId extends grpc.MethodDefinition<service_pb.WalletId, service_pb.Empty> {
    path: "/wallet.WalletService/DeleteWalletByWalletId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.WalletId>;
    requestDeserialize: grpc.deserialize<service_pb.WalletId>;
    responseSerialize: grpc.serialize<service_pb.Empty>;
    responseDeserialize: grpc.deserialize<service_pb.Empty>;
}
interface IWalletServiceService_IDeleteWalletsByUserId extends grpc.MethodDefinition<service_pb.UserId, service_pb.Empty> {
    path: "/wallet.WalletService/DeleteWalletsByUserId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.UserId>;
    requestDeserialize: grpc.deserialize<service_pb.UserId>;
    responseSerialize: grpc.serialize<service_pb.Empty>;
    responseDeserialize: grpc.deserialize<service_pb.Empty>;
}
interface IWalletServiceService_ICheckTransactionExists extends grpc.MethodDefinition<service_pb.TransactionId, service_pb.BooleanResult> {
    path: "/wallet.WalletService/CheckTransactionExists";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.TransactionId>;
    requestDeserialize: grpc.deserialize<service_pb.TransactionId>;
    responseSerialize: grpc.serialize<service_pb.BooleanResult>;
    responseDeserialize: grpc.deserialize<service_pb.BooleanResult>;
}
interface IWalletServiceService_IGetTransactionByTxId extends grpc.MethodDefinition<service_pb.TransactionId, service_pb.Transaction> {
    path: "/wallet.WalletService/GetTransactionByTxId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.TransactionId>;
    requestDeserialize: grpc.deserialize<service_pb.TransactionId>;
    responseSerialize: grpc.serialize<service_pb.Transaction>;
    responseDeserialize: grpc.deserialize<service_pb.Transaction>;
}
interface IWalletServiceService_IGetTransactionsByWalletId extends grpc.MethodDefinition<service_pb.GetTransactionsByWalletIdRequest, service_pb.Transactions> {
    path: "/wallet.WalletService/GetTransactionsByWalletId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.GetTransactionsByWalletIdRequest>;
    requestDeserialize: grpc.deserialize<service_pb.GetTransactionsByWalletIdRequest>;
    responseSerialize: grpc.serialize<service_pb.Transactions>;
    responseDeserialize: grpc.deserialize<service_pb.Transactions>;
}
interface IWalletServiceService_IAddTransaction extends grpc.MethodDefinition<service_pb.AddTransactionRequest, service_pb.Transaction> {
    path: "/wallet.WalletService/AddTransaction";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.AddTransactionRequest>;
    requestDeserialize: grpc.deserialize<service_pb.AddTransactionRequest>;
    responseSerialize: grpc.serialize<service_pb.Transaction>;
    responseDeserialize: grpc.deserialize<service_pb.Transaction>;
}
interface IWalletServiceService_IGetTransactionsByUserId extends grpc.MethodDefinition<service_pb.UserId, service_pb.Transactions> {
    path: "/wallet.WalletService/GetTransactionsByUserId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.UserId>;
    requestDeserialize: grpc.deserialize<service_pb.UserId>;
    responseSerialize: grpc.serialize<service_pb.Transactions>;
    responseDeserialize: grpc.deserialize<service_pb.Transactions>;
}
interface IWalletServiceService_IGetWalletsByUserIdsAndCoinId extends grpc.MethodDefinition<service_pb.GetWalletsByUserIdsAndCoinIdRequest, service_pb.Wallets> {
    path: "/wallet.WalletService/GetWalletsByUserIdsAndCoinId";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.GetWalletsByUserIdsAndCoinIdRequest>;
    requestDeserialize: grpc.deserialize<service_pb.GetWalletsByUserIdsAndCoinIdRequest>;
    responseSerialize: grpc.serialize<service_pb.Wallets>;
    responseDeserialize: grpc.deserialize<service_pb.Wallets>;
}
interface IWalletServiceService_IGetPreTransactionDetail extends grpc.MethodDefinition<service_pb.AddTransactionRequest, service_pb.PreTransactionDetail> {
    path: "/wallet.WalletService/GetPreTransactionDetail";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.AddTransactionRequest>;
    requestDeserialize: grpc.deserialize<service_pb.AddTransactionRequest>;
    responseSerialize: grpc.serialize<service_pb.PreTransactionDetail>;
    responseDeserialize: grpc.deserialize<service_pb.PreTransactionDetail>;
}
interface IWalletServiceService_IgetTransactionsByWalletIdAndUserIdAndPagination extends grpc.MethodDefinition<service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, service_pb.Transactions> {
    path: "/wallet.WalletService/getTransactionsByWalletIdAndUserIdAndPagination";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest>;
    requestDeserialize: grpc.deserialize<service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest>;
    responseSerialize: grpc.serialize<service_pb.Transactions>;
    responseDeserialize: grpc.deserialize<service_pb.Transactions>;
}
interface IWalletServiceService_IgetTransactionsByUserIdAndPagination extends grpc.MethodDefinition<service_pb.getTransactionsByUserIdAndPaginationRequest, service_pb.Transactions> {
    path: "/wallet.WalletService/getTransactionsByUserIdAndPagination";
    requestStream: false;
    responseStream: false;
    requestSerialize: grpc.serialize<service_pb.getTransactionsByUserIdAndPaginationRequest>;
    requestDeserialize: grpc.deserialize<service_pb.getTransactionsByUserIdAndPaginationRequest>;
    responseSerialize: grpc.serialize<service_pb.Transactions>;
    responseDeserialize: grpc.deserialize<service_pb.Transactions>;
}

export const WalletServiceService: IWalletServiceService;

export interface IWalletServiceServer {
    getAllCoins: grpc.handleUnaryCall<service_pb.Empty, service_pb.Coins>;
    getCoinById: grpc.handleUnaryCall<service_pb.CoinId, service_pb.Coin>;
    getCoinBySymbol: grpc.handleUnaryCall<service_pb.CoinSymbol, service_pb.Coin>;
    createWallet: grpc.handleUnaryCall<service_pb.CreateWalletRequest, service_pb.Wallet>;
    getBalanceByCoinIdAndUserId: grpc.handleUnaryCall<service_pb.GetWalletByCoinIdAndUserIdRequest, service_pb.Balance>;
    getWalletsByUserId: grpc.handleUnaryCall<service_pb.UserId, service_pb.Wallets>;
    getWalletByAddress: grpc.handleUnaryCall<service_pb.Address, service_pb.Wallet>;
    getWalletBySid: grpc.handleUnaryCall<service_pb.Sid, service_pb.Wallet>;
    getWalletByWalletId: grpc.handleUnaryCall<service_pb.WalletId, service_pb.Wallet>;
    updateBalance: grpc.handleUnaryCall<service_pb.WalletId, service_pb.Wallet>;
    deleteWalletByWalletId: grpc.handleUnaryCall<service_pb.WalletId, service_pb.Empty>;
    deleteWalletsByUserId: grpc.handleUnaryCall<service_pb.UserId, service_pb.Empty>;
    checkTransactionExists: grpc.handleUnaryCall<service_pb.TransactionId, service_pb.BooleanResult>;
    getTransactionByTxId: grpc.handleUnaryCall<service_pb.TransactionId, service_pb.Transaction>;
    getTransactionsByWalletId: grpc.handleUnaryCall<service_pb.GetTransactionsByWalletIdRequest, service_pb.Transactions>;
    addTransaction: grpc.handleUnaryCall<service_pb.AddTransactionRequest, service_pb.Transaction>;
    getTransactionsByUserId: grpc.handleUnaryCall<service_pb.UserId, service_pb.Transactions>;
    getWalletsByUserIdsAndCoinId: grpc.handleUnaryCall<service_pb.GetWalletsByUserIdsAndCoinIdRequest, service_pb.Wallets>;
    getPreTransactionDetail: grpc.handleUnaryCall<service_pb.AddTransactionRequest, service_pb.PreTransactionDetail>;
    getTransactionsByWalletIdAndUserIdAndPagination: grpc.handleUnaryCall<service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, service_pb.Transactions>;
    getTransactionsByUserIdAndPagination: grpc.handleUnaryCall<service_pb.getTransactionsByUserIdAndPaginationRequest, service_pb.Transactions>;
}

export interface IWalletServiceClient {
    getAllCoins(request: service_pb.Empty, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    getAllCoins(request: service_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    getAllCoins(request: service_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    getCoinById(request: service_pb.CoinId, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    getCoinById(request: service_pb.CoinId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    getCoinById(request: service_pb.CoinId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    getCoinBySymbol(request: service_pb.CoinSymbol, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    getCoinBySymbol(request: service_pb.CoinSymbol, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    getCoinBySymbol(request: service_pb.CoinSymbol, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    createWallet(request: service_pb.CreateWalletRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    createWallet(request: service_pb.CreateWalletRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    createWallet(request: service_pb.CreateWalletRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    getWalletsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getWalletByAddress(request: service_pb.Address, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletByAddress(request: service_pb.Address, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletByAddress(request: service_pb.Address, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletBySid(request: service_pb.Sid, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletBySid(request: service_pb.Sid, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletBySid(request: service_pb.Sid, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletByWalletId(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    getWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    updateBalance(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    updateBalance(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    updateBalance(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    deleteWalletByWalletId(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteWalletsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    deleteWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    checkTransactionExists(request: service_pb.TransactionId, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    checkTransactionExists(request: service_pb.TransactionId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    checkTransactionExists(request: service_pb.TransactionId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    getTransactionByTxId(request: service_pb.TransactionId, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    getTransactionByTxId(request: service_pb.TransactionId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    getTransactionByTxId(request: service_pb.TransactionId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    addTransaction(request: service_pb.AddTransactionRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    addTransaction(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    addTransaction(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    getTransactionsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    getPreTransactionDetail(request: service_pb.AddTransactionRequest, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    getPreTransactionDetail(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    getPreTransactionDetail(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
}

export class WalletServiceClient extends grpc.Client implements IWalletServiceClient {
    constructor(address: string, credentials: grpc.ChannelCredentials, options?: object);
    public getAllCoins(request: service_pb.Empty, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    public getAllCoins(request: service_pb.Empty, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    public getAllCoins(request: service_pb.Empty, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coins) => void): grpc.ClientUnaryCall;
    public getCoinById(request: service_pb.CoinId, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public getCoinById(request: service_pb.CoinId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public getCoinById(request: service_pb.CoinId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public getCoinBySymbol(request: service_pb.CoinSymbol, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public getCoinBySymbol(request: service_pb.CoinSymbol, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public getCoinBySymbol(request: service_pb.CoinSymbol, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Coin) => void): grpc.ClientUnaryCall;
    public createWallet(request: service_pb.CreateWalletRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public createWallet(request: service_pb.CreateWalletRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public createWallet(request: service_pb.CreateWalletRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    public getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    public getBalanceByCoinIdAndUserId(request: service_pb.GetWalletByCoinIdAndUserIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Balance) => void): grpc.ClientUnaryCall;
    public getWalletsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getWalletByAddress(request: service_pb.Address, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletByAddress(request: service_pb.Address, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletByAddress(request: service_pb.Address, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletBySid(request: service_pb.Sid, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletBySid(request: service_pb.Sid, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletBySid(request: service_pb.Sid, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletByWalletId(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public getWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public updateBalance(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public updateBalance(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public updateBalance(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallet) => void): grpc.ClientUnaryCall;
    public deleteWalletByWalletId(request: service_pb.WalletId, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteWalletByWalletId(request: service_pb.WalletId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteWalletsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public deleteWalletsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Empty) => void): grpc.ClientUnaryCall;
    public checkTransactionExists(request: service_pb.TransactionId, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    public checkTransactionExists(request: service_pb.TransactionId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    public checkTransactionExists(request: service_pb.TransactionId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.BooleanResult) => void): grpc.ClientUnaryCall;
    public getTransactionByTxId(request: service_pb.TransactionId, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getTransactionByTxId(request: service_pb.TransactionId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getTransactionByTxId(request: service_pb.TransactionId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletId(request: service_pb.GetTransactionsByWalletIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public addTransaction(request: service_pb.AddTransactionRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public addTransaction(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public addTransaction(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transaction) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserId(request: service_pb.UserId, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserId(request: service_pb.UserId, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getWalletsByUserIdsAndCoinId(request: service_pb.GetWalletsByUserIdsAndCoinIdRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Wallets) => void): grpc.ClientUnaryCall;
    public getPreTransactionDetail(request: service_pb.AddTransactionRequest, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    public getPreTransactionDetail(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    public getPreTransactionDetail(request: service_pb.AddTransactionRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.PreTransactionDetail) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByWalletIdAndUserIdAndPagination(request: service_pb.getTransactionsByWalletIdAndUserIdAndPaginationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, metadata: grpc.Metadata, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
    public getTransactionsByUserIdAndPagination(request: service_pb.getTransactionsByUserIdAndPaginationRequest, metadata: grpc.Metadata, options: Partial<grpc.CallOptions>, callback: (error: grpc.ServiceError | null, response: service_pb.Transactions) => void): grpc.ClientUnaryCall;
}
