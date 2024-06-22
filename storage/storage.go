package storage

var (
	PoolSize   int
	TaskIDPool map[int]struct{}
)

func InitTaskIDPool() {
	PoolSize = 10000 // should be configurable in dockerfile
	TaskIDPool = make(map[int]struct{}, PoolSize)

	for i := 0; i < PoolSize; i += 1 {
		TaskIDPool[i] = struct{}{}
	}
}

func GetTaskIDPool() map[int]struct{} {
	return TaskIDPool
}
