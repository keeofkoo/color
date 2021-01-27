package color

import (
	_fmt "fmt"
	"os"
	"strings"
	"sync"
)

var isDisabled bool
var pool *sync.Pool

func init() {
	_, ok := os.LookupEnv("ANSI_COLORS_DISABLED")
	isDisabled = ok
	pool = &sync.Pool{New: func() interface{} { return strings.Builder{} }}
}

// Format defines format type.
type Format int

const reset Format = iota

// Attributes
const (
	Bold Format = iota + 1
	Dark
	_
	Underline
	Blink
	_
	Reverse
	Concealed
)

// Colors
const (
	Gray Format = iota + 30
	Red
	Green
	Yellow
	Blue
	Magenta
	Cyan
	White
)

// Highlights
const (
	OnGray Format = iota + 40
	OnRed
	OnGreen
	OnYellow
	OnBlue
	OnMagenta
	OnCyan
	OnWhite
)

const tokenFmt = "\x1b[%dm"

func token(fmt Format) string {
	return _fmt.Sprintf(tokenFmt, fmt)
}

// Colored applies a series of formats to the string and returns the result.
func Colored(s string, fmts ...Format) string {
	// short path
	if isDisabled || len(fmts) == 0 {
		return s
	}

	// process the formats
	var (
		attrs     []Format
		color     Format
		highlight Format
	)
	for _, fmt := range fmts {
		switch fmt {
		case Bold, Dark, Underline, Blink, Reverse, Concealed:
			attrs = append(attrs, fmt)
		case Gray, Red, Green, Yellow, Blue, Magenta, Cyan, White:
			color = fmt
		case OnGray, OnRed, OnGreen, OnYellow, OnBlue, OnMagenta, OnCyan, OnWhite:
			highlight = fmt
		}
	}

	// build the string
	b := pool.Get().(strings.Builder)
	defer func() {
		b.Reset()
		pool.Put(b)
	}()
	for i := len(attrs) - 1; i >= 0; i-- {
		b.WriteString(token(attrs[i]))
	}
	if highlight != 0 {
		b.WriteString(token(highlight))
	}
	if color != 0 {
		b.WriteString(token(color))
	}
	b.WriteString(s)
	b.WriteString(token(reset))

	// ok
	return b.String()
}
