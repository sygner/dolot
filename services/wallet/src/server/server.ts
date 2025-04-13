import * as grpc from '@grpc/grpc-js';
import { IWalletServiceServer, IWalletServiceService, WalletServiceService } from '../../proto/api/service_grpc_pb';
import { Database } from '../database/database'
import { Repository } from '../repository/repository'
import { CoinRepository } from '../repository/coin.repository'
import { WalletService } from '../services/wallet';
import { WalletHandler } from '../handlers/wallet';
import { WalletRepository } from '../repository/wallet.repository';
// import { UserId, Address, Wallet } from '../../proto/api/service_pb';
import { Wallet } from '../models/wallet';
import { CoinService } from '../services/coin';
import { CoinHandler } from '../handlers/coin';
import { TransactionService } from '../services/transaction';
import { TransactionRepository } from '../repository/transaction.repository';
import { TransactionHandler } from '../handlers/transaction';
import * as fs from "fs"
import { parse } from 'yaml';
import { CustomError } from '../types/error';
const GreeterServiceS: IWalletServiceService | IWalletServiceServer | any = WalletServiceService;

async function main() {

  const file = fs.readFileSync('./config.yml', 'utf8');
  const config = parse(file); // or yaml.load(file) if using js-yaml

  const dbConfig = config.database;
  const serverConfig = config.server;

  const client = await new Database(
    dbConfig.host,
    dbConfig.port,
    dbConfig.user,
    dbConfig.password,
    dbConfig.database
  ).connect();


  const repository = new Repository(client);
  const coinRepository = new CoinRepository(client);
  const walletRepository = new WalletRepository(client);
  const transactionRepository = new TransactionRepository(client);

  const result = await coinRepository.getAllCoins();
  console.log(`\x1b[32m[SERVER]\x1b[0m Starting Server...`);

  const server = new grpc.Server();

  const walletService = new WalletService(walletRepository,coinRepository);
  const coinService = new CoinService(coinRepository);
  const transactionService = new TransactionService(transactionRepository, walletRepository)

  const walletHandler = new WalletHandler(walletService);
  const coinHandler = new CoinHandler(coinService);
  const transactionHandler = new TransactionHandler(transactionService);
  
  try {
    const mainWallet = await walletService.createWallet({
      user_id:0,
      coin_id:1
  })
 const done = await walletRepository.makeTheMainAccountIndexToZero()
 if (done) {
    mainWallet.id = 0
    console.log(`\x1b[32m[SERVER]\x1b[0m Main Account Index Set To Zero`,mainWallet);
 }else{
  return
 } 
  } catch (error) {
    console.log(`\x1b[31m[SERVER]\x1b[0m Error Setting Main Account Index To Zero `,error);
    if (error instanceof CustomError) {
      if (error.message != "Wallet already exists") {
        return
      }
    }
  }
 server.addService(GreeterServiceS, {
    "getAllCoins": coinHandler.getAllCoins,
    "getCoinById": coinHandler.GetCoinById,
    "getCoinBySymbol": coinHandler.GetCoinBySymbol,
    "createWallet": walletHandler.createWallet,
    "getWalletsByUserId": walletHandler.getWalletsByUserId,
    "getWalletByAddress": walletHandler.getWalletByAddress,
    "updateBalance": walletHandler.updateBalance,
    "getBalanceByCoinIdAndUserId": walletHandler.getBalanceByCoinIdAndUserId,
    "deleteWalletsByUserId": walletHandler.deleteWalletsByUserId,
    "deleteWalletByWalletId": walletHandler.deleteWalletByWalletId,
    "checkTransactionExists": transactionHandler.checkTransactionExists,
    "getTransactionByTxId": transactionHandler.getTransactionByTxId,
    "getTransactionsByWalletId": transactionHandler.getTransactionsByWalletId,
    "addTransaction": transactionHandler.AddTransaction,
    "getTransactionsByUserId": transactionHandler.getTransactionsByUserId,
    "getPreTransactionDetail":transactionHandler.getPreTransactionDetail,
    "getTransactionsByWalletIdAndUserIdAndPagination": transactionHandler.getTransactionsByWalletIdAndUserIdAndPagination,
    "getTransactionsByUserIdAndPagination": transactionHandler.getTransactionsByUserIdAndPagination,
    "getWalletBySid":walletHandler.getWalletBySid,
    "getWalletByWalletId":walletHandler.getWalletByWalletId,
    "getWalletsByUserIdsAndCoinId":walletHandler.getWalletsByUserIdsAndCoinId,
  });

  const port = `${serverConfig.host}:${serverConfig.port}`;
  server.bindAsync(port, grpc.ServerCredentials.createInsecure(), (err, bindPort) => {
    if (err) {
      return console.error('Error starting server:', err);
    }
    console.log(`\x1b[32m[SERVER]\x1b[0m Look to go ${port}...`);
    // server.start();
  });
}

main();
