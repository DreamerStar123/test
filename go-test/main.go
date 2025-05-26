package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	id, err := uuid.NewRandom()
	if err != nil {
		return
	}
	fmt.Print(id.String())
}
