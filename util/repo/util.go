package repo

import (
	"encoding/base64"
)

func NCErr(msg string) *NoCreds {
	return &NoCreds{msg:msg}
}

type NoCreds struct {
	msg string
}

func (n *NoCreds) Error() string {
	return n.msg
}

func BitzToBase64(bits []byte) string {
	return base64.StdEncoding.EncodeToString(bits)
}

func StrToBase64(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}