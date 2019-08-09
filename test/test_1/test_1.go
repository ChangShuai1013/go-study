package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*原写法，会造成map多线程写异常
type Ban struct {
	visitIPs map[string]time.Time
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]time.Time)}
}

func (o *Ban) visit(ip string) bool {
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = time.Now()
	return false
}*/

type Ban struct {
	visitIPs map[string]struct{}
}

func NewBan() *Ban {
	return &Ban{visitIPs: make(map[string]struct{})}
}

var mapMutex *sync.Mutex

func (o *Ban) visit(ip string) bool {
	mapMutex.Lock()
	defer mapMutex.Unlock()
	if _, ok := o.visitIPs[ip]; ok {
		return true
	}
	o.visitIPs[ip] = struct{}{}
	go o.invalidAfter3Min(ip)
	return false
}

func (o *Ban) invalidAfter3Min(ip string) {
	time.Sleep(time.Minute * 3)
	mapMutex.Lock()
	visitIPs := o.visitIPs
	delete(visitIPs, ip)
	o.visitIPs = visitIPs
	mapMutex.Unlock()
}

func main() {
	mapMutex = new(sync.Mutex)
	var success int64
	ban := NewBan()
	wg := new(sync.WaitGroup)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 100; j++ {
			/* 原写法
			go func() {
				ip := fmt.Sprint("192.168.1.%d", j)
				if !ban.visit(ip) {
					success++
				}
			}()*/
			wg.Add(1)
			ipEnd := j
			go func() {
				defer wg.Done()
				ip := fmt.Sprint("192.168.1.%d", ipEnd)
				if !ban.visit(ip) {
					atomic.AddInt64(&success, 1)
				}
			}()
		}
	}
	wg.Wait()
	fmt.Println("success:", success)
}
