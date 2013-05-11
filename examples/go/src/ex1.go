package main

import (
	"./com/roboprogs/batchgrinder/grinder"
)

// TODO:  process something
func main(
		) {
	grinder.Run(
			grinder.Loader {},
			grinder.Transformer {},
			grinder.Dumper {})
}


// vi: ts=4 sw=4
// *** EOF ****
