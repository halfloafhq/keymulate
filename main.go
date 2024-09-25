package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/halfloafhq/keymulate/internal/kbd"
)

func main() {
	switchPtr := flag.String("switch", "blue", "The sound of switches to output")
  flag.Parse()

	keyboards, err := kbd.GetKeyboards()
	if err != nil {
		log.Fatalf("Error finding keyboards: %s\n", err.Error())
	}

	events := kbd.GetEvents(keyboards)

	//listen to events
	kbd.Listen(*switchPtr, events)
}
