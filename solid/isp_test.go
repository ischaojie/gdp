package solid

import "testing"

func TestIsp(t *testing.T) {
	prettyGirl := PrettyGirl{"张子枫"}

	searcher := Searcher{prettyGirl}
	searcher.show()

}
