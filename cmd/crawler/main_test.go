package main

import "testing"

func TestDryRunRejectsDatabaseOnlyOperations(t *testing.T) {
	for _, tc := range []struct{ dry, recrawl, recategorize, reject bool }{
		{true, true, false, true}, {true, false, true, true}, {true, true, true, true},
		{true, false, false, false}, {false, true, false, false}, {false, false, true, false},
	} {
		if err := validateDryRun(tc.dry, tc.recrawl, tc.recategorize); (err != nil) != tc.reject {
			t.Fatalf("mode %+v returned %v", tc, err)
		}
	}
}
