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
    var numberOfOtherPeopleToBring = 0;

    this._guests.forEach((item, i) => {
      if (item.confirmationStatus == confirmationStatus) {
        resultList.push(item);
        numberOfOtherPeopleToBring = numberOfOtherPeopleToBring + item.otherPeopleThisGuestWillBring;
      }
    });

    return [resultList, numberOfOtherPeopleToBring];
  }

  getConfirmedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.confirmed);
    return resultList;
  };

  getInterestedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.interested);
    return resultList;
  };

  getDeclinedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.declined);
    return resultList;
  };

  getNumberOfConfirmedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.confirmed);

    return resultList.length + numberOfPeopleToBring;
  };

  getNumberOfInterestedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.interested);
    return resultList.length + numberOfPeopleToBring;
  };

  getNumberOfDeclinedGuests() {
    const [resultList, numberOfPeopleToBring] = this.getConfirmationInformation(CONFIRMATION_STATUS.declined);
    return resultList.length + numberOfPeopleToBring;
  };


}

/* --- EXPORT DEFINITIONS ---  */

// export classes and types for being able to require them in unit-testing
exports.Guest = Guest
exports.OpenHERTAEvent = OpenHERTAEvent
exports.CONFIRMATION_STATUS = CONFIRMATION_STATUS
