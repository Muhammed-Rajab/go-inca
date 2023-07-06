package lfu

// ////////////////////////////
// LFUCache
// ////////////////////////////
type LFUCache struct {
	capacity uint32
}

func CreateLFUCacheEmpty(capacity uint32) *LFUCache {
	return &LFUCache{capacity}
}
