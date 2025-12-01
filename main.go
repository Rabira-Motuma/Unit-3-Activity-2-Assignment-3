/*
 * Author: Rabira Motuma
 * Version: 1.0.0
 * Date: 2025-11-27
 * Fileoverview: This program calculates the minimum amount of coins used to give change.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// variables
	var centsChangeAsString string
	var centsChangeAsNumber int
	var toonies int
	var loonies int
	var quarters int
	var dimes int
	var nickels int
	var pennies int

	reader := bufio.NewReader(os.Stdin)

	// input
	fmt.Printf("How much cents is the change? ")
	centsChangeAsString, _ = reader.ReadString('\n')
	centsChangeAsString = strings.TrimSpace(centsChangeAsString)

	// process
	centsChangeAsNumber, _ = strconv.Atoi(centsChangeAsString)
	toonies = centsChangeAsNumber / 200
	centsChangeAsNumber %= 200
	loonies = centsChangeAsNumber / 100
	centsChangeAsNumber %= 100
	quarters = centsChangeAsNumber / 25
	centsChangeAsNumber %= 25
	dimes = centsChangeAsNumber / 10
	centsChangeAsNumber %= 10
	nickels = centsChangeAsNumber / 5
	centsChangeAsNumber %= 5
	pennies = centsChangeAsNumber / 1
	centsChangeAsNumber %= 1

	// output
	fmt.Println()
	fmt.Printf(`Your change is %d toonies, %d dollars, %d quarters, %d dimes, %d nickels, and %d cents.`, toonies, loonies, quarters, dimes, nickels, pennies)

	fmt.Println("\nDone.")
}
