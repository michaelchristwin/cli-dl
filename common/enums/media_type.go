package enums

type MediaType int

const (
	AUDIO MediaType = iota
	VIDEO
	SUBTITLES
	CLOSED_CAPTIONS
)
