// Author: sheppard(ysf1026@gmail.com) 2014-01-09

package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	name := "sheppard"
	u := NewUser(name, nil, nil, nil)
	if u.Name != name {
		t.Fail()
	}
}
