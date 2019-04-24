package aphcollection

// UniqueFloat32 remove duplicates from float32 slice
func UniqueFloat32(sl []float32) []float32 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[float32]int)
	var a []float32
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueFloat64 remove duplicates from float64 slice
func UniqueFloat64(sl []float64) []float64 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[float64]int)
	var a []float64
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueInt remove duplicates from int slice
func UniqueInt(sl []int) []int {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[int]int)
	var a []int
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueInt16 remove duplicates from int16 slice
func UniqueInt16(sl []int16) []int16 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[int16]int)
	var a []int16
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueInt32 remove duplicates from int32 slice
func UniqueInt32(sl []int32) []int32 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[int32]int)
	var a []int32
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueInt64 remove duplicates from int64 slice
func UniqueInt64(sl []int64) []int64 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[int64]int)
	var a []int64
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueInt8 remove duplicates from int8 slice
func UniqueInt8(sl []int8) []int8 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[int8]int)
	var a []int8
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueUint remove duplicates from uint slice
func UniqueUint(sl []uint) []uint {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[uint]int)
	var a []uint
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueUint16 remove duplicates from uint16 slice
func UniqueUint16(sl []uint16) []uint16 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[uint16]int)
	var a []uint16
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueUint32 remove duplicates from uint32 slice
func UniqueUint32(sl []uint32) []uint32 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[uint32]int)
	var a []uint32
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueUint64 remove duplicates from uint64 slice
func UniqueUint64(sl []uint64) []uint64 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[uint64]int)
	var a []uint64
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueUint8 remove duplicates from uint8 slice
func UniqueUint8(sl []uint8) []uint8 {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[uint8]int)
	var a []uint8
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}

// UniqueString remove duplicates from string slice
func UniqueString(sl []string) []string {
	if len(sl) == 1 {
		return sl
	}
	m := make(map[string]int)
	var a []string
	for _, v := range sl {
		if _, ok := m[v]; !ok {
			a = append(a, v)
			m[v] = 1
		}
	}
	return a
}
