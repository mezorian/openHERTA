const CONFIRMATION_STATUS = Object.freeze( {
  none: "none",
  confirmed: "confirmed",
  interested: "interested",
  declined: "declined"
});

class Guest {
  constructor() {
    this._firstName = "";
    this._lastName = "";
    this._confirmationStatus = CONFIRMATION_STATUS.none;
  };

  set firstName(newValue) {
    this._firstName = newValue;
  };

  set lastName(newValue) {
    this._lastName = newValue;
  };

  set confirmationStatus(newValue) {
    this._confirmationStatus = confirmationStatus;
  };

  confirm() {
    this._confirmationStatus = CONFIRMATION_STATUS.confirmed;
  };
  interest() {
    this._confirmationStatus = CONFIRMATION_STATUS.interested;
  };
  decline() {
    this._confirmationStatus = CONFIRMATION_STATUS.declined;
  };

}

class Event {
  constructor() {
    this._name = "";
    // TODO: more properties
    this._guests = [];
  };
  addGuest(newGuest) {
    this._guests.push(newGuest)
  };
}

var event = new Event();
var guest = new Guest();
guest.firstName = "Max";
guest.lastName = "Mustermann";
guest.confirm();
event.addGuest(guest);


// define event listeners


$("button").on("click", function() {
  alert("hello world");
} );
