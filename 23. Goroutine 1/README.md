# 23. Goroutine 1

### Parallelism & Concurrency

고루틴(Goroutine)은 특정한 함수를 Go 프로세스의 다른 Context들과 독립적으로 동시에 실행시킬 수 있는 방법을 제공한다. Thread와 같은 녀석이라고 생각할 수 있으나, 동시성에 대한 고루틴의 접근 방법은 Thread와는 조금 다르다.

먼저 __Parallelism__ 은 우리말로 __병렬성__ 이라고 한다. 각각의 CPU Core를 점유한 상태로 동시에 실행되는 Multi-thread 형태의 방식을 떠올리면 되겠다. 병렬성은 Multi-Core 등 Physical level의 조건이 충족되어야만 보장될 수 있다.

반면, __Concurrency__ 는 우리말로 __병행성__ 이라고 하는데, 이 녀석은 Physical level의 조건과는 상관 없이 프로세스 내 Logical level에서의 동시 실행을 의미한다. 두 개의 루틴이 동시에 독립적으로 실행되고 있어도, 실제로는 Time-sharing 등의 방법으로 하나의 Core만 사용할 수도 있는 것이다. 이 경우 두 루틴은 병행성을 충족하지만 병렬성을 가진 것은 아니게 된다.

고루틴은 기본적으로 _병행성만을 충족한다._ 즉, 두 개 이상의 고루틴이 동시에 실행되어도 물리적으로는 하나의 Core만을 사용할 수 있는 것이다. 그러나 여러 개의 루틴은 서로 의존성이 없어서, 어떤 루틴이 먼저 실행되고 나중에 실행될 지 알 수 없다. 다시 말하면, 여러 고루틴의 동작이 서로 의존 관계에 놓이게 되면 치명적인 상황이 발생할 수 있다는 것이다.

### Goroutine의 사용

마치 C에서 정의한 함수를 Thread에 올리는 작업처럼, Go에서도 정의한 함수를 고루틴에 올릴 수 있다. 아래와 같이 `go` 키워드를 이용한다.

```go
func test() {
    for {
        fmt.Println("TEST FUNCTION")
    }
}

func main() {
    go test()
    for {
        fmt.Println("MAIN FUNCTION")
    }
}
```

`test` 함수는 그 내부에서 문자열을 무한으로 출력하는 하나의 Context를 생성한다. 또한 `main` 함수도 문자열을 출력하는 독립적인 Context를 가지고 있다. `main` 함수에서 문자열의 출력을 시작하기 전에, `go` 키워드를 사용하여 `test` 함수의 실행을 고루틴에 인가했다. 이후엔 `main` 함수와 `test` 함수의 루프가 동시에 돌아가게 된다. 병행성을 가지고 두 함수가 동시에 실행되는 것이다.

중요한 사실은, 두 함수의 실행 순서가 동률이 아니라는 점이다. `main` 함수의 출력이 10번 이루어질 동안 `test` 함수의 출력은 3번 발생할 수도 있고, 아예 실행이 되지 않을 수도 있다. 고루틴들이 어떻게 시분할하여 Core를 점유하는지에 따라 매번 실행이 달라진다. 병행성의 특징이 나타나는 것이다.

(Goroutine 실행 Waiting)