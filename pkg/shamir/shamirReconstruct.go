package shamir

func ShamirReconstruct(shares []share) []byte {
	// TODO Assertions...
	var secret []byte

	for i := 0; i < len(shares[0].slices); i++ {
		var byteShares []singleByteShare
		for _, share := range shares {
			byteShares := append(byteShares,
				singleByteShare{
					shareIndex: share.shareIndex,
					share:      share.slices[i],
				})
			reconstructedByte := reconstructByte(byteShares)
			secret := append(secret, reconstructedByte)
		}
	}

	return secret
}

func reconstructByte(byteShares []singleByteShare) byte {
	// TODO
	return nil
}
