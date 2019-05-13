package schedule

import "sync"

const cycleCount = 1 << 31

type Poll struct {
	counter Counter
}

type Counter struct {
	count int
	mutex *sync.Mutex
}

func (counter *Counter) Inc() {
	counter.mutex.Lock()
	counter.count = (counter.count + 1) % cycleCount
	counter.mutex.Unlock()
}

func (counter *Counter) Get() int {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	ans := counter.count
	return ans
}

func (strategy *Poll) Init() {
	strategy.counter = Counter{count: 0, mutex: &sync.Mutex{}}
}

func (strategy *Poll) Choose(client string, servers []string) string {
	strategy.counter.Inc()
	n := strategy.counter.Get()
	length := len(servers)
	return servers[n%length]
}
