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

	name := parts[0]
	if !IsFieldExported(name) {
		return nil, fmt.Errorf("struct tag `%s` contains unexported field\nhint: rename %q to %q", tag, name, UpperCommonInitialism(name))
	}

	var inline bool
	var inlinePrefix string
	if len(parts) > 1 {
		if parts[1] != "inline" {
			return nil, fmt.Errorf("invalid struct tag %q\nhint: use get:\"CustomName\" or get:\"CustomName,inline\" or get:\"-\"", tag)
		}
		inline = true

		if len(parts) > 2 {
			inlinePrefix = parts[2]
			if !IsFieldExported(inlinePrefix) {
				return nil, fmt.Errorf("struct tag `%s` contains unexported prefix field\nhint: rename %q to %q", tag, inlinePrefix, UpperCommonInitialism(inlinePrefix))
			}
		}
	}

	return &Tag{Name: name, Inline: inline, InlinePrefix: inlinePrefix}, nil
}

func IsFieldExported(name string) bool {
	c, size := utf8.DecodeRuneInString(name)

	// Optional. True if string is empty.
	if size == 0 {
		return true
	}

	// If the string exists, it must be upper.
	return size > 0 && unicode.IsUpper(c)
}
