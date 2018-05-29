package main

// A buffered channel that we can send work requests on.
var WorkQueue = make(chan WorkRequest, 100)

type WorkRequest struct {
	Content string
}
