import { CoinRepository } from "../repository/coin.repository";
import { CustomError } from "../types/error";

export class CoinService{
    private coinRepository:CoinRepository;
    
    constructor(wp:CoinRepository){
        this.coinRepository = wp;
    }

    public async getCoins(){
        return await this.coinRepository.getAllCoins();
    }

    public async getCoinById(id:number){
        const coin = await this.coinRepository.getCoinById(id);

        if (!coin)
            throw new CustomError("Coin not found",404);
        
        return coin;
    }

    public async getCoinBySymbol(symbol:string){
        const coin = await this.coinRepository.getCoinBySymbol(symbol);

        if (!coin)
            throw new CustomError("Coin not found",404);
        
        return coin;
    }
}