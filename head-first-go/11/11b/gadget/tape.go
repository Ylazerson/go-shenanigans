// B"H

package gadget

import "fmt"

// -- -------------------------------------
type TapePlayer struct {
	Batteries string
}

func (t TapePlayer) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapePlayer) Stop() {
	fmt.Println("Stopped!")
}

// -- -------------------------------------
type TapeRecorder struct {
	Microphones int
}

func (t TapeRecorder) Play(song string) {
	fmt.Println("Playing", song)
}

func (t TapeRecorder) Record() {
	fmt.Println("Recording")
}

func (t TapeRecorder) Stop() {
	fmt.Println("Stopped!")
}

// -- -------------------------------------
// Won't satisfy the interface.
// -- -------------------------------------
type NoStopPlayer struct {
	Eyeballs string
}

func (t NoStopPlayer) Play(song string) {
	fmt.Println("Playing", song)
}

// -- -------------------------------------
// Won't satisfy the interface.
// -- -------------------------------------
type OddPlayer struct {
	Batteries string
}

func (t OddPlayer) Play(song int64) {
	fmt.Println("Playing song number", song)
}

func (t OddPlayer) Stop() {
	fmt.Println("Stopped!")
}
