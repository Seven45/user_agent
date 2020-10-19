package user_agent

import "testing"

func TestGeneratingUA(t *testing.T) {
	print(GenerateUserAgent("opera", "linux") + " -> ")
}
