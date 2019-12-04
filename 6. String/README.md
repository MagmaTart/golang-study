# 6. String

Go에서 문자열은 __read-only__ 인 `string` 자료형으로 다루어진다.\
기본적으로 문자열은 `byte`의 나열이다. 따라서 `string`은 `[]byte`로 캐스팅할 수 있다.\
Go의 모든 소스 코드 및 문자열은 UTF-8로 인코딩되며, 따라서 Go의 문자열 처리 방법도 모두 유니코드를 기반으로 한다.\
실제로 유니코드를 사용하는지 확인해보자.

```go
func unicode_test() {
	var str string = "안녕하세요"
	for idx, ch := range str {
		fmt.Println(idx, ch)
	}
}
```

실행 결과는 아래와 같았다.

```
0 50504
3 45397
6 54616
9 49464
12 50836
```

ch의 값으로 보아, 각 문자는 유니코드로 저장되는 것을 확인할 수 있다. `50504`는 hex로 `C548`이고, 유니코드의 `U+C548` 문자는 `안` 이다. 나머지 문자들도 같은 원리로 저장되어 있다.

`string`을 순회하는 for문을 돌 때는, 인덱스에 _UTF-8 바이트 위치_ 가 대입된다. '안녕하세요' 라는 문자열의 경우 유니코드로 문자열을 표현하는 데에 글자당 3바이트씩 필요하기 때문에, `idx`가 3씩 증가한다.\
만약 `string` 변수에 영어만 저장되어 있다면 `idx`는 1씩 증가할 것이다. 하나의 문자를 1바이트로 처리할 수 있기 때문이다.\
이를 통해 Go는 문자열에 문자들을 저장할 때, 문자 하나 당 몇 바이트를 할당하여 표현할 것인지 유동적으로 결정함을 알 수 있다.\
위 반복문에서의 `ch`와 같이 문자열에서 각 문자에 접근하면, 해당하는 유니코드 바이트를 `rune` 타입으로 참조한다.

`rune` 타입 변수는 `uint8`과 동일하므로, 각 문자가 위와 같이 숫자로 출력된다. 문자열의 각 문자를 출력하고 싶으면 `rune` 변수를 `string`으로 치환해주면 될 것이다.

```go
func str_print_test() {
	var str string = "스트링"
	for idx, ch := range str {
		fmt.Println(idx, string(ch))
	}
}
```

### 문자열의 바이트 레벨 처리

문자열에 대해 range로 루프를 돌면, 위에서 보았던 것과 같이 각 글자를 `rune` 타입으로 참조한다. 그러나 일반 배열을 돌듯이 루프를 돌면, 바이트 단위로 참조할 수 있다.

```go
func str_loop_test() {
	var str string = "퇴근"
	for i := 0; i < len(str); i++ {
		fmt.Printf("%d : %x\n", i, str[i])
	}
}
```

실행 결과는 아래와 같다.
```
0 : ed
1 : 87
2 : b4
3 : ea
4 : b7
5 : bc
```

이는 `string`이 결국 read-only인 `[]byte`이기 때문에 가능한 것이다. 따라서 `string`의 내용을 직접 수정하고 싶으면, `[]byte`로 캐스팅하여 수정한 후 다시 `string`으로 돌아오면 될 것이다.

```go
func str_modify_test() {
	var str string = "퇴근"
	b_str := []byte(str)
	b_str[2] += 5
	fmt.Println(string(b_str))
}
```

"퇴근" 이라는 문자열을 `[]byte`로 캐스팅해서, 세 번째 요소에 5를 더해준다. 현재 `str` 문자열은 하나의 문자 당 3바이트이므로, 다시 `string`으로 묶었을 때 첫번째 글자가 유니코드 순서상 5번째 다음 문자로 바뀔 것이다.\
결과는 예상대로, 문자열의 첫 글자가 유니코드 상에서 '퇴'의 5번 다음 문자인 '퇹'으로 치환된 것을 볼 수 있었다.

```
퇹근
```

### 문자열 Concatenate

문자열을 이어붙이는 작업은 Python과 같이 `+` 연산자를 이용하면 된다.

```go
func str_concat_test() {
	str1 := "테"
	str2 := "스트"
	fmt.Println(str1 + str2)
	str1 += str2
	fmt.Println(str1)
}
```

```
테스트
테스트
```