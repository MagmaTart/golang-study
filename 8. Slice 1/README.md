# 8. Slice 1

슬라이스는 C++의 `std::vector`와 비슷하다. 기본적으로 배열을 기반으로 만들어졌지만, 배열과 다르게 크기가 가변적으로 변할 수 있는 구조다.\
슬라이스는 _길이(length)_ 와 _용량(capacity)_ 을 가지고 동작한다.

### 슬라이스 선언

슬라이스의 선언은 아래와 같이, 배열의 길이 파트를 비워서 할 수 있다.

```go 
var slice []int
```

또한 슬라이스의 초기 길이를 정하고 싶을 경우, `make` 함수를 이용할 수 있다.

```go
slice := make([]int, 3)
```

이제 슬라이스의 길이만큼 배열처럼 사용할 수 있다.

```go
func slice_test1() {
    slice := make([]int, 3)
    slice[0] = 5
    slice[1] = 6
    slice[2] = 7
    for i := 0; i < len(slice); i++ {
        fmt.Println(slice[i])
    }
}
```

### 슬라이싱

파이썬에서의 리스트와 같이, Go의 슬라이스도 일정 부분만 떼어 참조하는 __슬라이싱__ 을 할 수 있다. 문법 또한 파이썬의 그것과 동일하다.

```go
func slice_test2() {
    var big []int{1, 2, 3, 4, 5}
    small := slice[1:4]
    for i := 0; i < len(big); i++ {
        fmt.Printf("%d ", big[i])
    }
    fmt.Println()
    for i := 0; i < len(small); i++ {
        fmt.Printf("%d ", small[i])
    }
}
```

```
1 2 3 4 5 
2 3 4
```

파이썬에서의 슬라이싱과 다른 점은, 음수 인덱스를 쓸 수 없다는 점이다.

### 요소 Append

슬라이스에 요소를 가변적으로 Append할 수 있다. string 슬라이스를 만들고 새로운 값들을 추가해보자.

```go
func print_string_slice(slice []string) {
    for i := 0; i < len(slice); i++ {
        fmt.Printf("%s ", slice[i])
    }
    fmt.Println()
}

func slice_test3() {
    strs := []string{"AAA", "BBB", "CCC"}
    print_string_slice(strs)
    strs = append(strs, "DDD", "EEE")
    print_string_slice(strs)
}
```

string 슬라이스를 초기화한 후, `append` 함수로 슬라이스에 새로운 문자열들을 추가했다.\
`append` 함수는 슬라이스와 그에 추가할 요소들을 가변인자로 입력받은 후, 모든 인자가 따라붙은 슬라이스를 반환한다.\
중요한 점은, 슬라이스의 자료형과 같은 자료형의 변수만 append할 수 있다는 점이다. `[]string` 슬라이스에는 `string` 변수만 붙일 수 있다.

슬라이스에 다른 슬라이스를 append하고 싶을 경우, 이어붙일 슬라이스를 각 요소로 모두 분해한 후 차례대로 모두 append해주어야 한다.\
아래와 같이 슬라이스에 다른 슬라이스를 이어붙일 수 있다.

```go
func slice_test4() {
    strs := []string{"AAA", "BBB", "CCC"}
    new := []string{"FFF", "GGG"}
    print_string_slice(strs)
    strs = append(strs, new...)
    print_string_slice(strs)
}
```

`append`의 두 번째 인자에 슬라이스가 들어갈 수 없으므로, 이어붙일 슬라이스를 분해해야 한다.\
위의 코드에서 보는 것과 같이, `...` 구문을 사용하여 슬라이스를 모두 분해한 것처럼 사용할 수 있다. `...` 구문은 마치 슬라이스의 모든 요소를 가변 인자로 나열한 것처럼 사용할 수 있게 해준다.
