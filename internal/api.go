package internal

import (
	"fmt"
	"io"
	"sync"
)

var (
	disableLogging      = false
	disableLoggingMutex = &sync.Mutex{}
)

func DisableLogging(b bool) {
	disableLoggingMutex.Lock()
	defer disableLoggingMutex.Unlock()

	disableLogging = b
}

func Fprintf(w io.Writer, format string, a ...any) {
	disableLoggingMutex.Lock()
	defer disableLoggingMutex.Unlock()

	if disableLogging {
		return
	}

	_, err := fmt.Fprintf(w, format, a...)
	if err != nil {
		return
	}
}

func Fprintln(w io.Writer, a ...any) {
	disableLoggingMutex.Lock()
	defer disableLoggingMutex.Unlock()

	if disableLogging {
		return
	}

	_, err := fmt.Fprintln(w, a...)
	if err != nil {
		return
	}
}
