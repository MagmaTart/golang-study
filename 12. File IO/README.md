# 12. File Input / Output

Go에서 파일 입출력은 `io.Reader`와 `io.Writer`를 기반으로 작업한다. 두 인터페이스는 기본적으로 바이트 단위로 파일에 읽고 쓸 수 있는 방법을 제공한다.\
C언어에서의 `fprintf` 등과 같이, Go에서도 `fmt` 패키지에 있는 `F` 접두사가 붙은 함수들을 이용하여 파일 입출력을 구현할 수 있다. 또한 파일 뿐만 아니라 표준 입출력, 소켓 등에 데이터를 기록할 때에도 사용 가능하다.

하나의 예를 들어보자. Go에서 표준 출력은 `os.Stdout`, 표준 입력은 `os.Stdin`이다.

```go
fmt.Fprintln(os.Stdout, "STDOUT TEST")
```

`Fprintln`은 첫 인자로 출력 스트림을 받고, 그 다음부터 가변인자로 출력할 데이터를 받는다. `os.Stdout`에 문자열을 출력했으므로 `fmt.Println`과 같이 터미널에 문자열이 출력된다.

### fmt 패키지에서 지원하는 입/출력 함수

`fmt` 패키지에서는 스트림을 열어 다양한 방법으로 입/출력을 수행할 수 있는 함수들을 제공하는데, C언어의 그것들과 같은 이름을 가지고 있는 경우가 많다.\
두 개의 정수를 입력받아 더한 결과를 출력하는 아래 예제를 보자.

```go
a, b := 0, 0
fmt.Fscanf(os.Stdin, "%d", &a)
fmt.Fscanf(os.Stdin, "%d", &b)
fmt.Fprintf(os.Stdout, "%d + %d = %d\n", a, b, a+b)	
```

C언어의 `fscanf`, `fprintf`와 쓰임새가 동일하다.\

또한 한 발짝 더 나아가, `sprintf`와 같은 문자열 대상 출력 함수도 지원된다.

```go
a, b := 0, 0
str := "10 + 20"
fmt.Sscanf(str, "%d + %d", &a, &b)
fmt.Printf("%d + %d = %d\n", a, b, a+b)     // 10 + 20 = 30
```

### 파일 읽기

파일 포인터는 `os.Open` 함수를 통해 얻어올 수 있다. `os.Open` 함수는 파일을 열면서 파일 포인터와 에러를 같이 반환한다. 에러가 `nil`일 경우 파일이 정상적으로 열렸다는 뜻이다.\
그 후 파일 포인터를 사용해서 `fmt.Fscanf` 등의 함수로 파일을 읽으면 된다.\
만약 파일을 여는 중이나 파일에 작업 중 에러가 발생하게 되어 해당 에러를 출력하면, 마치 C에서 `strerror` 함수를 이용하는 것과 같이 에러가 출력된다.

```go
func file_read_test() error {
    filename := "test.txt"
    file, err := os.Open(filename)
    if err != nil {		// 에러 발생시
        return err
    }
    defer file.Close()
    var read_str string
    if _, err = fmt.Fscanf(file, "%s", &read_str); err == nil {
        // Fscanf에서 에러 미 발생시
        fmt.Println(read_str)
    }
    return err
}
```

`defer` 키워드는 감싸고 있는 함수를 벗어날 때 호출할 함수를 미리 등록해놓는 기능을 한다. 따라서 이 함수가 리턴하면, `file.Close()`가 마지막으로 호출되어 파일을 자동으로 닫는다.\
`defer` 키워드의 수행은 호출의 역순으로 수행된다. 다시 말해서, 가장 먼저 `defer`로 등록한 함수가 가장 나중에 호출된다.

### 파일 쓰기

파일을 쓰는 것도 비슷한 구조로 이루어진다.

```go
func file_write_test() error {
    filename := "new.txt"
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()
    if _, err := fmt.Fprintln(file, "File writing test!"); err != nil {
        // Fprintln에서 에러 발생시
        return err
    }
    return err
}
```

### bufio 패키지 이용

그런데 위에서 소개하는 방법으로 파일을 읽고 쓰는 것은 한 가지 문제점이 있다. 예를 들어, 위의 예제에서 `%s`로 문자열을 읽어오려는 의도는 텍스트 파일에서 한 줄을 통째로 `string`에 저장하기 위함이었을 것이다. 하지만 `Fscanf`는 공백을 기준으로 문자열을 구분하기 때문에, 공백이 있을 경우 첫 공백 전의 단어만 읽어와 저장한다.

이러한 문제를 해결하기 위해서 _버퍼를 이용한 입출력_ 을 구현하여 사용할 필요성이 생기고, `bufio` 패키지의 인터페이스가 그것을 간단하게 구현하도록 도와준다.\
`bufio` 패키지는 자체로 Reader와 Writer 인터페이스를 제공하며, 각각 `io.Reader`와 `io.Writer`를 받아서 사용한다.

`bufio` 패키지의 Reader를 이용해 파일의 한 줄을 공백 포함 통째로 읽어 저장해보자.\
`import "bufio"`로 패키지를 import해주어야 한다.

```go
func file_bufio_read_test() error {
    filename := "test.txt"
    file, err := os.Open(filename)
    if err != nil {		// 에러 발생시
        return err
    }
    defer file.Close()

    reader := bufio.NewReader(file)
    if read_str, err := reader.ReadString('\n'); err == nil {
        fmt.Println(read_str)
    } else {
        return err
    }
    return err
}
```

`ReadString` 함수는 인자로 받은 delimeter가 나올 때 까지 스트림에서 읽어 버퍼에 저장하고 있다가, delimeter를 만나면 버퍼를 묶어 `string`으로 반환해준다.

`string`을 파일에 저장하는 것도 `bufio`로 해보자.

```go
func file_bufio_write_test() error {
    filename := "new.txt"
    file, err := os.Create(filename)
    if err != nil {
        return err
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    if _, err := writer.WriteString("bufio writing test!\n"); err != nil {
        return err
    }
    writer.Flush()
    return err
}
```

`WriteString` 함수는 인자로 `string`을 받아서, 일단 내부 버퍼에 쌓아놓고 있는다. 따라서 실제로 파일에 기록하려면 꼭 `Flush` 함수를 호출해주어야 한다.
