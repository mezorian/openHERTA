//Bring in the events library
const events = require("events")

//create a new event emitter
const emitter = new events.EventEmitter()

//listen for a particular event
emitter.on("itHappened", obj => {
  console.log("I'm responded to the emitted event")
  console.log(obj.message)
})

//emit the event
emitter.emit("itHappened", {
  message:
    "Hello, I'm an argument passed during the event to the event handler",
})
