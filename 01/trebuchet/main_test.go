package trebuchet

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestProcessOneLine(t *testing.T) {
	got := processInput(
		[]string{"one2three4five6seven"},
	)
	assert.Equal(t, 26, got)
}

func TestProcessTwoLines(t *testing.T) {
	got := processInput(
		[]string{
			"one2three4five6seven",
			"six6six",
		},
	)
	assert.Equal(t, 92, got)
}

func TestFullInput(t *testing.T) {
	main()
}
