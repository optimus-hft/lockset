package lockset

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

const (
	l1 = "lock1"
	l2 = "lock2"
)

func TestSet_Lock(t *testing.T) {
	s := New()

	t.Run("BasicLockUnlock", func(t *testing.T) {
		s.Lock(l1)
		s.Lock(l2)
		s.Unlock(l2)
		s.Unlock(l1)

		assert.Equal(t, 0, len(s.locks))
	})

	t.Run("ConcurrentLockUnlock", func(t *testing.T) {
		c := 0
		ch := make(chan struct{})
		go func() {
			s.Lock(l1)
			c++
			ch <- struct{}{}
			s.Unlock(l1)
		}()
		go func() {
			s.Lock(l1)
			defer s.Unlock(l1)

			c++
		}()
		go func() {
			<-ch // try to simulate acquiring a lock right when releasing it in another goroutine
			s.Lock(l1)
			defer s.Unlock(l1)

			c++
		}()

		time.Sleep(time.Second)
		s.Lock(l1)
		assert.Equal(t, 3, c)
		s.Unlock(l1)

		assert.Equal(t, 0, len(s.locks))
	})
}

func TestSet_TryLock(t *testing.T) {
	s := New()

	assert.Equal(t, true, s.TryLock(l1))
	assert.Equal(t, true, s.TryLock(l2))
	assert.Equal(t, false, s.TryLock(l1))
	assert.Equal(t, false, s.TryLock(l1))
	s.Unlock(l1)
	s.Unlock(l2)

	assert.Equal(t, true, s.TryLock(l1))
	assert.Equal(t, true, s.TryLock(l2))

	assert.Equal(t, 2, len(s.locks))

	s.Unlock(l1)
	s.Unlock(l2)

	assert.Equal(t, 0, len(s.locks))
}

func TestSet_Unlock(t *testing.T) {
	s := New()
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	s.Unlock(l1)
}
