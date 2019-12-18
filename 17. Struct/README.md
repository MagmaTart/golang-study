# 17. Struct

Go는 객체지향을 염두에 두고 설계한 언어가 아니다. 이에 대한 단적인 예시로, Go에는 __Class가 없다.__ 당연히 상속 등과 같은 기능도 존재하지 않는다.\
대신 Go는 __구조체(Struct)__ 와 __메서드(Method)__, __인터페이스(Interface)__ 를 통해 Go만의 방식으로 객체지향적인 기능을 지원한다.

### 구조체의 선언과 초기화, 참조

구조체는 여러 필드를 가지고 있는 구조로, C와 같은 다른 언어들에서의 구조체와 비슷한 역할을 한다. 간단한 구조체를 하나 선언해보자.

```go
var profile = struct {
    name string
    age int
    is_male bool
}
```

구조체의 정의 자체를 구조체 변수 선언과 동시에 하고 있다. 이는 번거로우므로, 당연히 정의된 구조체에 이름을 붙여줄 수 있다.

```go
type Person struct{
    name string
    age int
    is_male bool
}
```

`type` 키워드를 사용해 구조체에 이름을 붙여주었다. 이제 아래와 같이 구조체 변수를 선언하고 초기화할 수 있다.

```go
var profile1 = Person{"Jenny", 20, false}             // 필드의 순서대로 초기화 값을 설정
var profile2 = Person{"Robert", 25, true}
var profile3 = Person{name: "James", is_male: true}   // 원하는 필드만 초기화하고 싶을 경우, 콜론 사용
```

일부 필드만 초기화할 경우, 초기화 값이 전달되지 않은 필드는 해당 자료형의 기본값으로 설정된다.

구조체 필드의 참조는 다른 언어들과 같이 `.` 연산자를 사용하면 된다.

```go
age_sum := profile1.age + profile2.age + profile3.age
fmt.Println(age_sum)    // 45
```

### 구조체의 내장

사실 구조체 자체만 놓고 보면 큰 메리트가 없어 보인다. 하지만 구조체를 빛나게 하는  특징적인 기능 중 하나가 바로 __구조체 내장__ 이다. 구조체가 다른 구조체의 필드를 동시에 모두 가지고 있을 수 있게 함으로써, 마치 객체지향 언어의 상속처럼 구조체를 디자인할 수 있게 한다.

위에서 만든 `Person` 구조체를 재사용해서 _회원 정보_ 를 관리하는 구조체를 만들어보자.

```go
type Person struct{
    name string
    age int
    is_male bool
}

type Account struct{
    Person
    id, pw string
}

func main() {
    var acc1 Account
    acc1.name = "Jake"
    acc1.age = 37
    acc1.is_male = true
    acc1.id = "iamjake"
    acc1.pw = "jake8846"
}
```

`Account` 구조체에 `Person` 구조체를 내장하고 있다. 이후에 `Account` 구조체 변수를 선언하면, 해당 변수는 `Account`만의 필드인 `id, pw` 뿐만 아니라, `Person`의 필드인 `name, age, is_male`까지 자신의 필드처럼 사용할 수 있다.\
오해하지 말아야 할 점은, 이것이 _상속의 구현은 아니라는 점_ 이다. 어디까지나 구조상의 편리함을 위해 필드를 동시에 내장할 수 있도록 해주는 기능일 뿐이다.
