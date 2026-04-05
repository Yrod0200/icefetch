package helpers

type ANSI struct {
	Reset string

	// text sytles
	Bold      string
	Dim       string
	Italic    string
	Underline string
	Blink     string
	Reverse   string
	Hidden    string

	// fg
	Black   string
	Red     string
	Green   string
	Yellow  string
	Blue    string
	Magenta string
	Cyan    string
	White   string

	// bright fg
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string

	// bg
	BgBlack   string
	BgRed     string
	BgGreen   string
	BgYellow  string
	BgBlue    string
	BgMagenta string
	BgCyan    string
	BgWhite   string

	// bright bg
	BgBrightBlack   string
	BgBrightRed     string
	BgBrightGreen   string
	BgBrightYellow  string
	BgBrightBlue    string
	BgBrightMagenta string
	BgBrightCyan    string
	BgBrightWhite   string
}

var Color = ANSI{
	Reset: "\033[0m",

	Bold:      "\033[1m",
	Dim:       "\033[2m",
	Italic:    "\033[3m",
	Underline: "\033[4m",
	Blink:     "\033[5m",
	Reverse:   "\033[7m",
	Hidden:    "\033[8m",

	Black:   "\033[30m",
	Red:     "\033[31m",
	Green:   "\033[32m",
	Yellow:  "\033[33m",
	Blue:    "\033[34m",
	Magenta: "\033[35m",
	Cyan:    "\033[36m",
	White:   "\033[37m",

	BrightBlack:   "\033[90m",
	BrightRed:     "\033[91m",
	BrightGreen:   "\033[92m",
	BrightYellow:  "\033[93m",
	BrightBlue:    "\033[94m",
	BrightMagenta: "\033[95m",
	BrightCyan:    "\033[96m",
	BrightWhite:   "\033[97m",

	BgBlack:   "\033[40m",
	BgRed:     "\033[41m",
	BgGreen:   "\033[42m",
	BgYellow:  "\033[43m",
	BgBlue:    "\033[44m",
	BgMagenta: "\033[45m",
	BgCyan:    "\033[46m",
	BgWhite:   "\033[47m",

	BgBrightBlack:   "\033[100m",
	BgBrightRed:     "\033[101m",
	BgBrightGreen:   "\033[102m",
	BgBrightYellow:  "\033[103m",
	BgBrightBlue:    "\033[104m",
	BgBrightMagenta: "\033[105m",
	BgBrightCyan:    "\033[106m",
	BgBrightWhite:   "\033[107m",
}
