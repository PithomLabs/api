package password

import (
	"testing"
)

type testCase struct {
	password string
	Criteria
}

func TestValidate(t *testing.T) {
	testCases := []testCase{
		{
			password: "SomethingStrange1%",
			Criteria: Criteria{
				true,
				true,
				true,
				true,
				0,
			},
		},
		{
			password: "1211Some?",
			Criteria: Criteria{
				false,
				true,
				true,
				false,
				9,
			},
		},
	}
	for _, tc := range testCases {
		var res Criteria
		res = Validate(tc.password)
		if res != tc.Criteria {
			t.Errorf("\nInput was %v\nGot %v\nExpected %v", tc.password, res, tc.Criteria)
		}
	}
}
