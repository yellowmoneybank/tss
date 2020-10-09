package cheaterDetection

import (
	"math/big"
	"sort"

	"github.com/google/uuid"
	"moritzm-mueller.de/tss/pkg/shamir"
)

// returns the uuid of the cheater-shares. If the returning slice is nil, all shares are consistent.
func IsConsistent(shares []shamir.Share, t AntiCheat) []uuid.UUID {
	prime := (new(big.Int)).Set(&t.P)

	TNew := calcSum1(extractSecrets(shares), prime)

	var cheaters []uuid.UUID

	// sort by index
	sort.Slice(shares, func(i, j int) bool { return shares[i].ShareIndex < shares[j].ShareIndex })

	for i, share := range shares {
		if !isValidShare(t, TNew, i) {
			cheaters = append(cheaters, share.Id)
		}
	}

	return cheaters
}

func isValidShare(antiCheat AntiCheat, tNew *big.Int, index int) bool {
	T := (new(big.Int)).Set(&antiCheat.T)
	T.Sub(T, tNew)

	divisor := (new(big.Int)).Set(&antiCheat.P)
	divisor.Exp(divisor, big.NewInt(int64(2*index)), nil)

	T.DivMod(T, divisor, &antiCheat.P)

	return T.Int64() == 0
}
