package tennis_test

import (
	"testing"

	"github.com/chonlatee/simpleTennis/tennis"
)

type testSuit struct {
	a      int
	b      int
	expect string
}

func TestTennis(t *testing.T) {

	validTestCase := map[int]testSuit{
		// A WIN
		1: {
			a:      0,
			b:      0,
			expect: "LOVE - LOVE",
		},
		2: {
			a:      15,
			b:      0,
			expect: "FIFTEEN - LOVE",
		},
		3: {
			a:      30,
			b:      0,
			expect: "THIRTY - LOVE",
		},
		4: {
			a:      40,
			b:      0,
			expect: "FORTY - LOVE\t-> A WIN",
		},
		// B WIN
		5: {
			a:      0,
			b:      15,
			expect: "LOVE - FIFTEEN",
		},
		6: {
			a:      15,
			b:      15,
			expect: "FIFTEEN - FIFTEEN",
		},
		7: {
			a:      15,
			b:      30,
			expect: "FIFTEEN - THIRTY",
		},
		8: {
			a:      15,
			b:      40,
			expect: "FIFTEEN - FORTY\t-> B WIN",
		},
	}

	invalidTestCase := map[int]testSuit{
		1: {
			a:      10,
			b:      10,
			expect: "Invalid score player A expect `|0, 15, 30, 40|` got `10`",
		},
		2: {
			a:      15,
			b:      -1,
			expect: "Invalid score player B expect `|0, 15, 30, 40|` got `-1`",
		},
	}

	// valid condition
	for _, v := range validTestCase {
		actual, _ := tennis.ShowScore(v.a, v.b)
		if actual != v.expect {
			t.Errorf("Test failed, expected '%s', got '%s'", v.expect, actual)
		}
	}

	// invalid condition
	for _, v := range invalidTestCase {
		_, actual := tennis.ShowScore(v.a, v.b)

		// actual.Error() convert fmt.Errorf() to string
		if actual.Error() != v.expect {
			t.Errorf("Test failed, expected '%s', got '%s'", v.expect, actual)
		}
	}

}
