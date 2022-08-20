package channel

import (
	"runtime"
	"testing"
)

var numCPU = runtime.GOMAXPROCS(0)
var numCPU1 = runtime.NumCPU()

func TestServer(t *testing.T) {
	t.Log(numCPU)
	t.Log(numCPU1)
}
