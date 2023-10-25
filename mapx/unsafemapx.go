package mapx

// 非安全泛型map
type unsafeMapx[Key comparable, Value any] map[Key]Value

// 初始化
func newUnSafeMapx[Key comparable, Value any]() unsafeMapx[Key, Value] {
	return make(unsafeMapx[Key, Value])
}

// 取值
func (m unsafeMapx[Key, Value]) Get(key Key) Value {
	return m[key]
}

// 赋值
func (m unsafeMapx[Key, Value]) Set(key Key, value Value) {
	m[key] = value
}

// 删除
func (m unsafeMapx[Key, Value]) Delete(key Key) {
	delete(m, key)
}

// 长度
func (m unsafeMapx[Key, Value]) Len() int {
	return len(m)
}

// 遍历
func (m unsafeMapx[Key, Value]) Range(fn func(Key, Value)) {
	for k, v := range m {
		fn(k, v)
	}
}

// 获取所有key
func (m unsafeMapx[Key, Value]) Keys() []Key {
	keys := make([]Key, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// 是否存在
func (m unsafeMapx[Key, Value]) IsExist(key Key) bool {

	_, ok := m[key]
	return ok
}
