package util

type StringSet map[string]void

func NewStringSet(s []string) StringSet {
	result := make(StringSet, len(s))
	for _, v := range s {
		result[v] = voidVar
	}
	return result
}

func EmptyStringSet() StringSet {
	return make(StringSet, 0)
}

func (s *StringSet) Add(el string) {
	(*s)[el] = voidVar
}

func (s *StringSet) Remove(el string) {
	delete(*s, el)
}

func (s *StringSet) Pop() string {
	var el string
	for v := range *s {
		el = v
		break
	}

	delete(*s, el)

	return el
}

func (s *StringSet) Contains(el string) bool {
	if _, ok := (*s)[el]; ok {
		return true
	}

	return false
}

func (s *StringSet) Intersection(s2 *StringSet) StringSet {
	result := make(StringSet, len(*s))

	for el := range *s {
		if (*s2).Contains(el) {
			result[el] = voidVar
		}
	}

	return result
}

func (s *StringSet) Difference(s2 *StringSet) StringSet {
	result := make(StringSet, len(*s))

	for el := range *s {
		if !(*s2).Contains(el) {
			result[el] = voidVar
		}
	}

	return result
}

func (s *StringSet) AsSlice() []string {
	result := make([]string, len(*s))
	i := 0
	for k := range *s {
		result[i] = k
		i++
	}

	return result
}

func StringPermutations(L []string, r int) [][]string {
	if r == 1 {
		temp := make([][]string, 0)
		for _, rr := range L {
			t := make([]string, 0)
			t = append(t, rr)
			temp = append(temp, [][]string{t}...)
		}
		return temp
	}

	res := make([][]string, 0)
	for i := 0; i < len(L); i++ {
		perms := make([]string, 0)
		perms = append(perms, L[:i]...)
		perms = append(perms, L[i+1:]...)
		for _, x := range StringPermutations(perms, r-1) {
			t := append(x, L[i])
			res = append(res, [][]string{t}...)
		}
	}
	return res
}

func MapKeysAsStringSet(m map[string]string) StringSet {
	result := make(StringSet, len(m))
	for k := range m {
		result[k] = voidVar
	}

	return result
}

func MapValuesAsStringSet(m map[string]string) StringSet {
	result := make(StringSet, len(m))
	for _, v := range m {
		result[v] = voidVar
	}

	return result
}

func Disjoint(as1 []string, as2 []string) bool {
	for _, s1 := range as1 {
		for _, s2 := range as2 {
			if s1 == s2 {
				return false
			}
		}
	}

	return true
}
