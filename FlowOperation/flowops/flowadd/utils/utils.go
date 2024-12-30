package utils

import (
	"github.com/netrack/openflow/ofp"
	"log"
	"net"
)

func Uint32ToXMValue(value uint32) ofp.XMValue {
	buf := make([]byte, 4)
	buf[0] = byte(value >> 24)
	buf[1] = byte(value >> 16)
	buf[2] = byte(value >> 8)
	buf[3] = byte(value)
	return ofp.XMValue(buf)
}

func Uint16ToXMValue(value uint16) ofp.XMValue {
	buf := make([]byte, 2)
	buf[0] = byte(value >> 8)
	buf[1] = byte(value)
	return ofp.XMValue(buf)
}
func Uint8ToXMValue(value uint8) ofp.XMValue {
	buf := make([]byte, 1)
	buf[0] = byte(value)
	return ofp.XMValue(buf)
}

func MacStringTo6Byte(s string) [6]byte {
	hw, err := net.ParseMAC(s)
	if err != nil {
		log.Println("ERROR: parsing MAC:", err)
	}
	var arr [6]byte
	copy(arr[:], hw)
	return arr
}
