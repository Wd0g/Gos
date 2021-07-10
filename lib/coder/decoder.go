package coder

import (
	"net/url"

	"github.com/Wd0g/GoShell/lib/common"
)

func base64Decoder(val url.Values) url.Values {
	for k, _ := range val {
		newVal := common.Base64Decode(val.Get(k))
		val.Set(k, newVal)
	}
	return val
}

func base64Encoder(src string) string {
	return common.Base64Encode(src)
}

func plainDecoder(val url.Values) url.Values {
	return val
}

func plainEncoder(src string) string {
	return src
}

func NewDecoder(name string) func(val url.Values) url.Values {
	switch name {
	case "plaing":
		return plainDecoder
	case "base64":
		return base64Decoder
	default:
		return plainDecoder
	}
}

func NewEncoder(name string) func(string) string {
	switch name {
	case "plaing":
		return plainEncoder
	case "base64":
		return base64Encoder
	default:
		return plainEncoder
	}

}
