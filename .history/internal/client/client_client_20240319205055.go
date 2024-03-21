package client

import "errors"

const (
	Size1Kb  = 1024
	Size1Mb  = Size1Kb * 1024
	Size10Mb = Size1Mb * 10
)

var (
	ErrNoFormat = errors.New("no video format provided")
)

// DefaultClient type to use. No reason to change but you could if you wanted to.
var DefaultClient = AndroidClient
