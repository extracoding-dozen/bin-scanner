package scanner

import (
	"SignatureScanner/internal"
	"SignatureScanner/internal/file_position"
	"SignatureScanner/internal/prefix_tree"
	"os"
)

type SignatureScanner struct {
	signatures internal.SignatureContainer
}

func NewSignatureScanner() *SignatureScanner {
	return &SignatureScanner{
		signatures: prefix_tree.NewTree(),
	}
}

func (sc *SignatureScanner) Load(filepath string) error {
	return sc.signatures.LoadFromFile(filepath)
}

func (sc *SignatureScanner) Scan(filepath string) []file_position.Match {
	var results []file_position.Match
	data, err := os.ReadFile(filepath)
	if err != nil {
		return results
	}

	for i := 0; i < len(data); i++ {
		foundSig, _ := sc.signatures.FindFromPosition(data, int64(i))

		if foundSig != "" {
			posPtr := file_position.NewPosition(i, foundSig)
			results = append(results, *posPtr)
		}
	}

	return results
}
