package cheaterDetection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateCheckValues(t *testing.T) {
	assert.Equal(t, nil, CalculateCheckValues(nil))
}
