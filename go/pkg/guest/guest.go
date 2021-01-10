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
func (g *Guest) confirm() {
  g.ConfirmationStatus = Confirmed
}

func (g *Guest) interest() {
  g.ConfirmationStatus = Interested
}

func (g* Guest) decline() {
  g.ConfirmationStatus = Declined
  g.OtherPeopleThisGuestWillBring = 0
}
