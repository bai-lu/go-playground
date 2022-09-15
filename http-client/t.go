package httpclient

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

type Closer struct{}

type readCloser struct {
	reader *bytes.Buffer
}

func Request() (*http.Response, error) {
	resp, err := http.Get("http://galleryapi.micloud.xiaomi.net")
	if err != nil {
		panic(err)
	}
	buff := bytes.NewBuffer([]byte(""))

	io.Copy(buff, resp.Body)
	resp1 := new(http.Response)
	resp1.Body = &readCloser{
		reader: buff,
	}

	defer resp.Body.Close()
	return resp1, err

}

func (rc *readCloser) Read(b []byte) (int, error) {
	return rc.reader.Read(b)
}

func (rc *readCloser) Close() error {
	return nil
}

func Read() {
	resp, err := Request()
	if err != nil {
		panic(err)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	print(string(body))
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
