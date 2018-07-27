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
		//The maximum length of the prefix and suffix
		//which is also the offset of prefix in m.p
		offset := m.next[i-1]

		//compare current char and the next char of last max prefix
		if m.p[i] != m.p[offset] && offset > 0 {
			offset = m.next[offset-1]
		}

		//we have find a maximum prefix and suffix
		//	two situations:
		//		case 1:offset = 0 do nothing
		//		case 2:m.p[i]==m.p[offset]
		if m.p[i] == m.p[offset] {
			//obviously the length of maximum prefix and suffix should add 1
			offset++
		}

		//if offset = 0 ,no maximum prefix and suffix is found

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
