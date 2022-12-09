package files

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

const (
	StaticFilesDir = "/static"
	IndexFileName  = "index.html"
)

var (
	staticFiles string = ""
	hasIndex    bool   = false
)

func init() {
	staticFiles, hasIndex = getDirContent(StaticFilesDir)
}

func GetFiles() string {
	return staticFiles
}

func HasIndex() bool {
	return hasIndex
}

func getDirContent(dir string) (string, bool) {
	var buffer bytes.Buffer
	hasIdx := false
	err := filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			subPath := strings.TrimSpace(
				strings.TrimPrefix(
					strings.TrimPrefix(
						strings.TrimPrefix(path, dir),
						"\\"),
					"/"))

			if len(subPath) > 0 {
				buffer.WriteString(subPath)
				buffer.WriteString("\n")

				if subPath == IndexFileName {
					hasIdx = true
				}
			}
			return nil
		})

	if err != nil {
		fmt.Println(err.Error())
	}

	return strings.TrimSpace(buffer.String()), hasIdx
}
