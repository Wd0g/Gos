package decoder

import (
	"net/url"

	"github.com/Wd0g/GoShell/lib/common"
)

func Base64(val url.Values) url.Values {
	for k, _ := range val {
		newVal := common.Base64Decode(val.Get(k))
		val.Set(k, newVal)
	}
	return val
}

func Plain(val url.Values) url.Values {
	return val
}
