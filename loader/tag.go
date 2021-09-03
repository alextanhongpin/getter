package loader

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
	"unicode/utf8"
)

var (
	tagRe *regexp.Regexp
)

func init() {
	tagRe = regexp.MustCompile(`get:"([\w-,]+)"`)
}

type Tag struct {
	Name         string
	Skip         bool
	Inline       bool
	InlinePrefix string
}

func NewTag(tag string) (*Tag, error) {
	/*
		Possible tags combination:

		get:"CustomName"
		get:"CustomName,inline" // For struct
		get:"-"

	*/
	matches := tagRe.FindAllStringSubmatch(tag, -1)
	if len(matches) == 0 {
		return nil, nil
	}
	match := matches[0][1]
	if match == "-" {
		return &Tag{Skip: true}, nil
	}

	parts := strings.Split(match, ",")
	var inlinePrefix string
	var inline bool
	if len(parts) > 1 {
		if parts[1] != "inline" {
			return nil, fmt.Errorf("getter: invalid struct tag %q\nhint: use get:\"CustomName\" or get:\"CustomName,inline\" or get:\"-\"", tag)
		}
		inline = true
		if len(parts) > 2 {
			inlinePrefix = parts[2]
			initial, size := utf8.DecodeRuneInString(inlinePrefix)
			if size != 0 {
				if !unicode.IsUpper(initial) {
					return nil, fmt.Errorf("getter: struct tag `%s` contains unexported field\nhint: rename %q to %q", tag, inlinePrefix, UpperCommonInitialism(inlinePrefix))
				}
			}
		}
	}

	name := parts[0]
	initial, size := utf8.DecodeRuneInString(name)
	if size != 0 {
		if !unicode.IsUpper(initial) {
			return nil, fmt.Errorf("getter: struct tag `%s` contains unexported field\nhint: rename %q to %q", tag, name, UpperCommonInitialism(name))
		}
	}

	return &Tag{Name: name, Inline: inline, InlinePrefix: inlinePrefix}, nil
}
