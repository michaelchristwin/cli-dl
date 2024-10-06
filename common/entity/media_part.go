package entity

type MediaPart struct {
	MediaSegments []MediaSegment // Slice of MediaSegment
}

func NewMediaPart() *MediaPart {
	return &MediaPart{
		MediaSegments: make([]MediaSegment, 0),
	}
}
func (mp *MediaPart) Sum() float64 {
	totalDuration := 0.0
	for _, segment := range mp.MediaSegments {
		totalDuration += segment.Duration
	}
	return totalDuration
}
