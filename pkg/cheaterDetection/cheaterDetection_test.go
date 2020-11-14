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

		shares[3].Secrets = []uint16{1}
		fmt.Println(shares[3].ShareIndex)

		assert.Equal(t, []int{int(shares[3].ShareIndex)}, IsConsistent(shares, antiCheat))
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
		secret := "seeeeecret!!!!!"
		var threshold uint8 = 5
		numberOfShares := 10
		shares, err := shamir.SplitSecret([]byte(secret), numberOfShares, threshold)
		if err != nil {
			fmt.Println("ლ(ಠ_ಠ ლ)")
		}

		antiCheat, _ := CalculateCheckValue(shares)

		// manipulate share...
		shares[0].Secrets = []uint16{1, 23, 4}
		hashesNew := calcHashes(shares)
		tNew := calcTNew(&antiCheat.P, hashesNew)

		for i, share := range shares {
			if i == 0{
				assert.Falsef(t, isValidShare(antiCheat, tNew, int(share.ShareIndex)), "share with index %d is false negative", share.ShareIndex)

				continue
			}
			assert.Truef(t, isValidShare(antiCheat, tNew, int(share.ShareIndex)), "share with index %d is false positive", share.ShareIndex)
		}
	}
	{
		assert.True(t, isValidShare(AntiCheat{
			T: *big.NewInt(19037),
			P: *big.NewInt(11),
		}, big.NewInt(29289), 1))

		boolean := isValidShare(AntiCheat{
			T: *big.NewInt(19037),
			P: *big.NewInt(11),
		}, big.NewInt(29289), 3)
		assert.False(t, boolean)
	}
}

func Test_calcTNew(t *testing.T) {
	{
		hashes := map[int]big.Int{
			1: *big.NewInt(7),
			3: *big.NewInt(1),
		}

		tNew := calcTNew(big.NewInt(11), hashes)

		assert.Equal(t, big.NewInt(14648), tNew)
	}

	{
		tnew := calcTNew(big.NewInt(11), map[int]big.Int{
			1: *big.NewInt(7),
			3: *big.NewInt(2),
		})
		assert.Equal(t, big.NewInt(29289), tnew)
	}
	{
		tnew := calcTNew(big.NewInt(257), map[int]big.Int{
			1: *big.NewInt(10),
			6: *big.NewInt(12),
			3: *big.NewInt(11),
		})
		result, _ := new(big.Int).SetString("15083859530707885268837409", 10)
		assert.Equal(t, result, tnew)
	}
}
