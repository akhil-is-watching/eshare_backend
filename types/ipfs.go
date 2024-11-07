package types

type DocumentPackage struct {
	EncryptedData EncryptedData `json:"encryptedData"`
	MerkleProof   MerkleProof   `json:"merkleProof"`
	Metadata      Metadata      `json:"metadata"`
}

type EncryptedData struct {
	Ciphertext  string `json:"ciphertext"`
	IV          string `json:"iv"`
	OwnerPubKey string `json:"ownerPubKey"`
}

type MerkleProof struct {
	Proof       []string `json:"proof"`
	RootDocHash string   `json:"rootDocHash"`
	LeafHash    string   `json:"leafHash"`
}

type Metadata struct {
	Timestamp int64  `json:"timestamp"`
	Algorithm string `json:"algorithm"`
	FileName  string `json:"fileName"`
	Version   string `json:"version"`
}
