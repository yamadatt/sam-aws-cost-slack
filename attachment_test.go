package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tetsuya28/aws-cost-report/i18y"
)

func TestToAttachment(t *testing.T) {
	err := i18y.Init()
	assert.NoError(t, err)
	Language = "en"

	costs := []DailyCost{
		{Services: map[string]ServiceDetail{"svc": {CostAmount: 2, CostUnit: "USD", UsageAmount: 2, UsageUnit: ""}}},
	}

	atts, err := toAttachment(costs)
	assert.NoError(t, err)
	if len(atts) == 0 {
		t.Fatalf("no attachments returned")
	}

	a := atts[0]
	assert.Equal(t, "#ffffff", a.Color)
	fields := a.Fields
	assert.Equal(t, i18y.Translate(Language, "cost"), fields[0].Title)
	assert.Contains(t, fields[0].Value, "USD")
}

func TestToAttachmentInvalid(t *testing.T) {
	_, err := toAttachment([]DailyCost{})
	assert.Error(t, err)
}
