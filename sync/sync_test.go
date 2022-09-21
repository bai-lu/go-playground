package sync

import (
	"sync"
	"testing"
)

func TestA(t *testing.T) {
	sm := sync.Map{}
	sm.Store("a", 1)
	// json.Marshal(sm)

}
