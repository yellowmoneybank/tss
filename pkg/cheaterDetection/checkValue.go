package cheaterDetection

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"math/big"

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

	hashes := calcHashes(shares)

	c, err := rand.Int(rand.Reader, prime)
	if err != nil {
		return antiCheat, err
	}

	t := calcT(hashes, prime, c)

	antiCheat.T = *t

	return antiCheat, nil
}

func calcT(hashes map[int]big.Int, prime, c *big.Int) *big.Int {
	sum1 := calcSum1(hashes, prime)
	sum2 := calcSum2(c, prime, len(hashes))

	return sum1.Add(sum1, sum2)
}

func calcSum1(hashes map[int]big.Int, prime *big.Int) *big.Int {
	sum1 := big.NewInt(0)

	for index, hash := range hashes {
		// p^(2(i-1))
		p := (new(big.Int)).Exp(prime, big.NewInt(int64(2*(index-1))), nil)

		sum1.Add(sum1, p.Mul(p, &hash))
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
func decodeSecret(share shamir.Share) []byte {
	var secret []byte

	for _, i := range share.Slices {
		b := make([]byte, 2)
		binary.BigEndian.PutUint16(b, i)

		secret = append(secret, b[0], b[1])
	}

	return secret
}
