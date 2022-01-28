godaemon
========

Run golang app as background program, 以后台形式运行golang

## 安装：

```
go get github.com/surfaceyu/godeamon
```

## 示例1:
// with only godeamon flags
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
	godeamon "github.com/surfaceyu/godeamon"
)

func main() {	
	godeamon.App.Run()

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":7070", mux))
}
```
## 示例2:
with flags in your apps 
```go
package main

import (
	"fmt"
	"log"
	"net/http"

	flag "github.com/spf13/pflag"
	godeamon "github.com/surfaceyu/godeamon"
)

func main() {
	goDaemonF := flag.Bool("f", false, "run app with -f.")
	godeamon.App.Run()
	fmt.Println("=============>", *goDaemonF)

	mux := http.NewServeMux()
	mux.HandleFunc("/index", func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte("hello, golang!\n"))
	})
	log.Fatalln(http.ListenAndServe(":7070", mux))
}

## 运行

```
./example -d
~$ curl http://127.0.0.1:7070/index
hello, golang!
