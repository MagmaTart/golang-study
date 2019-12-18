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

어떤 타입이던지 메서드를 가질 수 있으므로, 리시버가 구조체이기만 하면 구조체 또한 메서드를 가질 수 있다. 이는 마치 구조체가 멤버 메서드를 가지고 있는 것 같은 기능을 하게 해준다.

저번 구조체 예제에서 보았던 `Person` 구조체에 대한 출력 기능을 가지는 메서드를  정의해보자.

```go
type Person struct{
    name string
    age int
    gender string
}

func (person Person) Print() {
    fmt.Println("Name is", person.name)
    fmt.Println("Age is", person.age)
    fmt.Println("Gender is", person.gender)
}

func main() {
    var me = Person{"Lee Soomin", 20, "Male"}
    me.Print()
}
```

### 포인터 리시버

기본적으로 Go의 함수 인자는 복사되어 넘어가듯이, 리시버도 __복사되어 메서드로 넘어간다.__ 이 말은 즉, 메서드 내부에서 리시버에 대한 수정이 있어도 메서드 외부에는 적용되지 않는다는 소리이다.\
Go에서는 인자와 같이 리시버 또한 포인터로 넘어갈 수 있다.

예를 들어 위의 `Person` 예제에 덧붙여 아래의 메서드를 작성했다고 치자. `Person`의 `name` 필드를 변경하는 것을 목적으로 한 메서드이다.

```go
func (person Person) change_name(new string) {
    person.name = new
}

func main() {
    var me = Person{"Lee Soomin", 20, "Male"}   // me.name : "Lee Soomin"
    me.change_name("MagmaTart")                 // me.name : "Lee Soomin"
}
```

`me.name`이 변경되는 것을 기대했으나 그렇지 않았다. 메서드에 리시버 구조체가 _복사되어_ 넘어갔기 때문이다. 원하는 대로 메서드가 필드를 직접 수정하게 하려면, 리시버에 포인터를 넘겨야 한다.

```go
func (person *Person) change_name(new string) {
    person.name = new
}

func main() {
    var me = Person{"Lee Soomin", 20, "Male"}   // me.name : "Lee Soomin"
    me.change_name("MagmaTart")                 // me.name : "MagmaTart"
}
```

구조체 변수의 포인터가 메서드로 넘어가므로, 메서드에서 구조체 변수를 수정한 사항이 메서드 외부에서도 그대로 적용된다. 마치 구조체의 멤버 메소드처럼 사용할 수 있는 것이다.

### 메서드의 Public 및 Private

객체지향의 접근 지정자들과 같은 기능을 하는 예약어가 Go에는 따로 없다. 대신 메서드, 변수, 함수 등의 접근 지정에는 한가지 규칙이 존재한다.\
__식별자의 이름이 대문자로 시작하면 Public, 소문자로 시작하면 Private이다.__ Public인 녀석만 모듈 외부에서 보고 사용할 수 있다.\
이는 메서드, 상수, 변수, 함수, Named Type 등등에 모두 적용되는 룰이다.
