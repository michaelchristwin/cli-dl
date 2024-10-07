package entity

import (
	"github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/utils"
	"hash"
	"hash/fnv"
	"log"
	"strconv"
	"time"
)

type MediaSegment struct {
	Index        int64
	Duration     float64
	Title        *string
	DateTime     *time.Time
	StartRange   *int64
	StopRange    *int64
	ExpectLength *int64
	EncryptInfo  EncryptInfo
	Url          string
	NameFromVar  *string
}

func (m *MediaSegment) CalculateStopRange() *int64 {
	if m.StartRange != nil && m.ExpectLength != nil {
		stopRange := *m.StartRange + *m.ExpectLength - 1
		return &stopRange
	}
	return nil
}

func (m *MediaSegment) isEquals(other *MediaSegment) bool {
	if other == nil {
		return false
	}

	// Compare primitive fields
	if m.Index != other.Index || m.Duration != other.Duration || m.Url != other.Url { // Compare Url directly
		return false
	}

	// Compare pointer fields (nullable in Go)
	if !utils.StringEquals(m.Title, other.Title) ||
		!utils.Int64Equals(m.StartRange, other.StartRange) ||
		!utils.Int64Equals(m.StopRange, other.StopRange) ||
		!utils.Int64Equals(m.ExpectLength, other.ExpectLength) {
		return false
	}

	return true
}

func hashPointerString(h hash.Hash64, s *string) {
	if s != nil {
		h.Write([]byte(*s))
	} else {
		h.Write([]byte("nil"))
	}
}
func hashString(h hash.Hash64, s string) {
	_, err := h.Write([]byte(s))
	if err != nil {
		log.Printf("Error writing to hash: %v", err)
	}
}

// Helper function to hash an int64
func hashInt64(h hash.Hash64, i *int64) {
	if i != nil {
		h.Write([]byte(strconv.FormatInt(*i, 10)))
	} else {
		h.Write([]byte("nil"))
	}
}

// Helper function to hash a float64
func hashFloat64(h hash.Hash64, f float64) {
	h.Write([]byte(strconv.FormatFloat(f, 'f', -1, 64)))
}

// GetHashCode generates a combined hash for all fields
func (m *MediaSegment) GetHashCode() int {
	h := fnv.New64a()

	// Hash each field
	hashInt64(h, &m.Index)
	hashFloat64(h, m.Duration)
	hashPointerString(h, m.Title)
	hashInt64(h, m.StartRange)
	hashInt64(h, m.StopRange)
	hashInt64(h, m.ExpectLength)
	hashString(h, m.Url)

	// Convert hash to int (take lower bits, ensure compatibility)
	return int(h.Sum64())
}

// Helper function to compare two *int64 values
