package commandline

import (
	"flag"
	"fmt"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

const Version string = "0.0.1"

// headerMap is a custom type that represents a map of headers.
type headerMap map[string]string

// headersVar is a wrapper for the headerMap to satisfy the flag.Value interface.
type headersVar struct {
	headers *headerMap
}

// String returns the string representation of the headerMap.
func (h *headersVar) String() string {
	// Return a formatted string representation of the headers.
	var result []string
	for key, value := range *h.headers {
		result = append(result, fmt.Sprintf("%s:%s", key, value))
	}
	return strings.Join(result, ", ")
}

// Set parses a key:value string and adds it to the headerMap.
func (h *headersVar) Set(value string) error {
	// Expecting key:value format
	parts := strings.SplitN(value, ":", 2)
	if len(parts) != 2 {
		return fmt.Errorf("invalid header format, expecting key:value, got: %s", value)
	}

	key := strings.TrimSpace(parts[0])
	val := strings.TrimSpace(parts[1])

	// Initialize the headerMap if it is nil
	if *h.headers == nil {
		*h.headers = make(headerMap)
	}

	// Add the key-value pair to the headerMap
	(*h.headers)[key] = val
	return nil
}

type stringSlice []string

func (s *stringSlice) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringSlice) Set(value string) error {
	*s = append(*s, value)
	return nil
}

var speedRegex = regexp.MustCompile(`([0-9\.]+)(M|K)`)

// Parse speed limit
func parseSpeedLimit(value string) (*int64, error) {
	input := strings.ToUpper(value)
	match := speedRegex.FindStringSubmatch(input)
	if match == nil {
		return nil, fmt.Errorf("invalid speed limit format: %s", input)
	}

	number, err := strconv.ParseFloat(match[1], 64)
	if err != nil {
		return nil, err
	}

	var speed int64
	switch match[2] {
	case "M":
		speed = int64(number * 1024 * 1024)
	case "K":
		speed = int64(number * 1024)
	default:
		return nil, fmt.Errorf("unknown unit: %s", match[2])
	}

	return &speed, nil
}

type speedFlag int64

func (s *speedFlag) String() string {
	return fmt.Sprintf("%d", *s)
}

func (s *speedFlag) Set(value string) error {
	parsedSpeed, err := parseSpeedLimit(value)
	if err != nil {
		return err
	}
	*s = speedFlag(*parsedSpeed)
	return nil
}

type Options struct {
	Input                  string
	TmpDir                 *string
	SaveDir                *string
	SaveName               *string
	SavePattern            *string
	UILanguage             *string
	UrlProcessorArgs       *string
	Keys                   *stringSlice
	KeyTextFile            string
	Headers                *headerMap
	LogLevel               string
	SubtitleFormat         string
	AutoSelect             bool
	SubOnly                bool
	ThreadCount            int
	DownloadRetryCount     int
	SkipMerge              bool
	SkipDownload           bool
	NoDateInfo             bool
	BinaryMerge            bool
	UseFFmpegConcatDemuxer bool
	DelAfterDone           bool
	AutoSubtitleFix        bool
	CheckSegementsCount    bool
	WriteMetaJson          bool
	AppendUrlParams        bool
	MP4RealTimeDecryption  bool
	UseShakaPackager       bool
	ForceAnsiConsole       bool
	NoAnsiColor            bool
	DecryptionBinaryPath   string
	FFmpegBinaryPath       string
	BaseUrl                string
	ConcurrentDownload     bool
	NoLog                  bool
	AdKeywords             *stringSlice
	MaxSpeed               *speedFlag
	UseSystemProxy         bool
}

func CommandInvoker() Options {
	opts := Options{
		Headers:          new(headerMap),
		Keys:             new(stringSlice),
		AdKeywords:       new(stringSlice),
		MaxSpeed:         new(speedFlag),
		TmpDir:           new(string),
		SaveDir:          new(string),
		SaveName:         new(string),
		SavePattern:      new(string),
		UILanguage:       new(string),
		UrlProcessorArgs: new(string),
	}

	flag.StringVar(&opts.Input, "input", "", "Input URL or file")
	flag.StringVar(opts.TmpDir, "tmp-dir", "", "Set directory for temporary files")
	flag.StringVar(opts.SaveDir, "save-dir", "", "Set ouput directory")
	flag.StringVar(opts.SavePattern, "save-pattern", "", "Set")
	flag.StringVar(opts.UILanguage, "ui-language", "", "")
	flag.StringVar(opts.UrlProcessorArgs, "urlprocessor-args", "", "")
	flag.Var(opts.Keys, "key", "Pass decryption key(s) to mp4decrypt/shaka-packager. format:\r\n--key KID1:KEY1 --key KID2:KEY2")
	flag.StringVar(&opts.KeyTextFile, "key-text-file", "", "")
	flag.Var(&headersVar{opts.Headers}, "H", "Specify headers in the format key:value")
	flag.Var(&headersVar{opts.Headers}, "header", "Specify headers in the format key:value")
	flag.StringVar(&opts.LogLevel, "log-level", "INFO", "Set log level")
	flag.StringVar(&opts.SubtitleFormat, "sub-format", "SRT", "")
	flag.BoolVar(&opts.AutoSelect, "auto-select", false, "")
	flag.BoolVar(&opts.SubOnly, "sub-only", false, "")
	flag.IntVar(&opts.ThreadCount, "thread-count", runtime.GOMAXPROCS(0), "")
	flag.IntVar(&opts.DownloadRetryCount, "download-retry-count", 3, "")
	flag.BoolVar(&opts.SkipMerge, "skip-merge", false, "")
	flag.BoolVar(&opts.SkipDownload, "skip-download", false, "")
	flag.BoolVar(&opts.NoDateInfo, "no-date-info", false, "")
	flag.BoolVar(&opts.BinaryMerge, "binary-merge", false, "")
	flag.BoolVar(&opts.UseFFmpegConcatDemuxer, "use-ffmpeg-concat-demuxer", false, "")
	flag.BoolVar(&opts.DelAfterDone, "del-after-done", true, "")
	flag.BoolVar(&opts.AutoSubtitleFix, "auto-subtitle-fix", true, "")
	flag.BoolVar(&opts.CheckSegementsCount, "check-segments-count", true, "")
	flag.BoolVar(&opts.WriteMetaJson, "write-meta-json", true, "")

	flag.BoolVar(&opts.AppendUrlParams, "append-url-params", false, "Description for append-url-params")
	flag.BoolVar(&opts.MP4RealTimeDecryption, "mp4-real-time-decryption", false, "Description for mp4-real-time-decryption")
	flag.BoolVar(&opts.UseShakaPackager, "use-shaka-packager", false, "Description for use-shaka-packager")
	flag.BoolVar(&opts.ForceAnsiConsole, "force-ansi-console", false, "Description for force-ansi-console")
	flag.BoolVar(&opts.NoAnsiColor, "no-ansi-color", false, "Description for no-ansi-color")
	flag.StringVar(&opts.DecryptionBinaryPath, "decryption-binary-path", "", "Path for decryption-binary")
	flag.StringVar(&opts.FFmpegBinaryPath, "ffmpeg-binary-path", "", "Path for FFmpeg-binary")
	flag.StringVar(&opts.BaseUrl, "base-url", "", "Base URL for the operation")
	flag.BoolVar(&opts.ConcurrentDownload, "concurrent-download", false, "Enable concurrent downloads")
	flag.BoolVar(&opts.NoLog, "no-log", false, "Disable logging")
	flag.Var(opts.AdKeywords, "ad-keyword", "Ad keywords (can specify multiple)")
	flag.Var(opts.MaxSpeed, "R", "Max download speed (in bytes/sec)")
	flag.Var(opts.MaxSpeed, "max-speed", "Max download speed (in bytes/sec)")
	flag.BoolVar(&opts.UseSystemProxy, "use-system-proxy", true, "")

	//Parse all flags
	flag.Parse()

	return opts
}
