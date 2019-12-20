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

### 빈 인터페이스


