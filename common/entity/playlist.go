package entity

type Playlist struct {
	Url               string
	Islive            bool
	RefreshIntervalMs float64
	TotalDuration     float64
	TargetDuration    *float64
	MediaInit         *MediaSegment
	MediaParts        []MediaPart
}

func (p *Playlist) GetTotalDuration() *float64 {
	totalDuration := 0.0

	// Sum the duration of all MediaParts
	for _, part := range p.MediaParts {
		totalDuration += part.Sum()
	}

	// Assign the total duration to the Playlist's TotalDuration field
	p.TotalDuration = totalDuration

	// Return a pointer to the total duration
	return &p.TotalDuration
}

func NewPlaylist() *Playlist {
	return &Playlist{
		RefreshIntervalMs: 15000,
	}
}
