package cheaterDetection

import (
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"math/big"
	"sort"

	"moritzm-mueller.de/tss/pkg/shamir"
)

// some random 257-Bit prime.
const primeString = "193589079713426316252469562119038007891098842698016891089969436916793696692793"

type AntiCheat struct {
	T big.Int
	P big.Int
}

func CalculateCheckValue(shares []shamir.Share) (AntiCheat, error) {
	var antiCheat AntiCheat

	prime, ok := new(big.Int).SetString(primeString, 10)
	if !ok {
		return antiCheat, errors.New("SetString failed")
	}

	antiCheat.P = *prime

	shareSecrets := extractSecrets(shares)

	t, err := calcT(shareSecrets, prime)
	if err != nil {
		return antiCheat, err
	}

	antiCheat.T = *t

	return antiCheat, nil
}

func extractSecrets(shares []shamir.Share) [][]byte {
	shareSecrets := make([][]byte, len(shares))

	// order by index
	sort.Slice(shares, func(i, j int) bool { return shares[i].ShareIndex < shares[j].ShareIndex })

	for i := range shareSecrets {
		shareSecrets[i] = decodeSecret(shares[i])
	}

	return shareSecrets
}

func calcT(shares [][]byte, prime *big.Int) (*big.Int, error) {
	c, err := rand.Int(rand.Reader, prime)
	if err != nil {
		return &big.Int{}, err
	}

	sum1 := calcSum1(shares, prime)
	sum2 := calcSum2(c, prime, len(shares))

	return sum1.Add(sum1, sum2), nil
}

func calcSum1(shares [][]byte, prime *big.Int) *big.Int {
	sum1 := big.NewInt(0)
	h := sha256.New()

	for i := 1; i <= len(shares); i++ {
		share := shares[i-1]

		hash := (new(big.Int)).SetBytes(h.Sum(share))

		// p^(2(i-1))
		p := (new(big.Int)).Exp(prime, big.NewInt(int64(2*(i-1))), nil)

		sum1.Add(sum1, hash.Mul(hash, p))
	}

	return sum1
}

func calcSum2(randomNumber *big.Int, prime *big.Int, numberOfShares int) *big.Int {
	sum2 := big.NewInt(0)

	for i := 1; i <= numberOfShares-1; i++ {
		c := (new(big.Int)).Set(randomNumber)

		// p = 2i - 1
		p := (new(big.Int)).Exp(prime, big.NewInt(int64((2*i)-1)), nil)

		sum2.Add(sum2, c.Mul(c, p))
	}

	return sum2
}

// shamir.Share is of type []uint16, so some conversion to []byte is due.
// TODO This code is architecture dependent.
func decodeSecret(share shamir.Share) []byte {
	var secret []byte

	for _, i := range share.Slices {
		h, l := uint8(i>>8), uint8(i&0xff)
		secret = append(secret, h, l)
	}

	return secret
}
