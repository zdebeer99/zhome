# zhome

Note this is a work in progress and not designed to be user friendly regarding setting up new equipment etc.

zhome is a home automation service that allows you to turn lights on and off using a web interface. Currently the service is running on my raspberry pi and supports the following devices.

* Arduino Firmata (not tested after latest changes)

* Arduino zIOBoard  
  Custom very simple communication protocol.

* Qwikswitch devices.  
  http://www.qwikswitch.co.za/

## Compiling

You will need to compile the client side first, see 'client/ReadMe.md'

Then you can compile the go project.

Raspberry Pi

You can update the build_pi.sh script and update_pi_config.sh script in the scripts folder to build the code for the raspberry and copy the files onto the raspberry pi.

## Project Overview

** Folders **

* Arduino - Contains the ino files for the arduino board.
* client - Contains the html and javascript files for the front end.
* pkg - contains the go source files for the backend server.
* scripts - some helper scripts for compiling and uploading to a raspberry pi. for development I use gin to continuesly reload the project while developing.


## LICENSE

MIT
