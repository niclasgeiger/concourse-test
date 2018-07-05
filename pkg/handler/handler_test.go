package handler

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculator_Do(t *testing.T) {
	testCases := map[string]struct {
		A              int
		B              int
		FunctionCalled string
		Operation      Operation
		ExpectedResult int
		ExpectedErr    error
	}{

		"add": {
			A:              1,
			B:              1,
			FunctionCalled: "Sum",
			Operation:      ADD,
			ExpectedResult: 2,
			ExpectedErr:    nil,
		},
		"sub": {
			A:              3,
			B:              7,
			FunctionCalled: "Sub",
			Operation:      SUB,
			ExpectedResult: -4,
			ExpectedErr:    nil,
		},
		"prod": {
			A:              6,
			B:              7,
			FunctionCalled: "Prod",
			Operation:      PROD,
			ExpectedResult: 42,
			ExpectedErr:    nil,
		},
		"unknown operation": {
			A:              13,
			B:              7,
			FunctionCalled: "none",
			Operation:      "bla",
			ExpectedResult: 0,
			ExpectedErr:    unknownOperationErr,
		},
	}
	for desc, test := range testCases {
		t.Run(desc, func(t *testing.T) {
			calc := new(calculator)
			actual, err := calc.Do(test.A, test.B, test.Operation)
			assert.Equal(t, test.ExpectedResult, actual)
			assert.Equal(t, test.ExpectedErr, err)
		})
	}
}
