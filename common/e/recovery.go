package e

import (
	"context"
	"gmimo/common/log"
)

// 提供给子协程的recovery
func Recovery() {
	if r := recover(); r != nil {
		log.Errorc(context.Background(), "panic:", r)
	}
}
