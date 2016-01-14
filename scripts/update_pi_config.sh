#!/usr/bin/env bash

ssh pi@10.0.0.120 "sudo service zhome stop"
scp ../config_release.yaml pi@10.0.0.120:"/home/pi/zhome_pi/config.yaml"
ssh pi@10.0.0.120 "sudo service zhome start"
