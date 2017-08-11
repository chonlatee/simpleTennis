package tennis_test

import (
	"testing"

	"github.com/chonlatee/simpleTennis/tennis"
)

type scoreTestSuit struct {
	scoreA int
	scoreB int
	expect string
}

func TestTennis(t *testing.T) {

	validTestCase := []scoreTestSuit{
		// A WIN
		{
			scoreA: 0,
			scoreB: 0,
			expect: "LOVE - LOVE",
		},
		{
			scoreA: 15,
			scoreB: 0,
			expect: "FIFTEEN - LOVE",
		},
		{
			scoreA: 30,
			scoreB: 0,
			expect: "THIRTY - LOVE",
		},
		{
			scoreA: 40,
			scoreB: 0,
			expect: "FORTY - LOVE\t-> A WIN",
		},
		// B WIN
		{
			scoreA: 0,
			scoreB: 15,
			expect: "LOVE - FIFTEEN",
		},
		{
			scoreA: 15,
			scoreB: 15,
			expect: "FIFTEEN - FIFTEEN",
		},
		{
			scoreA: 15,
			scoreB: 30,
			expect: "FIFTEEN - THIRTY",
		},
		{
			scoreA: 15,
			scoreB: 40,
			expect: "FIFTEEN - FORTY\t-> B WIN",
		},
	}

	invalidTestCase := []scoreTestSuit{
		{
			scoreA: 10,
			scoreB: 10,
			expect: "Invalid score player A expect `|0, 15, 30, 40|` got `10`",
		},
		{
			scoreA: 15,
			scoreB: -1,
			expect: "Invalid score player B expect `|0, 15, 30, 40|` got `-1`",
		},
	}

	// valid condition
	for _, v := range validTestCase {
		actual, _ := tennis.ShowScore(v.scoreA, v.scoreB)
		if actual != v.expect {
			t.Errorf("Test failed, expected '%s', got '%s'", v.expect, actual)
		}
	}

	// invalid condition
	for _, v := range invalidTestCase {
		_, actual := tennis.ShowScore(v.scoreA, v.scoreB)

		// actual.Error() convert fmt.Errorf() to string
		if actual.Error() != v.expect {
			t.Errorf("Test failed, expected '%s', got '%s'", v.expect, actual)
		}
	}

}
