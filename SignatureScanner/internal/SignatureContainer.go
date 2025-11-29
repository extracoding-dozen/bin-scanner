package internal

type SignatureContainer interface {
	FindFromPosition(source []byte, position int64) (string, error)
	LoadFromFile(filePath string) error
}
