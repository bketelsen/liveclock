package main

import (
	"github.com/bketelsen/liveclock/liveclock"
)

func main() {
	// Localtime Live Clock Cog
	lc := liveclock.NewLiveClock()
	lc.SetID("myLiveClock")
	lc.CogInit(nil)
	err := lc.Start()
	if err != nil {
		println("Encountered the following error when attempting to start the local liveclock cog: ", err)
	}
}
