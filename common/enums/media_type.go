package enums

import (
	"encoding/json"
)

type MediaType int

const (
	AUDIO MediaType = iota
	VIDEO
	SUBTITLES
	CLOSED_CAPTIONS
)

var MediaTypeStrings = map[MediaType]string{
	AUDIO:           "AUDIO",
	VIDEO:           "VIDEO",
	SUBTITLES:       "SUBTITLES",
	CLOSED_CAPTIONS: "CLOSED_CAPTIONS",
}

func (m MediaType) String() string {
	if str, exists := MediaTypeStrings[m]; exists {
		return str
	}
	return "Unknown MediaType"
}

// MarshalJSON marshals the MediaType to JSON
func (m MediaType) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.String())
}

func (m *MediaType) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	// Find the MediaType that matches the string
	for mediaType, str := range MediaTypeStrings {
		if str == s {
			*m = mediaType
			return nil
		}
	}
	return nil
}
