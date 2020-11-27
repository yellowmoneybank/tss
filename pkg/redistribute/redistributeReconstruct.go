package redistribute

import (
	"errors"
	"fmt"

	"moritzm-mueller.de/tss/pkg/secretSharing"
	"moritzm-mueller.de/tss/pkg/vsr"
)

func Reconstruct(redistShares []secretSharing.RedistShare) (secretSharing.Share, error) {
	err := sanityCheck(redistShares)
	if err != nil {
		return secretSharing.Share{}, err
	}
	if vsr.ValidVSRShare(redistShares) == false {
		// 	// //	return secretSharing.Share{}, errors.New("not a valid VSRShare")
	}

	reconstructedShare := secretSharing.Share{
		ID:         [16]byte{},
		Threshold:  redistShares[0].NewThreshold,
		ShareIndex: redistShares[0].NewIndex,
		Secrets:    []secretSharing.ByteShare{},
		Prime:      redistShares[0].Share.Prime,
		Q:          redistShares[0].Share.Q,
		G:          redistShares[0].Share.G,
	}

	for i := 0; i < len(redistShares[0].Share.Secrets); i++ {
		// build points
		var points []secretSharing.Point

		for j := 0; j < len(redistShares); j++ {
			points = append(points, secretSharing.Point{
				X: int(redistShares[j].OldIndex),
				Y: int(redistShares[j].Share.Secrets[i].Share),
			})
		}

		p := secretSharing.ReconstructPolynom(points, redistShares[0].Share.Prime)

		secret := p(0)

		if secret != float64(int(secret)) {
			fmt.Println("Houston, we have a problem")
		}

		// Set checkvalues for Feldman + VSR
		reconstructedShare.Secrets = append(reconstructedShare.Secrets, secretSharing.ByteShare{
			Share:       uint16(secret),
			GS:          redistShares[0].Share.Secrets[i].GS,
			CheckValues: nil,
		})
	}

	return reconstructedShare, nil
}

func contains(slice []uint16, element uint16) bool {
	for _, a := range slice {
		if a == element {
			return true
		}
	}

	return false
}

func sanityCheck(redistShares []secretSharing.RedistShare) error {
	// check, that all shares have a different old index and the same new index
	newIndex := redistShares[0].NewIndex

	var oldIndex []uint16

	for _, redistShare := range redistShares {
		if redistShare.NewIndex != newIndex {
			return errors.New("redistShares do not have the same new index")
		}

		if contains(oldIndex, redistShare.OldIndex) {
			return errors.New("redistShares do not have a unique oldIndex")
		}

		oldIndex = append(oldIndex, redistShare.OldIndex)
	}

	// check, that all shares have the same size
	size := len(redistShares[0].Share.Secrets)
	for _, redistShare := range redistShares {
		if len(redistShare.Share.Secrets) != size {
			return errors.New("redistShares do not have the same size")
		}
	}

	return nil
}

// helper function
func RedistSharesToShareholders(redistShares [][]secretSharing.RedistShare) ([]secretSharing.Share, error) {
	orderdredistShares := make(map[uint16][]secretSharing.RedistShare)

	for _, redistShare := range redistShares {
		for _, indexShare := range redistShare {
			orderdredistShares[indexShare.NewIndex] = append(orderdredistShares[indexShare.NewIndex], indexShare)
		}
	}

	var newShares []secretSharing.Share

	for _, shares := range orderdredistShares {
		share, err := Reconstruct(shares)
		if err != nil {
			return nil, err
		}

		newShares = append(newShares, share)
	}

	return newShares, nil
}
