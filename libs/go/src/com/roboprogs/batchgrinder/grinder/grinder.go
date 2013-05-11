package grinder

import (
	"log"
)

type (
	// function pointers to fetch input for a job
	Loader struct {
		// func ptrs
	}

	// function pointers to validate and transform data for a job
	Transformer struct {
		// func ptrs
	}

	// function pointers to create output for a job
	Dumper struct {
		// func ptrs
	}
)


// Drive processing of a batch job
func Run(
		loader Loader,
		transformer Transformer,
		dumper Dumper) {
	log.SetFlags( log.Ldate | log.Lmicroseconds | log.Lshortfile)  // TODO:  external config
	log.Printf( "Read header\n")
	log.Printf( "Process header\n")
	log.Printf( "Write header\n")
	log.Printf( "Read units\n")
	log.Printf( "Process units\n")
	log.Printf( "Write units\n")
	log.Printf( "Read trailer\n")
	log.Printf( "Process trailer\n")
	log.Printf( "Write trailer\n")
	log.Printf( "DONE\n")
}


// vi: ts=4 sw=4
// *** EOF ****
