package password

import "testing"

type testCase struct {
	name     string
	password string
	Criteria
}

func TestValidate(t *testing.T) {
	testCases := []testCase{
		{
			name:     "absolutely valid password",
			password: "SomethingStrange1%",
			Criteria: Criteria{
				Length:   true,
				Number:   true,
				Upper:    true,
				Special:  true,
				Position: 0,
			},
		},
		{
			name:     "short and invalid password",
			password: "1211Some?",
			Criteria: Criteria{
				Length:   false,
				Number:   true,
				Upper:    true,
				Special:  false,
				Position: 9,
			},
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var result Criteria
			result = Validate(tc.password)
			if result != tc.Criteria {
				t.Errorf("\nInput was %v\nGot %v\nExpected %v", tc.password, result, tc.Criteria)
			}
		})

	}
}

func TestCharacterSequence(t *testing.T) {
	for i := 0; i <= 10; i++ {
		ch := CharacterSequence()
		vc := Validate(ch)
		t.Run("Passing the validation", func(t *testing.T) {
			if vc != perfect {
				t.Errorf("Generated password didnt pass validation %v, %v", ch, vc)
			}
		})
	}
}

func TestWordsSequence(t *testing.T) {
	err := GenerateWordSlice()
	if err != nil {
		t.Errorf("Couldnt generate word slice %v", err)
	}
	for i := 0; i <= 10; i++ {
		ws := WordsSequence()

		vc := Validate(ws)
		lenws := len(ws)
		t.Run("Passing the validation", func(t *testing.T) {
			if vc != perfect {
				t.Errorf("\n Word sequence doesnt pass validation %v, %v", vc, ws)
			}
		})
		t.Run("Long enough", func(t *testing.T) {
			if lenws < 20 {
				t.Errorf("Expected len to be more than 20, got %v", lenws)
			}
		})
	}

}
