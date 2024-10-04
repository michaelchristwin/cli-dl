package entity

import (
	"strings"

	"github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/enums"
)

type EncryptInfo struct {
	Method enums.EncryptMethod
	Key    []byte
	IV     []byte
}

func NewEncryptInfo() *EncryptInfo {
	return &EncryptInfo{
		Method: enums.NONE,
	}
}

func NewEncryptInfoWithMethod(method string) *EncryptInfo {
	return &EncryptInfo{
		Method: ParseMethod(method),
	}
}
func ParseMethod(method string) enums.EncryptMethod {
	if method != "" {
		// Replace "-" with "_" and try to match the enum
		method = strings.ReplaceAll(method, "-", "_")
		method = strings.ToUpper(method)
		if m, exists := enums.MethodStringMap[method]; exists {
			return m
		}
	}
	return enums.UNKNOWN // Default to UNKNOWN if no match
}
