package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/mkch/writeclip/clipboard"
)

const VERSION = "0.1"

func main() {
	startLine := flag.Int("start", 1, "The starting line number. -1 for not adding line numbers.")
	flag.Parse()
	if flag.NArg() == 1 && flag.Arg(0) == "version" {
		fmt.Fprintf(os.Stdout, "writeclip %v\n", VERSION)
		return
	}

	str, err := clipboard.Text()
	if err != nil {
		log.Fatal(err)
	}
	if len(str) == 0 {
		return
	}
	str = format(str, *startLine)
	clipboard.SetText(str)
}

// format formats str.
// startLineNumber is the starting number used as the first
// line number. -1 means no line number.
func format(str string, startLineNumber int) string {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.ReplaceAll(str, "\t", "    ")
	lines := strings.Split(str, "\n")
	if startLineNumber != -1 {
		lineNumberLen := lineNumberStrLen(startLineNumber, lines)
		for i, line := range lines {
			lineNumber := fmt.Sprintf("%0"+strconv.Itoa(lineNumberLen)+"d", startLineNumber+i)
			lines[i] = lineNumber + "  " + line
		}
	}
	str = strings.Join(lines, "\r\n")
	return str
}

// lineNumberStrLen returns the max length of line number strings.
func lineNumberStrLen(start int, lines []string) int {
	len0 := len(strconv.Itoa(start))
	lenN := len(strconv.Itoa(start + len(lines) - 1))
	return max(len0, lenN)
}
