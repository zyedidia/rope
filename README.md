# Rope

[![Go Reference](https://pkg.go.dev/badge/github.com/zyedidia/rope.svg)](https://pkg.go.dev/github.com/zyedidia/rope)

An implementation of the Rope data structure, which is useful for storing large
arrays that are frequently mutated. The Rope is often used instead of a string
(array of characters), but can be used to replace any kind of array. This
implementation provides a Rope for arrays of bytes, but it should be easy to
modify it for other types of elements (this can be made generic in future
versions of Go).

See the Godoc for the available functions. Much of the implementation is based
on the Rope implementation [here](https://github.com/component/rope).
