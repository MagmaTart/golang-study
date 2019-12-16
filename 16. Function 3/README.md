# 16. Function 3

### Closure

클로저는 내부의 함수가 외부 함수의 Context에 접근하는 형태의 함수를 말한다. Go의 함수는 First-class Citizen이므로, 함수 내에 함수를 가지고 있을 수 있다. 아래와 같은 경우를 예로 들 수 있겠다.

```go
func stack_and_print() func(string) {
    var string_stack []string
    return func(str string) {
        string_stack = append(string_stack, str)
        for i := 0; i < len(string_stack); i++ {
            fmt.Printf("%s ", string_stack[i])
        }
        fmt.Println()
    }
} 

func main() {
    stacker := stack_and_print()
    stacker("Google")       // "Google"
    stacker("Apple")        // "Google Apple"
    stacker("Microsoft")    // "Google Apple Microsoft"
}
```

`stack_and_print` 함수는 `string` 슬라이스를 가지고 있으며, 내부 함수가 있고 그 함수를 반환한다.\
여기서 확인해야 할 부분은 `stack_and_print`의 내부 익명 함수가 외부의 `string_stack`을 사용하고 있다는 점이다. 내부 함수 안에서는 함수 외부에 있는 슬라이스에 새로운 문자열을 append하고 있다. 그리고 신기한 사실은, `stack_and_print` 함수를 여러번 호출하는 것처럼 보이지만 `string_stack`은 항상 같은 객체라는 것이다.

이러한 일이 가능한 이유는, `stack_and_print`와 그 내부 함수가 클로저인 구조이기 때문이다.\
`stacker := stack_and_print()` 구문 이후 `stacker`의 호출은 `stack_and_print`의 내부 함수 호출과 동일해진다. 그러나 내부 함수가 `stack_and_print`의 슬라이스를 참조하므로, Go 컴파일러는 `stack_and_print` 함수 내의 지역 변수들을 메모리에 올리고 클로저로써 따로 관리한다. 그래서 내부 함수를 여러번 호출해도 같은 슬라이스를 바라보고 있는 것이다.\
외부 함수의 Context가 함수의 스코프를 벗어나도 유지되는 것은 클로저의 주요한 특징이다.

### Iterator 및 Generator

클로저의 특성을 이용하면, Iterator 함수도 만들어낼 수 있다. 아래의 예제는 호출시마다 슬라이스의 요소를 하나씩 반환해주는 클로저를 생성한다.

```go
func slice_iterator(slc []int) func() int {
    idx := 0
    return func() int {
        ret := slc[idx]
        idx++
        idx %= len(slc)
        return ret
    }
}

func main() {
    slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
    next_element := slice_iterator(slc)
    for i := 0; i < 20; i++ {
        fmt.Println(next_element())
    }
}
```

`slice_iterator`의 내부 함수가 외부의 `idx` 변수를 참조하므로, 클로저가 된다. 따라서 `idx`는 함수를 벗어나도 사라지지 않고, 이를 이용하여 호출시마다 슬라이스의 다음 인덱스 요소를 반환하는 내부 함수를 구현하였다. 이렇게 이터레이터를 구현할 수 있다.

또한 같은 방법으로 이터레이터의 일종인 제너레이터도 만들 수 있다. 아래 코드는 1부터 10까지 더하는 코드인데, 더하기 위한 숫자는 제너레이터가 생성한다.

```go
func generator() func() int {
    num := 0
    return func() int {
	    num++
	    return num
    }
}

func main() {
    next_num := generator()
    total := 0
    for i := 0; i < 100; i++ {
        total += next_num()
    }
    fmt.Println("Sum 1 ~ 100 :", total)     // 5050
}
```

이 코드에서도 똑같이, `generator`의 내부 함수가 외부의 `num`을 참조하므로 클로저가 된다.
