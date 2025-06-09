package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	os.Setenv("SLACK_TOKEN", "test-token")
	os.Setenv("SLACK_CHANNEL", "test-channel")
	os.Setenv("LANGUAGE", "en")
	defer os.Unsetenv("SLACK_TOKEN")
	defer os.Unsetenv("SLACK_CHANNEL")
	defer os.Unsetenv("LANGUAGE")

	cfg, err := New()
	assert.NoError(t, err)
	assert.Equal(t, "test-token", cfg.SlackToken)
	assert.Equal(t, "test-channel", cfg.SlackChannel)
	assert.Equal(t, "en", cfg.Language)
}
