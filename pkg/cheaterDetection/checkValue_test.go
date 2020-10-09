package cheaterDetection

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"moritzm-mueller.de/tss/pkg/shamir"
)

func Test_CalculateCheckValue(t *testing.T) {
	{
		secret := "much secret"
		var threshold uint8 = 3
		numberOfShares := 5
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}

		antiCheat, err := CalculateCheckValue(shares)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}
		fmt.Println((antiCheat.T).String())
	}
}

func Test_calcSum1(t *testing.T) {
	assert.NotEmpty(t, calcSum1([][]byte{{1, 2, 3, 4, 5}, {1, 2, 3, 4, 5}}, big.NewInt(13)))
}

func Test_calcSum2(t *testing.T) {
	{
		assert.Equal(t, int64(1050), calcSum2(big.NewInt(3), big.NewInt(7), 3).Int64())
	}
}
