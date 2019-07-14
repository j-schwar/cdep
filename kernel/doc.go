// Package kernel defines various kernels which can be pieced together to form
// a pipeline.
//
// A kernel is defined as an asynchronous function which consumes items from an
// inbound channel, transforms them in some way, then emits new items on an
// outbound channel. The outbound channel is created and closed by the kernel
// itself.
package kernel
