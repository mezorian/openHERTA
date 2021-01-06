const {CONFIRMATION_STATUS} = require("../js/classes.js");
const {Guest} = require("../js/classes.js");


describe("Test Guest for correct attribute setting", function() {

  var guest;

  beforeEach(function() {
    guest = new Guest();
  });

  it("Guest has empty inialization variables", function() {
    expect(guest.firstName).toBe("");
    expect(guest.lastName).toBe("");
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.none);
    expect(guest.otherPeopleThisGuestWillBring).toBe(0);
  });

  it("Getters return the values which were set by setters", function() {
    guest.firstName = "Max";
    guest.lastName = "Mustermann";
    guest.confirmationStatus = CONFIRMATION_STATUS.confirmed;
    guest.otherPeopleThisGuestWillBring = 4;

    expect(guest.firstName).toBe("Max");
    expect(guest.lastName).toBe("Mustermann");
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.confirmed);
    expect(guest.otherPeopleThisGuestWillBring).toBe(4);
  });

  it("confirm() leads to a confirmationStatus confirmed", function() {
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.none);
    guest.confirm();

    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.confirmed);
  });

  it("interest() leads to a confirmationStatus confirmed", function() {
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.none);
    guest.interest();

    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.interested);
  });

  it("decline() leads to a confirmationStatus confirmed", function() {
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.none);
    guest.decline();

    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.declined);
  });

  it("confirmationStatus can be changed several times", function() {
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.none);
    guest.decline();
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.declined);
    guest.interest();
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.interested);
    guest.decline();
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.declined);
    guest.confirm();
    expect(guest.confirmationStatus).toBe(CONFIRMATION_STATUS.confirmed);
  });

});

describe("Test if guest can successfully bring someone", function() {

  var guest;

  beforeEach(function() {
    guest = new Guest();
    expect(guest.otherPeopleThisGuestWillBring).toBe(0);
  });

  it("otherPeopleThisGuestWillBring is set correctly after bringSomeone()", function() {
    guest.bringSomeone(3);

    expect(guest.otherPeopleThisGuestWillBring).toBe(3);
  });

  it("check if otherPeopleThisGuestWillBring can be set back to 0", function() {
    guest.bringSomeone(3);
    expect(guest.otherPeopleThisGuestWillBring).toBe(3);
    guest.bringSomeone(0);
    expect(guest.otherPeopleThisGuestWillBring).toBe(0);
  });

  it("otherPeopleThisGuestWillBring can be changed several times", function() {
    for (var i=0; i<=100; i++) {
      min = 0;
      max = 50;
      numberOfPeopleToBring = Math.round(Math.random() * (max - min) + min);
      guest.bringSomeone(numberOfPeopleToBring);
      expect(guest.otherPeopleThisGuestWillBring).toBe(numberOfPeopleToBring);
    }
  });

  it("otherPeopleThisGuestWillBring cannot be not-integer", function() {
    for (var i=0; i<=100; i++) {
      min = 0;
      max = 50;
      numberOfPeopleToBring = Math.random() * (max - min) + min;
      guest.bringSomeone(numberOfPeopleToBring);
      expect(guest.otherPeopleThisGuestWillBring).toBe(Math.round(numberOfPeopleToBring));
    }
  });

  it("otherPeopleThisGuestWillBring cannot be negative", function() {
    for (var i=0; i<=100; i++) {
      min = -50;
      max = -1;
      numberOfPeopleToBring = Math.round(Math.random() * (max - min) + min);
      guest.bringSomeone(numberOfPeopleToBring);
      expect(guest.otherPeopleThisGuestWillBring).toBe(0);
    }
  });

  // it("otherPeopleThisGuestWillBring cannot be positive if guest declined", function() {
  //   // bringing someone is okay if you confirm or are interested
  //   guest.confirm();
  //   guest.bringSomeone(3);
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(3);
  //   guest.interest();
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(3);
  //   guest.bringSomeone(5);
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(5);
  //
  //   // but if you decline, you cannot bring someone
  //   guest.decline();
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(0);
  //   guest.bringSomeone(5);
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(0);
  //
  //   // if you after some time decide to confirm again, you can again bring someone
  //   guest.confirm();
  //   guest.bringSomeone(1);
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(1);
  //   guest.interest();
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(1);
  //   guest.bringSomeone(7);
  //   expect(guest.otherPeopleThisGuestWillBring).toBe(7);
  // });

});
