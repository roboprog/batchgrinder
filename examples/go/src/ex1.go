package main

import (
	"fmt"

	"./com/roboprogs/batchgrinder/grinder"
)

// simulate a batch job, using framework to pipeline steps
func main(
		) {

	// pretend to read input
	load := func(
			num int) interface {} {
		if num > 6 {
			return nil
		}

		in_rec := fmt.Sprintf( "In rec %d", num)
		return in_rec
	}

	// pretend to process input and create output
	transform := func (
			in_data interface {},
			num int) interface {} {
		in_rec := in_data.( string)
		return in_rec + " X"
	}

	// pretend to write output
	dump := func (
			out_data interface {},
			num int) {
		out_rec := out_data.( string)
		fmt.Printf( "%d) %s\n", num, out_rec)
	}

	grinder.Run(
			grinder.Loader { &load, nil, nil },
			grinder.Transformer { &transform, nil, nil },
			grinder.Dumper { &dump, nil, nil })
}


// vi: ts=4 sw=4
// *** EOF ****
