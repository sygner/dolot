import { WalletService } from "../services/wallet";
import * as grpc from '@grpc/grpc-js';
import { Wallet as WalletModel } from '../models/wallet';

import { UserId, Address, Wallet, CreateWalletRequest, Wallets, UpdateBalanceRequest, WalletId, Empty, GetWalletByCoinIdAndUserIdRequest, Balance, Sid, GetWalletsByUserIdsAndCoinIdRequest } from '../../proto/api/service_pb';
import { IWalletServiceServer, IWalletServiceService, WalletServiceService } from '../../proto/api/service_grpc_pb';
import { CustomError } from "../types/error";
import { getUnixTimestamp } from "../utils/time";

export class WalletHandler {

    private walletService: WalletService;

    constructor(ws: WalletService) {
        this.walletService = ws;
        this.createWallet = this.createWallet.bind(this);
        this.getWalletsByUserId = this.getWalletsByUserId.bind(this);
        this.getWalletByAddress = this.getWalletByAddress.bind(this);
        this.updateBalance = this.updateBalance.bind(this);
        this.deleteWalletsByUserId = this.deleteWalletsByUserId.bind(this);
        this.deleteWalletByWalletId = this.deleteWalletByWalletId.bind(this);
        this.getBalanceByCoinIdAndUserId = this.getBalanceByCoinIdAndUserId.bind(this);
        this.getWalletBySid = this.getWalletBySid.bind(this);
        this.getWalletByWalletId = this.getWalletByWalletId.bind(this);
        this.getWalletsByUserIdsAndCoinId = this.getWalletsByUserIdsAndCoinId.bind(this);
    }

    async createWallet(
        call: grpc.ServerUnaryCall<CreateWalletRequest, Wallet>,
        callback: grpc.sendUnaryData<Wallet>
    ) {
        const reply = new Wallet();

        try {
            const request: WalletModel = {
                user_id: call.request.getUserId(),
                coin_id: call.request.getCoinId(),
            };

            const response = await this.walletService.createWallet(request);
            reply.setId(response.id!);
            reply.setAddress(response.address!);
            reply.setBalance(response.balance!);
            reply.setCoinId(response.coin_id!);
            reply.setUserId(response.user_id!);
            reply.setPublicKey(response.public_key!);
            reply.setPrivateKey(response.private_key!);
            reply.setSid(response.sid!);
            reply.setCreatedAt(getUnixTimestamp(response.created_at!).toString())
            reply.setUpdatedAt(getUnixTimestamp(response.updated_at!).toString())

            console.log("#4 - Wallet created successfully");
            callback(null, reply);
        } catch (error) {
            console.error('Error in createWallet:', error);

            if (error instanceof CustomError) {
                // Use the CustomError's toGrpcStatus method
                callback(error.toGrpcStatus(), null);
            } else {
                // Handle other unexpected errors
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async getWalletsByUserId(
        call: grpc.ServerUnaryCall<UserId, Wallets>,
        callback: grpc.sendUnaryData<Wallets>
    ) {
        console.log("RRRRRRRRRRRRRRRRRRRRR")
        try {
            const response = await this.walletService.getWalletsByUserId(call.request.getUserId());
            const wallets = new Wallets();
            response.forEach((w) => {
                const wallet = new Wallet();
                wallet.setId(w.id!);
                wallet.setCoinId(w.coin_id!);
                wallet.setUserId(w.user_id!);
                wallet.setAddress(w.address!);
                wallet.setBalance(w.balance!);
                wallet.setPrivateKey(w.private_key!);
                wallet.setPublicKey(w.public_key!);
                wallet.setMnemonic(w.mnemonic!);
                wallet.setSid(w.sid!);
                wallet.setCreatedAt(getUnixTimestamp(w.created_at!).toString());
                wallet.setUpdatedAt(getUnixTimestamp(w.updated_at!).toString());
                wallets.addWallets(wallet);
            });
            console.log("DDDDDDD ",response,wallets)
            callback(null, wallets);
        } catch (error) {
            console.error('Error in getWalletsByUserId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async getWalletByAddress(
        call: grpc.ServerUnaryCall<Address, Wallet>,
        callback: grpc.sendUnaryData<Wallet>
    ) {
        try {
            const response = await this.walletService.getWalletByAddress(call.request.getAddress());
            const wallet = new Wallet();
            wallet.setId(response.id!);
            wallet.setCoinId(response.coin_id!);
            wallet.setUserId(response.user_id!);
            wallet.setAddress(response.address!);
            wallet.setBalance(response.balance!);
            wallet.setPrivateKey(response.private_key!);
            wallet.setPublicKey(response.public_key!);
            wallet.setMnemonic(response.mnemonic!);
            wallet.setSid(response.sid!);
            wallet.setCreatedAt(getUnixTimestamp(response.created_at!).toString());
            wallet.setUpdatedAt(getUnixTimestamp(response.updated_at!).toString());
            callback(null, wallet);
        } catch (error) {
            console.error('Error in getWalletByAddress:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }
    
    public async getWalletBySid(
        call: grpc.ServerUnaryCall<Sid, Wallet>,
        callback: grpc.sendUnaryData<Wallet>
    ) {
        try {
            const response = await this.walletService.getWalletBySid(call.request.getSid());
            const wallet = new Wallet();
            wallet.setId(response.id!);
            wallet.setCoinId(response.coin_id!);
            wallet.setUserId(response.user_id!);
            wallet.setAddress(response.address!);
            wallet.setBalance(response.balance!);
            wallet.setPrivateKey(response.private_key!);
            wallet.setPublicKey(response.public_key!);
            wallet.setMnemonic(response.mnemonic!);
            wallet.setSid(response.sid!);
            wallet.setCreatedAt(getUnixTimestamp(response.created_at!).toString());
            wallet.setUpdatedAt(getUnixTimestamp(response.updated_at!).toString());
            callback(null, wallet);
        } catch (error) {
            console.error('Error in getWalletByAddress:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async getWalletByWalletId(
        call: grpc.ServerUnaryCall<WalletId, Wallet>,
        callback: grpc.sendUnaryData<Wallet>
    ) {
        try {
            const response = await this.walletService.getWalletByWalletId(call.request.getId());
            const wallet = new Wallet();
            wallet.setId(response.id!);
            wallet.setCoinId(response.coin_id!);
            wallet.setUserId(response.user_id!);
            wallet.setAddress(response.address!);
            wallet.setBalance(response.balance!);
            wallet.setPrivateKey(response.private_key!);
            wallet.setPublicKey(response.public_key!);
            wallet.setMnemonic(response.mnemonic!);
            wallet.setSid(response.sid!);
            wallet.setCreatedAt(getUnixTimestamp(response.created_at!).toString());
            wallet.setUpdatedAt(getUnixTimestamp(response.updated_at!).toString());
            callback(null, wallet);
        } catch (error) {
            console.error('Error in getWalletByAddress:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async updateBalance(
        call: grpc.ServerUnaryCall<WalletId, Wallet>,
        callback: grpc.sendUnaryData<Wallet>
    ) {
        try {
            const response = await this.walletService.getAndUpdateWalletBalance(call.request.getId());
            const wallet = new Wallet();
            wallet.setId(response.id!);
            wallet.setCoinId(response.coin_id!);
            wallet.setUserId(response.user_id!);
            wallet.setAddress(response.address!);
            wallet.setBalance(response.balance!);
            wallet.setPrivateKey(response.private_key!);
            wallet.setPublicKey(response.public_key!);
            wallet.setMnemonic(response.mnemonic!);
            wallet.setSid(response.sid!);
            wallet.setCreatedAt(getUnixTimestamp(response.created_at!).toString());
            wallet.setUpdatedAt(getUnixTimestamp(response.updated_at!).toString());
            callback(null, wallet);
        } catch (error) {
            console.error('Error in updateBalance:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async getBalanceByCoinIdAndUserId(
        call: grpc.ServerUnaryCall<GetWalletByCoinIdAndUserIdRequest, Balance>,
        callback: grpc.sendUnaryData<Balance>
    ) {
        try {
            const response = await this.walletService.getBalanceByCoinIdAndUserId(call.request.getCoinId(), call.request.getUserId());
            console.log("Latest ",response)
            callback(null, new Balance().setBalance(response));
        } catch (error) {
            console.error('Error in getBalanceByCoinIdAndUserId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async deleteWalletsByUserId(
        call: grpc.ServerUnaryCall<UserId, Empty>,
        callback: grpc.sendUnaryData<Empty>
    ) {
        try {
            await this.walletService.deleteWalletsByUserId(call.request.getUserId());
            callback(null, new Empty());
        } catch (error) {
            console.error('Error in deleteWalletsByUserId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async deleteWalletByWalletId(
        call: grpc.ServerUnaryCall<WalletId, Empty>,
        callback: grpc.sendUnaryData<Empty>
    ) {
        try {
            await this.walletService.deleteWalletByWalletId(call.request.getId());
            callback(null, new Empty());
        } catch (error) {
            console.error('Error in deleteWalletsByWalletId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }

    public async getWalletsByUserIdsAndCoinId(
        call: grpc.ServerUnaryCall<GetWalletsByUserIdsAndCoinIdRequest, Wallets>,
        callback: grpc.sendUnaryData<Wallets>
    ) {
        try {
            const response = await this.walletService.getWalletsByUserIdsAndCoinId(call.request.getUserIdsList(), call.request.getCoinId());
            const wallets = new Wallets();
            response.forEach((w) => {
                const wallet = new Wallet();
                wallet.setId(w.id!);
                wallet.setCoinId(w.coin_id!);
                wallet.setUserId(w.user_id!);
                wallet.setAddress(w.address!);
                wallet.setBalance(w.balance!);
                wallet.setPrivateKey(w.private_key!);
                wallet.setPublicKey(w.public_key!);
                wallet.setMnemonic(w.mnemonic!);
                wallet.setSid(w.sid!);
                wallet.setCreatedAt(getUnixTimestamp(w.created_at!).toString());
                wallet.setUpdatedAt(getUnixTimestamp(w.updated_at!).toString());
                wallets.addWallets(wallet);
            });
            callback(null, wallets);
        } catch (error) {
            console.error('Error in getWalletsByUserIdsAndCoinId:', error);
            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: 'An unexpected error occurred',
                    },
                    null
                );
            }
        }
    }
}