package xsync

import (
	"go-speed/global"
	"sync"
)

var WaitGroup waitGroupWrapper

var Quit bool
var Wait *sync.WaitGroup

func init() {
	Wait = new(sync.WaitGroup)
}

type waitGroupWrapper struct {
}

//同步执行
func (w waitGroupWrapper) Wrap(handler func(params ...interface{}), params ...interface{}) {
	Wait.Add(1)
	defer Wait.Done()
	defer func() {
		if err := recover(); err != nil {
			global.Logger.Error().Msgf("WaitGroupWrapper panic", err)
		}
	}()

	handler(params...)
}

//启动协程执行
func (w waitGroupWrapper) WrapGoroutine(handler func(params ...interface{}), params ...interface{}) {
	Wait.Add(1)

	go func() {
		defer Wait.Done()
		defer func() {
			if err := recover(); err != nil {
				global.Logger.Error().Msgf("WaitGroupWrapper panic", err)
			}
		}()

		handler(params...)
	}()
}
