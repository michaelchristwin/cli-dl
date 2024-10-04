package enums

type RoleType int

const (
	Subtitle RoleType = iota
	Main
	Alternate
	Supplementary
	Commentary
	Dub
	Description
	Sign
	Metadata
)
