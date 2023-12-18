package cc

import "sync"

func countL(dest *int, cnt int, wg *sync.WaitGroup, lock *sync.Mutex) {
	for i := 0; i < cnt; i++ {
		lock.Lock()
		*dest += 1
		lock.Unlock()
	}
	wg.Done()
}

func ConcurrentCount(count, bots int) int {
	dest := new(int)
	wg := new(sync.WaitGroup)
	lock := new(sync.Mutex)
	perCnt, lstCnt := count/bots, count%bots
	wg.Add(bots + 1)

	go countL(dest, lstCnt, wg, lock)
	for i := 0; i < bots; i++ {
		go countL(dest, perCnt, wg, lock)
	}

	wg.Wait()
	return *dest
}
