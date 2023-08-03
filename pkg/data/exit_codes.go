package data

var ExitCodes = (func() map[int]int {
	exitCodes := map[int]int{}
	for i, v := range StatusCodes {
		exitCodes[v] = i
	}
	return exitCodes
})()
