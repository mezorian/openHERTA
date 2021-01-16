/* GLOBAL VARS OF WEB APPLICATION */
ohEvent = new OpenHERTAEvent();

/* --- INITIALIZE WEBSITE WITH TESTING DATA ---  */
// confirmed guests

guest = new Guest();
guest.firstName = "John";
guest.lastName = "Stone";
guest.confirm();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Ponnappa";
guest.lastName = "Priya";
guest.confirm();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Mia";
guest.lastName = "Wong";
guest.confirm();
guest.bringSomeone(2);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Peter";
guest.lastName = "Stanbridge";
guest.confirm();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Natalie";
guest.lastName = "Lee-Walsh";
guest.confirm();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Ang";
guest.lastName = "Li";
guest.bringSomeone(2);
guest.confirm();
ohEvent.addGuest(guest);

// interested guests

guest = new Guest();
guest.firstName = "Nguta";
guest.lastName = "Ithya";
guest.interest();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Tamzyn";
guest.lastName = "French";
guest.interest();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Salome";
guest.lastName = "Simoes";
guest.interest();
guest.bringSomeone(1);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Trevor";
guest.lastName = "Virtue";
guest.interest();
guest.bringSomeone(3);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Tarryn";
guest.lastName = "Campbell-Gillies";
guest.interest();
guest.bringSomeone(2);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Eugenia";
guest.lastName = "Anders";
guest.interest();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Andrew";
guest.lastName = "Kazantzis";
guest.interest();
guest.bringSomeone(2);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Verona";
guest.lastName = "Blair";
guest.interest();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Jane";
guest.lastName = "Meldrum";
guest.interest();
guest.bringSomeone(2);
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Maureen";
guest.lastName = "M. Smith";
guest.interest();
ohEvent.addGuest(guest);

// declined guests

guest = new Guest();
guest.firstName = "Desiree";
guest.lastName = "Burch";
guest.decline();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Daly";
guest.lastName = "Harry";
guest.decline();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Hayman";
guest.lastName = "Andrews";
guest.decline();
ohEvent.addGuest(guest);

guest = new Guest();
guest.firstName = "Ruveni";
guest.lastName = "Ellawala";
guest.decline();
ohEvent.addGuest(guest);



/* --- EVENT LISTENERS ---  */


// CUSTOM events

function reloadData() {
  reloadGuests();
};

function reloadGuests() {
  confirmedGuests = ohEvent.getConfirmedGuests();
  interestedGuests = ohEvent.getInterestedGuests();
  declinedGuests = ohEvent.getDeclinedGuests();
  numberOfConfirmedGuests = ohEvent.getNumberOfConfirmedGuests();
  numberOfInterestedGuests = ohEvent.getNumberOfInterestedGuests();
  numberOfDeclinedGuests = ohEvent.getNumberOfDeclinedGuests();
  $("#content_guests_summary_confirmed_button").html(numberOfConfirmedGuests + "<br>Zusagen");
}

// DOM events

$("button").on("click", function() {
  alert("Hello World");
  reloadData();
});

$(".delete_button").on("click", function() {
  $(this).closest("li").remove();
});



/* --- MEDIA QUERIES ---  */

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
