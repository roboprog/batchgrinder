// simulate a batch job, using framework to pipeline steps

import com.roboprogs.batchgrinder.Grinder
import com.roboprogs.batchgrinder.grinder.Callbacks
import com.roboprogs.batchgrinder.grinder.Dumper
import com.roboprogs.batchgrinder.grinder.Loader
import com.roboprogs.batchgrinder.grinder.Transformer

// implement callback interfaces

// pretend to read input
def load = [ "unit" : {
		num ->
		if ( num > 6) {
			return null
		}

		return "In rec " + num
	}
]

// pretend to process input and create output
def transform = [ "unit" : {
		in_data, num ->
		return in_data + " X"
	}
]

// pretend to write output
def dump = [ "unit" : {
		out_data, num ->
		println "${num}) ${out_data}"
	}
]

def callbacks = [
	"load" : load as Loader,
	"transform" : transform as Transformer,
	"dump" : dump as Dumper,
]

// use the engine to drive the functions (closures) we defined above
Grinder.run( callbacks as Callbacks)


// vi: ts=4 sw=4
// *** EOF ****
