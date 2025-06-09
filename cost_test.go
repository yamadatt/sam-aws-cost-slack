package main

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/costexplorer"
	"github.com/stretchr/testify/assert"
	"github.com/tetsuya28/aws-cost-report/testdata"
)

func TestToCost(t *testing.T) {
	data, err := testdata.GetCostAndUsage()
	assert.NoError(t, err)
	if len(data.ResultsByTime) == 0 {
		t.Fatalf("no result")
	}

	g := data.ResultsByTime[0].Groups[0]
	detail, err := toCost(g)
	assert.NoError(t, err)
	assert.Equal(t, 0.0, detail.CostAmount)
	assert.Equal(t, "USD", detail.CostUnit)
	assert.Equal(t, 0.0, detail.UsageAmount)
	assert.Equal(t, "", detail.UsageUnit)

	// nil group should not fail
	d, err := toCost(&costexplorer.Group{})
	assert.NoError(t, err)
	assert.Equal(t, 0.0, d.CostAmount)
}
