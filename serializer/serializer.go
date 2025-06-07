package serializer

import "errors"

func Serialize(payload string) error {
	switch payload[0] {
	case '+':
	case '-':
	case ':':
	case '$':
	case '*':

	default:
		return errors.New("invalid data type determination")
	}

	return nil
}
