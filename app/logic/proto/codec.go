// Author: sheppard(ysf1026@gmail.com) 2014-01-06

package proto

import (
	"encoding/json"
	"reflect"
)

type Pack struct {
	Type string
	Data interface{}
}

func Encode(v interface{}) string {
	msgName := reflect.TypeOf(v).String()
	msgName = msgName[7:] //TODO better way to remove *proto.

	pack := &Pack{msgName, v}
	b, err := json.Marshal(pack)
	if err != nil {
		panic(err)
	}
	return string(b)
}
