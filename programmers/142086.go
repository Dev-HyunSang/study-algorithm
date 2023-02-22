// 가장 가까운 같은 글자
package programmers

import "sync"

func solution142086(s string) []int {
	slen := len(s)
	var answer []int
	mapp := &sync.Map{}

	for i := 0; i < slen; i++ {
		ap := s[i]
		if val, ok := mapp.Load(ap); ok {
			answer = append(answer, i-val.(int))
		} else {
			answer = append(answer, -1)
		}
		mapp.Store(ap, i)
	}

	return answer
}
