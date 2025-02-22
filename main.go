package main

import (
	"flag"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"

	"github.com/mkch/writeclip/clipboard"
)

const VERSION = "0.2"

func main() {
	startLine := flag.Int("start", 1, "The starting line number. -1 for not adding line numbers.")
	trim := flag.Bool("trim", true, "Trim the common leading spaces of each line.")
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
	str = format(str, *startLine, *trim)
	clipboard.SetText(str)
}

// format formats str.
// startLineNumber is the starting number used as the first
// line number. -1 means no line number.
// If trimCommonLeading is true, the common leading spaces
// in each line will be trimmed.
func format(str string, startLineNumber int, trim bool) string {
	str = strings.ReplaceAll(str, "\r\n", "\n")
	str = strings.ReplaceAll(str, "\t", "    ")
	lines := strings.Split(str, "\n")
	if trim {
		trimLeadingSpaces(lines)
	}
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

// trimLeadingSpaces trims the common leading spaces in each line.
func trimLeadingSpaces(lines []string) {
	if len(lines) == 0 {
		return
	}
	var common = math.MaxInt
	for _, line := range lines {
		n := len(line) - len(strings.TrimLeft(line, " "))
		common = min(n, common)
	}
	if common == 0 {
		return
	}
	for i, line := range lines {
		lines[i] = line[common:]
	}
}

// lineNumberStrLen returns the max length of line number strings.
func lineNumberStrLen(start int, lines []string) int {
	len0 := len(strconv.Itoa(start))
	lenN := len(strconv.Itoa(start + len(lines) - 1))
	return max(len0, lenN)
}
