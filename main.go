package main

import (
	"fmt"
	"os"
	"regexp"
	"runtime"
	"strings"
	"time"

	"github.com/algorand/go-algorand-sdk/crypto"
	"github.com/algorand/go-algorand-sdk/mnemonic"
)

func printStatus(countChannel chan int) {
	var count int = 0
	var start time.Time = time.Now()

	for {
		count += <-countChannel

		var end time.Time = time.Now()
		var diff float64 = (end.Sub(start).Seconds())

		if diff >= 1 {
			fmt.Printf("\raddrs/sec: %12d", count)
			count = 0
			start = time.Now()
		}
	}
}

func findMatch(needle *regexp.Regexp, total *uint64, foundAccount chan crypto.Account, countChannel chan int) {
	var account crypto.Account = crypto.GenerateAccount()

	for {
		if needle.MatchString(account.Address.String()) {
			foundAccount <- account
			break
		}

		account = crypto.GenerateAccount()
		countChannel <- 1
		*total += 1
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Need to specify address regexp\n")
		return
	}

	// Max number of goroutines
	var numCPU int = runtime.NumCPU()

	// Compile regexp pattern
	var needle *regexp.Regexp = regexp.MustCompile(strings.ToUpper(os.Args[1]))

	// Might get a lot of addresses in total
	var total uint64 = 0

	// Used to store the first match
	var foundAccount = make(chan crypto.Account)

	// Track addr/sec across goroutines
	var count = make(chan int)

	fmt.Printf("looking for address using pattern '%s'\n", needle)
	fmt.Print("??? addr/sec")

	for i := 0; i < numCPU; i++ {
		go findMatch(needle, &total, foundAccount, count)
	}

	// Start outputting the status
	go printStatus(count)

	// Wait for the first match, this will "stop" the goroutines
	var account crypto.Account = <-foundAccount

	// Newline to clear the status output
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
