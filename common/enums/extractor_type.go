package enums

type ExtractorType int

// Define enum values using iota
const (
	MPEG_DASH ExtractorType = iota
	HLS
	HTTP_LIVE
	MSS
)
