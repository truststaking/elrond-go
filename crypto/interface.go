package crypto

import (
	"github.com/ElrondNetwork/elrond-go-sandbox/hashing"
)

// KeyGenerator is an interface for generating different types of cryptographic keys
type KeyGenerator interface {
	GeneratePair() (PrivateKey, PublicKey)
	PrivateKeyFromByteArray(b []byte) (PrivateKey, error)
	PublicKeyFromByteArray(b []byte) (PublicKey, error)
}

// Key represents a crypto key - can be either private or public
type Key interface {
	// ToByteArray returns the byte array representation of the key
	ToByteArray() ([]byte, error)
}

// PrivateKey represents a private key that can sign data or decrypt messages encrypted with a public key
type PrivateKey interface {
	Key
	// Sign can be used to sign a message with the private key
	Sign(message []byte) ([]byte, error)
	// GeneratePublic builds a public key for the current private key
	GeneratePublic() PublicKey
}

// PublicKey can be used to encrypt messages
type PublicKey interface {
	Key
	// Verify signature represents the signed hash of the data
	Verify(data []byte, signature []byte) (bool, error)
}

// MultiSigner provides functionality for multi-signing a message
type MultiSigner interface {
	// NewMultiSiger instantiates another multiSigner of the same type
	NewMultiSiger(hasher hashing.Hasher, pubKeys []string, key PrivateKey, index uint16) (MultiSigner, error)
	// MultiSigVerifier Provides functionality for verifying a multi-signature
	MultiSigVerifier
	// CreateCommitment creates a secret commitment and the corresponding public commitment point
	CreateCommitment() (commSecret []byte, commitment []byte, err error)
	// AddCommitmentHash adds a commitment hash to the list with the specified position
	AddCommitmentHash(index uint16, commHash []byte) error
	// CommitmentHash returns the commitment hash from the list with the specified position
	CommitmentHash(index uint16) ([]byte, error)
	// SetCommitmentSecret sets the committment secret
	SetCommitmentSecret(commSecret []byte) error
	// CommitmentBitmap returns the bitmap with the set
	CommitmentBitmap() []byte
	// AddCommitment adds a commitment to the list with the specified position
	AddCommitment(index uint16, value []byte) error
	// Commitment returns the commitment from the list with the specified position
	Commitment(index uint16) ([]byte, error)
	// AggregateCommitments aggregates the list of commitments
	AggregateCommitments() ([]byte, error)
	// SetAggCommitment sets the aggregated commitment for the marked signers in bitmap
	SetAggCommitment(aggCommitment []byte, bitmap []byte) error
	// SignPartial creates a partial signature
	SignPartial() ([]byte, error)
	// SigBitmap returns the bitmap for the set partial signatures
	SigBitmap() []byte
	// AddSignPartial adds the partial signature of the signer with specified position
	AddSignPartial(index uint16, sig []byte) error
	// VerifyPartial verifies the partial signature of the signer with specified position
	VerifyPartial(index uint16, sig []byte) error
	// AggregateSigs aggregates all collected partial signatures
	AggregateSigs() ([]byte, error)
}

// MultiSigVerifier Provides functionality for verifying a multi-signature
type MultiSigVerifier interface {
	// SetMessage sets the message to be multi-signed upon
	SetMessage(msg []byte)
	// SetBitmap sets the bitmap for the participating signers
	SetSigBitmap([]byte) error
	// SetAggregatedSig sets the aggregated signature
	SetAggregatedSig([]byte) error
	// Verify verifies the aggregated signature
	Verify() error
}
