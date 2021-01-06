const {
  CONFIRMATION_STATUS
} = require("../js/classes.js");
const {
  Guest
} = require("../js/classes.js");
const {
  OpenHERTAEvent
} = require("../js/classes.js");


describe("Test OpenHERTAEvent for correct attribute setting", function() {

  var event;

  beforeEach(function() {
    event = new OpenHERTAEvent();
  });

  it("Guest has empty inialization variables", function() {
    expect(event.name).toBe("");
    expect(event.guests).toEqual([]);
    expect(event.guests.length).toBe(0);
  });

  it("Getters return the values which were set by setters", function() {

    // prepare guests
    guest1 = new Guest();
    guest1.firstName = "Max";
    guest1.lastName = "Mustermann";
    guest1.confirmationStatus = CONFIRMATION_STATUS.confirmed;
    guest1.otherPeopleThisGuestWillBring = 4;
    guest2 = new Guest();
    guest2.firstName = "Maria";
    guest2.lastName = "Musterfrau";
    guest2.confirmationStatus = CONFIRMATION_STATUS.declined;
    guest2.otherPeopleThisGuestWillBring = 0;
    guestsArray = [guest1, guest2];

    // set attributes of event
    event.name = "Test OpenHERTAEvent";
    event.guests = guestsArray;

    expect(event.name).toBe("Test OpenHERTAEvent");
    expect(event.guests).toEqual(guestsArray);
  });

});

describe("Test OpenHERTAEvent for adding and counting guests", function() {

  var event;

  beforeEach(function() {
    event = new OpenHERTAEvent();
  });

  it("Test that there are no guests at the very beginning", function() {
    // get guest confirmation information
    confirmedGuests = event.getConfirmedGuests();
    interestedGuests = event.getInterestedGuests();
    declinedGuests = event.getDeclinedGuests();
    numberOfConfirmedGuests = event.getNumberOfConfirmedGuests();
    numberOfInterestedGuests = event.getNumberOfInterestedGuests();
    numberOfDeclinedGuests = event.getNumberOfDeclinedGuests();

    // check if guest confirmation information is correct
    expect(confirmedGuests).toEqual([]);
    expect(interestedGuests).toEqual([]);
    expect(declinedGuests).toEqual([]);

    expect(numberOfConfirmedGuests).toEqual(0);
    expect(numberOfInterestedGuests).toEqual(0);
    expect(numberOfDeclinedGuests).toEqual(0);
  });

  it("Test that there are the right guests and correct number of guests after adding some", function() {
    // prepare guests
    guest1 = new Guest();
    guest1.firstName = "Max";
    guest1.lastName = "Mustermann";
    guest1.confirmationStatus = CONFIRMATION_STATUS.confirmed;
    guest1.otherPeopleThisGuestWillBring = 4;
    guest2 = new Guest();
    guest2.firstName = "Mia";
    guest2.lastName = "Wong";
    guest2.confirmationStatus = CONFIRMATION_STATUS.interested;
    guest2.otherPeopleThisGuestWillBring = 2
    guest3 = new Guest();
    guest3.firstName = "Peter";
    guest3.lastName = "Stanbridge";
    guest3.confirmationStatus = CONFIRMATION_STATUS.interested;
    guest3.otherPeopleThisGuestWillBring = 0
    guest4 = new Guest();
    guest4.firstName = "Maria";
    guest4.lastName = "Musterfrau";
    guest4.confirmationStatus = CONFIRMATION_STATUS.declined;
    guest4.otherPeopleThisGuestWillBring = 0;
    guest5 = new Guest();
    guest5.firstName = "John";
    guest5.lastName = "Stone";
    guest5.confirmationStatus = CONFIRMATION_STATUS.declined;
    guest5.otherPeopleThisGuestWillBring = 0;

    // add guests to event
    event.addGuest(guest1);
    event.addGuest(guest2);
    event.addGuest(guest3);
    event.addGuest(guest4);
    event.addGuest(guest5);

    // get guest confirmation information
    confirmedGuests = event.getConfirmedGuests();
    interestedGuests = event.getInterestedGuests();
    declinedGuests = event.getDeclinedGuests();
    numberOfConfirmedGuests = event.getNumberOfConfirmedGuests();
    numberOfInterestedGuests = event.getNumberOfInterestedGuests();
    numberOfDeclinedGuests = event.getNumberOfDeclinedGuests();

    // check if guest confirmation information is correct
    expect(confirmedGuests).toEqual([guest1]);
    expect(interestedGuests).toEqual([guest2, guest3]);
    expect(declinedGuests).toEqual([guest4, guest5]);

    expect(numberOfConfirmedGuests).toEqual(5); // (1 + 4)
    expect(numberOfInterestedGuests).toEqual(4); // (1 + 2) + 1
    expect(numberOfDeclinedGuests).toEqual(2); // (1 + 0) + (1 + 0)
  });
});
