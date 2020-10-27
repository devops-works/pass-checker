package main

import (
	"crypto/sha1"
	"flag"
	"fmt"
	"math/big"
	"syscall"

	"golang.org/x/crypto/ssh/terminal"

	passwordvalidator "github.com/lane-c-wagner/go-password-validator"
)

func main() {
	var g int64
	flag.Int64Var(&g, "g", 1000000000, "guesses per second (default: 1.000.000.000)")
	flag.Parse()

	p := getPassword()
	entropy := getEntropy(p)
	fmt.Printf("\nEntropy: %.3f bits\nTime before guaranteed successful crack : %s\n", entropy, getCrackDuration(entropy, g))
	fmt.Println(getSHA1Sum(p))
}

func getPassword() string {
	fmt.Print("Enter Password: ")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		panic(err)
	}
	return string(bytePassword)
}

func getEntropy(p string) float64 {
	return passwordvalidator.GetEntropy(p)
}

func getCrackDuration(e float64, g int64) string {
	ent := big.NewInt(int64(e))
	s := big.NewInt(0).Exp(big.NewInt(2), ent, big.NewInt(0)) // s = 2**e
	s = s.Div(s, big.NewInt(g))                               // s = s/g

	m := big.NewInt(0).Div(s, big.NewInt(60))
	s = s.Mod(s, big.NewInt(60))
	if m.String() == "0" {
		return s.String() + "s"
	}

	h := big.NewInt(0).Div(m, big.NewInt(60))
	m = m.Mod(m, big.NewInt(60))
	if h.String() == "0" {
		return m.String() + "m " + s.String() + "s"
	}

	d := big.NewInt(0).Div(h, big.NewInt(24))
	h = h.Mod(h, big.NewInt(24))
	if d.String() == "0" {
		return h.String() + "h " + m.String() + "m " + s.String() + "s"
	}

	y := big.NewInt(0).Div(d, big.NewInt(365))
	d = d.Mod(d, big.NewInt(365))
	if y.String() == "0" {
		return d.String() + "d " + h.String() + "h " + m.String() + "m " + s.String() + "s"
	}

	return y.String() + "y " + d.String() + "d " + h.String() + "h " + m.String() + "m " + s.String() + "s"
}

func getSHA1Sum(p string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(p)))
}
