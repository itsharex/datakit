// Generated automatically.  DO NOT HAND-EDIT.

package xterm

import "github.com/gdamore/tcell/v2/terminfo"

func init() {

	// X11 terminal emulator
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:          "xterm",
		Aliases:       []string{"xterm-debian"},
		Columns:       80,
		Lines:         24,
		Colors:        8,
		Bell:          "\a",
		Clear:         "\x1b[H\x1b[2J",
		EnterCA:       "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:        "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:    "\x1b[?12l\x1b[?25h",
		HideCursor:    "\x1b[?25l",
		AttrOff:       "\x1b(B\x1b[m",
		Underline:     "\x1b[4m",
		Bold:          "\x1b[1m",
		Dim:           "\x1b[2m",
		Italic:        "\x1b[3m",
		Blink:         "\x1b[5m",
		Reverse:       "\x1b[7m",
		EnterKeypad:   "\x1b[?1h\x1b=",
		ExitKeypad:    "\x1b[?1l\x1b>",
		SetFg:         "\x1b[3%p1%dm",
		SetBg:         "\x1b[4%p1%dm",
		SetFgBg:       "\x1b[3%p1%d;4%p2%dm",
		ResetFgBg:     "\x1b[39;49m",
		AltChars:      "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:      "\x1b(0",
		ExitAcs:       "\x1b(B",
		StrikeThrough: "\x1b[9m",
		Mouse:         "\x1b[M",
		SetCursor:     "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:   "\b",
		CursorUp1:     "\x1b[A",
		KeyUp:         "\x1bOA",
		KeyDown:       "\x1bOB",
		KeyRight:      "\x1bOC",
		KeyLeft:       "\x1bOD",
		KeyInsert:     "\x1b[2~",
		KeyDelete:     "\x1b[3~",
		KeyBackspace:  "\u007f",
		KeyHome:       "\x1bOH",
		KeyEnd:        "\x1bOF",
		KeyPgUp:       "\x1b[5~",
		KeyPgDn:       "\x1b[6~",
		KeyF1:         "\x1bOP",
		KeyF2:         "\x1bOQ",
		KeyF3:         "\x1bOR",
		KeyF4:         "\x1bOS",
		KeyF5:         "\x1b[15~",
		KeyF6:         "\x1b[17~",
		KeyF7:         "\x1b[18~",
		KeyF8:         "\x1b[19~",
		KeyF9:         "\x1b[20~",
		KeyF10:        "\x1b[21~",
		KeyF11:        "\x1b[23~",
		KeyF12:        "\x1b[24~",
		KeyBacktab:    "\x1b[Z",
		Modifiers:     1,
		AutoMargin:    true,
	})

	// xterm with 88 colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:          "xterm-88color",
		Columns:       80,
		Lines:         24,
		Colors:        88,
		Bell:          "\a",
		Clear:         "\x1b[H\x1b[2J",
		EnterCA:       "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:        "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:    "\x1b[?12l\x1b[?25h",
		HideCursor:    "\x1b[?25l",
		AttrOff:       "\x1b(B\x1b[m",
		Underline:     "\x1b[4m",
		Bold:          "\x1b[1m",
		Dim:           "\x1b[2m",
		Italic:        "\x1b[3m",
		Blink:         "\x1b[5m",
		Reverse:       "\x1b[7m",
		EnterKeypad:   "\x1b[?1h\x1b=",
		ExitKeypad:    "\x1b[?1l\x1b>",
		SetFg:         "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:         "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:       "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		ResetFgBg:     "\x1b[39;49m",
		AltChars:      "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:      "\x1b(0",
		ExitAcs:       "\x1b(B",
		StrikeThrough: "\x1b[9m",
		Mouse:         "\x1b[M",
		SetCursor:     "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:   "\b",
		CursorUp1:     "\x1b[A",
		KeyUp:         "\x1bOA",
		KeyDown:       "\x1bOB",
		KeyRight:      "\x1bOC",
		KeyLeft:       "\x1bOD",
		KeyInsert:     "\x1b[2~",
		KeyDelete:     "\x1b[3~",
		KeyBackspace:  "\u007f",
		KeyHome:       "\x1bOH",
		KeyEnd:        "\x1bOF",
		KeyPgUp:       "\x1b[5~",
		KeyPgDn:       "\x1b[6~",
		KeyF1:         "\x1bOP",
		KeyF2:         "\x1bOQ",
		KeyF3:         "\x1bOR",
		KeyF4:         "\x1bOS",
		KeyF5:         "\x1b[15~",
		KeyF6:         "\x1b[17~",
		KeyF7:         "\x1b[18~",
		KeyF8:         "\x1b[19~",
		KeyF9:         "\x1b[20~",
		KeyF10:        "\x1b[21~",
		KeyF11:        "\x1b[23~",
		KeyF12:        "\x1b[24~",
		KeyBacktab:    "\x1b[Z",
		Modifiers:     1,
		AutoMargin:    true,
	})

	// xterm with 256 colors
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:          "xterm-256color",
		Columns:       80,
		Lines:         24,
		Colors:        256,
		Bell:          "\a",
		Clear:         "\x1b[H\x1b[2J",
		EnterCA:       "\x1b[?1049h\x1b[22;0;0t",
		ExitCA:        "\x1b[?1049l\x1b[23;0;0t",
		ShowCursor:    "\x1b[?12l\x1b[?25h",
		HideCursor:    "\x1b[?25l",
		AttrOff:       "\x1b(B\x1b[m",
		Underline:     "\x1b[4m",
		Bold:          "\x1b[1m",
		Dim:           "\x1b[2m",
		Italic:        "\x1b[3m",
		Blink:         "\x1b[5m",
		Reverse:       "\x1b[7m",
		EnterKeypad:   "\x1b[?1h\x1b=",
		ExitKeypad:    "\x1b[?1l\x1b>",
		SetFg:         "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;m",
		SetBg:         "\x1b[%?%p1%{8}%<%t4%p1%d%e%p1%{16}%<%t10%p1%{8}%-%d%e48;5;%p1%d%;m",
		SetFgBg:       "\x1b[%?%p1%{8}%<%t3%p1%d%e%p1%{16}%<%t9%p1%{8}%-%d%e38;5;%p1%d%;;%?%p2%{8}%<%t4%p2%d%e%p2%{16}%<%t10%p2%{8}%-%d%e48;5;%p2%d%;m",
		ResetFgBg:     "\x1b[39;49m",
		AltChars:      "``aaffggiijjkkllmmnnooppqqrrssttuuvvwwxxyyzz{{||}}~~",
		EnterAcs:      "\x1b(0",
		ExitAcs:       "\x1b(B",
		StrikeThrough: "\x1b[9m",
		Mouse:         "\x1b[M",
		SetCursor:     "\x1b[%i%p1%d;%p2%dH",
		CursorBack1:   "\b",
		CursorUp1:     "\x1b[A",
		KeyUp:         "\x1bOA",
		KeyDown:       "\x1bOB",
		KeyRight:      "\x1bOC",
		KeyLeft:       "\x1bOD",
		KeyInsert:     "\x1b[2~",
		KeyDelete:     "\x1b[3~",
		KeyBackspace:  "\u007f",
		KeyHome:       "\x1bOH",
		KeyEnd:        "\x1bOF",
		KeyPgUp:       "\x1b[5~",
		KeyPgDn:       "\x1b[6~",
		KeyF1:         "\x1bOP",
		KeyF2:         "\x1bOQ",
		KeyF3:         "\x1bOR",
		KeyF4:         "\x1bOS",
		KeyF5:         "\x1b[15~",
		KeyF6:         "\x1b[17~",
		KeyF7:         "\x1b[18~",
		KeyF8:         "\x1b[19~",
		KeyF9:         "\x1b[20~",
		KeyF10:        "\x1b[21~",
		KeyF11:        "\x1b[23~",
		KeyF12:        "\x1b[24~",
		KeyBacktab:    "\x1b[Z",
		Modifiers:     1,
		AutoMargin:    true,
	})
}
