package enums

import (
	"encoding/json"
)

type EncryptMethod int

const (
	NONE EncryptMethod = iota
	AES_128
	AES_128_ECB
	SAMPLE_AES
	SAMPLE_AES_CTR
	CENC
	CHACHA20
	UNKNOWN
)

var EncryptMethodStrings = map[EncryptMethod]string{
	NONE:           "NONE",
	AES_128:        "AES-128",
	AES_128_ECB:    "AES-128_ECB",
	SAMPLE_AES:     "SAMPLE_AES",
	SAMPLE_AES_CTR: "SAMPLE_AES_CTR",
	CENC:           "CENC",
	CHACHA20:       "CHACHA20",
	UNKNOWN:        "UNKNOWN",
}

// String method to convert EncryptMethod to a string
func (e EncryptMethod) String() string {
	if str, exists := EncryptMethodStrings[e]; exists {
		return str
	}
	return "UNKNOWN"
}

var MethodStringMap = map[string]EncryptMethod{
	"NONE":           NONE,
	"AES_128":        AES_128,
	"AES_128_ECB":    AES_128_ECB,
	"SAMPLE_AES":     SAMPLE_AES,
	"SAMPLE_AES_CTR": SAMPLE_AES_CTR,
	"CENC":           CENC,
	"CHACHA20":       CHACHA20,
}

// MarshalJSON marshals the EncryptMethod to JSON
func (e EncryptMethod) MarshalJSON() ([]byte, error) {
	return json.Marshal(e.String())
}

func (e *EncryptMethod) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	// Find the EncryptMethod that matches the string
	for method, str := range EncryptMethodStrings {
		if str == s {
			*e = method
			return nil
		}
	}

	*e = UNKNOWN
	return nil
}
