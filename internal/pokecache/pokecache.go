package pokecache

import (
	"sync"
	"time"
)

type Cache struct { //exposed
	CacheMap map[string]cacheEntry
	mu       sync.Mutex
}

type cacheEntry struct { // not exposed
	createdAt time.Time
	Val       []byte
}

func NewCache(interval time.Duration) *Cache { // exposed
	localcache := &Cache{
		CacheMap: make(map[string]cacheEntry),
	}
	go localcache.readLoop(interval)
	return localcache
}

func (c *Cache) Add(key string, val []byte) { // exposed
	var newEntry cacheEntry
	newEntry.Val = val
	newEntry.createdAt = time.Now()

	c.mu.Lock()
	defer c.mu.Unlock()
	c.CacheMap[key] = newEntry

}

func (c *Cache) Get(key string) ([]byte, bool) { // exposed
	c.mu.Lock()
	defer c.mu.Unlock()
	value, ok := c.CacheMap[key]
	if !ok {
		return nil, false
	}
	return value.Val, true

}

func (c *Cache) readLoop(interval time.Duration) { // not exposed

	/*
		so basically, every interval
		Loop through all the values of the map
			compare if any value needs removal & remove it

	*/

	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {

		t := <-ticker.C

		//whenever the inteval is over we iterate
		//through the whole list
		for key, value := range c.CacheMap {

			//if the value has been there more than the defined interval time
			//remove entry
			if t.Sub(value.createdAt) > interval {
				c.mu.Lock()
				delete(c.CacheMap, key)
				c.mu.Unlock()

			}

		}

	}

}
