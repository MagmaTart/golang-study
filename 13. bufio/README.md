# 13. bufio 패키지

이전 챕터에서 소개하는 방법으로 파일을 읽고 쓰는 것은 한 가지 문제점이 있다. 예를 들어, 이전 챕터의 예제에서 `%s`로 문자열을 읽어오려는 의도는 텍스트 파일에서 한 줄을 통째로 `string`에 저장하기 위함이었을 것이다. 하지만 `Fscanf`는 공백을 기준으로 문자열을 구분하기 때문에, 공백이 있을 경우 첫 공백 전의 단어만 읽어와 저장한다.

이러한 문제를 해결하기 위해서 _버퍼를 이용한 입출력_ 을 구현하여 사용할 필요성이 생기고, `bufio` 패키지의 인터페이스가 그것을 간단하게 구현하도록 도와준다.\
`bufio` 패키지는 자체로 Reader와 Writer 인터페이스를 제공하며, 각각 `io.Reader`와 `io.Writer`를 받아서 사용한다.

### bufio.NewReader

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

### bufio.NewWriter

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

~작성 예정~