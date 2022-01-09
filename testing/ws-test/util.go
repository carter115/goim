package main

import "gmimo/common/util"

func GetToken() string {
	return util.GetUUID()
}

func GetDeviceId() string {
	return util.GetUUID()
}

func GetTerminal() string {
	return "WEB"
}
func GetVersion() string {
	return "0.0.testing-1"
}
