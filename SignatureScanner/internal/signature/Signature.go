package signature

import "errors"

type Signature struct {
	sign string
}

func NewSignature(sg string) *Signature {
	return &Signature{sign: sg}
}

func (sg *Signature) Compare(other *Signature) (bool, error) {
	if other == nil || sg.sign == "" || other.sign == "" {
		return false, errors.New("Null pointer or empty strings")
	}
	if len(sg.sign) != len(other.sign) {
		return false, nil
	}
	for ind := range sg.sign {
		if sg.sign[ind] != '?' && other.sign[ind] != '?' && sg.sign[ind] != other.sign[ind] {
			return false, nil
		}
	}
	return true, nil
}

func (sg *Signature) Pack(source []string) error {
	if source == nil {
		return errors.New("Null pointer")
	}
	resSign := ""
	for _, elem := range source {
		resSign += elem
	}
	sg.sign = resSign
	return nil
}

func Unpack(textSign string) []string {
	var res []string
	for i := 0; i < len(textSign)-1; i += 2 {
		bt := string(textSign[i]) + string(textSign[i+1])
		res = append(res, bt)
	}
	return res
}
