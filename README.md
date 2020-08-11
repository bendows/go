# go
go packages

```bash
bendows@bendows-G41MT-S2:/home/test-log$ find
.
./main.go
bendows@bendows-G41MT-S2:/home/test-log$ go mod init some-module-name
go: creating new go.mod: module some-module-name
bendows@bendows-G41MT-S2:/home/test-log$ go build
go: finding module for package github.com/bendows/go
go: found github.com/bendows/go in github.com/bendows/go v0.0.0-20200704011029-35f2948f1512
bendows@bendows-G41MT-S2:/home/test-log$ ./some-module-name 
2020-08-11 21:59:57 main.go:11: hello
2020-08-11 21:59:57 main.go:12: hello
hello there
bendows@bendows-G41MT-S2:/home/test-log$ cat main.go 
package main

import (
	"fmt"

	logger "github.com/bendows/go"
)

func main() {
	logger.LogOn = true
	logger.Loginfo.Println("hello")
	logger.Logerror.Println("hello")
	fmt.Println("hello there")
}
bendows@bendows-G41MT-S2:/home/test-log$ 
```
