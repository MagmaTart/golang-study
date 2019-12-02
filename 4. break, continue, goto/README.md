# 4. break, continue, goto

1. Go에서 break는 for 루프를 탈출할 때나, switch-case 구문을 탈출할 때 사용한다.

```go
func break_test() {
    for num := 0; num < 100; num++ {
        fmt.Println("num :", num)
        if num >= 50 {
            break
        }
    }
}
```

2. continue는 for 루프의 조건 검사부로 돌아갈 때에 사용한다.

```go
func continue_test() {
    for num := 0; num < 100; num++ {
        if num%2 == 0 {
            continue
        }
        fmt.Println("num :", num)
    }
}
```

3. goto는 임의의 문장으로 이동하기 위해 사용한다. C언어에서의 goto와 동일하다.

```go
func goto_test() {
    for num := 0; num < 100; num++ {
        fmt.Println("num :", num)
        if num >= 70 {
            goto END
        }
    }
END:
    fmt.Println("JUMPED END")
}
```

4. break문 뒤에 레이블을 붙이면, 해당 레이블 바로 밑의 블럭을 통째로 빠져나오게 된다.\
아래 코드에서 `break LOOP_A` 구문은 이중 for문 안에 들어가 있지만, `LOOP_A` 레이블이 바깥 for문을 묶고 있으므로 이중 for문을 한번에 탈출하게 된다.

```go
func break_label_test() {
LOOP_A:
    for a := 0; a < 100; a++ {
LOOP_B:
        for b := 0; b < 100; b++ {
            fmt.Printf("%d * %d = %d\n", a, b, a*b)
            if a >= 10 && b >= 50 {
                break LOOP_A
            } else if b >= 50 {
                break LOOP_B
            }
        }
    }
}
```


