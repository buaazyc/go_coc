package goroutine

import (
	"log"
	"runtime"
)

// GoWithRecover 启动一个协程, 处理函数为handle, 自动recover
func GoWithRecover(handle func()) {
	go func() {
		defer CatchRecover(handle)
		handle()
	}()
}

// CatchRecover panic捕捉
func CatchRecover(handle func()) {
	if err := recover(); err != nil {
		buf := make([]byte, 2048)
		n := runtime.Stack(buf, false)
		stackInfo := string(buf[:n])
		log.Printf("exec panic error: %v , error: %v", stackInfo, err)
		handle()
	}
}
