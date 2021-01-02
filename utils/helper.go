package utils

import (
	"fmt"
	"testing"

	"gotest.tools/v3/assert"
)

func Checkf(t *testing.T, comparison assert.BoolOrComparison, testCase interface{}) {
	assert.Check(t, comparison, fmt.Sprintf("%+v", testCase))
}
