package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	// print a welcome message
	intro()

	// Create a channel to indicate when the user wants to quit
	doneChan := make(chan bool)

	// start a goroutine to read user input and run program
	go readUserInput(os.Stdin, doneChan)

	// block until donechan gets a val
	<-doneChan

	// close the channel
	close(doneChan)

	// goodbye message
	fmt.Println("Goodbye!")

}

func readUserInput(in io.Reader, doneChan chan bool) {
	scanner := bufio.NewScanner(in)

	for {
		res, done := checkInput(scanner)
		if done {
			doneChan <- true
			return
		}

		fmt.Println(res)
		cursor()
	}
}

func checkInput(scanner *bufio.Scanner) (string, bool) {
	scanner.Scan()

	if strings.EqualFold(scanner.Text(), "q") {
		return "", true
	}

	no, err := strconv.Atoi(scanner.Text())
	if err != nil {
		return "", true
	}
	_, res := isPrime(no)
	return res, false

}

func intro() {
	fmt.Println("Check the no is prime or not ")
	fmt.Println("Enter a number if you want to quit please enter q")
	cursor()
}

func cursor() {
	fmt.Print("-->  ")
}

func isPrime(n int) (bool, string) {
	if n == 0 || n == 1 {
		return false, fmt.Sprintf("%d is not prime, by definition", n)
	}

	if n < 0 {
		return false, fmt.Sprintf("%d prime no cannot be negative", n)
	}

	for i := 2; i <= n/2; i++ {
		if n%2 == 0 {
			return false, fmt.Sprintf("%d is not prime no as it is divisible by %d", n, i)
		}
	}

	return true, fmt.Sprintf("%d is a prime number", n)
}
