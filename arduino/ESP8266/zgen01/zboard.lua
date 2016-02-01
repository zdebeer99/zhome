-------------------------------
--helper functions
-------------------------------
function split(text, delimiter)
  local result = { }
  local from  = 1
  local delim_from, delim_to = string.find( text, delimiter, from  )
  while delim_from do
    table.insert( result, string.sub( text, from , delim_from-1 ) )
    from  = delim_to + 1
    delim_from, delim_to = string.find( text, delimiter, from  )
  end
  table.insert( result, string.sub( text, from  ) )
  return result
end

-------------------------------
--input output functions
-------------------------------
pinmodes = {}
pinstates = {}
accesspoints = {}

--get available ap every minute.
tmr.alarm(0,60 * 1000, 1, function()
  wifi.sta.getap(1, function(d) accesspoints = d end)
end)

function setmode(pin, mode)
  local p = tonumber(pin)
  if mode=="input" then
    pinmodes[p] = "input"
    gpio.mode(p, gpio.INPUT)
  else
    pinmodes[p] = "output"
    gpio.mode(p, gpio.OUTPUT)
  end
end

function getmode(pin)
  local p = tonumber(pin)
  return pinmodes[p]
end

function setpin(pin, mode)
  local p = tonumber(pin)
  if mode=="on" then
    pinstates[p] = "on"
    gpio.write(p, gpio.HIGH)
  else
    pinstates[p] = "off"
    gpio.write(p, gpio.LOW)
  end
end

function getpin(pin)
  --local p = tonumber(pin)
  if gpio.read(pin)==gpio.HIGH then
    return "on"
  else
    return "off"
  end
end

-------------------------------
--HTTP Functions
-------------------------------

-- /cmd/pin/value
function route(path,args)
  local data = {status="OK",statusCode=0};
  local cmd = path[2];
  if (cmd=="") then
    data.aplist = accesspoints
    return data
  end
  --from here all commands requires at least 1 argument.
  local arg1 = path[3];
  local arg2 = path[4];
  if (arg1==nil) or (arg1=="") then
    data.status="Command '"..cmd.."' requires at least 1 argument. example: /"..cmd.."/pin"
    data.statusCode = 1
    return data
  end
  if (cmd=="mode") then
    data.pin = arg1
    if arg2~=nil then
      setmode(arg1, arg2)
      return data
    else
      data.value = getmode(arg1)
      return data
    end
  end
  if (cmd=="pin") then
    data.pin = arg1
    if arg2~=nil then
      setpin(arg1, arg2)
      return data
    else
      data.value = getpin(arg1)
      return data
    end
  end
  data.status="Invalid Command"
  data.statusCode = 1
  return data
end

srv=net.createServer(net.TCP)
srv:listen(80,function(conn)
    conn:on("receive", function(client,request)
        local buf = "";
        local _, _, method, path, vars = string.find(request, "([A-Z]+) (.+)?(.+) HTTP");
        if(method == nil)then
            _, _, method, path = string.find(request, "([A-Z]+) (.+) HTTP");
        end
        local args = {}
        if (vars ~= nil)then
            for k, v in string.gmatch(vars, "(%w+)=(%w+)&*") do
                args[k] = v
            end
        end
        print(path)
        local arrpath = split(path,"/")
        data = route(arrpath, args)
        --response
        client:send("HTTP/1.1 200/OK\r\nServer: zHome\r\nContent-Type: application/json\r\nCache-Control: no-cache no-store\r\nExpires: -1\r\n\r\n")
        local resp = cjson.encode(data)
        client:send(resp);
        client:close();
        collectgarbage();
    end)
end)
