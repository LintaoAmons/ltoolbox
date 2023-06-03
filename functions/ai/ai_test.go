package ai

import "testing"

func TestAskChatGPT(t *testing.T) {
	a := AskChatGPT("Test Test", "token")
	println(a)
}
