package entity

import (
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/michaelchristwin/N_M3U8DL-RE-go.git/common/enums"
)

type StreamSpec struct {
	MediaType       *enums.MediaType
	GroupId         *string
	Language        *string
	Name            *string
	Default         *enums.Choice
	SkippedDuration *float64
	MSSData         *MSSData
	Bandwidth       *int
	Codecs          *string
	Resolution      *string
	FrameRate       *float64
	Channels        *string
	Extension       *string
	Role            *enums.RoleType
	VideoRange      *string
	Characteristics *string
	PublishTime     *time.Time
	AudioId         *string
	VideoId         *string
	SubtitleId      *string
	PeriodId        *string
	Url             string
	OriginalUrl     string
	Playlist        *Playlist
	SegmentsCount   int
}

func (s *StreamSpec) GetSegmentsCount() *int {
	segmentsCount := 0
	if s.Playlist != nil {
		for _, parts := range s.Playlist.MediaParts {

			segmentsCount += len(parts.MediaSegments)
		}
	}
	s.SegmentsCount = segmentsCount
	return &s.SegmentsCount
}

func (s *StreamSpec) ToShortString() *string {
	prefixStr, returnStr, encStr, bandwidthStr, channelsStr := "", "", "", "", ""
	switch *s.MediaType {
	case enums.AUDIO:
		{
			prefixStr = fmt.Sprintf("[deepskyblue3]Aud[/] %s", encStr)
			if s.Bandwidth != nil {
				bandwidthStr = fmt.Sprintf("%d Kbps", *s.Bandwidth/1000)
			}
			if s.Channels != nil {
				channelsStr = fmt.Sprintf("%s CH", *s.Channels)
			}
			d := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s",
				*s.GroupId,
				bandwidthStr,
				*s.Name,
				*s.Codecs,
				*s.Language,
				channelsStr,
				*s.Role.String(),
			)
			returnStr = html.EscapeString(d)
		}
	case enums.SUBTITLES:
		{
			prefixStr = fmt.Sprintf("[deepskyblue3_1]Sub[/] %s", encStr)
			d := fmt.Sprintf("%s | %s | %s | %s | %s", *s.GroupId, *s.Language, *s.Name, *s.Codecs, *s.Role.String())
			returnStr = html.EscapeString(d)
		}
	default:
		{
			var frameRateStr string
			prefixStr = fmt.Sprintf("[aqua]Vid[/] %s", encStr)
			if s.Bandwidth != nil {
				bandwidthStr = fmt.Sprintf("%d Kbps", *s.Bandwidth/1000)
			}
			if s.FrameRate != nil {
				frameRateStr = fmt.Sprintf("%v", s.FrameRate)
			}
			d := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s", *s.Resolution, bandwidthStr, *s.GroupId, frameRateStr, *s.Codecs, *s.VideoRange, *s.Role.String())
			returnStr = html.EscapeString(d)
		}
	}
	returnStr = prefixStr + returnStr
	returnStr = strings.TrimSpace(strings.TrimSuffix(returnStr, "|"))
	for strings.Contains(returnStr, "|  |") {
		returnStr = strings.ReplaceAll(returnStr, "|  |", "|")
	}
	return &returnStr
}
