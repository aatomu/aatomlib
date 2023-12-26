package utils

import (
	"fmt"
	"reflect"
	"sync"
)

type ExMap struct {
	Sm sync.Map
}

// 排他的Mapを入手
func NewExMap() *ExMap {
	return &ExMap{}
}

func (m *ExMap) Write(key string, value any) {
	m.Sm.Store(key, value)
}

func (m *ExMap) Read(key string, v any) (err error) {
	rv := reflect.ValueOf(v)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return fmt.Errorf("invaild value %v", reflect.TypeOf(v))
	}
	value, ok := m.Sm.Load(key)
	if !ok {
		return fmt.Errorf("value not found")
	}
	v = &value
	return
}

func (m *ExMap) Check(key string) (ok bool) {
	_, ok = m.Sm.Load(key)
	return
}

func (m *ExMap) Delete(key string) {
	m.Sm.Delete(key)
}
