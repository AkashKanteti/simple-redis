package serializer

import "fmt"

func deserializeString(payload string) string {
	return fmt.Sprintf("+%s\r\n", payload)
}

func deserializeError(payload string) string {
	return fmt.Sprintf("-%s\r\n", payload)
}

func deserializeInteger(payload int) string {
	return fmt.Sprintf(":%d\r\n", payload)
}
