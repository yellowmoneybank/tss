package feldman

import (
	"math/big"

	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func IsValidShare(share secretSharing.Share) bool {
	for i, secret := range share.Secrets {
		if !isValidSecret(secret, share.G, share.Q, int(share.ShareIndex)) {
			println("secret",i," is false")
			return false
		}
	}

	return true
}

func isValidSecret(secret secretSharing.ByteShare, generator int, q int, index int) bool {
	g, s := big.NewInt(int64(generator)), big.NewInt(int64(secret.Share))
	g.Exp(g, s, big.NewInt(int64(q)))

	product := big.NewInt(int64(secret.CheckValues[0]))

	for i := 1; i < len(secret.CheckValues); i++ {
		exp := big.NewInt(int64(index))
		exp.Exp(exp, big.NewInt(int64(i)), nil)

		ga := big.NewInt(int64(secret.CheckValues[i]))
		ga.Exp(ga, exp, nil)

		product.Mul(product, ga)
	}

	product.Mod(product, big.NewInt(int64(q)))
	x := product.Cmp(g)
	return x == 0
}

func CalculateCheckValues(g, q int, coefficients []int) []int {
	qBig := big.NewInt(int64(q))

	var checkValues []int
	for _, coefficient := range coefficients {
		checkValue := big.NewInt(int64(g))
		checkValue.Exp(checkValue, big.NewInt(int64(coefficient)), qBig)

		checkValues = append(checkValues, int(checkValue.Int64()))
	}

	return checkValues
}
