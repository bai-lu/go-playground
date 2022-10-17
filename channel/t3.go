package channel

func A() {
	ch := make(chan result, 3)
	ch <- result{"a", nil}

	// a := <-ch
	println(len(ch), cap(ch))
}
