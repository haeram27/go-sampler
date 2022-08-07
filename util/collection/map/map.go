package maphelper

type Map[K comparable, V any] map[K]V

func (m Map[K, V]) Contains(key K) bool {
	_, ok := m[key]
	return ok
}

func (m Map[K, V]) GetDefault(key K, dft V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return dft
}

func (m Map[K, V]) PutAll(m2 map[K]V) {
	for k, v := range m2 {
		m[k] = v
	}
}
