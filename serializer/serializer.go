package serializer

import "fmt"

func SerializeString(payload string) string {
	return fmt.Sprintf("+%s\r\n", payload)
}

func serializeError(payload string) string {
	return fmt.Sprintf("-%s\r\n", payload)
}

func serializeInteger(payload int) string {
	return fmt.Sprintf(":%d\r\n", payload)
}

func serializeBulkStrings(payload string) string {
	return fmt.Sprintf("$%d\r\n%s\r\n", len(payload), payload)
}

func SerializeArray(payload []string) string {
	result := fmt.Sprintf("*%d\r\n", len(payload))
	for _, item := range payload {
		result = result + serializeBulkStrings(item)
	}

	return result
}
