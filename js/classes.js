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
    this._otherPeopleThisGuestWillBring = 0;
  };

  // bring other people

  // bringing someone is only allowed for
  // - guests that have currently NOT declined
  // - guests that want to bring positive numbers of other people
  //   --> negative numbers will be set to 0
  // - guests that want to bring integer numbers of other people
  //   --> non-integer numbers will be rounded
  bringSomeone(numberOfPeopleToBring) {
    if (this.confirmationStatus == CONFIRMATION_STATUS.declined) {
      this._otherPeopleThisGuestWillBring = 0;
    } else {
      if (numberOfPeopleToBring < 0) {
        this._otherPeopleThisGuestWillBring = 0;
      } else {
        this._otherPeopleThisGuestWillBring = Math.round(numberOfPeopleToBring);
      }
    }


  }
}


class OpenHERTAEvent {
  // constructor
  constructor() {
    this._name = "";
    // TODO: more properties
    this._guests = [];

    //this._changeConfirmationStatusEvent = new Event('event');
    //elem.addEventListener('changeConfirmationStatus', function (e) { console.log("caught event") }, false);
  };

  // setters
  set name(newValue) {
    this._name = newValue;
  };

  set guests(newValue) {
    this._guests = newValue;
  }

  // getters
  get name() {
    return this._name;
  };

  get guests() {
    return this._guests;
  }

  addGuest(newGuest) {
    this._guests.push(newGuest)
  };

  // get guest confirmation information
  getConfirmationInformation(confirmationStatus) {
    var resultList = [];

    //elem.dispatchEvent(this._changeConfirmationStatusEvent);

    return resultList;
  }

  getConfirmedGuests() {
    return this.getConfirmationInformation(CONFIRMATION_STATUS.confirmed);
  };

  getInterestedGuests() {
    return this.getConfirmationInformation(CONFIRMATION_STATUS.interested);
  };

  getDeclinedGuests() {
    return this.getConfirmationInformation(CONFIRMATION_STATUS.declined);
  };

  getNumberOfConfirmedGuests() {
    return this.getConfirmedGuests().length;
  };

  getNumberOfInterestedGuests() {
    return this.getInterestedGuests().length;
  };

  getNumberOfDeclinedGuests() {
    return this.getDeclinedGuests().length;
  };


}

/* --- EXPORT DEFINITIONS ---  */

// export classes and types for being able to require them in unit-testing
exports.Guest = Guest
exports.OpenHERTAEvent = OpenHERTAEvent
exports.CONFIRMATION_STATUS = CONFIRMATION_STATUS
