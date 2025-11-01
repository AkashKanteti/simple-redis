package deserializer

import (
	"errors"
	"fmt"
	"strconv"
)

const CRLF = "\r\n"

type Deserializer struct {
	Cmd    string
	result *interface{}
}

func Deserialize(payload string) (interface{}, error) {
	switch payload[0] {
	case '+', '-':
		result, err := getString(payload)
		if err != nil {
			return Deserializer{}, err
		}
		return Deserializer{
			Cmd:    result,
			result: nil,
		}, nil
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
}

func handleArray(payload string) (interface{}, error) {
	idx := 1 // for forwarding type signal

	// re-use string method
	sizeString, err := getString(payload)
	if err != nil {
		return nil, err
	}

	size, err := strconv.Atoi(sizeString)
	if err != nil {
		return nil, err
	}

	// base cases for nil and 0 strings
	if size == 0 {
		return "", nil
	} else if size == -1 {
		return nil, nil
	}

	// size and CLRF
	idx += len(sizeString) + 4

	result := make([]string, size)
	currIdx := 0
	for _ = range size {
		currElement, nextIdx, err := handleBulkString(payload, idx)
		fmt.Printf("%v\n", currElement)
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
		if k+2 > len(payload) {
			break
		}

		if payload[k:k+2] == CRLF {
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
	// re-use string method
	sizeString, err := getString(payload[idx:])
	if err != nil {
		return nil, -1, err
	}
	idx += 1

	size, err := strconv.Atoi(sizeString)
	if err != nil {
		return nil, -1, err
	}

	// base cases for nil and 0 strings
	if size == 0 {
		return "", idx + 5, nil
	} else if size == -1 {
		return nil, idx + 2, nil
	}

	// size and CLRF
	idx += len(sizeString) + 4
	result := payload[idx : idx+size]

	// string payload length
	idx += size

	// check ending CLRF
	if !checkCrlf(payload, idx) {
		return nil, -1, errors.New("invalid starting CRLF")
	}

	// return next element start index
	idx += 4
	return result, idx, nil
}

func checkCrlf(payload string, idx int) bool {
	if payload[idx:idx+4] == CRLF {
		return true
	}
	return false
}
