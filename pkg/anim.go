package pkg

import (
	"fmt"
	"time"
)

func StartSpinner() (stop func()) {
	done := make(chan struct{})
	go func() {
		frames := []string{"|", "/", "-", "\\"}
		i := 0
		ticker := time.NewTicker(250 * time.Millisecond)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				fmt.Printf("\r%s", frames[i%len(frames)])
				i++
			}
		}
	}()
	return func() {
		close(done)
		fmt.Print("\r\033[K")
	}
}
