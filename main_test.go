package main

import (
	"reflect"
	"testing"
)

func TestEven(t *testing.T) {
	testCases := []struct {
		sString        string
		tString        string
		expectedResult bool
		testCaseID     string
	}{
		{
			sString:        "anagram",
			tString:        "nagaram",
			expectedResult: true,
			testCaseID:     "test1",
		},
		{
			sString:        "rat",
			tString:        "tarr",
			expectedResult: false,
			testCaseID:     "test2",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.testCaseID, func(t *testing.T) {
			result := isAnagram(tc.sString, tc.tString)
			if !reflect.DeepEqual(result, tc.expectedResult) {
				t.Errorf("Test Case %s failed. Expected: %v, Got: %v", tc.testCaseID, tc.expectedResult, result)
			}
		})
	}
}
