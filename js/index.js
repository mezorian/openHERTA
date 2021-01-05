const CONFIRMATION_STATUS = Object.freeze({
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
});

$(".delete_button").on("click", function() {
  $(this).closest("li").remove();
});

function processMediaQueries(mediaQuery) {
  // If media query matches (small screen)
  if (mediaQuery.matches) {
    // move content_discussion div into content_container_2
    $("#content_discussion").detach().appendTo($("#content_container_2"))

    // If media query does not match (wide screen)
  } else {
    // and if content_discussion is currently a child of content_container_2
    if ($("#content_discussion").parent().length > 0) {
      if ($("#content_discussion").parent().attr("id") === "content_container_2") {
        // move content_discussion div back into content_container_1
        $("#content_discussion").detach().appendTo($("#content_container_1"))
      }
    }
  }
}

var mediaQuery = window.matchMedia("(max-width: 1100px)")
processMediaQueries(mediaQuery) // Call listener function at run time
mediaQuery.addListener(processMediaQueries) // Attach listener function on state changes
