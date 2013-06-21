/*
    BatchGrinder: batch processing framework

    Copyright (C) 2013, Robin R Anderson
    roboprog@yahoo.com
    PO 1608
    Shingle Springs, CA 95682

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU Lesser General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU Lesser General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package grinder

import (
	"log"
)

type (
	// Chose not to use interfaces, to allow maximum flexibility
	// for the library user  --  not all methods need to be implemented,
	// they can belong to whatever structure makes sense for your app
	// (an object for doing all input, or, an object for doing all header operations, for instance),
	// nor do the routines even need to be actual methods.
	//  - or -
	// "Dude, not everything's an object"

	// Functions to fetch input for a job.
	Loader struct {
		// callback to get the main processing units
		Unit func ( int) interface {}

		// callback to get the [file] header
		Header func () interface {}

		// callback to get the [file] trailer
		Trailer func () interface {}
	}

	// Functions to validate and transform data for a job
	Transformer struct {
		// callback to transform / validate the main processing units
		Unit func ( interface {}, int) interface {}

		// callback to transform / validate the [file] header
		Header func ( interface {}) interface {}

		// callback to transform / validate the [file] trailer
		Trailer func ( interface {}) interface {}
	}

	// Functions to create output for a job
	Dumper struct {
		// callback to put the main processing units
		Unit func ( interface {}, int)

		// callback to put the [file] header
		Header func ( interface {})

		// callback to put the [file] trailer
		Trailer func ( interface {})
	}

	// Job callbacks
	Callbacks struct {
		Load Loader
		Transform Transformer
		Dump Dumper
	}
)


// Drive processing of a batch job
func Run(
		callbacks Callbacks) {
	log.SetFlags( log.Ldate | log.Lmicroseconds | log.Lshortfile)  // TODO:  external config
	log.Printf( "BEGIN\n")

	// TODO:  set number of CPUs available for program

	proc_hdr( callbacks.Load, callbacks.Transform, callbacks.Dump)

	proc_units( callbacks.Load, callbacks.Transform, callbacks.Dump)

	proc_tlr( callbacks.Load, callbacks.Transform, callbacks.Dump)

	log.Printf( "DONE\n")
}

// process header, if any
func proc_hdr(
		load Loader,
		transform Transformer,
		dump Dumper) {
	in_hdr := func () interface {} {
		if load.Header != nil {
			log.Printf( "Read header\n")
			return load.Header()
		}
		return nil
	} ()

	out_hdr := func () interface {} {
		if transform.Header != nil {
			log.Printf( "Process header\n")
			return transform.Header( in_hdr)
		}
		return nil
	} ()

	if dump.Header != nil {
		log.Printf( "Write header\n")
		dump.Header( out_hdr)
	}
}

// process "units" (customers, statements, messages, whatever)
func proc_units(
		load Loader,
		transform Transformer,
		dump Dumper) {
	loaded_units := make( chan interface {}, 3)  // TODO: configure
	transformed_units := make( chan interface {}, 3)  // TODO: configure
	eof := make( chan string)

	go load_units( load, loaded_units)
	go transform_units( transform, loaded_units, transformed_units)
	// TODO:  allow multiple transformers (concurrent, not just pipelined)
	go dump_units( dump, transformed_units, eof)
	<- eof
}

// load "units" (customers, statements, messages, whatever)
func load_units(
		load Loader,
		loaded_units chan interface {}) {
	num := 0
	for {
		num++
		in_unit := load.Unit( num)
		// TODO: check for errors
		loaded_units <- in_unit
		if in_unit == nil {
			return
		}

		log.Printf( "Read unit %d\n", num)
		// TODO: periodic progress message
	}
}

// transform "units" (customers, statements, messages, whatever)
func transform_units(
		transform Transformer,
		loaded_units chan interface {},
		transformed_units chan interface {}) {
	num := 0
	for {
		num++
		in_unit := ( <- loaded_units)
		if in_unit == nil {
			transformed_units <- nil
			return
		}

		out_unit := transform.Unit( in_unit, num)
		// TODO: check for errors
		log.Printf( "Processed unit %d\n", num)
		// TODO: periodic progress message
		transformed_units <- out_unit
	}
}

// dump "units" (customers, statements, messages, whatever)
func dump_units(
		dump Dumper,
		transformed_units chan interface {},
		eof chan string) {
	num := 0
	for {
		num++
		out_unit := ( <- transformed_units)
		if out_unit == nil {
			eof <- "EOF"
			return
		}
		dump.Unit( out_unit, num)
		// TODO: check for errors
		log.Printf( "Wrote unit %d\n", num)
		// TODO: periodic progress message
	}
}

// process trailer, if any
func proc_tlr(
		load Loader,
		transform Transformer,
		dump Dumper) {
	in_tlr := func () interface {} {
		if load.Trailer != nil {
			log.Printf( "Read trailer\n")
			return load.Trailer()
		}
		return nil
	} ()

	out_tlr := func () interface {} {
		if transform.Trailer != nil {
			log.Printf( "Process trailer\n")
			return transform.Trailer( in_tlr)
		}
		return nil
	} ()

	if dump.Trailer != nil {
		log.Printf( "Write trailer\n")
		dump.Trailer( out_tlr)
	}
}


// vi: ts=4 sw=4
// *** EOF ****
