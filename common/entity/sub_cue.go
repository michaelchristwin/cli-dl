package entity

import (
	"hash"
	"hash/fnv"
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

func (s *SubCue) GetHashCode() int {
	h := fnv.New64a()
	localHashInt64(h, int64(s.StartTime))
	localHashInt64(h, int64(s.EndTime))

	hashString(h, s.Settings)
	hashString(h, s.Payload)
	return int(h.Sum64())
}

func localHashInt64(h hash.Hash64, i int64) {
	var buf [8]byte
	for j := 0; j < 8; j++ {
		buf[j] = byte(i >> (56 - j*8))
	}
	h.Write(buf[:])
}
