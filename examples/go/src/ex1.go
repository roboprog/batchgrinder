package main

import (
	"fmt"

	"./com/roboprogs/batchgrinder/grinder"
)

// simulate a batch job, using framework to pipeline steps
func main () {

	callbacks := grinder.Callbacks {}

	// pretend to read input
	callbacks.Load.Unit = func(
			num int) interface {} {
		if num > 6 {
			return nil
		}

		in_rec := fmt.Sprintf( "In rec %d", num)
		return in_rec
	}

	// pretend to process input and create output
	callbacks.Transform.Unit = func (
			in_data interface {},
			num int) interface {} {
		in_rec := in_data.( string)
		return in_rec + " X"
	}

	// pretend to write output
	callbacks.Dump.Unit = func (
			out_data interface {},
			num int) {
		out_rec := out_data.( string)
		fmt.Printf( "%d) %s\n", num, out_rec)
	}

	grinder.Run( callbacks)
}


// vi: ts=4 sw=4
// *** EOF ****
