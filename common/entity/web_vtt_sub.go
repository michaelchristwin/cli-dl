package entity

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	tsMapRegex    = regexp.MustCompile(`X-TIMESTAMP-MAP.*`)
	tsValueRegex  = regexp.MustCompile(`MPEGTS:(\d+)`)
	splitRegex    = regexp.MustCompile(`\s`)
	vttClassRegex = regexp.MustCompile(`<c\..*?>([\\s\\S]*?)</c>`)
)

type WebVttSub struct {
	Cues            []SubCue
	MpegtsTimestamp int64
}

func Parse(text string, baseTimestamp int64) (*WebVttSub, error) {
	trimmedText := strings.TrimSpace(text)
	if !strings.HasPrefix(trimmedText, "WEBVTT") {
		return nil, fmt.Errorf("bad vtt")
	}

	webSub := &WebVttSub{}
	lines := strings.Split(text, "\n")

	// Handle timestamp map
	if tsMapRegex.MatchString(text) {
		matches := tsValueRegex.FindStringSubmatch(tsMapRegex.FindString(text))
		if len(matches) > 1 {
			timestamp, err := strconv.ParseInt(matches[1], 10, 64)
			if err == nil {
				webSub.MpegtsTimestamp = timestamp
			}
		}
	}

	var needPayload bool
	var timeLine string
	var payloads []string

	for _, line := range lines {
		if strings.Contains(line, " --> ") {
			needPayload = true
			timeLine = strings.TrimSpace(line)
			continue
		}

		if needPayload {
			if strings.TrimSpace(line) == "" {
				payload := strings.Join(payloads, "\n")
				if strings.TrimSpace(payload) == "" {
					payloads = nil
					continue
				}

				timeComponents := splitRegex.Split(strings.ReplaceAll(timeLine, "-->", ""), -1)
				var nonEmptyComponents []string
				for _, comp := range timeComponents {
					if comp != "" {
						nonEmptyComponents = append(nonEmptyComponents, comp)
					}
				}

				if len(nonEmptyComponents) >= 2 {
					startTime := convertToTS(nonEmptyComponents[0])
					endTime := convertToTS(nonEmptyComponents[1])
					style := ""
					if len(nonEmptyComponents) > 2 {
						style = strings.Join(nonEmptyComponents[2:], " ")
					}

					webSub.Cues = append(webSub.Cues, SubCue{
						StartTime: startTime,
						EndTime:   endTime,
						Payload:   removeClassTag(strings.Join(payloads, "")),
						Settings:  style,
					})
				}

				payloads = nil
				needPayload = false
			} else {
				payloads = append(payloads, strings.TrimSpace(line))
			}
		}
	}

	if baseTimestamp != 0 {
		for i := range webSub.Cues {
			startMs := webSub.Cues[i].StartTime.Milliseconds()
			if startMs-baseTimestamp >= 0 {
				webSub.Cues[i].StartTime = time.Duration(startMs-baseTimestamp) * time.Millisecond
				webSub.Cues[i].EndTime = time.Duration(webSub.Cues[i].EndTime.Milliseconds()-baseTimestamp) * time.Millisecond
			} else {
				break
			}
		}
	}

	return webSub, nil
}

func removeClassTag(text string) string {
	if vttClassRegex.MatchString(text) {
		matches := vttClassRegex.FindAllStringSubmatch(text, -1)
		var result strings.Builder
		for _, match := range matches {
			if len(match) > 1 {
				result.WriteString(match[1])
				result.WriteString(" ")
			}
		}
		return strings.TrimSpace(result.String())
	}
	return text
}

func convertToTS(str string) time.Duration {
	str = strings.TrimSpace(str)

	// Handle "s" suffix for seconds
	if strings.HasSuffix(str, "s") {
		sec, err := strconv.ParseFloat(str[:len(str)-1], 64)
		if err == nil {
			return time.Duration(sec * float64(time.Second))
		}
	}

	// Handle standard timestamp format
	str = strings.ReplaceAll(str, ",", ".")
	parts := strings.Split(str, ".")

	var ms int64
	if len(parts) > 1 {
		ms, _ = strconv.ParseInt(parts[1], 10, 64)
	}

	timeParts := strings.Split(parts[0], ":")
	var duration time.Duration

	for i := len(timeParts) - 1; i >= 0; i-- {
		value, _ := strconv.ParseInt(timeParts[i], 10, 64)
		duration += time.Duration(value) * time.Duration(pow(60, len(timeParts)-1-i)) * time.Second
	}

	return duration + time.Duration(ms)*time.Millisecond
}

func pow(x, y int) int64 {
	result := int64(1)
	for i := 0; i < y; i++ {
		result *= int64(x)
	}
	return result
}

func (w *WebVttSub) String() string {
	var sb strings.Builder
	for _, cue := range w.getCues() {
		fmt.Fprintf(&sb, "%02d:%02d:%02d.%03d --> %02d:%02d:%02d.%03d %s\n",
			int(cue.StartTime.Hours()), int(cue.StartTime.Minutes())%60, int(cue.StartTime.Seconds())%60, cue.StartTime.Milliseconds()%1000,
			int(cue.EndTime.Hours()), int(cue.EndTime.Minutes())%60, int(cue.EndTime.Seconds())%60, cue.EndTime.Milliseconds()%1000,
			cue.Settings)
		sb.WriteString(cue.Payload)
		sb.WriteString("\n\n")
	}
	return sb.String()
}

func (w *WebVttSub) getCues() []SubCue {
	var result []SubCue
	for _, cue := range w.Cues {
		if cue.Payload != "" {
			result = append(result, cue)
		}
	}
	return result
}

func (w *WebVttSub) ToVtt() string {
	return "WEBVTT\n\n" + w.String()
}

func (w *WebVttSub) ToSrt() string {
	var sb strings.Builder
	index := 1

	for _, cue := range w.getCues() {
		fmt.Fprintf(&sb, "%d\n", index)
		fmt.Fprintf(&sb, "%02d:%02d:%02d,%03d --> %02d:%02d:%02d,%03d\n",
			int(cue.StartTime.Hours()), int(cue.StartTime.Minutes())%60, int(cue.StartTime.Seconds())%60, cue.StartTime.Milliseconds()%1000,
			int(cue.EndTime.Hours()), int(cue.EndTime.Minutes())%60, int(cue.EndTime.Seconds())%60, cue.EndTime.Milliseconds()%1000)
		sb.WriteString(cue.Payload)
		sb.WriteString("\n\n")
		index++
	}

	srt := sb.String()
	if strings.TrimSpace(srt) == "" {
		return "1\n00:00:00,000 --> 00:00:01,000"
	}

	return srt
}
