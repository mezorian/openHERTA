/* --- EVENT LISTENERS ---  */

// TODO : currently the log out button is temp. used as a refresh button
$("body").on("click", "#log_out_button", function() {
  reloadAll()
});


/* --- RELOAD FUNCTIONS --- */

function reloadHeader() {
  reloadHTMLDiv("header", "/getHeader");
};

function reloadMap() {
  reloadHTMLDiv("content_map", "/getContentMap");
};

function reloadHeaderDetails() {
  reloadHTMLDiv("content_header_details", "/getContentHeaderDetails");
};

function reloadGuest() {
  reloadHTMLDiv("content_guest", "/getContentGuest");
};

function reloadBringSomething() {
  reloadHTMLDiv("content_bring_something", "/getContentBringSomething");
};

function reloadGuests() {
  reloadHTMLDiv("content_guests", "/getGuests")
};

function reloadDiscussion() {
  reloadHTMLDiv("content_discussion", "/getContentDiscussion");
};

function reloadAll() {
  // header
  reloadHeader();
  // right container
  reloadMap();
  reloadHeaderDetails();
  reloadGuest();
  reloadBringSomething();
  // left container
  reloadGuests();
  reloadDiscussion();
};

// API calls

function reloadHTMLDiv(DivIDToReplace, APIEndpoint) {
  fetch(APIEndpoint)
  .then(response => response.text())
  .then((response) => {
      //console.log(response)
      $("#"+DivIDToReplace).replaceWith(response);
  })
  .catch(err => console.log(err))
}
