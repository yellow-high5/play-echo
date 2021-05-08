# Play Echo

execute below command for hotreloading server.
```sh
go get -u github.com/cespare/reflex
reflex -r '(\.go$|go\.mod)' -s go run server.go
```