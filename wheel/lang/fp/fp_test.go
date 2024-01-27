package fp

import (
	"testing"
)

func TestFP2(t *testing.T) {
	fn := func() {
		println("fn")
	}
	fn = nil
	fn()
}

func TestFP(t *testing.T) {
	noStopLoop()()
}
