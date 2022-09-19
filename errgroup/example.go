package errgroup

import (
	// "crypto/rand"

	"fmt"
	"math/rand"

	"golang.org/x/sync/errgroup"
)

var wg errgroup.Group

func F1() {

	subF := func() error {
		num := rand.Intn(10)
		fmt.Println(num)
		if num > 5 {
			return fmt.Errorf("available num %d", num)
		}
		return nil
	}
	for i := 0; i < 10; i++ {
		wg.Go(subF)
	}
	err := wg.Wait()
	fmt.Println(err)
}
