package main

import (
	"github.com/mezorian/openHERTA/go/pkg/guest"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGuestIsEmptyAfterCreation(t *testing.T) {
	var g guest.Guest
  assert.Equal(t, g.FirstName, "", "Error : variable should be empty")
	assert.Equal(t, g.LastName, "", "Error : variable should be empty")
	assert.Equal(t, g.ConfirmationStatus, guest.None, "Error : variable should be None")
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0, "Error : variable should be zero")
}
