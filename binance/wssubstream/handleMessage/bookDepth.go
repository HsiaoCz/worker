package handlemessage

import "fmt"

func HandleOrderBookDepth(message map[string]any) {
	fmt.Println(message)
}
