# 13. bufio 패키지

이전 챕터에서 소개하는 방법으로 파일을 읽고 쓰는 것은 한 가지 문제점이 있다. 예를 들어, 이전 챕터의 예제에서 `%s`로 문자열을 읽어오려는 의도는 텍스트 파일에서 한 줄을 통째로 `string`에 저장하기 위함이었을 것이다. 하지만 `Fscanf`는 공백을 기준으로 문자열을 구분하기 때문에, 공백이 있을 경우 첫 공백 전의 단어만 읽어와 저장한다.

이러한 문제를 해결하기 위해서 _버퍼를 이용한 입출력_ 을 구현하여 사용할 필요성이 생기고, `bufio` 패키지의 인터페이스가 그것을 간단하게 구현하도록 도와준다.\
`bufio` 패키지는 자체로 Reader와 Writer 인터페이스를 제공하며, 각각 `io.Reader`와 `io.Writer`를 받아서 사용한다.

### bufio.Reader

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

### bufio.Writer

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

### bufio.Scanner

`bufio.Scanner`는 `bufio.Reader`처럼 스트림에서 읽어오는 인터페이스이다. 그러나 더 쓰기 간편하며, 더욱 복잡한 리딩 패턴을 쉽게 작성할 수 있다는 장점이 있다.\
파일 스트림을 넣으면 기본적으로 줄 단위로 읽어준다. 줄을 읽으면서 마지막 개행 문자는 자동으로 무시해주기 때문에 신경쓸 점이 줄어든다.

```go
func file_bufio_scanner_test() error {
    filename := "test.txt"
    file, err := os.Open(filename)
    if err != nil {		// 에러 발생시
        return err
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        fmt.Println(scanner.Text())
    }
    if err := scanner.Err(); err != nil {
        return err
    }
    return err
}
```

`Scanner`가 파일을 읽는 도중에 생긴 에러는 아래의 `if` 절에서 처리하고 있다.

그럼 Scanner를 어떻게 커스터마이징하여 사용할 수 있는지 알아보자.\
`Scanner`의 동작은 전적으로 `Scanner.Split(func)` 함수에서 등록된 함수를 기반으로 한다. 입력 스트림을 토크나이징하고 스캐닝의 각 결과로 반환할 방법을 담은 함수를 정의하고, 그 함수를 `Scanner.Split`에 인자로 넘겨 등록해주면 된다.\
위에서 봤듯이, 함수를 따로 등록해주지 않을 경우 기본으로 사용하는 토크나이저 함수는 줄 단위로 끊어 반환하는 기능을 한다.

또한 `bufio` 패키지에는 기본적으로 구현되어 있는 몇 가지의 토크나이저 함수가 있다. 그 중 하나인 `bufio.ScanWords`를 사용하여 줄 단위가 아닌 단어 단위로 끊어서 읽어 보자. 위 코드의 `scanner` 선언 라인 아래에 다음의 한 줄을 추가한다.

```go
scanner.Split(bufio.ScanWords)
```

코드를 실행해보면, 스페이스와 개행 문자를 단어를 나누는 Delimeter로 가정하고 단어 단위로 나누어서 스캐닝 결과를 리턴하는 것을 볼 수 있다.\
`Split` 함수에는 구현 모양만 맞춘다면 어떤 함수던지 등록할 수 있다. 이에 대한 자세한 예제는 아래의 링크를 참고하자.\
[golang official bufio example](https://golang.org/src/bufio/example_test.go)