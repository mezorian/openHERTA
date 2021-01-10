package guest

type ConfimationStatus int

const (
  None ConfimationStatus = iota
  Confirmed ConfimationStatus = iota
  Interested ConfimationStatus = iota
  Declined ConfimationStatus = iota
)


type Guest struct {
  FirstName string
  LastName string
  ConfirmationStatus ConfimationStatus
  OtherPeopleThisGuestWillBring int
}

// change the confirmation status
func (g *Guest) Confirm() {
  g.ConfirmationStatus = Confirmed
}

func (g *Guest) Interest() {
  g.ConfirmationStatus = Interested
}

func (g* Guest) Decline() {
  g.ConfirmationStatus = Declined
  g.OtherPeopleThisGuestWillBring = 0
}


// bring other people

// bringing someone is only allowed for
// - guests that have currently NOT declined
// - guests that want to bring positive numbers of other people
//   --> negative numbers will be set to 0
// - guests that want to bring integer numbers of other people
//   --> non-integer numbers will be rounded
func (g* Guest) BringSomeone(numberOfPeopleToBring int) {
  if (g.ConfirmationStatus == Declined) {
    g.OtherPeopleThisGuestWillBring = 0
  } else {
    if (numberOfPeopleToBring < 0) {
      g.OtherPeopleThisGuestWillBring = 0
    } else {
      g.OtherPeopleThisGuestWillBring = numberOfPeopleToBring
    }
  }
}
