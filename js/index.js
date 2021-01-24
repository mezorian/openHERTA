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

// API calls

function getGuests() {
  fetch("/getGuests")
  .then(response => response.text())
  .then((response) => {
      console.log(response)
      $('#content_guests').replaceWith(response);
  })
  .catch(err => console.log(err))
}


// DOM events

$("body").on("click", "#content_guest_details_confirm_button", function() {
  getGuests();
});

$("body").on("click", "#content_guests_summary_confirmed_button", function() {
    $("#content_guests_summary_confirmed_button").addClass("gray_button_clicked");
    $("#content_guests_summary_interested_button").removeClass("gray_button_clicked");
    $("#content_guests_summary_declined_button").removeClass("gray_button_clicked");

    $("#content_guests_list_confirmed").removeClass("hidden");
    $("#content_guests_list_interested").addClass("hidden");
    $("#content_guests_list_declined").addClass("hidden");
    console.log("confirmed");
});

$("body").on("click", "#content_guests_summary_interested_button", function() {
    $("#content_guests_summary_confirmed_button").removeClass("gray_button_clicked");
    $("#content_guests_summary_interested_button").addClass("gray_button_clicked");
    $("#content_guests_summary_declined_button").removeClass("gray_button_clicked");

    $("#content_guests_list_confirmed").addClass("hidden");
    $("#content_guests_list_interested").removeClass("hidden");
    $("#content_guests_list_declined").addClass("hidden");
    console.log("interested");
});

$("body").on("click", "#content_guests_summary_declined_button", function() {
    $("#content_guests_summary_confirmed_button").removeClass("gray_button_clicked");
    $("#content_guests_summary_interested_button").removeClass("gray_button_clicked");
    $("#content_guests_summary_declined_button").addClass("gray_button_clicked");

    $("#content_guests_list_confirmed").addClass("hidden");
    $("#content_guests_list_interested").addClass("hidden");
    $("#content_guests_list_declined").removeClass("hidden");
    console.log("declined");
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
