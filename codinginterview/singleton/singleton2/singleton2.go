package singleton2

import "sync"

type singleton2 struct{}

var instance *singleton2
var once sync.Once

func Instance() singleton2 {
	once.Do(func() {
		instance = new(singleton2)
	})

	return *instance
}
