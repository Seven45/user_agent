package user_agent

import "testing"

func TestGeneratingUA(t *testing.T) {
	print(generateUserAgent("opera", "linux") + " -> ")
}
