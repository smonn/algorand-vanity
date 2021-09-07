package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Need to specify address prefix\n")
		return
	}

	start := time.Now()
	needle := regexp.MustCompile(strings.ToUpper(os.Args[1]))
	account := crypto.GenerateAccount()

	// Unlikely we hit more than a 32-bit integer for count as we reset every second
	var count uint = 0
	// Might get a lot of addresses in total
	var total uint64 = 0

	fmt.Printf("looking for prefix %s\n", needle)
	fmt.Print("??? addr/sec")

	for {
		if needle.MatchString(account.Address.String()) {
			break
		}

		account = crypto.GenerateAccount()
		count += 1
		total += 1
		end := time.Now()
		diff := end.Sub(start).Seconds()

		if diff >= 1 {
			start = time.Now()
			fmt.Printf("\r%d addr/sec", count)
			count = 0
		}
	}

	fmt.Println()
	m, err := mnemonic.FromPrivateKey(account.PrivateKey)

	if err != nil {
		fmt.Printf("error: %s\n", err.Error())
		return
	}

	fmt.Printf("checked %d addresses\n", total)
	fmt.Printf("account mnemonic: %s\n", m)
	fmt.Printf("account address: %s\n", account.Address)
}
