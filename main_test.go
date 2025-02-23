package main

import "testing"

func TestTrimLeadingSpaces(t *testing.T) {
	lines := []string{
		"  abc",
		"    def",
		"  ghi",
	}
	trimLeadingSpaces(lines)
	if lines[0] != "abc" || lines[1] != "  def" || lines[2] != "ghi" {
		t.Fatal(lines)
	}

	lines = []string{
		"  abc",
		"",
		"    def",
		"  ghi",
	}
	trimLeadingSpaces(lines)
	if lines[0] != "abc" || lines[1] != "" || lines[2] != "  def" || lines[3] != "ghi" {
		t.Fatal(lines)
	}

	lines = []string{
		"  abc",
		"|",
		"    def",
		"  ghi",
	}
	trimLeadingSpaces(lines)
	if lines[0] != "  abc" || lines[1] != "|" || lines[2] != "    def" || lines[3] != "  ghi" {
		t.Fatal(lines)
	}
}
