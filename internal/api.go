package internal

import (
	"fmt"
	"io"
)

var (
	disableLogging = false
)

func DisableLogging(b bool) {
	disableLogging = b
}

func Fprintf(w io.Writer, format string, a ...any) {
	if disableLogging {
		return
	}

	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		return
	}
}

func Fprintln(w io.Writer, a ...any) {
	if disableLogging {
		return
	}

	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		return
	}
}
