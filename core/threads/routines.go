package threads

/**
 * Created by frankieci on 2022/3/28 9:36 pm
 */

// GoFunc defines a running goroutine function
type GoFunc = func()

// GoSafe runs the given fn using another goroutine, recovers if fn panics.
func GoSafe(fn GoFunc) {
	go RunSafe(fn)
}

func RunSafe(fn GoFunc) {
	defer Recover()
	fn()
}
