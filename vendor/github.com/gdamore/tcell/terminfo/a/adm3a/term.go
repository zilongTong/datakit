// Generated automatically.  DO NOT HAND-EDIT.

package adm3a

import "github.com/gdamore/tcell/terminfo"

func init() {

	// lsi adm3a
	terminfo.AddTerminfo(&terminfo.Terminfo{
		Name:        "adm3a",
		Columns:     80,
		Lines:       24,
		Bell:        "\a",
		Clear:       "\x1a$<1/>",
		PadChar:     "\x00",
		SetCursor:   "\x1b=%p1%' '%+%c%p2%' '%+%c",
		CursorBack1: "\b",
		CursorUp1:   "\v",
		KeyUp:       "\v",
		KeyDown:     "\n",
		KeyRight:    "\f",
		KeyLeft:     "\b",
	})
}
