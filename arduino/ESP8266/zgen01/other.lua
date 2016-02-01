--set wifi host name
if (wifi.sta.sethostname("NodeMCU") == true) then
    print("hostname was successfully changed")
else
    print("hostname was not changed")
end

--setup ap
-- wifi.ap.config({ssid="nodemcu",pwd="1234567890"})
-- wifi.ap.dhcp.config({start="192.168.1.100"})
-- wifi.ap.dhcp.start()
-- cfg =
-- {
--     ip="192.168.1.1",
--     netmask="255.255.255.0",
--     gateway="192.168.1.1"
-- }
-- wifi.ap.setip(cfg)
