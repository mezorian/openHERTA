//const {EventTarget} = require('events');

//const event = new Event('build');

// Listen for the event.
//function handler(event) {console.log("caught")}

class MyTarget extends EventTarget {}
const target = new MyTarget();
target.addEventListener('foo', handler, { capture: true });  // first

//addEventListener('build', function (e) {console.log("caught") }, false);

// Dispatch the event.
//dispatchEvent(event);
