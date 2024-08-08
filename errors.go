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
