# 5. Data types

Go의 데이터 타입은 넓게 정리해서 아래와 같은 것들이 있다.
- 정수 데이터 타입 : `int`
- 실수 데이터 타입 : `float`
- 복소수 데이터 타입 : `complex`
- 기타 타입 : `byte`, `rune`

### 정수 데이터 타입

- `int8` : 8비트 Signed integer
-  `int16` : 16비트 Signed integer 
-  `int32` : 32비트 Signed integer
-  `int64` : 64비트 Signed integer
-  `uint8` : 8비트 Unsigned integer
-  `uint16` : 16비트 Unsigned integer
-  `uint32` : 32비트 Unsigned integer
-  `uint64` : 64비트 Unsigned integer

### 실수 데이터 타입

- `float32` : 32비트 Signed float
- `float64` : 64비트 Signed float

### 복소수 데이터 타입

- `complex64` : 32비트 실수부와 32비트 허수부로 구성
- `complex238` : 64비트 실수부와 64비트 허수부로 구성

### 기타 타입

- `byte` : `uint8`과 동일한 데이터 타입. 1바이트의 raw 데이터 저장을 위함.
- `rune` : UTF-8 Character를 저장하는 타입.

### 데이터 타입 확인

변수 선언 시 그냥 `int`로만 선언할 경우, 컴파일 시에 운영체제에 따라 몇 비트 정수를 사용할 것인지 자동으로 결정한다. 나의 경우 64비트 프로세서를 사용하기 때문에 int64를 사용하는 것을 확인할 수 있었다.\
float 타입 변수도 마찬가지로 타입을 명시하지 않고 선언해 놓으면, 자동으로 몇 비트 실수형을 사용할지 결정한다.\
실제로 64비트 타입을 사용하는지 확인해보자. `unsafe` 패키지의 `Sizeof` 함수를 사용하면 변수가 사용하고 있는 메모리 크기를 확인할 수 있다.

```go
func type_test1() {
    a := 100
    b := 35.465

    fmt.Println("Size of int :", unsafe.Sizeof(a))
    fmt.Println("Size of float :", unsafe.Sizeof(b))
}
```

둘다 결과가 8으로 나왔다. int와 float 둘 다 64비트 타입을 쓰고 있다는 이야기다.

이번에는 `byte`와 `rune`을 확인해보자.

```go
func type_test2() {
    var b byte = 0x9C
    var r1 rune = '민'
    var r2 rune = '\uAD7F'

    fmt.Println("Size of byte :", unsafe.Sizeof(b))
    fmt.Println("Size of rune :", unsafe.Sizeof(r1), unsafe.Sizeof(r2))
}
```

`byte`는 1바이트였고, `rune`은 4바이트였다. 유니코드 문자를 표현하기에 적합한 데이터 타입인 듯 하다.

마지막으로 complex 데이터 타입을 확인해보자.\
complex는 흔히 아는 a+bi 형식으로 선언할 수도 있고, `complex` 함수를 이용해 선언할 수도 있다.\
또한 `real` 함수로 complex의 실수부를, `imag` 함수로 complex의 허수부를 가져올 수 있다.

```go
func type_test3() {
    var com64 complex64 = 5.54 + 2.71i
    var com128 complex128 = 7.01 + 3.4512e-10i
    var com64_2 complex64 = complex(5.54, 2.71)			// complex 함수를 이용해 선언 가능
    var com128_2 complex128 = complex(7.01, 3.4512e-10)	// complex 함수를 이용해 선언 가능

    com64_real := real(com64)
    com64_imag := imag(com64)
    com128_real := real(com128)
    com128_imag := imag(com128)

    fmt.Println("Complex64 :", unsafe.Sizeof(com64), unsafe.Sizeof(com64_2))
    fmt.Println("Complex128 :", unsafe.Sizeof(com128), unsafe.Sizeof(com128_2))
    fmt.Println("Complex64 real, imag :", unsafe.Sizeof(com64_real), unsafe.Sizeof(com64_imag))
    fmt.Println("Complex128 real, imag :", unsafe.Sizeof(com128_real), unsafe.Sizeof(com128_imag))
}
```

결과는 아래와 같았다. `complex64`는 32비트 실수 / 허수부, `complex128`은 64비트 실수 / 허수부를 각각 가지고 있음을 확인하였다.

```
Complex64 : 8 8
Complex128 : 16 16
Complex64 real, imag : 4 4
Complex128 real, imag : 8 8
```

### Renamed Type

다른 언어들에서와 같이, Go에서도 각 타입에 사용자 정의 이름을 붙여줄 수 있다.\
아래와 같이, `type` 키워드를 이용한다.

```go
type MyInteger int
var a, b MyInteger = 100, 200
fmt.Println(a + b)    // 300
```

`MyInteger`는 `int`로 치환된다. `type` 키워드로 설정한 이름은 해당 스코프 내에서만 유효하다.

C에서의 타입 재정의와 다른 점이 있다면, 아래와 같은 상황이 있다.

```go
type Apple int
type Banana int
var a Apple = 100
var b Banana = 200
fmt.Println(a+b)    // Error!
```

같은 `int` 타입을 각각 다른 이름인 `Apple`, `Banana`로 재정의했다. 분명 둘이 같은 `int` 타입이었으므로, `Apple` 타입과 `Banana` 타입 변수를 더했을 때 문제가 없을 것으로 예상하였다. 그러나 Go의 컴파일러는, `Apple`과 `Banana`가 다른 타입이라고 컴파일 에러를 일으킨다. 이는 Go가 재정의한 타입에 대해서도 엄격한 검사를 시행한다는 점을 볼 수 있는 부분이다. 이는 프로그래머가 더 안전한 코드를 짤 수 있도록 도울 것이다.