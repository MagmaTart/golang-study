# 18. Method

__함수에 리시버가 붙으면 메서드이다.__ 메서드 안에서는 전달받은 인자와 같이, 리시버에 대한 참조도 가능하다. 다른 언어들에서의 메서드와 같이, 리시버에 대해 작용하는 서브루틴을 정의할 수 있다.

### 자료형에 대한 메서드

__Go의 기본 자료형에 대한 메서드는 만들 수 없다.__ `int` 형의 리시버를 가지는 메서드는 만들 수 없다는 이야기이다.

정사각형의 한 변의 길이를 나타내는 자료형인 `Length`를 `int` 타입의 재정의로 등록하자. 이제 정사각형의 넓이를 구하는 함수를 만들고 싶다. 기존에는 이렇게 작성했었다. 

```go
type Length int

func area(n Length) Length {
	return n * n
}

func main() {
    var length Length = 7
	area_result := area(length)  // 25
}
```

`area` 함수는 `Length` 값 하나를 인자로 받아 제곱한 값을 출력한다.\
이 함수를 `Length` 형 변수를 리시버로 받는 메소드로 변경해보자.

```go
func (n Length) area() Length {
    return n * n
}

func main() {
    var length Length = 15
    area_result = length.area()
}
```

메서드를 만들 때는, 리시버의 이름과 자료형만 함수 이름 앞에 정의해주면 된다. `Length` 타입의 리시버를 받는 메서드이므로, 어떤 `Length` 타입 변수던지 호출 가능하다.

### 구조체에 대한 메서드

