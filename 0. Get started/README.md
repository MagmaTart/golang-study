# 0. Get Started

### 초기 환경 세팅
Go 설치 후 초기 환경 세팅을 진행한다.
- GOROOT : GO의 실행 바이너리 파일이 위치한 Path.
- GOPATH : 현재 GO 파일을 컴파일하여 실행하는 Path.

별도의 설정이 없다면 GOROOT는 /usr/local/go로 기본 설정된다.

### 코드 컴파일 및 실행
- gofmt : GO의 표준 스타일 가이드에 맞도록 코드를 자동 수정해주는 툴
- goimports : 자동으로 import 경로를 추론하여 코드 상단에 추가해주는 툴
나 같은 경우는 goimports가 없어서, `sudo apt-get install golang-golang-x-tools`로 설치해줬다.

```
gofmt -w main.go
goimports -w main.go
go run main.go
```