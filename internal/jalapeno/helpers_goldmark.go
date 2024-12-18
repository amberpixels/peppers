package jalapeno

import (
	"bytes"

	mdtext "github.com/yuin/goldmark/text"
)

// contentFromLines returns the content of a node that is a lines holder
// each line is concatenated into a single byte slice
func contentFromLines(v interface {
	Lines() *mdtext.Segments
}, source []byte) []byte {
	lines := v.Lines()
	content := make([]byte, 0)
	for iLine := 0; iLine < lines.Len(); iLine++ {
		line := lines.At(iLine)
		content = append(content, line.Value(source)...)
	}

	content = bytes.TrimSpace(content)

	return content
}

func contentFromSegments(segments *mdtext.Segments, source []byte) []byte {
	content := make([]byte, 0)
	for i := 0; i < segments.Len(); i++ {
		s := segments.At(i)
		content = append(content, s.Value(source)...)
	}

	return content
}
