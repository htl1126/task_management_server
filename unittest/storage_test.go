package unittest

import (
	"task-manage-api/model"
	"task-manage-api/storage"
	"testing"
)

func TestInitTaskIDPool(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(s.PoolSize) // poolsize should be from dockerfile
	s.InitTaskIDPool()
	pool := s.GetTaskIDPool()

	if len(pool) != s.PoolSize {
		t.Errorf("Size of initialized task ID pool should be %d", s.PoolSize)
	}
}

func TestIsTaskIDPoolFull(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(3)
	s.InitTaskIDPool()

	if s.IsTaskIDPoolFull() != true {
		t.Errorf("IsTaskIDPoolFull() didn't run as expected")
	}
}

func TestIsTaskIDPoolEmpty1(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(0)
	s.InitTaskIDPool()

	if s.IsTaskIDPoolEmpty() != true {
		t.Errorf("IsTaskIDPoolEmpty() didn't run as expected")
	}
}

func TestIsTaskIDPoolEmpty2(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(1)
	s.InitTaskIDPool()

	_ = s.GetTaskID()
	if s.IsTaskIDPoolEmpty() != true {
		t.Errorf("IsTaskIDPoolEmpty() didn't run as expected")
	}
}

func TestGetTaskIDReturnValidID(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(3)
	s.InitTaskIDPool()

	taskID := s.GetTaskID()
	if !(0 <= taskID && taskID < 3 && s.GetTaskIDPoolSize() == 2) {
		t.Errorf("GetTaskID() didn't return a valid task ID")
	}
}

func TestGetTaskIDReturnInvalidID(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(0)
	s.InitTaskIDPool()

	taskID := s.GetTaskID()
	if taskID != -1 {
		t.Errorf("GetTaskID() didn't return an invalid task ID when the task ID pool was exhausted")
	}
}

func TestWriteToTaskPool(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(3)
	s.InitTaskIDPool()

	taskItem := model.TaskItem{
		ID:     0,
		Name:   "task 0",
		Status: model.Incompleted,
	}
	s.WriteToTaskPool(taskItem)

	taskPool := s.GetTaskPool()
	if taskPool[0] != taskItem {
		t.Errorf("WriteToTaskPool() didn't work correctly")
	}
}

func TestTaskPoolHasID1(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(3)
	s.InitTaskIDPool()

	taskItem := model.TaskItem{
		ID:     0,
		Name:   "task 0",
		Status: model.Incompleted,
	}
	s.WriteToTaskPool(taskItem)

	if s.TaskPoolHasID(0) != true {
		t.Errorf("TaskPoolHasID() didn't work correctly")
	}
}

func TestTaskPoolHasID2(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(3)
	s.InitTaskIDPool()

	taskItem := model.TaskItem{
		ID:     0,
		Name:   "task 0",
		Status: model.Incompleted,
	}
	s.WriteToTaskPool(taskItem)

	if s.TaskPoolHasID(1) != false {
		t.Errorf("TaskPoolHasID() didn't work correctly")
	}
}

func TestRecycleTaskID(t *testing.T) {
	s := storage.NewStorageInstance()
	s.SetTaskPoolSize(1)
	s.InitTaskIDPool()

	taskID := s.GetTaskID()
	oldTaskIDPoolSize := s.GetTaskIDPoolSize()
	s.RecycleTaskID(taskID)
	newTaskIDPoolSize := s.GetTaskIDPoolSize()

	if !(taskID == 0 && oldTaskIDPoolSize == 0 && newTaskIDPoolSize == 1) {
		t.Errorf("RecycleTaskID() didn't work correctly")
	}
}
