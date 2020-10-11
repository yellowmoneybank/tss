package cheaterDetection

import (
	"crypto/sha256"
	"errors"
	"log"
	"math/big"

	"moritzm-mueller.de/tss/pkg/shamir"
)

// returns the uuid of the cheater-shares. If the returning slice is nil, all shares are consistent.
func IsConsistent(shares []shamir.Share, t AntiCheat) []int {
	prime := (new(big.Int)).Set(&t.P)

	hashes := calcHashes(shares)

	TNew, err := calcTNew(shares, prime, hashes)
	if err != nil {
		log.Println("calcTNew failed")
	}

	var cheaters []int

	for _, share := range shares {
		if !isValidShare(t, TNew, int(share.ShareIndex)) {
			cheaters = append(cheaters, int(share.ShareIndex))
		}
	}

	return cheaters
}

func isValidShare(antiCheat AntiCheat, tNew *big.Int, index int) bool {
	T := (new(big.Int)).Set(&antiCheat.T)
	T.Sub(T, tNew)

	divisor := (new(big.Int)).Set(&antiCheat.P)
	divisor.Exp(divisor, big.NewInt(int64(2*(index-1))), nil)

	T = T.Div(T, divisor)
	T.Mod(T, &antiCheat.P)

	return T.Int64() == 0
}

func calcTNew(shares []shamir.Share, prime *big.Int, hashes map[int]big.Int) (*big.Int, error) {
	if len(shares) != len(hashes) {
		return &big.Int{}, errors.New("there should be as many shares as hashes")
	}

	tNew := big.NewInt(0)

	for _, share := range shares {
		exp := big.NewInt(2 * (int64(share.ShareIndex) - 1))
		p := (new(big.Int)).Set(prime)
		p.Exp(p, exp, nil)

		hash := hashes[int(share.ShareIndex)]

		p.Mul(p, &hash)
		tNew = tNew.Add(tNew, p)
	}

	return tNew, nil
}

func calcHashes(shares []shamir.Share) map[int]big.Int {
	h := sha256.New()
	hashes := make(map[int]big.Int)

	for _, share := range shares {
		hash := h.Sum(decodeSecret(share))
		hashes[int(share.ShareIndex)] = *(new(big.Int)).SetBytes(hash)
	}

	return hashes
}
