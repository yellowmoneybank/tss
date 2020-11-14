package secretSharing

import "github.com/google/uuid"

type Share struct {
	ID          uuid.UUID
	Threshold   uint8
	ShareIndex  uint16
	Secrets     []ByteShare
	Prime, Q, G int
}

type ByteShare struct {
	Share uint16
	// at contains g^s, g^a_1 ... for Feldman's VSS
	CheckValues []uint16
}
