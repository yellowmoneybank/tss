package shamir

import (
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ShamirReconstruct(t *testing.T) {
	{
		secret := []byte{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}

		splits, _ := SplitSecret(secret, 3, 3)
		reconstruct, _ := Reconstruct(splits)
		assert.Equal(t, secret, reconstruct)
	}

	{
		secret, _ := ioutil.ReadFile("./sample.txt")
		splits, _ := SplitSecret(secret, 5, 3)
		reconstruct, _ := Reconstruct(splits)
		assert.Equal(t, secret, reconstruct)
	}
}

