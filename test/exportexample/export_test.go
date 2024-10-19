package exportexample

// x is a variable in the package exportexample
// x is not exported, so it is not visible outside the package
// export it by renaming it to X.
var X = x

func (s sample) Y() string {
	return s.y
}
