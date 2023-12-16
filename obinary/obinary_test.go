package obinary_test

import (
	"bytes"
	// "fmt"
	"github.com/kumsumit/orientgo/obinary/rw"
	// "path/filepath"
	// "reflect"
	// "runtime"
	"testing"
)

// equals fails the test if exp is not equal to act.
// func equals(tb testing.TB, exp, act interface{}) {
// 	if !reflect.DeepEqual(exp, act) {
// 		_, file, line, _ := runtime.Caller(1)
// 		fmt.Printf("\033[31m%s:%d:\n\n\texp: %#v\n\n\tgot: %#v\033[39m\n\n",
// 			filepath.Base(file), line, exp, act)
// 		tb.FailNow()
// 	}
// }

func TestReadErrorResponseWithSingleException(t *testing.T) {
	buf := new(bytes.Buffer)
	bw := rw.NewWriter(buf)
	bw.WriteByte(byte(1)) // indicates continue of exception class/msg array
	bw.WriteStrings("org.foo.BlargException", "wibble wibble!!")
	bw.WriteByte(byte(0)) // indicates end of exception class/msg array
	bw.WriteBytes([]byte("this is a stacktrace simulator\nEOL"))
}

func TestReadErrorResponseWithMultipleExceptions(t *testing.T) {
	buf := new(bytes.Buffer)
	bw := rw.NewWriter(buf)
	bw.WriteByte(byte(1)) // indicates more exceptions to come
	bw.WriteStrings("org.foo.BlargException", "Too many blorgles!!")
	bw.WriteByte(byte(1)) // indicates more exceptions to come
	bw.WriteStrings("org.foo.FeebleException", "Not enough juice")
	bw.WriteByte(byte(1)) // indicates more exceptions to come
	bw.WriteStrings("org.foo.WobbleException", "Orbital decay")
	bw.WriteByte(byte(0)) // indicates end of exceptions
	bw.WriteBytes([]byte("this is a stacktrace simulator\nEOL"))
}
