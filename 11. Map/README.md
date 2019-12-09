# 11. Map

### Map의 생성 및 초기화
Map은 Key와 Value의 쌍을 순서와 상관없이 저장하는 구조로, 해시테이블로 구현된다.\
Go에서 맵은 다음과 같은 형식으로 정의할 수 있다.

```go
// var <name> map[KeyType]ValueType
var mymap map[int]string
```

Key의 타입과 Value의 타입을 명시하여 맵을 선언한다.\
슬라이스와 같이 저렇게 선언만 해 놓으면 nil map인 상태이므로, `make` 함수를 이용하여 맵을 생성해야 한다.\
또는 아래 두 번째 줄처럼, 선언과 동시에 생성할 수도 있다.

```go
mymap = make(map[int]string)    // make로 새 맵 생성 및 할당
mymap2 := map[int]string{}      // 새 맵의 생성과 동시에 할당
```

Map의 초기화는 아래와 같이 Key-Value 쌍들을 나열해서 할 수 있다.

```go
mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC"}
```

### Map 참조, 삽입 및 삭제

맵에 접근하고 참조하는 법은 파이썬과 비슷하다. 한 가지 다른 점은, 맵에서 특정 키를 이용해 값을 빼올 때 _해당 키의 존재 여부_ 를 같이 알 수 있다는 점이다.\
아래와 같이 맵을 참조할 수 있다.

```go
mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC"}
value, exists := mymap[2]    // value : "BBB", exists : true
value, exists = mymap[4]     // value : "", exists : false
mymap[4] = "DDD"
value, exists = mymap[4]     // value : "DDD", exists : true
```

직접 확인해보니, 없는 키를 참조할 경우 반환되는 값은 Map value의 자료형의 기본값이다. string일 경우 빈 문자열(`""`), int일 경우 `0`과 같은 식이다.

맵에서 특정 Key를 지울 때는 `delete` 함수를 사용한다.

```go
mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
// delete(map, key)
delete(mymap, 3)    // {1: "AAA", 2: "BBB", 4: "DDD"}
```

아예 Key-Value 쌍이 삭제되므로, Map의 길이가 1 줄어들게 된다.\
Map의 길이는 슬라이스와 같이 `len` 함수로 볼 수 있으며, Key-Value 쌍의 개수를 출력한다.

```go
mymap := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
fmt.Println(len(mymap))    // 4
```

맵 또한 슬라이스와 같이 `range`를 이용해 Iterating 할 수 있다. 맵을 Iterating 할 경우, 반복마다 Key와 Value 쌍이 하나씩 리턴된다. Map은 순서가 없으므로, 코드 실행시마다 리턴 순서는 바뀔 수 있다.

```go
func print_map(map_to_print map[int]string) {
    for key, value := range(map_to_print) {
        fmt.Printf("%d : %s\n", key, value)
    }
}
```

### Map의 비교

특히 테스트를 작성할 때, 두 개의 맵이 같은지 비교해야 할 일이 생길 수도 있다. Map은 순서가 없는 구조이기 때문에 테스트 작성이 쉽지 않은데, `reflect.DeepEqual`을 사용하면 _완전히 동일한 Key-Value 테이블을 가지고 있는지_ 확인 가능하다.

```go
func map_equal_test() {
    mymap1 := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
    mymap2 := map[int]string{1: "AAA", 2: "BBB", 3: "CCC", 4: "DDD"}
    mymap3 := map[int]string{1: "EEE", 2: "FFF", 3: "GGG", 4: "HHH"}

    fmt.Println("mymap1 is equal to mymap2? : ", reflect.DeepEqual(mymap1, mymap2))     // true
    fmt.Println("mymap1 is equal to mymap3? : ", reflect.DeepEqual(mymap1, mymap3))     // false
}
```
