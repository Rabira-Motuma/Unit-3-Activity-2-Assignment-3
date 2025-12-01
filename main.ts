/**
 * @author Rabira Motuma
 * @version 1.0.0
 * @date 2025-11-28
 * @fileoverview This calculates the minimum amount of coins used to give change.
 */

// variables
let centsChangeAsString: string;
let centsChangeAsNumber: number;
let toonies: number;
let loonies: number;
let quarters: number;
let dimes: number;
let nickels: number;
let pennies: number;

// input
centsChangeAsString = prompt("How much cents is the change? ") || "0";

// process
centsChangeAsNumber = parseInt(centsChangeAsString, 10);
toonies = Math.floor(centsChangeAsNumber / 200);
centsChangeAsNumber %= 200;
loonies = Math.floor(centsChangeAsNumber / 100);
centsChangeAsNumber %= 100;
quarters = Math.floor(centsChangeAsNumber / 25);
centsChangeAsNumber %= 25;
dimes = Math.floor(centsChangeAsNumber / 10);
centsChangeAsNumber %= 10;
nickels = Math.floor(centsChangeAsNumber / 5);
centsChangeAsNumber %= 5;
pennies = Math.floor(centsChangeAsNumber / 1);
centsChangeAsNumber %= 1;

// output
console.log("\n")
console.log(
  `Your change is: ${toonies} toonies, ${loonies} dollars, ${quarters} quarters, ${dimes} dimes, ${nickels} nickles, and ${pennies} cents.`)

console.log("\nDone.")
