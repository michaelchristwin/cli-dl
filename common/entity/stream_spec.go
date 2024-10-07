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
	// Declare all variables at the top
	prefixStr, returnStr, encStr, bandwidthStr, channelsStr, frameRateStr := "", "", "", "", "", ""

	// Add nil checks for crucial pointers
	if s.MediaType == nil {
		empty := ""
		return &empty
	}

	switch *s.MediaType {
	case enums.AUDIO:
		prefixStr = fmt.Sprintf("[deepskyblue3]Aud[/] %s", encStr)
		if s.Bandwidth != nil {
			bandwidthStr = fmt.Sprintf("%d Kbps", *s.Bandwidth/1000)
		}
		if s.Channels != nil {
			channelsStr = fmt.Sprintf("%s CH", *s.Channels)
		}

		// Safely handle potentially nil pointers
		groupID := safeDeref(s.GroupId, "")
		name := safeDeref(s.Name, "")
		codecs := safeDeref(s.Codecs, "")
		language := safeDeref(s.Language, "")
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		d := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s",
			groupID,
			bandwidthStr,
			name,
			codecs,
			language,
			channelsStr,
			role,
		)
		returnStr = html.EscapeString(d)

	case enums.SUBTITLES:
		prefixStr = fmt.Sprintf("[deepskyblue3_1]Sub[/] %s", encStr)

		// Safely handle potentially nil pointers
		groupID := safeDeref(s.GroupId, "")
		language := safeDeref(s.Language, "")
		name := safeDeref(s.Name, "")
		codecs := safeDeref(s.Codecs, "")
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}

		d := fmt.Sprintf("%s | %s | %s | %s | %s",
			groupID,
			language,
			name,
			codecs,
			role,
		)
		returnStr = html.EscapeString(d)

	default:
		prefixStr = fmt.Sprintf("[aqua]Vid[/] %s", encStr)
		if s.Bandwidth != nil {
			bandwidthStr = fmt.Sprintf("%d Kbps", *s.Bandwidth/1000)
		}
		if s.FrameRate != nil {
			frameRateStr = fmt.Sprintf("%v", s.FrameRate)
		}

		// Safely handle potentially nil pointers
		resolution := safeDeref(s.Resolution, "")
		groupID := safeDeref(s.GroupId, "")
		codecs := safeDeref(s.Codecs, "")
		videoRange := safeDeref(s.VideoRange, "")
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}

		d := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s",
			resolution,
			bandwidthStr,
			groupID,
			frameRateStr,
			codecs,
			videoRange,
			role,
		)
		returnStr = html.EscapeString(d)
	}

	returnStr = prefixStr + returnStr
	returnStr = strings.TrimSpace(strings.Trim(returnStr, "|"))

	// Simplify multiple pipe replacements
	for strings.Contains(returnStr, "| |") {
		returnStr = strings.ReplaceAll(returnStr, "| |", "|")
	}

	return &returnStr
}

func (s *StreamSpec) ToShortShortString() string {
	var prefixStr, returnStr, encStr string

	switch *s.MediaType {
	case enums.AUDIO:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[deepskyblue3]Aud[/] %s", encStr)
		d := fmt.Sprintf("%d Kbps | %s | %s | %sCH | %s",
			optionalBandwidth(s.Bandwidth),
			safeDeref(s.Name, ""),
			safeDeref(s.Language, ""),
			safeDeref(s.Channels, ""),
			role,
		)
		returnStr = html.EscapeString(d)

	case enums.SUBTITLES:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[deepskyblue3_1]Sub[/] %s", encStr)
		d := fmt.Sprintf("%s | %s | %s | %s",
			safeDeref(s.Language, ""),
			safeDeref(s.Name, ""),
			safeDeref(s.Codecs, ""),
			role,
		)
		returnStr = html.EscapeString(d)

	default:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[aqua]Vid[/] %s", encStr)
		d := fmt.Sprintf("%s | %d Kbps | %.2f | %s | %s",
			safeDeref(s.Resolution, ""),
			optionalBandwidth(s.Bandwidth),
			optionalFloat(s.FrameRate),
			safeDeref(s.VideoRange, ""),
			role,
		)
		returnStr = html.EscapeString(d)
	}

	returnStr = cleanUpString(prefixStr + strings.Trim(returnStr, "|"))
	return returnStr
}

func (s *StreamSpec) ToString() string {
	var prefixStr, returnStr, encStr string
	segmentsCountStr := formatSegmentsCount(s.SegmentsCount)

	// Check for encryption across media segments
	if s.Playlist != nil && hasEncryption(s.Playlist) {
		encStr = fmt.Sprintf("[red]*%s[/] ", getEncryptionMethods(s.Playlist))
	}

	switch *s.MediaType {
	case enums.AUDIO:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[deepskyblue3]Aud[/] %s", encStr)
		d := fmt.Sprintf("%s | %d Kbps | %s | %s | %s | %sCH | %s | %s",
			safeDeref(s.GroupId, ""),
			optionalBandwidth(s.Bandwidth),
			safeDeref(s.Name, ""),
			safeDeref(s.Codecs, ""),
			safeDeref(s.Language, ""),
			safeDeref(s.Channels, ""),
			segmentsCountStr,
			role,
		)
		returnStr = html.EscapeString(d)

	case enums.SUBTITLES:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[deepskyblue3_1]Sub[/] %s", encStr)
		d := fmt.Sprintf("%s | %s | %s | %s | %s | %s | %s",
			safeDeref(s.GroupId, ""),
			safeDeref(s.Language, ""),
			safeDeref(s.Name, ""),
			safeDeref(s.Codecs, ""),
			safeDeref(s.Characteristics, ""),
			segmentsCountStr,
			role,
		)
		returnStr = html.EscapeString(d)

	default:
		role := ""
		if s.Role != nil {
			roleStr := (*s.Role).String()
			if roleStr != nil {
				role = *roleStr
			}
		}
		prefixStr = fmt.Sprintf("[aqua]Vid[/] %s", encStr)
		d := fmt.Sprintf("%s | %d Kbps | %s | %.2f | %s | %s | %s | %s",
			safeDeref(s.Resolution, ""),
			optionalBandwidth(s.Bandwidth),
			safeDeref(s.GroupId, ""),
			optionalFloat(s.FrameRate),
			safeDeref(s.Codecs, ""),
			safeDeref(s.VideoRange, ""),
			segmentsCountStr,
			role,
		)
		returnStr = html.EscapeString(d)
	}

	// Clean up string and handle trailing pipes
	returnStr = cleanUpString(prefixStr + strings.Trim(returnStr, "|"))

	// Calculate total duration
	if s.Playlist != nil {
		totalDuration := s.Playlist.TotalDuration
		returnStr += fmt.Sprintf(" | ~%s", FormatTime(int(totalDuration)))
	}

	return returnStr
}

func FormatTime(seconds int) string {
	// Stub for formatting time, similar to C#'s time formatting
	return fmt.Sprintf("%d:%02d", seconds/60, seconds%60)
}

func optionalBandwidth(b *int) int {
	if b != nil {
		return *b / 1000
	}
	return 0
}

func optionalFloat(f *float64) float64 {
	if f != nil {
		return *f
	}
	return 0
}

func cleanUpString(input string) string {
	input = strings.TrimSpace(strings.Trim(input, "|"))
	for strings.Contains(input, "|  |") {
		input = strings.ReplaceAll(input, "|  |", "|")
	}
	return strings.TrimRight(input, "|")
}

func formatSegmentsCount(count int) string {
	if count == 0 {
		return ""
	}
	if count > 1 {
		return fmt.Sprintf("%d Segments", count)
	}
	return "1 Segment"
}

func hasEncryption(playlist *Playlist) bool {
	for _, part := range playlist.MediaParts {
		for _, segment := range part.MediaSegments {
			if segment.EncryptInfo.Method.String() != "None" {
				return true
			}
		}
	}
	return false
}

func getEncryptionMethods(playlist *Playlist) string {
	methods := make(map[enums.EncryptMethod]bool)
	for _, part := range playlist.MediaParts {
		for _, segment := range part.MediaSegments {
			if segment.EncryptInfo.Method.String() != "None" {
				methods[segment.EncryptInfo.Method] = true
			}
		}
	}
	uniqueMethods := []string{}
	for method := range methods {
		uniqueMethods = append(uniqueMethods, method.String())
	}
	return strings.Join(uniqueMethods, ",")
}

func safeDeref(s *string, defaultValue string) string {
	if s == nil {
		return defaultValue
	}
	return *s
}
