import { LCDClient, MsgSend, MnemonicKey, Coins } from '@terra-money/feather.js';
import { CustomError } from '../types/error';
import { WalletRepository } from '../repository/wallet.repository';
import { Pagination } from '@terra-money/feather.js/dist/client/lcd/APIRequester';

export const MAX_BALANCE_RETRIES = 5;
// const terra = new LCDClient({
//     URL: 'https://terra-testnet-api.polkachu.com/',
//     chainID: 'pisco-1',
// });

export const chainID = 'localterra';  // Ensure the correct chain ID is used


export const terra = new LCDClient({
  localterra: {
    lcd: 'http://74.208.195.154:2560/',
    chainID: 'localterra',
    isClassic: true,
    gasAdjustment: 1.75,
    gasPrices: { uluna: 0.015 },
    prefix: 'terra', // bech32 prefix, used by the LCD to understand which is the right chain to query
  },
});

export async function GetLunaBalanceByCoinIdAndUserId(walletRepository: WalletRepository, coinId: number, userId: number): Promise<number> {
  // Fetch wallet by coinId and userId from the repository
  const response = await walletRepository.getWalletByCoinIdAndUserId(coinId, userId);
  // console.log("GetLunaBalanceByCoinIdAndUserId ", response)

  // Check if wallet exists
  if (!response) {
    throw new CustomError('Wallet not found', 404);
  }

  let balance = 0;
  let retries = 0;
  let balanceResult: any;

  // Retry logic for fetching the balance
  while (retries < MAX_BALANCE_RETRIES) {
    try {
      // Fetch the balance with a timeout wrapper
      balanceResult = await withTimeout(terra.bank.balance(response.address!), 5000); // 5-second timeout
      balance = parseBalance(balanceResult); // Parse the balance if needed
      break; // Exit loop if balance fetch is successful
    } catch (balanceError) {
      retries++;
      console.error(`Attempt ${retries} failed:`, balanceError);

      if (retries >= MAX_BALANCE_RETRIES) {
        throw new CustomError('Failed to fetch balance after retries', 500);
      }

      // Wait before retrying
      await delay(1500); // Retry after 1.5 seconds
    }
  }

  // Assuming balanceResult[0] returns a Coins instance
  const coins = balanceResult[0]; // Retrieve the Coins instance from the response
  let ulunaBalance;
  // Access the _coins property of the Coins instance
  // Ensure _coins exist and contains uluna
  if (coins && coins._coins && coins._coins['uluna']) {
    ulunaBalance = coins._coins['uluna'];
  } else {
    console.warn("uluna balance not found, setting to zero.", balanceResult);
    ulunaBalance = { amount: "0" };
  }

  const latestBalance = parseFloat(ulunaBalance.amount);
  if (response.id || response.id == 0 ) {
    try {
      console.log("Update",latestBalance)
      const resultw = await walletRepository.updateBalance(response.id, latestBalance);
      // console.log(resultw)
    } catch (error){
      console.log(error)
    }
    return latestBalance;
  } else {
    console.log("Wallet ID is missing coin id and user id ", coinId, userId,response)
    throw new CustomError('Wallet ID is missing', 400);
  }
  // Access the _coins property of the Coins instance
  // let ulunaBalance = coins ? coins._coins['uluna'] : 0; // Get the uluna balance from the coins
  // if (!ulunaBalance)
  //   ulunaBalance = 0
  // console.log("DDD ",ulunaBalance)
  // // If ulunaBalance exists, parse and update the wallet's balance
  // if (ulunaBalance) {
  //   const latestBalance = parseFloat(ulunaBalance.amount); // Convert string to number

  //   // Ensure the wallet ID exists and update the balance in the repository
  //   if (response.id) {
  //     await walletRepository.updateBalance(response.id, latestBalance);
  //   } else {
  //     throw new CustomError('Wallet ID is missing', 400); // Handle missing wallet ID
  //   }

  //   return latestBalance; // Return the updated balance
  // } else {
  //   // Handle case where uluna balance doesn't exist
  //   console.log('uluna balance does not exist');
  //   throw new CustomError('uluna balance not found in the coins object', 400);
  // }
}


export async function GetLunaBalanceByAddress(walletRepository: WalletRepository, address: string): Promise<number> {
  // Fetch wallet by coinId and userId from the repository
  const response = await walletRepository.getWalletByAddress(address);
  // console.log("GetLunaBalanceByAddress ", response)


  // Check if wallet exists
  if (!response) {
    throw new CustomError('Wallet not found', 404);
  }

  let balance = 0;
  let retries = 0;
  let balanceResult: any;

  // Retry logic for fetching the balance
  while (retries < MAX_BALANCE_RETRIES) {
    try {
      // Fetch the balance with a timeout wrapper
      balanceResult = await withTimeout(terra.bank.balance(response.address!), 5000); // 5-second timeout
      balance = parseBalance(balanceResult); // Parse the balance if needed
      break; // Exit loop if balance fetch is successful
    } catch (balanceError) {
      retries++;
      console.error(`Attempt ${retries} failed:`, balanceError);

      if (retries >= MAX_BALANCE_RETRIES) {
        throw new CustomError('Failed to fetch balance after retries', 500);
      }

      // Wait before retrying
      await delay(1500); // Retry after 1.5 seconds
    }
  }

  // Log the balance result (for debugging)

  // Assuming balanceResult[0] returns a Coins instance
  const coins = balanceResult[0]; // Retrieve the Coins instance from the response
  let ulunaBalance;
  // Access the _coins property of the Coins instance
  // Ensure _coins exist and contains uluna
  if (coins && coins._coins && coins._coins['uluna']) {
    ulunaBalance = coins._coins['uluna'];
  } else {
    console.warn("uluna balance not found, setting to zero.");
    ulunaBalance = { amount: "0" };
  }

  const latestBalance = parseFloat(ulunaBalance.amount);
  if (response.id || response.id == 0 ) {
    await walletRepository.updateBalance(response.id, latestBalance);
    return latestBalance;
  } else {
    console.log("Wallet ID is missing address", address,response)
    throw new CustomError('Wallet ID is missing', 400);
  }
}

function delay(ms: number): Promise<void> {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

function withTimeout<T>(promise: Promise<T>, timeout: number): Promise<T> {
  return new Promise((resolve, reject) => {
    const timer = setTimeout(() => {
      reject(new Error('Operation timed out'));
    }, timeout);

    promise
      .then((result) => {
        clearTimeout(timer);
        resolve(result);
      })
      .catch((err) => {
        clearTimeout(timer);
        reject(err);
      });
  });
}

function parseBalance(balanceResult: [Coins, Pagination]): number {
  return balanceResult[1].total
}

