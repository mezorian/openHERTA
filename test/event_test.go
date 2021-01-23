package main

import (
	"testing"

	"github.com/mezorian/openHERTA/go/pkg/event"
	"github.com/stretchr/testify/assert"
)

func TestEventForCorrectAttributeSetting(t *testing.T) {
	var e event.Event
	assert.Equal(t, e.Name, "", "Error : variable should be empty")
	assert.Equal(t, len(e.Guests), 0, "Error : variable should be zero")
}
