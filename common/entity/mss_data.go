package entity

type MSSData struct {
	FourCC             string
	CodecPrivateData   string
	Type               string
	Timesacle          int
	SamplingRate       int
	Channels           int
	BitsPerSample      int
	NalUnitLengthField int
	Duration           int64
	IsProtection       bool
	ProtectionSystemID string
	ProtectionData     string
}
