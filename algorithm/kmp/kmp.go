package algorithm

type KMP struct {
	next []int
	p    []rune
}

func New(key string) *KMP {
	m := &KMP{
		next: make([]int, len(key)),
		p:    []rune(key),
	}

	//init
	m.build()
	return m
}

//build creates a KMP for match
func (m *KMP) build() {
	for i := 1; i < len(m.p); i++ {
		//set offset as last matches word,which means
		//letters betwwen p[0:offset] and p[i-offset:i] are the same
		offset := m.next[i-1]

		for offset > 0 && m.p[i] != m.p[offset] {
			offset = m.next[offset-1]
		}

		//if next letter are the same,offset + 1
		if m.p[i] == m.p[offset] {
			offset++
		}

		m.next[i] = offset
	}
}

//Match returns the offset
func (m *KMP) Match(text string) (int, bool) {
	b := []rune(text)
	offset := 0
	//begin to search
	for i := range b {
		//find the matched position
		if offset == len(m.p) {
			return i - offset + 1, true
		}

		//find part of matched position
		if b[i] == m.p[offset] && offset < len(m.p) {
			offset++
			//fail to search,move
		} else if b[i] != m.p[offset] && offset > 0 {
			offset = offset - m.next[offset]
		} else {
			//do nothing
		}
	}

	return -1, false
}
