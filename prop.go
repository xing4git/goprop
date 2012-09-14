/**
 * Package prop attempts to provider a similar usage with Properties in java.
 */
package prop

import (
	"os"
	"io"
	"strings"
	"bufio"
	"errors"
	"strconv"
)

/**
 * Load loads properties from propPath
 */
func Load(propPath string) (prop map[string]string, err error) {
	prop = make(map[string]string)
	file, err := os.Open(propPath)
	if err != nil {
		return
	}
	defer file.Close()

	count := 0
	reader := bufio.NewReader(file)
	for {
		count++
		linebytes, isPrefix, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				err = nil
			}
			break
		}
		if isPrefix {
			err = errors.New("contains too long line")
			break
		}
		line := string(linebytes)
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "#") {
			continue
		}
		equalNum := strings.Count(line, "=")
		if equalNum != 1 {
			err = errors.New("invalid format at line " + strconv.Itoa(equalNum))
			break
		}
		equalIndex := strings.Index(line, "=")
		key := strings.TrimSpace(string(line[0:equalIndex]))
		value := strings.TrimSpace(string(line[equalIndex+1:]))
		if key == "" || value == "" {
			err = errors.New("invalid format at line " + strconv.Itoa(equalNum))
		}
		if _, ok := prop[key]; ok {
			err = errors.New("contains duplicate key: " + key)
		}
		prop[key] = value
	}
	return
}
