# 15. Function 2

Go에서 함수는 __일급 객체(First-class Citizen)__ 이다. 즉, Go에서는 함수를 변수에 할당할 수 있으며, 함수를 함수의 인자로써 넘길 수 있고, 함수가 함수를 리턴할 수 있다. 일급 객체에 대한 자세한 내용은 다음 블로그 글을 참고하자 : [1급 객체란?](https://medium.com/@lazysoul/functional-programming-%EC%97%90%EC%84%9C-1%EA%B8%89-%EA%B0%9D%EC%B2%B4%EB%9E%80-ba1aeb048059)

위의 사실을 기억하고 관련된 내용을 하나씩 살펴보자.

### Function literals

함수 리터럴(Function literals)은 __익명 함수__ 라고도 부른다. 말그대로 이름이 없는 함수이다.\
두 integer를 더하는 함수를 정의해보자.

```go
func add(a, b int) int {
    return a + b
}
```

이제 아래와 같이 `add` 함수를 호출할 수 있다.

```go
result := add(10, 20)   // result : 30
```

이제 `add`와 같은 기능을 하는 함수 리터럴을 만들어 보자. 단순히 함수의 이름을 생략해주면 된다.

```go
func(a, b int) int {
    return a + b
}
```

그러나 이 상태로는 아무 것도 할 수 없다. 익명 함수는 이름이 없으므로 호출이 불가능하기 때문이다.\
그러나 함수는 _일급 객체_ 이므로, 다음과 같이 익명 함수를 사용할 수 있다.

```go
anonymous_add := func(a, b int) int {
    return a+b
}
result := anonymous_add(10, 20)
```

두 정수를 더하는 익명 함수를 정의하여 `anonymous_add`에 할당해주고 있다. 이후부터 `anonymous_add`는 그 자체로 함수이자 변수가 되어, 호출할 수 있으며 인자로써 다른 함수로 입력될 수 있다.

### Higher-order Function

__Higher-order Function(고차 함수)__ 은 함수형 프로그래밍 패러다임에서 나온 개념으로, __함수를 파라미터로 전달받거나, 함수를 리턴하는 함수__ 를 의미한다. 이러한 코드 작성이 가능하려면 기본적으로 함수가 First-class Citizen 이어야 하며, 따라서 Go 코드는 Higher-order Function을 구현 가능하다.

먼저 함수를 리턴하는 함수의 경우를 보자. 아래의 함수는 성과 이름을 각각 받아서 합친 `string`을 출력한다.

```go
func append_two_strings(str1 string) func(string) string {
    return func(str2 string) string {
        return str1 + " " + str2
    }
}

func main() {
    first_name := "Soomin"
    last_name := "Lee"
    name := append_two_strings(first_name)(last_name)   // name : "Soomin Lee"
}
```

`append_two_strings` 함수는 `string` 하나를 입력받는다. 함수의 반환형이 `func(string) string {...}` 이므로, 또 다른 `string`을 입력받는 _함수를 리턴_ 한다. 따라서 `append_two_strings(first_name)` 가 리턴한 함수를 다시 호출할 수 있다.\
그러므로 아래와 같이 작성할 수도 있다.

```go
appender := append_two_strings(first_name)
name := appender(last_name)
fmt.Println(name)
```

이제 함수를 인자로 받는 함수의 경우를 보자. 아래의 함수는 입력된 값의 세제곱을 반환한다.

```go
func cubic(n int) int {
    return n * n * n
}

func get_cubic(calculator func(int) int, n int) int {
    return calculator(n)
}

func main() {
    result := get_cubic(cubic, 3)
    fmt.Println(result)
}
```

`get_cubic`은 `func(int) int` 모양의 함수를 인자로 받고, 내부적으로 `calculator`라는 이름으로 사용한다. `get_cubic`의 첫 번째 인자에는 `func(int) int` 형의 어떤 함수라도 올 수 있다. 입력 값의 세제곱을 반환하는 코드이므로, calculator 인자에는 입력된 `int` 값의 세제곱을 반환하는 `cubic` 함수를 전달했다. 따라서 `get_cubic`은 `cubic(n)`의 결과를 리턴하게 된다.

### Higher-order function의 고차원 추상화



[여기 참고](https://www.golangprograms.com/higher-order-functions-in-golang.html)
