package unittest

import (
	"task-manage-api/storage"
	"testing"
)

func TestInitTaskIDPool(t *testing.T) {
	storage.InitTaskIDPool()
	pool := storage.GetTaskIDPool()

	if len(pool) != 10000 { // 10000 should be from dockerfile
		t.Errorf("Size of Task ID Pool should be %d", 10000)
	}
}
