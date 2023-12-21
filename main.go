package main

import (
	"fmt"

	"github.com/Infatzmut/cryptip/decrypt"
	"github.com/Infatzmut/cryptip/encrypt"
)

func main() {
	encryptedStr := encrypt.Nimbus("kodekloud")
	fmt.Println(encryptedStr)
	fmt.Println(decrypt.Nimbus(encryptedStr))
}
