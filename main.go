package main

import (
	"fmt"
	"log"

	"github.com/halfloafhq/keymulate/internal/kbd"
)

func main() {
	keyboards, err := kbd.GetKeyboards()
	if err != nil {
		log.Fatalf("Error finding keyboards: %s\n", err.Error())
	}

  fmt.Println(keyboards)
}
