package dp

import "testing"

func TestAdapter(t *testing.T) {

	adaptee := adapteeImpl{}

	target := NewAdapter(&adaptee)
	// 你看，本来是adaptee的方法，现在被适配到target上了
	res := target.Request()

	if res != "I'm Adaptee method" {
		t.Fatalf("wrong res: %s", res)
	}
}
