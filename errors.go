package graph

// InvalidArgumentErr describes argument parameters that a receiving method
// considers invalid or incorrectly specified.
type InvalidArgumentErr struct {
	arg    string
	reason string
}

func (e InvalidArgumentErr) Error() string {
	return "invalid argument: " + e.arg + " (" + e.reason + ")"
}

// DirectedGraphErr describes an operation that is not supported on directed graphs.
type DirectedGraphErr struct{}

func (e DirectedGraphErr) Error() string {
	return "err: operation not supported on directed graphs"
}
