package lockset

import (
	"sync"
)

type lockKey any

type lock struct {
	mu      sync.Mutex
	counter uint
}

// Set provides dynamic mutexes based on lock keys. Each key is locked and unlocked separately and does not affect other keys.
// Instead of protecting everything with a giant mutex, Different parts of code can be protected by a tiny mutex in isolation to provide more throughput and concurrency.
type Set struct {
	mu    sync.Mutex
	locks map[lockKey]*lock
}

// Lock can be used to acquire a lock. If mutex is already locked, it will block the caller until mutex becomes unlocked.
func (s *Set) Lock(name lockKey) {
	s.mu.Lock()
	l, ok := s.locks[name]
	if !ok {
		l = &lock{}
		s.locks[name] = l
	}
	l.counter++
	s.mu.Unlock()

	l.mu.Lock()
}

// TryLock can be used to acquire a lock. If mutex is not locked, It will lock the mutex and returns true. If mutex is already locked, It will not block the caller and returns false.
func (s *Set) TryLock(name lockKey) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, ok := s.locks[name]; ok {
		return false // if lock already exists, it means it's locked by another process, we cannot lock it.
	}

	l := &lock{
		counter: uint(1),
	}
	s.locks[name] = l
	l.mu.Lock()

	return true
}

// Unlock can be used to unlock an already locked mutex. If mutex is not locked, Calling Unlock will panic.
func (s *Set) Unlock(name lockKey) {
	s.mu.Lock()
	l, ok := s.locks[name]
	if !ok {
		panic("unlocked an unlock mutex")
	}
	l.counter--
	if l.counter == 0 {
		delete(s.locks, name)
	}
	s.mu.Unlock()

	l.mu.Unlock()
}

// New creates a new lock set.
func New() *Set {
	return &Set{
		locks: make(map[lockKey]*lock),
	}
}
