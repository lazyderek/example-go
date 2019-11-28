package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseToml(t *testing.T) {
	assert.NoError(t, ParseToml())
}
