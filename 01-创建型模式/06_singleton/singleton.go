package singleton

import "sync"

var singleton *Singleton
var once sync.Once

// Singleton 单例模式类
type Singleton struct{}

// GetInstance 获取单例模式对象
func GetInstance() *Singleton {
	once.Do(func() {
		singleton = new(Singleton)
	})
	return singleton
}
