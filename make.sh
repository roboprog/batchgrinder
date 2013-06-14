#!/bin/sh -x

# Brute force build for batch library project.
# Feel free to use make, or ant, or maven
# as appropriate for the language in which you choose to work.
# However, this works, even on a Mac, if not on Windows.
# I'm not writing for Windows.

# This script assumes that Go and Groovy compilers are on the PATH.

( cd examples/go/bin ; \
	go build ../src/ex1.go )

( cd libs/groovy/src ; \
	find . -name '*.groovy' | xargs groovyc -d ../../../examples/groovy/bin  )
( cd examples/groovy/src ; \
	find . -name '*.groovy' | xargs groovyc -cp . -d ../bin  )


# vi: ts=4 sw=4
# *** EOF ***
