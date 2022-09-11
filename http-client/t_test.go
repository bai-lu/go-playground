package httpclient

import "testing"

func TestHttp(t *testing.T) {
	Request()
}

func TestHttpWithClose(t *testing.T) {
	HttpWithClose()
}

func TestHttpWithoutClose(t *testing.T) {
	HttpWithoutClose()
}
