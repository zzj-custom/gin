package task

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"sync"
)

var (
	taskPool     *ants.Pool
	taskPoolOnce sync.Once
)

func Pool() *ants.Pool {
	taskPoolOnce.Do(func() {
		pool, err := ants.NewPool(150)
		if err != nil {
			fmt.Printf("创建ant_pool失败，err: %v", err)
			return
		}
		taskPool = pool
	})
	return taskPool
}
