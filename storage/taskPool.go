package storage

import "task-manage-api/model"

// in memory storage
var StorageMgr Storage

type Storage struct {
	PoolSize   int
	TaskIDPool map[int]struct{}
	TaskItems  map[int]model.TaskItem
}

func NewStorageInstance() Storage {
	return Storage{}
}

func (s *Storage) SetTaskPoolSize(size int) {
	s.PoolSize = size
}

func (s *Storage) InitTaskIDPool() {
	s.TaskIDPool = make(map[int]struct{}, s.PoolSize)
	s.TaskItems = make(map[int]model.TaskItem)

	for i := 0; i < s.PoolSize; i += 1 {
		s.TaskIDPool[i] = struct{}{}
	}
}

func (s Storage) GetTaskIDPool() map[int]struct{} {
	return s.TaskIDPool
}

func (s Storage) GetTaskItems() map[int]model.TaskItem {
	return s.TaskItems
}

func (s Storage) GetTaskIDPoolSize() int {
	return len(s.TaskIDPool)
}

func (s Storage) IsTaskIDPoolFull() bool {
	return len(s.TaskIDPool) == s.PoolSize
}

func (s Storage) IsTaskIDPoolEmpty() bool {
	return len(s.TaskIDPool) == 0
}

func (s *Storage) GetTaskID() int {
	var taskID int

	if s.IsTaskIDPoolEmpty() {
		return -1
	}

	for i := range s.TaskIDPool {
		taskID = i
		break
	}

	delete(s.TaskIDPool, taskID)
	return taskID
}

func (s *Storage) WriteToTaskPool(item model.TaskItem) { // need to check if id has already existed?
	s.TaskItems[item.ID] = item // need to test?
}

func (s Storage) TaskPoolHasID(id int) bool { // need to test
	if _, ok := s.TaskItems[id]; ok {
		return true
	}
	return false
}

func (s Storage) GetTaskItemByID(id int) model.TaskItem {
	return s.TaskItems[id]
}
