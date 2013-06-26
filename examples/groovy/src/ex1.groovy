// simulate a batch job, using framework to pipeline steps

import com.roboprogs.batchgrinder.Grinder
import com.roboprogs.batchgrinder.grinder.Callbacks
import com.roboprogs.batchgrinder.grinder.Dumper
import com.roboprogs.batchgrinder.grinder.Loader
import com.roboprogs.batchgrinder.grinder.Transformer

// TODO:  define and implement callback interfaces
def callbacks = [
	"load" : [:] as Loader,
	"transform" : [:] as Transformer,
	"dump" : [:] as Dumper,
]

Grinder.run( callbacks as Callbacks)


// vi: ts=4 sw=4
// *** EOF ****
