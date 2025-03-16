// GENERATED CODE -- DO NOT EDIT!

'use strict';
var grpc = require('@grpc/grpc-js');
var service_pb = require('./service_pb.js');

function serialize_wallet_AddTransactionRequest(arg) {
  if (!(arg instanceof service_pb.AddTransactionRequest)) {
    throw new Error('Expected argument of type wallet.AddTransactionRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_AddTransactionRequest(buffer_arg) {
  return service_pb.AddTransactionRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Address(arg) {
  if (!(arg instanceof service_pb.Address)) {
    throw new Error('Expected argument of type wallet.Address');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Address(buffer_arg) {
  return service_pb.Address.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Balance(arg) {
  if (!(arg instanceof service_pb.Balance)) {
    throw new Error('Expected argument of type wallet.Balance');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Balance(buffer_arg) {
  return service_pb.Balance.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_BooleanResult(arg) {
  if (!(arg instanceof service_pb.BooleanResult)) {
    throw new Error('Expected argument of type wallet.BooleanResult');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_BooleanResult(buffer_arg) {
  return service_pb.BooleanResult.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Coin(arg) {
  if (!(arg instanceof service_pb.Coin)) {
    throw new Error('Expected argument of type wallet.Coin');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Coin(buffer_arg) {
  return service_pb.Coin.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_CoinId(arg) {
  if (!(arg instanceof service_pb.CoinId)) {
    throw new Error('Expected argument of type wallet.CoinId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_CoinId(buffer_arg) {
  return service_pb.CoinId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_CoinSymbol(arg) {
  if (!(arg instanceof service_pb.CoinSymbol)) {
    throw new Error('Expected argument of type wallet.CoinSymbol');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_CoinSymbol(buffer_arg) {
  return service_pb.CoinSymbol.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Coins(arg) {
  if (!(arg instanceof service_pb.Coins)) {
    throw new Error('Expected argument of type wallet.Coins');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Coins(buffer_arg) {
  return service_pb.Coins.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_CreateWalletRequest(arg) {
  if (!(arg instanceof service_pb.CreateWalletRequest)) {
    throw new Error('Expected argument of type wallet.CreateWalletRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_CreateWalletRequest(buffer_arg) {
  return service_pb.CreateWalletRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Empty(arg) {
  if (!(arg instanceof service_pb.Empty)) {
    throw new Error('Expected argument of type wallet.Empty');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Empty(buffer_arg) {
  return service_pb.Empty.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_GetTransactionsByWalletIdRequest(arg) {
  if (!(arg instanceof service_pb.GetTransactionsByWalletIdRequest)) {
    throw new Error('Expected argument of type wallet.GetTransactionsByWalletIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_GetTransactionsByWalletIdRequest(buffer_arg) {
  return service_pb.GetTransactionsByWalletIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_GetWalletByCoinIdAndUserIdRequest(arg) {
  if (!(arg instanceof service_pb.GetWalletByCoinIdAndUserIdRequest)) {
    throw new Error('Expected argument of type wallet.GetWalletByCoinIdAndUserIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_GetWalletByCoinIdAndUserIdRequest(buffer_arg) {
  return service_pb.GetWalletByCoinIdAndUserIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_GetWalletsByUserIdsAndCoinIdRequest(arg) {
  if (!(arg instanceof service_pb.GetWalletsByUserIdsAndCoinIdRequest)) {
    throw new Error('Expected argument of type wallet.GetWalletsByUserIdsAndCoinIdRequest');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_GetWalletsByUserIdsAndCoinIdRequest(buffer_arg) {
  return service_pb.GetWalletsByUserIdsAndCoinIdRequest.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Sid(arg) {
  if (!(arg instanceof service_pb.Sid)) {
    throw new Error('Expected argument of type wallet.Sid');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Sid(buffer_arg) {
  return service_pb.Sid.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Transaction(arg) {
  if (!(arg instanceof service_pb.Transaction)) {
    throw new Error('Expected argument of type wallet.Transaction');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Transaction(buffer_arg) {
  return service_pb.Transaction.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_TransactionId(arg) {
  if (!(arg instanceof service_pb.TransactionId)) {
    throw new Error('Expected argument of type wallet.TransactionId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_TransactionId(buffer_arg) {
  return service_pb.TransactionId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Transactions(arg) {
  if (!(arg instanceof service_pb.Transactions)) {
    throw new Error('Expected argument of type wallet.Transactions');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Transactions(buffer_arg) {
  return service_pb.Transactions.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_UserId(arg) {
  if (!(arg instanceof service_pb.UserId)) {
    throw new Error('Expected argument of type wallet.UserId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_UserId(buffer_arg) {
  return service_pb.UserId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Wallet(arg) {
  if (!(arg instanceof service_pb.Wallet)) {
    throw new Error('Expected argument of type wallet.Wallet');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Wallet(buffer_arg) {
  return service_pb.Wallet.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_WalletId(arg) {
  if (!(arg instanceof service_pb.WalletId)) {
    throw new Error('Expected argument of type wallet.WalletId');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_WalletId(buffer_arg) {
  return service_pb.WalletId.deserializeBinary(new Uint8Array(buffer_arg));
}

function serialize_wallet_Wallets(arg) {
  if (!(arg instanceof service_pb.Wallets)) {
    throw new Error('Expected argument of type wallet.Wallets');
  }
  return Buffer.from(arg.serializeBinary());
}

function deserialize_wallet_Wallets(buffer_arg) {
  return service_pb.Wallets.deserializeBinary(new Uint8Array(buffer_arg));
}


var WalletServiceService = exports.WalletServiceService = {
  getAllCoins: {
    path: '/wallet.WalletService/GetAllCoins',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.Empty,
    responseType: service_pb.Coins,
    requestSerialize: serialize_wallet_Empty,
    requestDeserialize: deserialize_wallet_Empty,
    responseSerialize: serialize_wallet_Coins,
    responseDeserialize: deserialize_wallet_Coins,
  },
  getCoinById: {
    path: '/wallet.WalletService/GetCoinById',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.CoinId,
    responseType: service_pb.Coin,
    requestSerialize: serialize_wallet_CoinId,
    requestDeserialize: deserialize_wallet_CoinId,
    responseSerialize: serialize_wallet_Coin,
    responseDeserialize: deserialize_wallet_Coin,
  },
  getCoinBySymbol: {
    path: '/wallet.WalletService/GetCoinBySymbol',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.CoinSymbol,
    responseType: service_pb.Coin,
    requestSerialize: serialize_wallet_CoinSymbol,
    requestDeserialize: deserialize_wallet_CoinSymbol,
    responseSerialize: serialize_wallet_Coin,
    responseDeserialize: deserialize_wallet_Coin,
  },
  createWallet: {
    path: '/wallet.WalletService/CreateWallet',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.CreateWalletRequest,
    responseType: service_pb.Wallet,
    requestSerialize: serialize_wallet_CreateWalletRequest,
    requestDeserialize: deserialize_wallet_CreateWalletRequest,
    responseSerialize: serialize_wallet_Wallet,
    responseDeserialize: deserialize_wallet_Wallet,
  },
  getBalanceByCoinIdAndUserId: {
    path: '/wallet.WalletService/GetBalanceByCoinIdAndUserId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.GetWalletByCoinIdAndUserIdRequest,
    responseType: service_pb.Balance,
    requestSerialize: serialize_wallet_GetWalletByCoinIdAndUserIdRequest,
    requestDeserialize: deserialize_wallet_GetWalletByCoinIdAndUserIdRequest,
    responseSerialize: serialize_wallet_Balance,
    responseDeserialize: deserialize_wallet_Balance,
  },
  getWalletsByUserId: {
    path: '/wallet.WalletService/GetWalletsByUserId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.UserId,
    responseType: service_pb.Wallets,
    requestSerialize: serialize_wallet_UserId,
    requestDeserialize: deserialize_wallet_UserId,
    responseSerialize: serialize_wallet_Wallets,
    responseDeserialize: deserialize_wallet_Wallets,
  },
  getWalletByAddress: {
    path: '/wallet.WalletService/GetWalletByAddress',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.Address,
    responseType: service_pb.Wallet,
    requestSerialize: serialize_wallet_Address,
    requestDeserialize: deserialize_wallet_Address,
    responseSerialize: serialize_wallet_Wallet,
    responseDeserialize: deserialize_wallet_Wallet,
  },
  getWalletBySid: {
    path: '/wallet.WalletService/GetWalletBySid',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.Sid,
    responseType: service_pb.Wallet,
    requestSerialize: serialize_wallet_Sid,
    requestDeserialize: deserialize_wallet_Sid,
    responseSerialize: serialize_wallet_Wallet,
    responseDeserialize: deserialize_wallet_Wallet,
  },
  getWalletByWalletId: {
    path: '/wallet.WalletService/GetWalletByWalletId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.WalletId,
    responseType: service_pb.Wallet,
    requestSerialize: serialize_wallet_WalletId,
    requestDeserialize: deserialize_wallet_WalletId,
    responseSerialize: serialize_wallet_Wallet,
    responseDeserialize: deserialize_wallet_Wallet,
  },
  updateBalance: {
    path: '/wallet.WalletService/UpdateBalance',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.WalletId,
    responseType: service_pb.Wallet,
    requestSerialize: serialize_wallet_WalletId,
    requestDeserialize: deserialize_wallet_WalletId,
    responseSerialize: serialize_wallet_Wallet,
    responseDeserialize: deserialize_wallet_Wallet,
  },
  deleteWalletByWalletId: {
    path: '/wallet.WalletService/DeleteWalletByWalletId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.WalletId,
    responseType: service_pb.Empty,
    requestSerialize: serialize_wallet_WalletId,
    requestDeserialize: deserialize_wallet_WalletId,
    responseSerialize: serialize_wallet_Empty,
    responseDeserialize: deserialize_wallet_Empty,
  },
  deleteWalletsByUserId: {
    path: '/wallet.WalletService/DeleteWalletsByUserId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.UserId,
    responseType: service_pb.Empty,
    requestSerialize: serialize_wallet_UserId,
    requestDeserialize: deserialize_wallet_UserId,
    responseSerialize: serialize_wallet_Empty,
    responseDeserialize: deserialize_wallet_Empty,
  },
  checkTransactionExists: {
    path: '/wallet.WalletService/CheckTransactionExists',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.TransactionId,
    responseType: service_pb.BooleanResult,
    requestSerialize: serialize_wallet_TransactionId,
    requestDeserialize: deserialize_wallet_TransactionId,
    responseSerialize: serialize_wallet_BooleanResult,
    responseDeserialize: deserialize_wallet_BooleanResult,
  },
  getTransactionByTxId: {
    path: '/wallet.WalletService/GetTransactionByTxId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.TransactionId,
    responseType: service_pb.Transaction,
    requestSerialize: serialize_wallet_TransactionId,
    requestDeserialize: deserialize_wallet_TransactionId,
    responseSerialize: serialize_wallet_Transaction,
    responseDeserialize: deserialize_wallet_Transaction,
  },
  getTransactionsByWalletId: {
    path: '/wallet.WalletService/GetTransactionsByWalletId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.GetTransactionsByWalletIdRequest,
    responseType: service_pb.Transactions,
    requestSerialize: serialize_wallet_GetTransactionsByWalletIdRequest,
    requestDeserialize: deserialize_wallet_GetTransactionsByWalletIdRequest,
    responseSerialize: serialize_wallet_Transactions,
    responseDeserialize: deserialize_wallet_Transactions,
  },
  addTransaction: {
    path: '/wallet.WalletService/AddTransaction',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.AddTransactionRequest,
    responseType: service_pb.Transaction,
    requestSerialize: serialize_wallet_AddTransactionRequest,
    requestDeserialize: deserialize_wallet_AddTransactionRequest,
    responseSerialize: serialize_wallet_Transaction,
    responseDeserialize: deserialize_wallet_Transaction,
  },
  getTransactionsByUserId: {
    path: '/wallet.WalletService/GetTransactionsByUserId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.UserId,
    responseType: service_pb.Transactions,
    requestSerialize: serialize_wallet_UserId,
    requestDeserialize: deserialize_wallet_UserId,
    responseSerialize: serialize_wallet_Transactions,
    responseDeserialize: deserialize_wallet_Transactions,
  },
  getWalletsByUserIdsAndCoinId: {
    path: '/wallet.WalletService/GetWalletsByUserIdsAndCoinId',
    requestStream: false,
    responseStream: false,
    requestType: service_pb.GetWalletsByUserIdsAndCoinIdRequest,
    responseType: service_pb.Wallets,
    requestSerialize: serialize_wallet_GetWalletsByUserIdsAndCoinIdRequest,
    requestDeserialize: deserialize_wallet_GetWalletsByUserIdsAndCoinIdRequest,
    responseSerialize: serialize_wallet_Wallets,
    responseDeserialize: deserialize_wallet_Wallets,
  },
};

exports.WalletServiceClient = grpc.makeGenericClientConstructor(WalletServiceService);
