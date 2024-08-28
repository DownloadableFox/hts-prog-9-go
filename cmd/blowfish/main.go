package main

import (
	"fmt"

	"github.com/downloadablefox/hts-prog-9/internal"
)

func main() {
	text := "hello world!"
	key := "mykey"

	blowfish := internal.NewBlowfish(false, false)
	blowfish.SetKey(key)

	encrypted := blowfish.Encrypt(text)
	fmt.Printf("Encrypted: %s\n", encrypted)

	decrypted, err := blowfish.Decrypt(encrypted)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	fmt.Printf("Decrypted: %s\n", decrypted)
}
