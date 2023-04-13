package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

func TestPrompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("Could create a pipe: %w", err.Error())
	}
	os.Stdout = w
	prompt()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if string(out) != "->" {
		t.Errorf("incorrect promt : expected -> but got %s", string(out))
	}
}
func TestIntro(t *testing.T) {
	oldOut := os.Stdout
	r, w, err := os.Pipe()
	if err != nil {
		t.Fatal("Update the page", err.Error())
	}
	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut
	out, _ := io.ReadAll(r)
	if string(out) != "Is it Prime?, ------------, Enter a whole number, and we'll tell you if it is a prime number or not. Enter q to quit." {
		t.Errorf("Don't get the stirngs as a start %s", string(out))
	}
}
func TestCheckNumbers(t *testing.T) {
	quitScanner := bufio.NewScanner(strings.NewReader("q\n"))
	msg, quit := checkNumbers(quitScanner)
	if msg != "" || quit != true {
		t.Errorf("checkNumbers(quitScanner) = (%s, %t), expected ('', true)", msg, quit)
	}

	nonNumScanner := bufio.NewScanner(strings.NewReader("abc\n"))
	msg, quit = checkNumbers(nonNumScanner)
	if msg != "Please enter a whole number!" || quit != false {
		t.Errorf("checkNumbers(nonNumericScanner) = (%s, %t), expected ('Please enter a whole number!', false)", msg, quit)
	}

	numScanner := bufio.NewScanner(strings.NewReader("42\n"))
	msg, quit = checkNumbers(numScanner)
	if msg != "Number is prime" || quit != false {
		t.Errorf("checkNumbers(numericScanner) = (%s, %t), expected ('Number is prime', false)", msg, quit)
	}

	errScanner := bufio.NewScanner(strings.NewReader("1000\n"))
	msg, quit = checkNumbers(errScanner)
	if msg != "Error checking if number is prime" || quit != false {
		t.Errorf("checkNumbers(errorScanner) = (%s, %t), expected ('Error checking if number is prime', false)", msg, quit)
	}
}
func TestReadUserInput(t *testing.T) {
	doneChan := make(chan bool)

	var stdin bytes.Buffer

	stdin.Write([]byte("1\nq\n"))

	go readUserInput(&stdin, doneChan)
	<-doneChan
	close(doneChan)
}

func TestIsPrime(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2!"},
		{"zero", 0, false, "0 is not prime, by definition!"},
		{"one", 1, false, "1 is not prime, by definition!"},
		{"negative number", -11, false, "Negative numbers are not prime, by definition!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}

		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}
}
