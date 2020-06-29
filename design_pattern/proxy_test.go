package dp

import "testing"

func TestProxy(t *testing.T) {

	sub := &Proxy{Subject{}}

	res := sub.Do()

	if res != "before:do real:after" {
		t.Fatalf("subject get wrong res, %s", res)
	}
}
