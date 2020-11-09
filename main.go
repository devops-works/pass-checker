package main

import (
	"bufio"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"syscall"

	passwordvalidator "github.com/lane-c-wagner/go-password-validator"
	log "github.com/sirupsen/logrus"

	"golang.org/x/crypto/ssh/terminal"
)

var (
	// Version holds build version
	Version string
	// BuildDate holds build date
	BuildDate string
)

func main() {
	var g int64
	flag.Int64Var(&g, "g", 1000000000, "guesses per second (default: 1000000000)")
	flag.Parse()

	fmt.Printf("pass-checker %s (built %s)\n", Version, BuildDate)

	stat, _ := os.Stdin.Stat()
	var p string

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// the input is piped
		r := bufio.NewReader(os.Stdin)

		var output []rune
		for {
			input, _, err := r.ReadRune()
			if err != nil && err == io.EOF {
				break
			}
			output = append(output, input)
		}

		p = strings.Replace(string(output), "\n", "", -1)
	} else {
		// get the input from the terminal
		p = getPassword()
	}

	entropy := getEntropy(p)

	fmt.Printf("Entropy: %.3f bits\nTime before guaranteed successful crack : %s\n", entropy, getCrackDuration(entropy, g))

	callAPI(getSHA1Sum(p))
}

func getPassword() string {
	fmt.Print("Enter password:\n")
	bytePassword, err := terminal.ReadPassword(int(syscall.Stdin))
	if err != nil {
		log.Fatalf("error reading password from stdin: %v", err)
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

	mil := big.NewInt(0).Div(y, big.NewInt(1000))
	y = y.Mod(y, big.NewInt(1000))
	if mil.String() == "0" {
		return y.String() + "y " + d.String() + "d " + h.String() + "h " + m.String() + "m " + s.String() + "s"
	}

	return m.String() + "mil " + y.String() + "y " + d.String() + "d " + h.String() + "h " + m.String() + "m " + s.String() + "s"
}

func getSHA1Sum(p string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(p)))
}

func callAPI(h string) {
	r, err := http.Get("https://api.pwnedpasswords.com/range/" + h[0:5])
	if err != nil {
		log.Errorf("an error occured while contacting API : %v", err)
		return
	}
	defer r.Body.Close()

	if r.StatusCode != 200 {
		log.Errorf("HTTP status code is not 200")
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Errorf("an error occured while reading response body : %v", err)
		return
	}

	if strings.Contains(string(body), strings.ToUpper(h)[5:]) {
		fmt.Println("Your password's hash exists in HaveIBeenPwned database.")
	}
}
