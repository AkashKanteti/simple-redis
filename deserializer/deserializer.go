package deserializer

import (
	"errors"
	"strconv"
)

const CRLF = `\r\n`

func deserialize(payload string) (interface{}, error) {
	switch payload[0] {
	case '+':
		idx, err := fetchCrlf(payload)
		if err != nil {
			return "", err
		}
		return payload[1:idx], nil
	case '-':
		idx, err := fetchCrlf(payload)
		if err != nil {
			return "", err
		}
		return payload[1:idx], nil
	case ':':
		idx, err := fetchCrlf(payload)
		if err != nil {
			return "", err
		}
		
		if payload[1:idx] == "" {
			return "", errors.New("empty int request")
		}

		return strconv.Atoi(payload[1:idx])
	case '$':
	case '*':

	default:
		return "", errors.New("invalid data type determination")
	}

	return "", nil
}

func fetchCrlf(payload string) (int, error) {
	for k := range payload {
		if k+4 > len(payload) {
			break
		}
		if payload[k:k+4] == CRLF {
			return k, nil
		}
	}
	return -1, errors.New("data doesn't contain CRLF")
}
