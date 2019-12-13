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

~작성 예정~