package main

import (
	"fmt"

	"github.com/pborman/uuid"
)

func main() {
	fmt.Printf("hello world!: %s", uuid.NewRandom().String())
}
