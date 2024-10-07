package entity

import (
	"time"

	"github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/utils"
)

type SubCue struct {
	StartTime time.Duration
	EndTime   time.Duration
	Payload   string
	Settings  string
}

func (s *SubCue) isEquals(other *SubCue) bool {
	if other == nil {
		return false
	}

	// Compare Settings and Payload as strings (null check included)
	if !utils.StringEquals(&s.Settings, &other.Settings) ||
		!utils.StringEquals(&s.Payload, &other.Payload) {
		return false
	}

	// Compare StartTime and EndTime as time.Duration
	if s.StartTime != other.StartTime || s.EndTime != other.EndTime {
		return false
	}

	return true
}
