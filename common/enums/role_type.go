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

var roleTypeToString = map[RoleType]string{
	Subtitle:      "Subtitle",
	Main:          "Main",
	Alternate:     "Alternate",
	Supplementary: "Supplementary",
	Commentary:    "Commentary",
	Dub:           "Dub",
	Description:   "Description",
	Sign:          "Sign",
	Metadata:      "Metadata",
}

// String returns the string representation of the RoleType.
func (r RoleType) String() *string {
	if roleName, ok := roleTypeToString[r]; ok {
		return &roleName
	}
	unknown := "Unknown RoleType"
	return &unknown // Return pointer to "Unknown RoleType" if not found
}
