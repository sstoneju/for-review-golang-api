# 프로젝트 실행 방법
golang을 먼저 설치해준다.
```sh
$ go get .
$ go run .
```

https://github.com/cosmtrek/air
```sh
# air package를 설치해준다.
$ go install github.com/cosmtrek/air@latest
# air를 실행한다.
$ air
```

```sh
$ go install github.com/swaggo/swag/cmd/swag@latest
$ swag init # 코드를 읽어서 swagger를 생성한다.
```