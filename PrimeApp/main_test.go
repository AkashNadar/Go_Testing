package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
	"testing"
)

func Test_isPrime(t *testing.T) {
	// res, msg := isPrime(0)

	// if res {
	// 	t.Errorf("Expected false")
	// }

	// if msg != fmt.Sprintf("%d is not prime, by definition", 0) {
	// 	t.Errorf("Wrong returned msg %s", msg)
	// }

	// Table format
	testCases := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number"},
		{"not_prime", 8, false, "8 is not prime no as it is divisible by 2"},
		{"no is 0", 0, false, "0 is not prime, by definition"},
		{"no is -ve", -10, false, "-10 prime no cannot be negative"},
	}

	for _, test := range testCases {
		isPrime, msg := isPrime(test.testNum)

		if test.expected && !isPrime {
			t.Errorf("Expected true value but got false")
		}

		if !test.expected && isPrime {
			t.Errorf("Expected false but got true")
		}

		if test.msg != msg {
			t.Errorf("Expected %s \n Got %s", test.msg, msg)
		}
	}

}

func Test_cursor(t *testing.T) {
	// save the copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	cursor()

	// close the writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our cursor() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if string(out) != "-->  " {
		t.Errorf("Expected '-->' but got '%s'", string(out))
	}
}

func Test_intro(t *testing.T) {
	// save the copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	intro()

	// close the writer
	_ = w.Close()

	// reset os.Stdout to what it was before
	os.Stdout = oldOut

	// read the output of our cursor() func from our read pipe
	out, _ := io.ReadAll(r)

	// perform our test
	if !strings.Contains(string(out), "Enter a number") {
		t.Errorf("Expected not found but got '%s'", string(out))
	}
}

func Test_checkInput(t *testing.T) {
	testCases := []struct {
		name    string
		testNum string
		exit    bool
		msg     string
	}{
		{"prime", "7", false, "7 is a prime number"},
		{"not_prime", "8", false, "8 is not prime no as it is divisible by 2"},
		{"no is 0", "0", false, "0 is not prime, by definition"},
		{"no is -ve", "-10", false, "-10 prime no cannot be negative"},
		{"not a whole num", "th", true, ""},
		{"q is pressed", "q", true, ""},
	}

	for _, test := range testCases {
		input := strings.NewReader(fmt.Sprintf("%v", test.testNum))
		reader := bufio.NewScanner(input)
		res, exit := checkInput(reader)

		if !strings.EqualFold(res, test.msg) {
			t.Errorf("expected %s, got %s\n", test.msg, res)
		}

		if test.exit && !exit {
			t.Errorf("Expected true")
		}

		if !test.exit && exit {
			t.Errorf("Expected false")
		}
	}
}

func Test_readUserInput(t *testing.T) {
	// to test this func we need channel and an instance of io.Reader
	doneChan := make(chan bool)

	// create a reference to a bytes.Buffer
	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}
