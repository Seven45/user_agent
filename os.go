package user_agent

import (
	"math/rand"
	"strconv"
)

func getWindows() string {
	etc := []string{"WOW64", "Win64; x64"}
	ver := []string{"10.0", "6." + strconv.Itoa(rand.Intn(4))}
	main := "Windows NT "

	version := ver[rand.Intn(1)]

	randBool := rand.Intn(2) != 1
	if version == "10.0" || randBool {
		version += "; " + etc[rand.Intn(2)]
	}

	return main + version
}

func getMacOs() string {
	main := "Macintosh; Intel Mac OS X 10_"
	sub := strconv.Itoa(rand.Intn(5) + 10)
	if sub != "14" {
		sub += "_" + strconv.Itoa(rand.Intn(7))
	} else {
		sub += "_" + strconv.Itoa(rand.Intn(3))
	}

	return main + sub
}

func getLinux() string {
	ver := []string{"x86_64", "i686", "i686 on x86_64"}
	main := "X11; Linux "

	return main + ver[rand.Intn(len(ver))]
}

func getRandomOs() string {
	os := []func() string{getWindows, getMacOs, getLinux}
	return os[rand.Intn(len(os))]()
}
