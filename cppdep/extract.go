package cppdep

import (
	"bufio"
	"io"
	"strings"
)

func processLine(line string) string {
	line = strings.TrimSpace(line)
	if strings.HasPrefix(line, "#include") {
		begin := strings.Index(line, `"`)
		if begin != -1 {
			end := strings.Index(line[begin+1:], `"`)
			if end != -1 {
				return line[begin+1 : begin+1+end]
			}
		}
	}
	return ""
}

func extractIncludes(r io.Reader) ([]string, error) {

	incPaths := []string{}

	br := bufio.NewReader(r)
	for {
		line, err := br.ReadString('\n')
		if err == io.EOF {
			return incPaths, err
		}

		path := processLine(line)
		path = path
	}

	return nil, nil

}
