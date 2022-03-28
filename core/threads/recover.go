package threads

import "log"

/**
 * Created by frankieci on 2022/3/28 9:37 pm
 */

func Recover(cleanups ...func()) {
	for _, cleanup := range cleanups {
		cleanup()
	}
	if p := recover(); p != nil {
		log.Println(p)
	}
}
