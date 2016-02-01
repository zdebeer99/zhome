--wifi connection setup is located in user.lua
--user.lua example, this example connects using a static ip. replace the username and password with your ap values.
-- wifi.setmode(wifi.STATION)
-- wifi.sta.config("username","password",1)

--connect to AP
print("Welcome to ZHOME")

dofile("user.lua")
dofile("zboard.lua")

-- tmr.alarm(0,10000, 1, function()
--    if wifi.sta.status()==5 then
--       print('connected ip: ',wifi.sta.getip())
--       --gpio.write(3,gpio.LOW)
--       tmr.stop(0)
--     else
--       print("connecting to AP...",wifi.sta.status())
--       --gpio.write(3,gpio.HIGH)
--    end
-- end)
-- wifi.sta.eventMonReg(wifi.STA_IDLE, function() print("STATION_IDLE") end)
-- wifi.sta.eventMonReg(wifi.STA_CONNECTING, function() print("STATION_CONNECTING") end)
-- wifi.sta.eventMonReg(wifi.STA_WRONGPWD, function() print("STATION_WRONG_PASSWORD") end)
-- wifi.sta.eventMonReg(wifi.STA_APNOTFOUND, function() print("STATION_NO_AP_FOUND") end)
-- wifi.sta.eventMonReg(wifi.STA_FAIL, function() print("STATION_CONNECT_FAIL") end)
-- wifi.sta.eventMonReg(wifi.STA_GOTIP, function() print("STATION_GOT_IP. IP", wifi.sta.getip()) end)
-- wifi.sta.eventMonStart(10*1000)
