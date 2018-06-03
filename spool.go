package gomarkov

import "sync"

type spool struct {
	stringMap map[string]int
	sync.RWMutex
}

func (s *spool) add(str string) int {
	s.RLock()
	index, ok := s.stringMap[str]
	s.RUnlock()
	if ok {
		return index
	}
	s.Lock()
	defer s.Unlock()
	index, ok = s.stringMap[str]
	if ok {
		return index
	}
	index = len(s.stringMap)
	s.stringMap[str] = index
	return index
}

func (s *spool) get(str string) (int, bool) {
	index, ok := s.stringMap[str]
	return index, ok
}
