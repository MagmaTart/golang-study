# 3. Control Flow

Go는 기본적으로 if, switch, for, select의 Control Flow를 제공한다. select같은 경우는 바로 이해하기 조금 어려워서, 나머지 세 가지 제어문의 사용법을 보고 가려고 한다.

### Go Operators
Go에서의 연산자는 대부분이 C와 비슷하다. 아래와 같은 연산자들을 지원한다.
- Arithmetic Operators : `+`, `-`, `*`, `/`, `%`, `++`, `--`
- Relational Operators : `==`, `!=`, `>`, `>=`, `<`, `<=`
- Logical Operators : `&&`, `||`, `!` (C언어와 같이 정수 자료형에 대해서도 동작한다.)
- Bitwise Operators : `&`, `|`, `^`, `<<`, `>>`
- Assignment Operators : `=`, `+=`, `-=`, `*=`, ...
- Pointer Operators : `*`, `&`

증감 연산자 (`++`, `--`)는 무조건 피연산자 뒤에서만 사용할 수 있다(후위증감).\
이외에는 C언어와 대부분 동일한 기능을 가지는 연산자들이며, 포인터 연산까지 지원한다.

### if
Go에서 if문은 아래와 같이 사용한다. 조건의 형식은 여태 보던 언어들과 비슷하다.

```go
a := 10
if a > 5 {
    fmt.Println("a is greater than 5")
}
```

조건식에 괄호가 사용되지 않는다는 점만 기억하면 될 듯 하다.

else if도 같은 형식으로 사용한다.

```go
a := 70
if a < 60 {
    fmt.Println("a < 60")
} else if a > 60 && a < 70 {
    fmt.Println("60 < a < 70")
} else {
    fmt.Println("a > 70")
}
```

또한 __Optional Statement__ 라고 해서, 조건식의 평가 직전에 간단한 Statement를 실행할 수 있다. 이 Statement에서 선언된 변수 등은 현재 if문의 Scope 안에서만 사용할 수 있다.\

```go
a := 5
if num := a*a; num > 20 {
    fmt.Println("num is greater than 20")
}
```

위의 코드에서는, if문의 조건식 실행 직전에 `num := a*a` 코드를 먼저 실행시켜 `num` 변수를 선언한다. `num` 변수는 따라오는 if문의 scope 안에서만 유효하고, if문을 탈출하면 소멸된다.

### switch

C언어에서의 switch와 기본적으로 같은 기능을 하지만, 훨씬 다양한 사용 바리에이션이 존재한다.

1. Go의 switch에서는 하나의 case문에 여러 값을 지정할 수 있다.

```go
num := 3
switch num {
case 1, 2:
    fmt.Println("A")
case 3, 4:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

2. switch 키워드 뒤에는 변수 뿐만 아니라, 아래와 같이 Expression도 들어갈 수 있다.\
case 키워드에도 마찬가지로 값이 아니라 조건문을 넣을 수 있다. 이 경우 조건을 만족하는 case문을 실행하게 된다.

```go
num := 2
switch mul := num * 2; {
case mul < 4:
    fmt.Println("mul < 4")
case mul >= 4 && mul < 6:
    fmt.Println("4 <= mul < 6")
default:
    fmt.Println("mul >= 6")
}
```

3. switch 키워드 뒤에 Expression이 없어도 된다.

```go
num := 10
switch {
case num < 10:
    fmt.Println("num < 10")
case num >= 10 && num < 20:
    fmt.Println("10 <= num < 20")
default:
    fmt.Println("num > 20")
}
```

4. break가 없어도 자동으로 해당 case만 실행하고 건너뛴다.\
다음 case로 흐르게 하려면 __fallthrough__ 구문을 사용해야 한다.

```go
num := 1
switch num {
case 1:
    fmt.Println("num : 1")
    fallthrough
case 2:
    fmt.Println("num : 2")
    fallthrough
default:
    fmt.Println("num : 3")
}
```

### for

익히 알고 있는 for문의 형태는 Go에서 아래와 같이 구현한다.

```go
sum := 0
for i := 1; i <= 100; ++i {
    sum += i
}
```

C언어에서의 while과 같이, for문에서 조건식만 사용해 루프를 구성할 수도 있다.

```go
sum := 0
i := 1
for i <= 100 {
    sum += i
    ++i
}
```

조건식까지 생략하여 무한루프를 구성할 수 있다.

```go
for {
    fmt.Println("Hello)
}
```

배열, 슬라이스와 같은 컬렉션은 Iterating하기 위한 for-range 구문도 존재한다.\
아래와 같이 구성하면, 마치 파이썬의 enumerate와 같이 각 요소와 인덱스를 차례대로 순회한다.

```go
strs := []string{"AAA", "BBB", "CCC"}
for idx, name := range strs {
    fmt.Println(idx, name)
}
```