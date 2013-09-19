go-collada
==========

A Go package for working  with Collada V1.5 (*.dae) documents.

See Collada 1.5 Specification
http://www.khronos.org/files/collada_spec_1_5.pdf

IMPORTANT
=========

Currently, not all of the schema has been completed, classes with
TODO annotated in the struct definition imply that there is content missing

TESTING
=======

The package test suit validates the correctness of the schema by importing/exporting
and comparing against an external collada document.

More complex scenes can be added to the test suite to validate completness of the implementation

KNOWN ISSUES
============

- Partially Complete Schema

Only a subset of all classes have complete definitions,
this will cause the importer to ignore xml which do not match the struct definitions

- Order of anonymous arrays

In the Node struct, a sequence of transform operations with a strict order is defined,
however due to the nature of the golang xml unmarshal decoder, this ordering is lost.

Possible solutions invlove manually parsing the node content, but this will add significant
complexity to the parsing logic, which is at the moment completely declaritive using `xml` tags
