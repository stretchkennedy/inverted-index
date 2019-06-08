package index

import (
	"unicode"
)

type scanner struct {
	data *string
	eof bool
	start, stop int
}

// TODO: better scanner
func newScanner(data *string) *scanner {
	return &scanner{
		data: data,
	}
}

func (s *scanner) Scan() bool {
	if s.eof {
		return false
	}
	var (
		start int
		length int
		r rune
	)
	for start, r = range (*s.data)[s.stop:] {
		if unicode.IsLetter(r) {
			break
		}
	}
	for length, r = range (*s.data)[s.stop+start:] {
		if !unicode.IsLetter(r) {
			break
		}
	}
	s.start = s.stop + start
	s.stop = s.stop + start + length
	if len(*s.data) - 1 == s.stop {
		s.eof = true
	}
	return length != 0
}

func (s *scanner) Start() int {
	return s.start
}

func (s *scanner) Stop() int {
	return s.stop
}

func (s *scanner) Text() string {
	return (*s.data)[s.start:s.stop]
}
