package httpclient

import (
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

func Request() {
	res, err := http.Get("http://galleryapi.micloud.xiaomi.net")
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	print(string(body))
	time.Sleep(10 * time.Second)
}

func HttpWithClose() {
	for i := 0; i < 20; i++ {
		resp, err := http.Get("http://galleryapi.micloud.xiaomi.net")
		if err != nil {
			fmt.Println(err)
			return
		}
		resp.Body.Close()
	}
	fmt.Println("goroutine num is", runtime.NumGoroutine())
	time.Sleep(10 * time.Second)
}

func HttpWithoutClose() {
	for i := 0; i < 20; i++ {
		_, err := http.Get("http://galleryapi.micloud.xiaomi.net")
		if err != nil {
			fmt.Println(err)
			return
		}
	}
	fmt.Println("goroutine num is ", runtime.NumGoroutine())
	time.Sleep(10 * time.Second)
}
