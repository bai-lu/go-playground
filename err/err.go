package err

import "time"

func triggerPass0() {
	defer func() {
		println("defer triggerPass0")
		recover()
		println("recoverd")
	}()
	triggerPass1()
	time.Sleep(time.Second * 1)

}

func triggerPass1() {
	defer func() {
		println("defer triggerPass1")
	}()
	triggerPanic()
	time.Sleep(time.Second * 1)
}

func triggerPanic() {
	defer func() {
		time.Sleep(time.Second * 3)
		println("defer triggerPanic")

	}()
	panic("trigger err")
}

func main() {
	go triggerPass0()
	// go triggerPanic()
	time.Sleep(time.Second * 10)
}
