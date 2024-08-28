package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/downloadablefox/hts-prog-9/internal"
)

var (
	SudokuString string
	ChiperText   string
)

func init() {
	flag.StringVar(&SudokuString, "sudoku", "", "Sudoku string")
	flag.StringVar(&ChiperText, "chiper", "", "Chiper text")
	flag.Parse()

	if SudokuString == "" && ChiperText == "" {
		flag.Usage()
		os.Exit(1)
	}
}

func decrypt(sudoku internal.Sudoku) {
	fmt.Printf("Solution: %s\n", sudoku.String())

	// Get sha1 of the solution as the key
	key := internal.SHA1(sudoku.String())
	fmt.Printf("Key: [%x]\n", key)

	// Decrypt the cyphertext from base64
	chiper, err := internal.DecodeBase64(ChiperText)
	if err != nil {
		fmt.Println("error: invalid base64 chiper text")
		os.Exit(1)
	}
	fmt.Printf("Chiper: [%x]\n", chiper)

	// Decrypt the chiper text using the key
	blowfish := internal.NewBlowfish(true, true)
	blowfish.SetKey(fmt.Sprintf("%x", key))
	decrypted, err := blowfish.Decrypt(ChiperText)
	if err != nil {
		fmt.Println("error: failed to decrypt chiper text")
		os.Exit(1)
	}

	fmt.Printf("Decrypted: %s\n", string(decrypted))
}

func main() {
	sudoku, err := internal.SudukoFromString(SudokuString)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Sudoku: %+v\n", sudoku)

	// 60 second timer to solve the sudoku
	timer := time.NewTimer(60 * time.Second)

	solved := make(chan internal.Sudoku, 5)
	go sudoku.Solve(solved)

	for {
		select {
		case s := <-solved:
			decrypt(s)
		case <-timer.C:
			os.Exit(1)
		}
	}
}
