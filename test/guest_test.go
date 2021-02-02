package main

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/mezorian/openHERTA/go/pkg/guest"
)

func TestGuestIsEmptyAfterCreation(t *testing.T) {
	var g guest.Guest
	assert.Equal(t, g.FirstName, "", "Error : variable should be empty")
	assert.Equal(t, g.LastName, "", "Error : variable should be empty")
	assert.Equal(t, g.ConfirmationStatus, guest.None, "Error : variable should be None")
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0, "Error : variable should be zero")
}

func TestConfirmLeadsToConfirmationaStatusConfirmed(t *testing.T) {
	var g guest.Guest
	g.Confirm()
	assert.Equal(t, g.ConfirmationStatus, guest.Confirmed)
}

func TestInterestLeadsToConfirmationaStatusInterested(t *testing.T) {
	var g guest.Guest
	g.Interest()
	assert.Equal(t, g.ConfirmationStatus, guest.Interested)
}

func TestDeclineLeadsToConfirmationaStatusDeclined(t *testing.T) {
	var g guest.Guest
	g.Decline()
	assert.Equal(t, g.ConfirmationStatus, guest.Declined)
}

func TestConfirmationStatusCanBeChangedSeveralTimes(t *testing.T) {
	var g guest.Guest
	assert.Equal(t, g.ConfirmationStatus, guest.None)
	g.Decline()
	assert.Equal(t, g.ConfirmationStatus, guest.Declined)
	g.Interest()
	assert.Equal(t, g.ConfirmationStatus, guest.Interested)
	g.Decline()
	assert.Equal(t, g.ConfirmationStatus, guest.Declined)
	g.Confirm()
	assert.Equal(t, g.ConfirmationStatus, guest.Confirmed)
}

func TestIfGuestCanSuccessfullyBringSomeone(t *testing.T) {
	var g guest.Guest
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0)

	// test for :
	//  - otherPeopleThisGuestWillBring is set correctly after bringSomeone()
	//  - check if otherPeopleThisGuestWillBring can be set back to 0
	//  - otherPeopleThisGuestWillBring can be changed several times
	var testValues = []int{
		3, 10, 56, 16, 58,
		8, 43, 90, 42, 53,
		86, 84, 2, 71, 29,
		60, 23, 83, 62, 64,
		31, 6, 99, 90, 17,
		85, 43, 0, 99, 7,
		18, 88, 82, 99, 98,
		56, 97, 46, 8, 80,
		60, 0, 75, 69, 95,
		31, 17, 8, 67, 83,
		50, 38, 30, 1, 59,
		83, 0, 40, 8, 9,
		22, 27, 74, 72, 16,
		23, 15, 59, 1, 85,
		10, 61, 1, 29, 80,
		74, 74, 48, 2, 68,
		21, 30, 79, 73, 31,
		34, 53, 96, 22, 32,
		99, 79, 74, 7, 12,
		8, 89, 66, 3, 34,
	}

	for _, testValue := range testValues {
		g.BringSomeone(testValue)
		assert.Equal(t, g.OtherPeopleThisGuestWillBring, testValue)
	}

}

func TestIfOtherPeopleThisGuestWillBringCannotBeNegative(t *testing.T) {
	var g guest.Guest
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0)

	var testValues = []int{
		-56, -67, -85, -37, -100,
		-76, -95, -46, -77, -17,
		-37, -1, -17, -61, -3,
		-4, -80, -10, -80, -41,
		-87, -98, -19, -38, -22,
		-50, -82, -91, -87, -60,
		-100, -51, -28, -5, -42,
		-30, -28, -39, -18, -13,
		-29, -19, -65, -37, -72,
		-77, -91, -88, -7, -34,
		-100, -13, -87, -35, -52,
		-50, -78, -88, -13, -76,
		-7, -74, -81, -88, -49,
		-42, -44, -75, -95, -5,
		-52, -18, -98, -21, -22,
		-40, -80, -18, -68, -67,
		-76, -61, -30, -8, -39,
		-49, -67, -85, -38, -42,
		-58, -77, -88, -39, -82,
		-57, -23, -51, -64, -5,
	}

	for _, testValue := range testValues {
		g.BringSomeone(testValue)
		assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0)
	}

}

func TestOtherPeopleThisGuestWillBringCannotBePositiveIfGuestDeclined(t *testing.T) {
	var g guest.Guest

	// bringing someone is okay if you confirm or are interested
	g.Confirm()
	g.BringSomeone(3)
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 3)
	g.Interest()
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 3)
	g.BringSomeone(5)
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 5)

	// but if you decline, you cannot bring someone
	g.Decline()
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0)
	g.BringSomeone(5)
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 0)

	// if you after some time decide to confirm again, you can again bring someone
	g.Confirm()
	g.BringSomeone(1)
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 1)
	g.Interest()
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 1)
	g.BringSomeone(7)
	assert.Equal(t, g.OtherPeopleThisGuestWillBring, 7)
}
