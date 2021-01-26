/* --- EVENT LISTENERS ---  */

// TODO : currently the log out button is temp. used as a refresh button
$("body").on("click", "#log_out_button", function() {
  reloadData()
});


/* --- RELOAD FUNCTIONS --- */

function reloadData() {
  // header
  reloadHTMLDiv("header", "/getHeader")
  // right container
  reloadHTMLDiv("content_map", "/getContentMap");
  reloadHTMLDiv("content_header_details", "/getContentHeaderDetails")
  reloadHTMLDiv("content_guest", "/getContentGuest")
  reloadHTMLDiv("content_bring_something", "/getContentBringSomething")
  // left container
  reloadHTMLDiv("content_guests", "/getGuests")
  reloadHTMLDiv("content_discussion", "/getContentDiscussion")
};

// API calls

function reloadHTMLDiv(DivIDToReplace, APIEndpoint) {
  fetch(APIEndpoint)
  .then(response => response.text())
  .then((response) => {
      console.log(response)
      $("#"+DivIDToReplace).replaceWith(response);
  })
  .catch(err => console.log(err))
}
