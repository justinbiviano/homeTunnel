package main

import (
	"fmt"

	"github.com/justinbiviano/homeTunnel/crypto"
)

func main() {
	fmt.Println("Nothing Started Here")
	public, private, _ := crypto.GeneratePrivatePublicKeys()
	fmt.Println(public, private)
}
