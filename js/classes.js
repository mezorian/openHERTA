/* --- TYPE DEFINITIONS ---  */

const CONFIRMATION_STATUS = Object.freeze({
  none: "none",
  confirmed: "confirmed",
  interested: "interested",
  declined: "declined"
});


/* --- CLASS DEFINITIONS ---  */

class Guest {
  // constructor
  constructor() {
    this._firstName = "";
    this._lastName = "";
    this._confirmationStatus = CONFIRMATION_STATUS.none;
    this._otherPeopleThisGuestWillBring = 0;
  };

  // setters
  set firstName(newValue) {
    this._firstName = newValue;
  };

  set lastName(newValue) {
    this._lastName = newValue;
  };

  set confirmationStatus(newValue) {
    this._confirmationStatus = newValue;
  };

  set otherPeopleThisGuestWillBring(newValue) {
    this._otherPeopleThisGuestWillBring = newValue;
  };

  // getters
  get firstName() {
    return this._firstName;
  };

  get lastName() {
    return this._lastName;
  };

  get confirmationStatus() {
    return this._confirmationStatus;
  };

  get otherPeopleThisGuestWillBring() {
    return this._otherPeopleThisGuestWillBring;
  };

  // change the confirmation status
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



exports.Guest = Guest
exports.CONFIRMATION_STATUS = CONFIRMATION_STATUS
