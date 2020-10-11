package cheaterDetection

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"
	"moritzm-mueller.de/tss/pkg/shamir"
)

func TestIsConsistent(t *testing.T) {
	secret := "much secret, very secure"

	var threshold uint8 = 5

	numberOfShares := 100

	{
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}

		antiCheat, _ := CalculateCheckValue(shares)
		assert.Empty(t, IsConsistent(shares, antiCheat))
	}
	{
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}
		antiCheat, _ := CalculateCheckValue(shares)

		shares[3].Slices = []uint16{1}

		shares[4].Slices = []uint16{1}

		shares[5].Slices = []uint16{1}

		assert.Equal(t, []int{int(shares[3].ShareIndex), int(shares[4].ShareIndex), int(shares[5].ShareIndex)}, IsConsistent(shares, antiCheat))
	}
}

func Test_isValidShare(t *testing.T) {
	{
		antiCheat := AntiCheat{
			T: *big.NewInt(10),
			P: *big.NewInt(13),
		}

		tNew := big.NewInt(7)

		assert.False(t, isValidShare(antiCheat, tNew, 1))
	}
	{
		secret := "1"
		var threshold uint8 = 5
		numberOfShares := 7
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}

		antiCheat, _ := CalculateCheckValue(shares)

		tNew := calcSum1(calcHashes(shares), &antiCheat.P)

		assert.True(t, isValidShare(antiCheat, tNew, 7))
	}
	{
		validShare := isValidShare(AntiCheat{
			T: *big.NewInt(19037),
			P: *big.NewInt(11),
		}, big.NewInt(14648), 1)
		assert.True(t, validShare)
	}
	{
		secret := "much secret, very secure"
		var threshold uint8 = 5
		numberOfShares := 10
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}

		antiCheat, _ := CalculateCheckValue(shares)

		// manipulate share...
		shares[0].Slices = []uint16{1, 23, 4}
		tNew, _ := calcTNew(shares, &antiCheat.P, calcHashes(shares))

		for i, share := range shares {
			if i == 0 {
				assert.False(t, isValidShare(antiCheat, tNew, int(share.ShareIndex)))
			}
			assert.True(t, isValidShare(antiCheat, tNew, int(share.ShareIndex)))
		}
	}
}

func Test_calcTNew(t *testing.T) {
	shares := []shamir.Share{
		{
			ID:         [16]byte{},
			ShareIndex: 1,
			Slices:     []uint16{2},
		},
		{
			ID:         [16]byte{},
			ShareIndex: 3,
			Slices:     []uint16{8},
		},
	}
	hashes := map[int]big.Int{
		1: *big.NewInt(7),
		3: *big.NewInt(1),
	}

	tnew, _ := calcTNew(shares, big.NewInt(11), hashes)

	assert.Equal(t, big.NewInt(14648), tnew)
}
