# 20. Interface 1

__인터페이스(Interface)__ 는 메서드의 집합체이다. 구조체가 필드만을 가지고 있을 수 있었다면, 인터페이스는 메소드만을 가지고 있을 수 있다.\
인터페이스는 _특정 타입이 구현해야만 하는 메서드_ 들을 정의한다. 이후에 특정 함수가 특정 인터페이스에 맞는 타입의 변수만을 받도록 정의될 경우, 그 인터페이스의 모든 메서드가 구현된 타입만 해당 함수로 넘어갈 수 있다.\
아래의 코드에서 인터페이스의 기본적인 정의 방법을 볼 수 있다.

```go
type IFileManager interface {
    Read() string
    Write(data string) error
}
```

`IFileManager` 인터페이스는 `Read`와 `Write`, 두 가지의 메서드를 정의하고 있다. 이후에 이 인터페이스에 맞는 타입을 만들려면, 두 메서드를 모두 구현해야 한다.

위의 인터페이스를 따르는 두 개의 구조체 타입을 만들어 보자.

```go
type CSVFileManager struct {
	filename string
}

type TXTFileManager struct {
	filename string
}

func (fm TXTFileManager) Read() string {
	return "Dummy String TXT"
}

func (fm TXTFileManager) Write(data string) error {
	var err error = nil
	fmt.Println("WRITE :", data)
	return err
}

func (fm CSVFileManager) Read() string {
	return "Dummy String CSV"
}

func (fm CSVFileManager) Write(data string) error {
	var err error = nil
	fmt.Println("WRITE :", data)
	return err
}
```

예제용 코드이므로, 이름은 `Read, Write`이지만 더미 코드로 만들었다. 일단 각각의 함수는 당연히 아래와 같이 호출 가능하다.

```go
var csv = CSVFileManager{"Dummy_Filename1"}
var txt = TXTFileManager{"Dummy_Filename2"}
fmt.Println(csv.Read())     // "Dummy String CSV"
fmt.Println(txt.Read())     // "Dummy String TXT"
csv.Write("TEST_CSV")       // "WRITE : TEST_CSV"
txt.Write("TEST_TXT")       // "WRITE : TEST_TXT"
```

여기서 `CSVFileManager`와 `TXTFileManager`는 `IFileManager` 인터페이스에서 정의한 시그니처에 맞는 각각의 `Read, Write` 함수를 가지고 있으므로, 인터페이스에 맞는 구조체 타입이다.\
이제 __인터페이스에 맞는 타입을 인자로 받는 함수__ 를 정의해보자.

```go
func FileManagerTest(manager IFileManager) {
    fmt.Println(manager.Read())
    manager.Write("Write test")
}
```

함수 `FileManagerTest`는 인자로 `IFileManager` 인터페이스를 넘겨받는다. 이 때 이 함수로 넘길 수 있는 인자는 `IFileManager`의 __모든 메서드가 구현된 타입__ 들이 된다. 따라서 위의 `CSVFileManager, TXTFileManager`는 이 함수에 인자로 넘길 수 있다.

```go
FileManagerTest(csv)
FileManagerTest(txt)
```

이제 새로운 구조체를 하나 더 만들고, `Read` 함수 하나만 구현해보자. 이 구조체는 `IFileManager` 인터페이스에 맞지 않을 것이다.

```go
type JSONFileManager struct {
	filename string
}

func (fm JSONFileManager) Read() string {
	return "Dummy String JSON"
}
```

이후 `JSONFileManager` 구조체 변수를 `FileManagerTest`에 인자로 넘기려하면 에러가 발생한다. `JSONFileManager` 구조체가 `IFileManager` 인터페이스에 맞지 않기 때문이다.

```go
var json = JSONFileManager{"Dummy_Filename3"}
FileManagerTest(json)       // 에러 발생!
```

위의 예제들을 통해, 인터페이스는 어떠한 타입들이 어떤 메서드들을 가지고 있어야 하는지 정의하는 역할임을 알 수 있다.

### Go에서 인터페이스의 네이밍 관례

Go에서는 인터페이스의 이름을 붙일 때, 인터페이스의 주요 메서드 이름에 `er`을 붙여 표현하는 경우가 많다.\
예를 들어, `Read` 메서드를 정의하고 있는 인터페이스는 `io.Reader`이고, `Write` 메서드를 정의하고 있는 있는 인터페이스는 `io.Writer`이다.\
string으로 표현될 수 있는 타입들은 모두 `String()` 메서드를 가지고 있고, 모두 `fmt.Stringer` 인터페이스와 호환된다. 그래서 `fmt.Println` 등의 메서드에서 사용할 수 있었던 것이다.\
반대로 사용자가 직접 정의한 타입에 `string`을 반환하는 `String()` 메서드만 만들어 주면, 그 타입은 `fmt.Stringer` 인터페이스에 호환되는 타입이 된다.
