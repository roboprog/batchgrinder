package main

import (
	"fmt"

	"./com/roboprogs/batchgrinder/grinder"
)

// simulate a batch job, using framework to pipeline steps
func main () {

	loader := grinder.Loader {}
	transformer := grinder.Transformer {}
	dumper := grinder.Dumper {}

	// pretend to read input
	loader.Unit = func(
			num int) interface {} {
		if num > 6 {
			return nil
		}

		in_rec := fmt.Sprintf( "In rec %d", num)
		return in_rec
	}

	// pretend to process input and create output
	transformer.Unit = func (
			in_data interface {},
			num int) interface {} {
		in_rec := in_data.( string)
		return in_rec + " X"
	}

	// pretend to write output
	dumper.Unit = func (
			out_data interface {},
			num int) {
		out_rec := out_data.( string)
		fmt.Printf( "%d) %s\n", num, out_rec)
	}

	grinder.Run( loader, transformer, dumper)
}


// vi: ts=4 sw=4
// *** EOF ****
