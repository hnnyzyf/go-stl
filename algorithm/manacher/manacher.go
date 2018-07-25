package algorithm

type manacher struct {
	s rune
	d rune
}

func NewManacher(start rune, delimiter rune) *manacher {
	return &manacher{
		s: start,
		d: delimiter,
	}
}

//Matchers will return the offset and the
func (m *manacher) GetLPS(text string) string {
	b := []rune(text)

	//create s
	s := make([]rune, 2*len(b)+2)
	s[0] = m.s
	s[len(s)-1] = m.d
	for i := range b {
		s[2*i+1] = m.d
		s[2*i+2] = b[i]
	}

	//create p
	p := make([]int, len(s))
	//record the offset of LPS
	id := 0
	//mx is the offset of right board of LPS
	mx := id + p[id]
	offset := 0
	for i := 1; i < len(s); i++ {

		//there are 2 situations:
		//		case 1:i<mx p[i] = min(mx-i,p[j])
		//		case 2:i>=mx,we do nothing
		if i < mx {
			p[i] = min(mx-i, p[2*id-i])
		} else {
			p[i] = 1
		}

		//find the LPS when i is the center
		for i+p[i] < len(s) && s[i-p[i]] == s[i+p[i]] {
			p[i]++
		}

		//update mx and id if the right board of LPS of i is bigger than the right board of LPS of id
		if i+p[i] > mx {
			id = i
			mx = id + p[id]

		}

		//record
		if p[offset] < p[i] {
			offset = i
		}
	}

	temp := s[offset-p[offset]+1 : offset+p[offset]]
	substring := make([]rune, 0)
	for i := range temp {
		if temp[i] != m.d {
			substring = append(substring, temp[i])
		}
	}

	return string(substring)
}

//min returns the minum number
func min(var1 int, var2 int) int {
	if var1 < var2 {
		return var1
	} else {
		return var2
	}
}
