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
