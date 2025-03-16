export function randomNumberBetween(min:number, max:number):number{
    return Math.floor(Math.random() * (max - min + 1) + min);
}

export function randomAlphabet(length: number = 1): string {
    const alphabet = 'abcdefghijklmnopqrstuvwxyz';
    let result = '';

    for (let i = 0; i < length; i++) {
        const randomIndex = randomNumberBetween(0, alphabet.length - 1);
        result += alphabet[randomIndex];
    }

    return result;
}

export function randomAlphanumeric(length: number = 1): string {
    const characters = 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789';
    let result = '';

    for (let i = 0; i < length; i++) {
        const randomIndex = randomNumberBetween(0, characters.length - 1);
        result += characters[randomIndex];
    }

    return result;
}
