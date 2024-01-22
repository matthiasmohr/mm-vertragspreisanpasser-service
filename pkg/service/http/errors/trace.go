package errors

import (
	"runtime"
)

const (
	// Second level call means that we need the name of a function which indirectly calls withTrace.
	secondLevelTraceFunctionCall = 2
)

// withTrace returns the name of the function.
// For example, if stackFramesToSkip is equal to 2, then:
// if funcA() calls funcB(), and funcB() calls withTrace(1) => "funcB" is returned
// if funcA() calls funcB(), and funcB() calls withTrace(2) => "funcA" is returned.
func withTrace(stackFramesToSkip int) string {
	pc, _, _, ok := runtime.Caller(stackFramesToSkip)
	if !ok {
		return ""
	}

	fn := runtime.FuncForPC(pc)

	return fn.Name()
}
