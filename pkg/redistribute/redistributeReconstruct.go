package redistribute

import (
	"errors"
	"fmt"

	"moritzm-mueller.de/tss/pkg/secretSharing"
)

func Reconstruct(redistShares []RedistShare) (secretSharing.Share, error) {
	err := sanityCheck(redistShares)
	if err != nil {
		return secretSharing.Share{}, err
	}

	reconstructedShare := secretSharing.Share{
		ID:         [16]byte{},
		Threshold:  redistShares[0].newThreshold,
		ShareIndex: redistShares[0].newIndex,
		Secrets:    []secretSharing.ByteShare{},
		Prime:      redistShares[0].share.Prime,
		Q:          redistShares[0].share.Q,
		G:          redistShares[0].share.G,
	}

	for i := 0; i < len(redistShares[0].share.Secrets); i++ {
		// build points
		var points []secretSharing.Point

		for j := 0; j < len(redistShares); j++ {
			points = append(points, secretSharing.Point{
				X: int(redistShares[j].oldIndex),
				Y: int(redistShares[j].share.Secrets[i].Share),
			})
		}

		p := secretSharing.ReconstructPolynom(points, redistShares[0].share.Prime)

		secret := p(0)

		if secret != float64(int(secret)) {
			fmt.Println("Houston, we have a problem")
		}

		reconstructedShare.Secrets = append(reconstructedShare.Secrets, secretSharing.ByteShare{
			Share: uint16(secret),
			// TODO
			CheckValues: []uint16{},
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

func sanityCheck(redistShares []RedistShare) error {
	// check, that all shares have a different old index and the same new index
	newIndex := redistShares[0].newIndex

	var oldIndex []uint16

	for _, redistShare := range redistShares {
		if redistShare.newIndex != newIndex {
			return errors.New("redistShares do not have the same new index")
		}

		if contains(oldIndex, redistShare.oldIndex) {
			return errors.New("redistShares do not have a unique oldIndex")
		}

		oldIndex = append(oldIndex, redistShare.oldIndex)
	}

	// check, that all shares have the same size
	size := len(redistShares[0].share.Secrets)
	for _, redistShare := range redistShares {
		if len(redistShare.share.Secrets) != size {
			return errors.New("redistShares do not have the same size")
		}
	}

	return nil
}
