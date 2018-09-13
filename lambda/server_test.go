package lambda

import (
	"testing"
)

func TestNewTile(t *testing.T) {
	type testCase struct {
		code string
		err  string
	}

	tests := map[string]testCase{
		"Print 1": {
			code: "print(1)",
		},
		"err": {
			code: "pr(1)",
			err: `Traceback (most recent call last):
  File "<string>", line 1, in <module>
NameError: name 'pr' is not defined`,
		},
	}

	fn := func(tc testCase) func(*testing.T) {
		return func(*testing.T) {
			r := RunCode(tc.code)
			if r.Error != tc.err {
				println(r.Error)
				t.Fatal("Expected: %s, got: %s", tc.err, r.Error) //TODO finish this test
			}
		}
	}

	for name, tc := range tests {
		t.Run(name, fn(tc))
	}

}
