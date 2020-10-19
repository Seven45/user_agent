package user_agent_local

import "testing"

func TestGeneratingUA(t *testing.T) {
	print(generateUserAgent("opera", "linux") + " -> ")
}
