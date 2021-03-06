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

package com.roboprogs.batchgrinder

import groovy.util.logging.*

import com.roboprogs.batchgrinder.grinder.Callbacks
import com.roboprogs.batchgrinder.grinder.Dumper
import com.roboprogs.batchgrinder.grinder.Loader
import com.roboprogs.batchgrinder.grinder.Transformer

/**
 * Namespace for library entry point.
 */
@Log
class Grinder {

	/** accept the input configuration and start a batch job */
	public static void run(
		Callbacks callbacks) {
		// TODO:  figure out how to configure logger *from here*,
		//  without modifying prop file in JRE itself!

		// TODO:  set CPU *quota* for program
		//  (will try to eat them all, otherwise)

		log.info "BEGIN"

		procHdr callbacks.load, callbacks.transform, callbacks.dump

		procUnits callbacks.load, callbacks.transform, callbacks.dump

		procTlr callbacks.load, callbacks.transform, callbacks.dump

		log.info "DONE"
	}

	/** process header, if any */
	private static void procHdr(
			Loader load,
			Transformer transform,
			Dumper dump) {
		// TODO
	}

	/** process "units" (customers, statements, messages, whatever) */
	private static void procUnits(
			Loader load,
			Transformer transform,
			Dumper dump) {
		// TODO:  create and use queues with pipelined worker threads

		def num = 0
		for (;;) {
			num ++
			def in_data = load.unit num
			if ( in_data == null) {
				break  // === done ===
			}

			log.info "Read unit ${num}"
			def out_data = transform.unit in_data, num
			log.info "Processed unit ${num}"
			dump.unit out_data, num
			log.info "Wrote unit ${num}"
		}
	}

	/** process trailer, if any */
	private static void procTlr(
			Loader load,
			Transformer transform,
			Dumper dump) {
		// TODO
	}

}


// vi: ts=4 sw=4
// *** EOF ****
