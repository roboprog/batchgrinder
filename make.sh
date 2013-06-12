#!/bin/sh -x

# Brute force build for batch library project.
# Feel free to use make, or ant, or maven
# as appropriate for the language in which you choose to work.
# However, this works, even on a Mac, if not on Windows.
# I'm not writing for Windows.

# This script assumes that Go and Groovy compilers are on the PATH.

( cd examples/go/bin ; \
	go build ../src/ex1.go )

( cd examples/groovy/src ; \
	groovyc -d ../bin ex1.groovy )


# vi: ts=4 sw=4
# *** EOF ***
