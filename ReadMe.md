# zhome

MIT

This is a work in progress.

zhome is a home automation service that allows you to turn lights on and off using a web interface. Currently the service is running on my raspberry pi and supports the following devices.

* Arduino Firmata (not tested after latest changes)

* Arduino zIOBoard  
  Custom communication protocol. see './arduino/zIOBoard/' for details.

* Qwikswitch devices.  
  http://www.qwikswitch.co.za/


## Project Overview

### Features

* Basic Device Control from a web app.
* Scheduler
* [Planned] Log Sensor input to database.
* [Planned] Log All io comms to database and calculated electricity usage.

zhome uses go, coffeescript and jade

### Folders

* Arduino - Contains the .ino files for the arduino board.
* client - Contains the html and javascript files for the front end.
* pkg - contains the go source files for the back-end server.
* scripts - some helper scripts for compiling and uploading to a raspberry pi. For development I use gin to reload the project while developing.


## Getting Started

1. Build the client code, see './client/ReadMe.md'
2. Build the service. 'go build ./'

or if you have 'gin' installed simply run 'gin'

**Build for Raspberry Pi**

Change the ip in './scripts/build_pi.sh' and './scripts/update_pi_config.sh' scripts to build for the raspberry and copy the files onto the raspberry pi.

**Installing as a service**

copy './scripts/zhome.service' to '/etc/systemd/system'

```bash
sudo cp ./scripts/zhome.service /etc/systemd/system
sudo systemctl enable zhome
sudo service zhome start
```
