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
