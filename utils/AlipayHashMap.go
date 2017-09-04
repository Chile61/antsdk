package utils

import (
	"fmt"
)

// AlipayHashMap 参数结构
type AlipayHashMap struct {
	mMap   map[string]string
	Length int
}

// NewAlipayHashMap 生成空的参数结构
func NewAlipayHashMap() *AlipayHashMap {
	return &AlipayHashMap{mMap: make(map[string]string), Length: 0}
}

// Put Put
func (hm *AlipayHashMap) Put(key string, value interface{}) {
	var strValue string
	if value == nil {
		strValue = ""
	} else {
		strValue = fmt.Sprintf("%v", value)
	}

	hm.mMap[key] = strValue
	hm.Length = len(hm.mMap)
}

// Get Get
func (hm *AlipayHashMap) Get(key string) string {
	return hm.mMap[key]
}

// Remove Remove
func (hm *AlipayHashMap) Remove(key string) {
	delete(hm.mMap, key)
}

// GetMap GetMap
func (hm *AlipayHashMap) GetMap() map[string]string {
	return hm.mMap
}
