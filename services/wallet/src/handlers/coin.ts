import { Empty, Coin, Coins, CoinId, CoinSymbol } from "../../proto/api/service_pb";
import { CoinService } from "../services/coin";
import * as grpc from "@grpc/grpc-js";
import { CustomError } from "../types/error";

export class CoinHandler {

    private coinService: CoinService;

    constructor(ws: CoinService) {
        this.coinService = ws;
        this.getAllCoins = this.getAllCoins.bind(this);
        this.GetCoinById = this.GetCoinById.bind(this);
        this.GetCoinBySymbol = this.GetCoinBySymbol.bind(this);
    }

    async getAllCoins(
        call: grpc.ServerUnaryCall<Empty, Coins>,
        callback: grpc.sendUnaryData<Coins>
    ) {
        try {
            const response = await this.coinService.getCoins();
            const coins = new Coins();
            response.forEach((c)=>{
                const coin = new Coin();
                coin.setCoinId(c.id);
                coin.setCurrencyName(c.currency_name);
                coin.setCurrencySymbol(c.currency_symbol);
                coins.addCoins(coin);
            });
            callback(null, coins);
        } catch (error) {
            console.error("Error in getAllCoins:", error);

            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    async GetCoinById(
        call: grpc.ServerUnaryCall<CoinId, Coin>,
        callback: grpc.sendUnaryData<Coin>
    ) {
        try {
            const response = await this.coinService.getCoinById(call.request.getCoinId());
            const coin = new Coin();
            coin.setCoinId(response.id);
            coin.setCurrencyName(response.currency_name);
            coin.setCurrencySymbol(response.currency_symbol);
            callback(null, coin);
        } catch (error) {
            console.error("Error in GetCoinById:", error);

            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    }

    // get coin by symbol
    async GetCoinBySymbol(
        call: grpc.ServerUnaryCall<CoinSymbol, Coin>,
        callback: grpc.sendUnaryData<Coin>
    ) {
        try {
            const response = await this.coinService.getCoinBySymbol(call.request.getSymbol());
            const coin = new Coin();
            coin.setCoinId(response.id);
            coin.setCurrencyName(response.currency_name);
            coin.setCurrencySymbol(response.currency_symbol);
            callback(null, coin);
        } catch (error) {
            console.error("Error in GetCoinBySymbol:", error);

            if (error instanceof CustomError) {
                callback(error.toGrpcStatus(), null);
            } else {
                callback(
                    {
                        code: grpc.status.UNKNOWN,
                        details: "An unexpected error occurred",
                    },
                    null
                );
            }
        }
    } 


}
