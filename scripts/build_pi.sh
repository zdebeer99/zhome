#!/usr/bin/env bash
rm -r ./release/zhome_pi
mkdir ./release/zhome_pi
env GOOS=linux GOARCH=arm go build -v -o ./release/zhome_pi/zhome_pi ./server/

cp ./server/config_release.yaml ./release/zhome_pi/config.yaml
cp ./zhome.service ./release/zhome_pi/
cp -r ./server/static ./release/zhome_pi/static
ssh pi@10.0.0.120 "sudo service zhome stop"
scp -r ./release/zhome_pi/ pi@10.0.0.120:"/home/pi/"
ssh pi@10.0.0.120 "sudo service zhome start"
