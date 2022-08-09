package example

import "fmt"

func CreateMessage(name string) (message string, err error) {
	if name == "" {
		// TODO : エラーメッセージ考える
		err = fmt.Errorf("invalid name")
	}
	message = fmt.Sprintf("Hello, %s", name)
	return
}
