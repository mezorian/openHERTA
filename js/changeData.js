/* --- EVENT LISTENERS ---  */

// TODO : currently the log out button is temp. used as a refresh button
$("body").on("focusout", "#content_guest_details_first_name", function() {
  console.log("got ya");
  form = document.getElementById("guest_details_name_form");
  formData = new FormData(form);
  console.log(form)
  fetch('/updateGuestDetailsName',{
    method: "post",
    body: formData,
  })
  .then(response => response.text())
  .then((response) => {
      console.log(response)
  })
  .catch(err => console.log(err))
});
