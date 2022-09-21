package net

import (
	"context"
	"fmt"
	"net"
	"testing"
)

func TestA(t *testing.T) {
	res, err := net.DefaultResolver.LookupNS(context.TODO(), "i.mi.com")
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%+v", res)
	t.Log(net.LookupIP("c3-mc-sre00.bj"))
}

func TestB(t *testing.T) {
	var slice []int
	// var slice = []int{}
	if slice == nil {
		println("slice in nil")
		slice = append(slice, 0)
	}
	fmt.Println(slice)
}
