package logrotate

import (
	"io"
	"os"
	"time"
)

func Open(name string, n int64) *os.File {
	var r, w, _ = os.Pipe()
	go func() {
		for {
			var f, _ = os.Create(name)
			defer f.Close()
			if _, err := io.CopyN(f, r, n); err != nil {
				return
			}
			os.Rename(name, name+"-"+time.Now().String())
		}
	}()
	return w
}
