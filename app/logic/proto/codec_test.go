// Author: sheppard(ysf1026@gmail.com) 2014-01-23

package proto

import (
	"fmt"
	"testing"
)

type data struct {
	V1 int
	V2 string
}

func TestCodec(t *testing.T) {
	// Encode
	d := &data{12, "haha"}
	str := Encode(d)
	fmt.Println(d, str)

	// Decode
	d2 := Decode(str)
	fmt.Println(d2)
}
