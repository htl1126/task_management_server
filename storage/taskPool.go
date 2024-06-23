package storage

import "task-manage-api/model"

// in memory storage
var StorageMgr Storage

type Storage struct {
	PoolSize   int
	TaskIDPool map[int]struct{}
	TaskPool   map[int]model.TaskItem
}

func NewStorageInstance() Storage {
	return Storage{}
}

func (s *Storage) SetTaskPoolSize(size int) {
	s.PoolSize = size
}

func (s *Storage) InitTaskIDPool() {
	s.TaskIDPool = make(map[int]struct{}, s.PoolSize)
	s.TaskPool = make(map[int]model.TaskItem)

	for i := 0; i < s.PoolSize; i += 1 {
		s.TaskIDPool[i] = struct{}{}
	}
}

func (s Storage) GetTaskIDPool() map[int]struct{} {
	return s.TaskIDPool
}

func (s Storage) GetTaskPool() map[int]model.TaskItem {
	return s.TaskPool
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

func (s *Storage) WriteToTaskPool(item model.TaskItem) {
	s.TaskPool[item.ID] = item
}

func (s Storage) TaskPoolHasID(id int) bool {
	if _, ok := s.TaskPool[id]; ok {
		return true
	}
	return false
}

func (s Storage) GetTaskItemByID(id int) model.TaskItem {
	return s.TaskPool[id]
}

func (s *Storage) DeleteFromTaskPool(id int) {
	delete(s.TaskPool, id)
}

func (s *Storage) RecycleTaskID(id int) {
	s.TaskIDPool[id] = struct{}{}
}
