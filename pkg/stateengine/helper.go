package stateengine

import (
	"strconv"
)

func StrToByte(num string) (byte, error) {
	pin, err := strconv.ParseUint(num, 10, 8)
	if err != nil {
		return 0, err
	}
	return byte(pin), nil
}

func StrToInt(num string) (int, error) {
	pin, err := strconv.ParseInt(num, 10, strconv.IntSize)
	if err != nil {
		return 0, err
	}
	return int(pin), nil
}
