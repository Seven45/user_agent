package user_agent

import (
	"math/rand"
	"strings"
	"time"
)

func generateUserAgent(browser, os string) string {
	rand.Seed(time.Now().UnixNano())

	var resultBrowser string
	switch strings.ToLower(browser) {
	case "opera":
		resultBrowser = getOpera()
	case "firefox":
		resultBrowser = getFirefox()
	case "chrome":
		resultBrowser = getChrome()
	default:
		resultBrowser = getRandomBrowser()
	}

	var resultOs string
	switch strings.ToLower(os) {
	case "windows":
		resultOs = getWindows()
	case "macos":
		resultOs = getMacOs()
	case "linux":
		resultOs = getLinux()
	default:
		resultOs = getRandomOs()
	}

	return strings.ReplaceAll(resultBrowser, "%PLAT%", resultOs)
}
