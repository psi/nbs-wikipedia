# nbs-wikipedia

Utilities to import to MongoDB and run calculations on a Wikipedia data set.

Build with `make`. This is the first thing I've ever written in Go
beyond a "Hello, World!", so it's quite rough, particularly in terms of
figuring out code sharing between programs.

There's also a `release` target in the Makefile to bundle up a simple
tarball, which is pulled in and installed via Chef. In the real world,
release would better versioned and likely tied into CI.
