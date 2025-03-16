import { WalletRepository } from '../repository/wallet.repository';
import { Wallet } from '../models/wallet';
import { LCDClient, MsgSend, MnemonicKey, Coins } from '@terra-money/feather.js';
import { randomAlphanumeric, randomNumberBetween, randomAlphabet } from '../utils/rand';
import { CustomError } from '../types/error';
import { Pagination } from '@terra-money/feather.js/dist/client/lcd/APIRequester';
import { MAX_BALANCE_RETRIES, terra, GetLunaBalanceByCoinIdAndUserId } from "./contant"
import { CoinRepository } from '../repository/coin.repository';

export class WalletService {
    private walletRepository: WalletRepository;
    private coinRepository: CoinRepository;

    constructor(wp: WalletRepository,cp:CoinRepository) {
        this.walletRepository = wp;
        this.coinRepository = cp;
    }

    public async getWalletsByUserId(userId: number) {
        const coins = await this.coinRepository.getAllCoins();
        for (const coin of coins) {
            switch (coin.id) {
                case 1:
                    await this.getBalanceByCoinIdAndUserId(coin.id, userId);
                    break;
            }
        }
        
        const response = await this.walletRepository.getWalletsByUserId(userId);

        if (!response)
            throw new CustomError('Wallet not found', 404);

        return response;
    }

    public async getWalletsByPublicKey(publicKey: string) {
        const response = await this.walletRepository.getWalletByPublicKey(publicKey);

        if (!response)
            throw new CustomError('Wallet not found', 404);
        return response;
    }

    public async getWalletByAddress(address: string) {
        const response = await this.walletRepository.getWalletByAddress(address);

        if (!response)
            throw new CustomError('Wallet not found', 404);

        return response;
    }

    public async getWalletByWalletId(walletId: number) {
        const response = await this.walletRepository.getWalletByWalletId(walletId);

        if (!response)
            throw new CustomError('Wallet not found', 404);

        return response;
    }

    public async getWalletBySid(sid: string) {
        const response = await this.walletRepository.getWalletBySid(sid);

        if (!response)
            throw new CustomError('Wallet not found', 404);

        return response;
    }

    public async createWallet(wallet: Wallet): Promise<Wallet> {
        const exists = await this.walletRepository.checkWalletExists(wallet.user_id!, wallet.coin_id!);

        if (exists)
            throw new CustomError('Wallet already exists', 409);

        let resultWallet;
        switch (wallet.coin_id) {
            case 1:
                resultWallet = await CreateLunaWallet(wallet);
                break;
            default:
                throw new CustomError('Coin not supported', 400);
        }

        return await this.walletRepository.createWallet(resultWallet);
    }

    public async getBalanceByCoinIdAndUserId(coinId: number, userId: number): Promise<number> {
        const response = await this.walletRepository.getWalletByCoinIdAndUserId(coinId, userId);
        if (!response)
            throw new CustomError("Wallet not found", 404)
        switch (coinId) {
            case 1:
                return await GetLunaBalanceByCoinIdAndUserId(this.walletRepository, coinId, userId);
            default:
                throw new CustomError('Coin not supported', 400);
        }
    }

    public async updateWallet(wallet: Wallet) {
        return this.walletRepository.updateWallet(wallet);
    }

    public async updateBalance(walletId: number, balance: number) {
        return this.walletRepository.updateBalance(walletId, balance);
    }

    public async deleteWalletByWalletId(walletId: number) {
        return this.walletRepository.deleteWalletByWalletId(walletId);
    }

    public async deleteWalletsByUserId(userId: number) {
        return this.walletRepository.deleteWalletsByUserId(userId);
    }

    public async deleteWalletByAddress(address: string) {
        return this.walletRepository.deleteWalletByAddress(address);
    }

    public async getAndUpdateWalletBalance(walletId: number) {
        const response = await this.getWalletByWalletId(walletId)
        let balanceResult = await GetLunaBalanceByCoinIdAndUserId(this.walletRepository, response.coin_id!, response.user_id!)
        response.balance = balanceResult
        return response;
    }

    public async getWalletsByUserIdsAndCoinId(userIds: Array<number>, coinId: number) {
        return this.walletRepository.getWalletsByUserIdsAndCoinId(userIds, coinId);
    }
}

async function CreateLunaWallet(wallet: Wallet): Promise<Wallet> {
    const mk = new MnemonicKey();

    console.log("##1");
    const createdWallet = terra.wallet(mk);
    const address = createdWallet.key.accAddress("terra"); // Default prefix is 'terra'
    const publicKey = mk.publicKey!.pubkeyAddress("terra").toString(); // Correct way to access public key
    const privateKey = mk.privateKey.toString('hex');
    console.log("##2");

    console.log('Public Key:', publicKey.toString());
    console.log('Address:', address);
    console.log('Private Key:', privateKey);

    console.log("##3");

    let balance = 0;
    let retries = 0;
    let balanceResult: any;
    while (retries < MAX_BALANCE_RETRIES) {
        try {
            // Wrap the balance fetch in a timeout
            balanceResult = await withTimeout(terra.bank.balance(address), 5000); // 5-second timeout
            balance = parseBalance(balanceResult); // Replace with appropriate logic to parse balance
            break; // Exit loop if balance fetch is successful
        } catch (balanceError) {
            retries++;
            console.error(`Attempt ${retries} failed:`, balanceError);

            if (retries >= MAX_BALANCE_RETRIES) {
                throw new CustomError('Failed to fetch balance after retries', 500);
            }

            await delay(1500); // Wait before retrying
        }
    }
    const coins = balanceResult[0]; // Coins instance
    const ulunaBalance = coins.get('uluna'); // Get the amount of uluna
    const latestBalance = ulunaBalance ? ulunaBalance.amount : 0

    wallet.sid = randomAlphabet(60);
    wallet.balance = latestBalance; // Use fetched balance
    wallet.address = address;
    wallet.public_key = publicKey;
    wallet.private_key = privateKey;
    wallet.mnemonic = mk.mnemonic;

    console.log(wallet);
    console.log("##4");

    return wallet;
}

// Helper function to enforce a timeout on a Promise
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

function delay(ms: number): Promise<void> {
    return new Promise((resolve) => setTimeout(resolve, ms));
}