package deserializer

import (
	"errors"
	"strconv"
)

const CRLF = `\r\n`

func deserialize(payload string) (interface{}, error) {
	switch payload[0] {
	case '+', '-':
		return getString(payload)
	case ':':
		result, err := getString(payload)
		if err != nil {
			return nil, err
		}
		if result == "" {
			return "", errors.New("empty int request")
		}

		return strconv.Atoi(result)
	case '$':
		result, _, err := handleBulkString(payload, 0)
		return result, err
	case '*':
		return handleArray(payload)

	default:
		return "", errors.New("invalid data type determination")
	}

	return "", nil
}

func handleArray(payload string) (interface{}, error) {
	idx := 1

	if payload[idx:idx+2] == "-1" {
		return nil, nil
	}

	resultLen := int(payload[idx+1])
	if resultLen == 0 {
		return []string{}, nil
	}
	idx += 1

	// check starting CRLF
	if !checkCrlf(payload, idx) {
		return nil, errors.New("invalid starting CRLF")
	}
	idx += 4

	result := make([]string, resultLen)
	currIdx := 0
	for _ = range resultLen {
		currElement, nextIdx, err := handleBulkString(payload, idx)
		if err != nil {
			return nil, err
		}
		result[currIdx] = currElement.(string)
		currIdx += 1
		idx = nextIdx
	}
	return result, nil
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

func getString(payload string) (string, error) {
	idx, err := fetchCrlf(payload)
	if err != nil {
		return "", err
	}
	return payload[1:idx], nil
}

func handleBulkString(payload string, idx int) (interface{}, int, error) {
	idx += 1 // for forwarding type signal
	if payload[idx:idx+2] == "-1" {
		return nil, idx + 2, nil
	}

	resultLen := int(payload[idx])
	idx += 1

	resultLen += 1 // 0 based sizes

	// check starting CRLF
	if !checkCrlf(payload, idx) {
		return nil, idx, errors.New("invalid starting CRLF")
	}
	idx += 4

	if resultLen == 0 {
		return "", idx, nil
	}

	//check actual payload
	result := payload[idx+5 : idx+5+resultLen]

	return result, idx + 5 + resultLen, nil
}

func checkCrlf(payload string, idx int) bool {
	if payload[idx:idx+4] == CRLF {
		return true
	}
	return false
}
