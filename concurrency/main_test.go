package concurrency

import "testing"

func TestTaskManager(t *testing.T) {
	TaskManager()

}

func TestSay(t *testing.T) {
	go Say("world")
	Say("hello")

}

func TestCutSlice(t *testing.T) {
	CutSlice()
}
