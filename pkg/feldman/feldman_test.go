package feldman

import (
	"testing"

	"moritzm-mueller.de/tss/pkg/secretSharing"
	//	"io/ioutil"
	// "reflect"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSecret(t *testing.T) {
	secret := secretSharing.ByteShare{
		Share:       20,
		CheckValues: []uint16{48, 27, 28, 53},
	}

	assert.True(t, isValidSecret(secret, 3, 59, 1))
}

func TestCalculateCheckValues(t *testing.T) {
	values := CalculateCheckValues(3, 59, []int{15, 3, 12, 19})
	assert.Equal(t, []int{48, 27, 28, 53}, values)
}
