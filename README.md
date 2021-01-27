# openHERTA
```


                        _   _  ___________ _____ ___  
                       | | | ||  ___| ___ \_   _/ _ \
  ___  _ __   ___ _ __ | |_| || |__ | |_/ / | |/ /_\ \
 / _ \| '_ \ / _ \ '_ \|  _  ||  __||    /  | ||  _  |
| (_) | |_) |  __/ | | | | | || |___| |\ \  | || | | |
 \___/| .__/ \___|_| |_\_| |_/\____/\_| \_| \_/\_| |_/
      | |                       \
      |_|                        \        __________________________________________                                 
                                  \      /                                          \
                                    ----|   "Das Leben ist hart aber ich bin Herta"  |
                                         \__________________________________________/
```
[![Build Status](https://travis-ci.com/mezorian/openHERTA.svg?branch=development)](https://travis-ci.com/mezorian/openHERTA) [![Coverage Status](https://coveralls.io/repos/github/mezorian/openHERTA/badge.svg?branch=development)](https://coveralls.io/github/mezorian/openHERTA?branch=development)
![GitHub milestone](https://img.shields.io/github/milestones/progress-percent/mezorian/openHerta/1) ![GitHub milestone](https://img.shields.io/github/milestones/progress/mezorian/openHerta/1)
[![Go Report Card](https://goreportcard.com/badge/github.com/mezorian/openHERTA)](https://goreportcard.com/report/github.com/mezorian/openHERTA)

open**HERTA** is an **open**source **h**omemade **e**vent o**r**ganiza**t**ion **a**pplication which was designed to have a simple self-hostable tool to get around / smash Facebook-Events.

The name HERTA is **for sure not** inspired by the character **Herta Müller** aka "Amazonenkönigin" of the **"Känguru"**-novels by **Marc-Uwe Kling**...

Regardless of these interesting and absolutely random parallels in naming, we are pretty sure that events organized with openHERTA will be at least as cool and crazy as all partys in Hertas "pub" in her apartment,

--------------------

`"Incidentally, I am of the opinion that Facebook must be smashed!"` - Note of the kangaroo

# Status of the project
The project is currently still under development / not yet usable in a productive environment.
For more information please also refer to the below progress badges and the roadmap.
The first non-testing release will be v1.0.0 .

![GitHub milestone](https://img.shields.io/github/milestones/progress-percent/mezorian/openHerta/1) ![GitHub milestone](https://img.shields.io/github/milestones/progress/mezorian/openHerta/1)

## Roadmap

### v0.1.0 First front-end design finished
- [x] create first html / css design
- [x] create login.html
- [x] create event.html
- [x] finish first design

### v0.2.0 first implementation of JS Front-End and Go Backend
- [x] implement unit-tests for go structs
- [x] add travis support (build / unit-testing)
- [x] add coveralls support (coverage)
- [x] finish go webserver implementation which serves static files
- [ ] finish go webserver API endpoints which handle data changes in front-end
- [ ] finish go webserver API endpoints which handle reload requests from front-end
- [x] finish java-script implementation for animations in event.html
- [ ] finish java-script implementation to trigger API endpoints when data is changed
- [x] finish java-script implementation to reload data from API endpoints

### v0.3.0 data persistence implementation
- [ ] add implementation to persist and load all data in / from database

### v0.4.0 login organizers implementation
- [ ] add login for event organizers

### v0.5.0 login users implementation
- [ ] add login for event guests

# Usage
Start the web-server
```
cd openHERTA
go run go/cmd/*.go
```

Open the following link to get to the landing page:
http://localhost:8080
