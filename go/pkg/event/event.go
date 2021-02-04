package event

import (
	"github.com/mezorian/openHERTA/go/pkg/guest"
)

type Event struct {
	ID             int
	Name           string
	InfoTitle      string
	Date           string
	Time           string
	AddressString  string
	AddressGPSLong float64
	AddressGPSLat  float64
	Description    string
	Guests         []*guest.Guest
}

func (e *Event) AddGuest(newGuest *guest.Guest) {
	e.Guests = append(e.Guests, newGuest)
}

// get guest confirmation information
func (e *Event) GetConfirmationInformation(confirmationStatus guest.ConfimationStatus) ([]*guest.Guest, int) {
	resultList := make([]*guest.Guest, 0)
	numberOfOtherPeopleToBring := 0
	for _, g := range e.Guests {
		if g.ConfirmationStatus == confirmationStatus {
			resultList = append(resultList, g)
			numberOfOtherPeopleToBring += g.OtherPeopleThisGuestWillBring
		}
	}

	return resultList, numberOfOtherPeopleToBring
}

func (e *Event) GetConfirmedGuests() []*guest.Guest {
	resultList, _ := e.GetConfirmationInformation(guest.Confirmed)
	return resultList
}

func (e *Event) GetInterestedGuests() []*guest.Guest {
	resultList, _ := e.GetConfirmationInformation(guest.Interested)
	return resultList
}

func (e *Event) GetDeclinedGuests() []*guest.Guest {
	resultList, _ := e.GetConfirmationInformation(guest.Declined)
	return resultList
}

func (e *Event) GetNumberOfConfirmedGuests() int {
	resultList, numberOfPeopleToBring := e.GetConfirmationInformation(guest.Confirmed)
	return len(resultList) + numberOfPeopleToBring
}

func (e *Event) GetNumberOfInterestedGuests() int {
	resultList, numberOfPeopleToBring := e.GetConfirmationInformation(guest.Interested)
	return len(resultList) + numberOfPeopleToBring
}

func (e *Event) GetNumberOfDeclinedGuests() int {
	resultList, numberOfPeopleToBring := e.GetConfirmationInformation(guest.Declined)
	return len(resultList) + numberOfPeopleToBring
}
