# 21. Interface 2

### 인터페이스 내장

마치 구조체처럼, 인터페이스도 다른 인터페이스를 안에 내장할 수 있다. 그럴 경우, 내장된 인터페이스의 메서드까지 모두 구현해야 한다.\
내장할 때는 구조체와 같이, 인터페이스 정의 시에 다른 인터페이스의 이름을 넣어주면 된다.

```go
type Animal interface {
    walk()
}

type Human interface {
    Animal
    speak()
}

type Person struct {
    name string
}

func (p Person) eat() {
    fmt.Println("Person : Eat")
}

func (p Person) speak() {
    fmt.Println("Person : Speak")
}

func HumanTest(h Human) {
    h.eat()
    h.speak()
}

func main() {
    var person = Person{"Soomin Lee"}
    HumanTest(person)
}
```

`Human` 인터페이스가 `Animal` 인터페이스를 내장하고 있다. 이후에 어떤 타입이 `Human` 인터페이스를 따르게 하려면, `Human` 인터페이스의 함수들뿐만 아니라 `Animal` 인터페이스의 함수들도 모두 구현해주어야 한다. 어느 쪽의 하나라도 구현되어있지 않은 타입은 `Human` 인터페이스와 호환되지 않는다.

### 빈 인터페이스와 Type Assertion

인터페이스에 아무 메서드도 정의하지 않으면, 모든 타입의 객체를 사용할 수 있다.

```go
func getType(arg interface{}) {
}
```

받은 인자가 어떤 타입의 변수인지 확인하는 코드를 아래처럼 작성할 수 있을 것이다. 아래같은 코드를 __타입 스위치(Type Switch)__ 라고 한다. 변수의 타입에 따라 다른 분기를 수행하는 경우이다.

```go
func getType(arg interface{}) {
    switch arg.(type) {
    case int:
        fmt.Println("Argument is Int")
    case float32, float64:
        fmt.Println("Argument is Float")
    case string:
        fmt.Println("Argument is String")
    }
}

func main() {
    i := 50
    f := 5.791
    s := "Hello"

    getType(i)
    getType(f)
    getType(s)
}
```

`getType` 함수는 어떤 타입의 인자던지 받을 수 있다. 대신 타입이 결정되어 있지 않기 때문에, 저 인자 `arg`를 함수 내에서 그대로 사용할 수는 없다. 그래서 빈 인터페이스 위치에 넘어온 인자의 타입을 명시적으로 결정해주어야 하는데, 이것을 __형 단언(Type Assertion)__ 이라고 한다. Type Assertion으로 인자의 타입을 결정한 후에 인자를 해당 타입에 맞게 사용할 수 있다.

```go
func printString(arg interface{}) {
    str := arg.(string)
    fmt.Println(str)
}

func main() {
    s := "Hello"
    printString(s)
}
```

`printString` 함수는 빈 인터페이스 인자를 받고 있다. 그러나 함수 내부에서 `arg` 인자는 무조건 `string`일 것이라고 Type Assertion을 하고 `arg` 인자를 `string`으로써 `str`에 대입한다.

여기서, 위에서 보았던 Type Switch가 Type Assertion과 동일한 문법임을 볼 수 있다. `type` 키워드를 사용해 Type Assertion을 작성하면, switch-case 구문에서 자료형의 타입에 따라 다른 분기로 넘어가도록 코드를 작성할 수 있다. 이렇게 Type Switch를 구현했던 것이다.

Type Switch는 당연히 프로그래머가 정의한 자료형에 대해서도 구현할 수 있다.

```go
type Rect struct {
    width, height float64
}

type Circle struct {
    radius float64
}

func area(shape interface{}) float64 {
    area := 0.0
    switch shape.(type) {
    case Rect:
        rect := shape.(Rect)
        area = rect.width * rect.height
    case Circle:
        circle := shape.(Circle)
        area = circle.radius * circle.radius * 3.141592
    }
    return area
}

func main() {
    var rect = Rect{5.0, 6.0}
    var circle = Circle{3.0}

    fmt.Println("Rect area :", area(rect))
    fmt.Println("Circle area :", area(circle))
}
```
