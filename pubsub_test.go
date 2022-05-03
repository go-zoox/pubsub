package pubsub

import (
	"sync"
	"testing"
	"time"
)

func TestEventEmitter(t *testing.T) {
	e := New()
	lock := sync.Mutex{}
	count := 0
	e.Subscribe("send.notify", SubscribeFunc(func(payload any) {
		lock.Lock()
		count++
		lock.Unlock()
		t.Log("send.notify", payload)
	}))

	e.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1)
		go func() {
			e.Publish("send.notify", index)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(10 * time.Millisecond)
	if count != 10 {
		t.Error("count should be 10, but", count)
	}
}

func TestOnce(t *testing.T) {
	e := New()
	lock := sync.Mutex{}
	count := 0
	e.SubscribeOnce("test", SubscribeFunc(func(payload any) {
		lock.Lock()
		count++
		lock.Unlock()
		t.Log("test", payload)
	}))

	e.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1)
		go func() {
			e.Publish("test", index)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(10 * time.Millisecond)

	if count != 1 {
		t.Error("count should be 1, but", count)
	}
}

func TestOff(t *testing.T) {
	e := New()
	lock := sync.Mutex{}
	count := 0
	fn := SubscribeFunc(func(payload any) {
		lock.Lock()
		count++
		lock.Unlock()
		t.Log("test", payload)
	})
	unsubscribe := e.Subscribe("test", fn)

	e.Start()

	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1)
		go func() {
			e.Publish("test", index)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(10 * time.Millisecond)

	if count != 10 {
		t.Error("count should be 10")
	}

	// e.UnSubscribe("test", fn)
	unsubscribe()

	wg = &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		index := i
		wg.Add(1)
		go func() {
			e.Publish("test", index)
			wg.Done()
		}()
	}

	wg.Wait()
	time.Sleep(10 * time.Millisecond)

	if count != 10 {
		t.Error("count should be 10, but", count)
	}
}
