- `go generate ./cmd/...`
- `go fmt ./cmd/...`
- `go build ./cmd/...`
- `./mytest`

prints:
```text
First's Get with int: 10
Second's Get with string: value

First's GetAny with int: 666
Second's GetAny with string: any value

First's PrintIntOrStringElements with int:
 - elem 1 : 10
 - elem 2 : 11

Second's PrintIntOrStringElements with string:
 - elem 1 : value1
 - elem 2 : value2

First's PrintAnyElements with int:
 - elem 1 : 10
 - elem 2 : 11

Second's PrintAnyElements with string:
 - elem 1 : valueAny1
 - elem 2 : valueAny2
```