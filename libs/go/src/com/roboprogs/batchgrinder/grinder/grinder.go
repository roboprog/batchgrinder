package grinder

import (
	"log"
)

type (
	// Function pointers to fetch input for a job.
	// Chose not to use an interface, to allow maximum flexibility
	// for the library user  --  not all methods need to be implemented,
	// they can belong to whatever structure makes sense for your app
	// (object for doing all input, or, object for doing all header operations, for instance),
	// nor do the routines even need to be actual methods.
	Loader struct {
		// callback to get the main processing units
		Unit * func ( int) * interface {}

		// callback to get the [file] header
		Header * func () * interface {}

		// callback to get the [file] trailer
		Trailer * func () * interface {}
	}

	// function pointers to validate and transform data for a job
	Transformer struct {
		// callback to transform / validate the main processing units
		Unit * func ( * interface {}, int) * interface {}

		// callback to transform / validate the [file] header
		Header * func ( * interface {}) * interface {}

		// callback to transform / validate the [file] trailer
		Trailer * func ( * interface {}) * interface {}
	}

	// function pointers to create output for a job
	Dumper struct {
		// callback to put the main processing units
		Unit * func ( * interface {}, int)

		// callback to put the [file] header
		Header * func ( * interface {})

		// callback to put the [file] trailer
		Trailer * func ( * interface {})
	}
)


// Drive processing of a batch job
func Run(
		load Loader,
		transform Transformer,
		dump Dumper) {
	log.SetFlags( log.Ldate | log.Lmicroseconds | log.Lshortfile)  // TODO:  external config
	log.Printf( "BEGIN\n")

	proc_hdr( load, transform, dump)

	proc_units( load, transform, dump)

	proc_tlr( load, transform, dump)

	log.Printf( "DONE\n")
}

// process header, if any
func proc_hdr(
		load Loader,
		transform Transformer,
		dump Dumper) {
	in_hdr := func () * interface {} {
		if load.Header != nil {
			log.Printf( "Read header\n")
			return ( *load.Header)()
		}
		return nil
	} ()

	out_hdr := func () * interface {} {
		if transform.Header != nil {
			log.Printf( "Process header\n")
			return ( *transform.Header)( in_hdr)
		}
		return nil
	} ()

	if dump.Header != nil {
		log.Printf( "Write header\n")
		( *dump.Header)( out_hdr)
	}
}

// process "units" (customers, statements, messages, whatever)
func proc_units(
		load Loader,
		transform Transformer,
		dump Dumper) {
	loaded_units := make( chan * interface {})
	transformed_units := make( chan * interface {})

	go load_units( load, loaded_units)
	go transform_units( transform, loaded_units, transformed_units)
	go dump_units( dump, transformed_units)
	// TODO:  wait for EOF, continue
}

// load "units" (customers, statements, messages, whatever)
func load_units(
		load Loader,
		loaded_units chan * interface {}) {
	log.Printf( "Read units\n")
	num := 0
	for {
		num++
		// TODO: periodic progress message
		in_unit := ( *load.Unit)( num)
		// TODO: check for errors / EOF
		loaded_units <- in_unit
	}
}

// transform "units" (customers, statements, messages, whatever)
func transform_units(
		transform Transformer,
		loaded_units chan * interface {},
		transformed_units chan * interface {}) {
	log.Printf( "Process units\n")
	num := 0
	for {
		num++
		// TODO: periodic progress message
		in_unit := ( <- loaded_units)
		// TODO: EOF
		out_unit := ( *transform.Unit)( in_unit, num)
		// TODO: check for errors
		transformed_units <- out_unit
	}
}

// dump "units" (customers, statements, messages, whatever)
func dump_units(
		dump Dumper,
		transformed_units chan * interface {}) {
	log.Printf( "Write units\n")
	num := 0
	for {
		num++
		// TODO: periodic progress message
		out_unit := ( <- transformed_units)
		// TODO: EOF
		( *dump.Unit)( out_unit, num)  // TODO:  counter
		// TODO: check for errors
	}
}

// process trailer, if any
func proc_tlr(
		load Loader,
		transform Transformer,
		dump Dumper) {
	in_tlr := func () * interface {} {
		if load.Trailer != nil {
			log.Printf( "Read trailer\n")
			return ( *load.Trailer)()
		}
		return nil
	} ()

	out_tlr := func () * interface {} {
		if transform.Trailer != nil {
			log.Printf( "Process trailer\n")
			return ( *transform.Trailer)( in_tlr)
		}
		return nil
	} ()

	if dump.Trailer != nil {
		log.Printf( "Write trailer\n")
		( *dump.Trailer)( out_tlr)
	}
}


// vi: ts=4 sw=4
// *** EOF ****
