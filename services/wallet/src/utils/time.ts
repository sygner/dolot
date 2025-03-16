export function getUnixTimestamp(date: Date | string): number {
    console.log(date)
    const parsedDate = typeof date === "string" ? new Date(date) : date;
    if (!(parsedDate instanceof Date) || isNaN(parsedDate.getTime())) {
        throw new Error("Invalid date object or string passed to getUnixTimestamp");
    }
    return Math.floor(parsedDate.getTime() / 1000);
}

export function Delay(ms: number) {
    return new Promise( resolve => setTimeout(resolve, ms) );
}