package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Reading from Standard Input
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSuffix(text, "\n")
	splitString := strings.Split(text, " ")

	// Extracting the inputs and typecasting them
	amountToWithdraw, accountBalance := splitString[0], splitString[1]
	amountToWithdrawInt, err := strconv.ParseInt(amountToWithdraw, 0, 64)
	// fmt.Println(i)
	// amountToWithdrawInt, err := strconv.Atoi(amountToWithdraw)
	if err != nil {
		fmt.Println(err)
	}

	accountBalanceInt, err := strconv.ParseFloat(accountBalance, 64)
	// accountBalanceInt, err := strconv.Atoi(accountBalance)
	if err != nil {
		fmt.Println(err)
	}

	// fmt.Printf("Amount to Withdraw: %v and Available balance is: %v\n", amountToWithdrawInt, accountBalanceInt)
	if amountToWithdrawInt%5 == 0 {
		if float64(amountToWithdrawInt) < accountBalanceInt {
			// Successful Transaction
			fmt.Println(accountBalanceInt - float64(amountToWithdrawInt) - 0.50)
		} else {
			// Insufficient Funds
			fmt.Println(accountBalanceInt)
		}
	} else {
		// Incorrect Withdrawal Amount
		// fmt.Println(accountBalanceInt)
		s := fmt.Sprintf("%.2f", accountBalanceInt)
		fmt.Println(s)

		// fmt.Println(math.Round(accountBalanceInt*100) / 100)
	}
}
