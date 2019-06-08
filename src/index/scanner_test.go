package index

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestScannerScans(t *testing.T) {
	input := `
When I was,
A young boy,
My Father
`
	scanner := newScanner(&input)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	assert.Equal(t, []string{
		"When", "I", "was",
		"A", "young", "boy",
		"My", "Father",
	}, words)
}

func TestScannerScansEmpty(t *testing.T) {
	input := "\t\t\t   \t \t" + `

`
	scanner := newScanner(&input)
	words := []string{}
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	assert.Equal(t, []string{}, words)
}
