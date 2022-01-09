package util

// 简单的自定义set结构
type Set map[string]struct{}

func (s Set) Push(keys ...string) {
	for _, key := range keys {
		s[key] = struct{}{}
	}
}

func (s Set) Pop(key string) string {
	delete(s, key)
	return key
}

func (s Set) Contain(key string) bool {
	for k, _ := range s {
		if key == k {
			return true
		}
	}
	return false
}

func (s Set) List() (li []string) {
	for k, _ := range s {
		li = append(li, k)
	}
	return
}
