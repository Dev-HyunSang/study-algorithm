# programmers

## 가장 가까운 글자
처음으로 불어봤습니다. [코딩테스트 - 가장 가까운 같은 글자
](https://velog.io/@qkqk2938/%EC%BD%94%EB%94%A9%ED%85%8C%EC%8A%A4%ED%8A%B8)를 참고하여 풀었습니다.
### 문제 설명
문자열 s가 주어졌을 때, s의 각 위치마다 자신보다 앞에 나왔으면서, 자신과 가장 가까운 곳에 있는 같은 글자가 어디 있는지 알고 싶습니다.  
예를 들어, s="banana"라고 할 때,  각 글자들을 왼쪽부터 오른쪽으로 읽어 나가면서 다음과 같이 진행할 수 있습니다.

b는 처음 나왔기 때문에 자신의 앞에 같은 글자가 없습니다. 이는 -1로 표현합니다.  
a는 처음 나왔기 때문에 자신의 앞에 같은 글자가 없습니다. 이는 -1로 표현합니다.  
n은 처음 나왔기 때문에 자신의 앞에 같은 글자가 없습니다. 이는 -1로 표현합니다.  
a는 자신보다 두 칸 앞에 a가 있습니다. 이는 2로 표현합니다.  
n도 자신보다 두 칸 앞에 n이 있습니다. 이는 2로 표현합니다.  
a는 자신보다 두 칸, 네 칸 앞에 a가 있습니다. 이 중 가까운 것은 두 칸 앞이고, 이는 2로 표현합니다.  
따라서 최종 결과물은 [-1, -1, -1, 2, 2, 2]가 됩니다.  

문자열 s이 주어질 때, 위와 같이 정의된 연산을 수행하는 함수 solution을 완성해주세요.

### 제한 사항
- 1 ≤ s의 길이 ≤ 10,000
  - s은 영어 소문자로만 이루어져 있습니다.

### 입 · 출력 예
- `"banana"` = `[-1, -1, -1, 2, 2, 2]`
- `"foobar"` = `[-1, -1, 1, -1, -1, -1]`


```go
func solution(s string) []int {
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
```

## 짝수와 홀수

### 문제 설명
정수 num이 짝수일 경우 "Even"을 반환하고 홀수인 경우 "Odd"를 반환하는 함수, solution을 완성해주세요.

### 제한 조건
- num은 int 범위의 정수
- 0은 짝수입니다.

| num | return |
|-----|-------|
| 3   | "Odd"|
|4 | "Even"|

```go
func solution(num int) string {
    if num % 2 == 0 {
        return "Even"
    }
    return "Odd"
}
```