// simulate a batch job, using framework to pipeline steps

import com.roboprogs.batchgrinder.Grinder
import com.roboprogs.batchgrinder.grinder.Callbacks

// TODO:  define and implement callback interfaces
def callbacks = [:]

Grinder.run( callbacks as Callbacks)


// vi: ts=4 sw=4
// *** EOF ****
