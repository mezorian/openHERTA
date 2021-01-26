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

/* --- INITIALIZE WEBSITE WITH TESTING DATA ---  */
// confirmed guests

func (e *Event) CreateDummyData() {

	e.Name = "Lorem ipsum Grillabend"
	e.InfoTitle = "Lorem ipsum Grillabend"
	e.Date = "Samstag, 20. Juni 2021"
	e.Time = "18:00 Uhr"
	e.AddressString = "Hinten im Garten"
	e.AddressGPSLat = 52.297652
	e.AddressGPSLong = -10.0383357
	// Todo
	e.Description = ""

	dummyGuest := new(guest.Guest)
	dummyGuest.FirstName = "John"
	dummyGuest.LastName = "Stone"
	dummyGuest.Confirm()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ponnappa"
	dummyGuest.LastName = "Priya"
	dummyGuest.Confirm()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Mia"
	dummyGuest.LastName = "Wong"
	dummyGuest.Confirm()
	dummyGuest.BringSomeone(2)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Peter"
	dummyGuest.LastName = "Stanbridge"
	dummyGuest.Confirm()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Natalie"
	dummyGuest.LastName = "Lee-Walsh"
	dummyGuest.Confirm()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ang"
	dummyGuest.LastName = "Li"
	dummyGuest.BringSomeone(2)
	dummyGuest.Confirm()
	e.AddGuest(dummyGuest)

	// interested guests

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Nguta"
	dummyGuest.LastName = "Ithya"
	dummyGuest.Interest()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Tamzyn"
	dummyGuest.LastName = "French"
	dummyGuest.Interest()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Salome"
	dummyGuest.LastName = "Simoes"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(1)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Trevor"
	dummyGuest.LastName = "Virtue"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(3)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Tarryn"
	dummyGuest.LastName = "Campbell-Gillies"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Eugenia"
	dummyGuest.LastName = "Anders"
	dummyGuest.Interest()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Andrew"
	dummyGuest.LastName = "Kazantzis"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Verona"
	dummyGuest.LastName = "Blair"
	dummyGuest.Interest()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Jane"
	dummyGuest.LastName = "Meldrum"
	dummyGuest.Interest()
	dummyGuest.BringSomeone(2)
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Maureen"
	dummyGuest.LastName = "M. Smith"
	dummyGuest.Interest()
	e.AddGuest(dummyGuest)

	// declined guests

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Desiree"
	dummyGuest.LastName = "Burch"
	dummyGuest.Decline()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Daly"
	dummyGuest.LastName = "Harry"
	dummyGuest.Decline()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Hayman"
	dummyGuest.LastName = "Andrews"
	dummyGuest.Decline()
	e.AddGuest(dummyGuest)

	dummyGuest = new(guest.Guest)
	dummyGuest.FirstName = "Ruveni"
	dummyGuest.LastName = "Ellawala"
	dummyGuest.Decline()
	e.AddGuest(dummyGuest)
}
