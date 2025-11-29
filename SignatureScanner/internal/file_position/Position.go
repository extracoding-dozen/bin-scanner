package file_position

type Match struct {
	Offset    int
	signature string
}

func NewPosition(offs int, sig string) *Match {
	return &Match{Offset: offs, signature: sig}
}

func (p *Match) GetSignature() string {
	return p.signature
}
