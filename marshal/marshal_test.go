package marshal

import (
	"testing"
)

// var numCPU = runtime.GOMAXPROCS(0)
// var numCPU1 = runtime.NumCPU()

var Data = initData()
var MarshaledData = jsonMarshal(Data)

func BenchmarkMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sprintfMarshal(Data)
	}
}

func TestStringBuilderMarshal(t *testing.T) {
	stringBuilderfMarshal(Data)
}

func BenchmarkStringBuilderMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringBuilderfMarshal(Data)
	}
}

func BenchmarkFormatMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		formatMarshal(Data)
	}
}

func BenchmarkItoaMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ItoaMarshal(Data)
	}
}

// result:
// BenchmarkMarshal-10    	    1675	    692724 ns/op	 4852857 B/op	    3005 allocs/op

func BenchmarkMarshalByJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonMarshal(Data)
	}
}

// result:
// BenchmarkMarshalByJson-10    	   14377	     81721 ns/op	   49359 B/op	       3 allocs/op

func TestStringUnMarshal(t *testing.T) {
	t.Log(stringUnMarshal(MarshaledData))

}

func TestGjsonUnMarshal(t *testing.T) {
	t.Log(gjsonUnMarshal(MarshaledData))
}

func BenchmarkGjsonUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gjsonUnMarshal(MarshaledData)
	}
}

// result:
// BenchmarkGjsonUnMarshal-10    	   12086	     97824 ns/op	  174256 B/op	      11 allocs/op

func BenchmarkStringUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringUnMarshal(MarshaledData)
	}
}

// result:
// BenchmarkStringUnMarshal-10    	   16905	     69977 ns/op	   16480 B/op	       3 allocs/op

func BenchmarkJsonUnMarshal(b *testing.B) {
	for i := 0; i < b.N; i++ {
		jsonUnMarshal(MarshaledData)
	}
}

// result:
// BenchmarkJsonUnMarshal-10    	    6775	    173832 ns/op	   52768 B/op	    1004 allocs/op
