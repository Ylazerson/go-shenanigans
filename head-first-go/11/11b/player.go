// B"H

package main

import "github.com/Ylazerson/go-shenanigans/head-first-go/11/11b/gadget"

// -- -------------------------------------
type Player interface {
	Play(string)
	Stop()
}

// -- -------------------------------------
func tryOutTapePlayer(t gadget.TapePlayer) {
	t.Play("Hashem is here!")
	t.Stop()
}

// -- -------------------------------------
func TryOutGadget(player Player) {

	player.Play("Test Track")
	player.Stop()

	// -- ---------------------------------
	// 99.99% if time you won't need this:
	// type assertion:
	recorder, ok := player.(gadget.TapeRecorder)

	if ok {
		recorder.Record()
	}

}

// -- -------------------------------------
func main() {

	// -- ---------------------------------
	// Create/set a TapePlayer struct:
	var tPlayer gadget.TapePlayer
	tPlayer.Batteries = "Duracell"

	// -- ---------------------------------
	tryOutTapePlayer(tPlayer)

	// -- ---------------------------------
	// Create/set a TapeRecorder struct:
	var tRecorder gadget.TapeRecorder
	tRecorder.Microphones = 2

	// -- ---------------------------------
	// Works:
	TryOutGadget(tPlayer)
	TryOutGadget(tRecorder)

	// -- ---------------------------------
	var ancientPlayer gadget.NoStopPlayer
	ancientPlayer.Eyeballs = "Brown"

	// Won't work:
	// TryOutGadget(ancientPlayer)

	// -- ---------------------------------
	var oddPlayer gadget.OddPlayer
	oddPlayer.Batteries = "AA"

	// Won't work:
	// TryOutGadget(oddPlayer)

}
