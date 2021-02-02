package main

import (
	"reflect"
	"testing"

	"github.com/mezorian/openHERTA/go/pkg/event"
	"github.com/mezorian/openHERTA/go/pkg/guest"
	"github.com/stretchr/testify/assert"
)

// Test Event for correct attribute setting

func TestEventHasEmptyInializationVariables(t *testing.T) {
	var e event.Event
	assert.Equal(t, e.Name, "", "Error : variable should be empty")
	assert.Equal(t, len(e.Guests), 0, "Error : variable should be zero")
}

// Test Event for adding and counting guests
func TestThatThereAreNoGuestsAtTheVeryBeginning(t *testing.T) {
	var e event.Event

	// get guest confirmation information
	confirmedGuests := e.GetConfirmedGuests()
	interestedGuests := e.GetInterestedGuests()
	declinedGuests := e.GetDeclinedGuests()
	numberOfConfirmedGuests := e.GetNumberOfConfirmedGuests()
	numberOfInterestedGuests := e.GetNumberOfInterestedGuests()
	numberOfDeclinedGuests := e.GetNumberOfDeclinedGuests()

	// check if guest confirmation information is correct
	assert.Equal(t, len(confirmedGuests), 0)
	assert.Equal(t, len(interestedGuests), 0)
	assert.Equal(t, len(declinedGuests), 0)

	assert.Equal(t, numberOfConfirmedGuests, 0)
	assert.Equal(t, numberOfInterestedGuests, 0)
	assert.Equal(t, numberOfDeclinedGuests, 0)
}

func TestThatThereAreTheRightGuestsAndCorrectNumberOfGuestsAfterAddingSome(t *testing.T) {
	var e event.Event

	// prepare guests
	var guest1 guest.Guest
	guest1.FirstName = "Max"
	guest1.LastName = "Mustermann"
	guest1.ConfirmationStatus = guest.Confirmed
	guest1.OtherPeopleThisGuestWillBring = 4
	var guest2 guest.Guest
	guest2.FirstName = "Mia"
	guest2.LastName = "Wong"
	guest2.ConfirmationStatus = guest.Interested
	guest2.OtherPeopleThisGuestWillBring = 2
	var guest3 guest.Guest
	guest3.FirstName = "Peter"
	guest3.LastName = "Stanbridge"
	guest3.ConfirmationStatus = guest.Interested
	guest3.OtherPeopleThisGuestWillBring = 0
	var guest4 guest.Guest
	guest4.FirstName = "Maria"
	guest4.LastName = "Musterfrau"
	guest4.ConfirmationStatus = guest.Declined
	guest4.OtherPeopleThisGuestWillBring = 0
	var guest5 guest.Guest
	guest5.FirstName = "John"
	guest5.LastName = "Stone"
	guest5.ConfirmationStatus = guest.Declined
	guest5.OtherPeopleThisGuestWillBring = 0

	// add guests to event
	e.AddGuest(&guest1)
	e.AddGuest(&guest2)
	e.AddGuest(&guest3)
	e.AddGuest(&guest4)
	e.AddGuest(&guest5)

	// get guest confirmation information
	confirmedGuests := e.GetConfirmedGuests()
	interestedGuests := e.GetInterestedGuests()
	declinedGuests := e.GetDeclinedGuests()
	numberOfConfirmedGuests := e.GetNumberOfConfirmedGuests()
	numberOfInterestedGuests := e.GetNumberOfInterestedGuests()
	numberOfDeclinedGuests := e.GetNumberOfDeclinedGuests()

	// check if guest confirmation information is correct

	assert.True(t, reflect.DeepEqual(confirmedGuests, []*guest.Guest{&guest1}))
	assert.True(t, reflect.DeepEqual(interestedGuests, []*guest.Guest{&guest2, &guest3}))
	assert.True(t, reflect.DeepEqual(declinedGuests, []*guest.Guest{&guest4, &guest5}))

	assert.Equal(t, numberOfConfirmedGuests, 5)  // (1 + 4)
	assert.Equal(t, numberOfInterestedGuests, 4) // (1 + 2) + (1 + 0)
	assert.Equal(t, numberOfDeclinedGuests, 2)   // (1 + 0) + (1 + 0)

	// guest one is only interested
	e.Guests[0].Interest()

	// get guest confirmation information
	confirmedGuests = e.GetConfirmedGuests()
	interestedGuests = e.GetInterestedGuests()
	declinedGuests = e.GetDeclinedGuests()
	numberOfConfirmedGuests = e.GetNumberOfConfirmedGuests()
	numberOfInterestedGuests = e.GetNumberOfInterestedGuests()
	numberOfDeclinedGuests = e.GetNumberOfDeclinedGuests()

	// check if guest confirmation information is correct
	assert.Equal(t, len(confirmedGuests), 0)
	assert.True(t, reflect.DeepEqual(interestedGuests, []*guest.Guest{&guest1, &guest2, &guest3}))
	assert.True(t, reflect.DeepEqual(declinedGuests, []*guest.Guest{&guest4, &guest5}))

	assert.Equal(t, numberOfConfirmedGuests, 0)  // (0 + 0)
	assert.Equal(t, numberOfInterestedGuests, 9) // (1 + 4) + (1 + 2) + (1 + 0)
	assert.Equal(t, numberOfDeclinedGuests, 2)   // (1 + 0) + (1 + 0)

	// guest one is only interested
	e.Guests[0].Decline()
	e.Guests[1].Confirm()
	e.Guests[4].Confirm()
	e.Guests[4].BringSomeone(2)

	// get guest confirmation information
	confirmedGuests = e.GetConfirmedGuests()
	interestedGuests = e.GetInterestedGuests()
	declinedGuests = e.GetDeclinedGuests()
	numberOfConfirmedGuests = e.GetNumberOfConfirmedGuests()
	numberOfInterestedGuests = e.GetNumberOfInterestedGuests()
	numberOfDeclinedGuests = e.GetNumberOfDeclinedGuests()

	// check if guest confirmation information is correct
	assert.True(t, reflect.DeepEqual(confirmedGuests, []*guest.Guest{&guest2, &guest5}))
	assert.True(t, reflect.DeepEqual(interestedGuests, []*guest.Guest{&guest3}))
	assert.True(t, reflect.DeepEqual(declinedGuests, []*guest.Guest{&guest1, &guest4}))

	assert.Equal(t, numberOfConfirmedGuests, 6)  // (1 + 2) + (1 + 2)
	assert.Equal(t, numberOfInterestedGuests, 1) // (1 + 0)
	assert.Equal(t, numberOfDeclinedGuests, 2)   // (1 + 0) + (1 + 0)
}
