# 19. Const, Enum

### Const 선언

Go에서 상수(Const)를 선언할 때에는 `const` 키워드를 사용한다. 아래 코드에서 보이는 것과 같이, 타입은 생략할수도 있고 명시적으로 지정해줄 수도 있다.

```go
const CA = 100
const CB int = 200
const CS = "MagmaTart"
```

아래 코드와 같이, 상수들을 괄호로 묶어 한번에 선언할 수도 있다.

```go
const (
    CC = 1234
    CD = 5678
    CS2 = "Soomin"
)
```

### Enum

Go에는 명시적으로 __Enum의 선언을 담당하는 키워드가 없다.__ Enum을 사용하고 싶을 때는 상수들을 순서대로 정의해서 사용해야 한다.\
기본적으로는 상수를 묶어 선언하는 방법으로 Enum처럼 정의할 수 있다.

```go
type Color int
const (
    RED   Color = 0
    GREEN Color = 1
    BLUE  Color = 2
)
```

그러나 일일이 상수의 값을 모두 결정하는 것은 힘든 일이다. 마치 C에서의 Enum처럼 자동으로 값을 붙이려면 `iota` 키워드를 이용하면 된다.

```go
const (
    RED   Color = iota      // 0
    GREEN Color = iota      // 1
    BLUE  Color = iota      // 2
)
```

`iota`는 상수 선언 블럭의 가장 위에서 0으로 초기화된다. 그리고 상수 선언에서 한 번 대입될 때마다 값을 1씩 증가시켜주는 동작을 한다. 그래서 직접 선언해주었던 것과 동일한 상수들을 만들 수 있었던 것이다.\
만약 이후로 나오는 상수들의 타입이 같다면, 아래 코드에서처럼 `iota` 키워드를 생략 가능하다.

```go
const (
    RED   Color = iota
    GREEN
    BLUE
)
```

`iota` 키워드는 더욱 복잡한 Enum의 선언도 가능하게 한다. 아래의 코드를 보자. 100점부터 60점 사이로 각 학점의 기준 점수를 Enum으로 만들었다.

```go
type Grade int
const (
    A Grade = 100 - (10 * iota)     // 100
    B                               // 90
    C                               // 80
    D                               // 70
    F                               // 60
)
```

맨 처음 상수 `A`를 선언할 때 `iota` 값을 이용해 특정한 계산식을 만들어주었다. 이후의 상수들에 대해 대입문을 생략할 경우, 마치 엑셀에서 하나의 셀을 쭉 끄는것처럼 __같은 게산식이 반복적으로 적용된다.__ 물론 iota는 상수 하나를 선언할 때마다 1씩 증가한다.

`iota` 값이 0일때를 무시하고 싶으면, 즉 Enum의 상수 값을 1부터 시작하고 싶으면 `_` 키워드를 사용한다.

```go
const (
    _           = iota
    RED   Color = iota      // 1
    GREEN                   // 2
    BLUE                    // 3
)
```

맨 첫 줄의 `_` 상수는 무시된다.
