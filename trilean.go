package sgdb

// trilean false = hidden, true = visible, any = both
type trilean int

const (
	False trilean = iota
	True
	Any
)

// TODO: Url parser to support t/f/any states.
