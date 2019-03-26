# quick start
### install golang 
https://golang.org/doc/install

### setting go root & path

```shell 
export GOROOT=$HOME/go
export PATH=$PATH:$GOPATH/bin
```  

### install tool packages
```shell
go get -u golang.org/x/tools/cmd/goimports
go get -u golang.org/x/tools/cmd/gorename
go get -u github.com/sqs/goreturns
go get -u github.com/nsf/gocode
go get -u github.com/alecthomas/gometalinter
go get -u github.com/zmb3/gogetdoc
go get -u github.com/rogpeppe/godef
go get -u golang.org/x/tools/cmd/guru
```

### install vgo package management & install dependency packages
https://github.com/golang/vgo
```shell
go get -u golang.org/x/vgo

// install dependency packages
dep ensure
```

### install live reload and task runner tool
https://github.com/oxequa/realize

```shell 
go get github.com/oxequa/realize
```

### init realize config to project
```shell
realize init
```

###### sample for realize config
```shell 
// cat .realize.yaml

settings:
  files:
    outputs:
      status: false
      path: ""
      name: .r.outputs.log
    logs:
      status: false
      path: ""
      name: .r.logs.log
    errors:
      status: false
      path: ""
      name: .r.errors.log
  legacy:
    force: false
    interval: 0s
schema:
  - name: gapi
    path: /Users/tk.kim/go/src/gapi
    commands:
      fmt:
        status: true
      install:
        status: true
      build:
        status: true
      run:
        status: true
    watcher:
      paths:
      - /
      extensions:
      - go
      ignored_paths:
      - .git
      - .realize
      - vendor
```

### run it !!!
```shell
realize start

[09:56:48][gapi] : Install started
[09:56:51][gapi] : Install completed in 2.865 s
[09:56:51][gapi] : Build started
[09:56:52][gapi] : Build completed in 1.196 s
[09:56:52][gapi] : Running..
[09:56:52][gapi] : [GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
[09:56:52][gapi] :  - using env:      export GIN_MODE=release
[09:56:52][gapi] :  - using code:     gin.SetMode(gin.ReleaseMode)
[09:56:52][gapi] : [GIN-debug] POST   /ptn/login                --> gapi/controllers.(*Auth).Login-fm (3 handlers)
[09:56:52][gapi] : [GIN-debug] POST   /ptn/logout               --> gapi/controllers.(*Auth).Logout-fm (4 handlers)
[09:56:52][gapi] : [GIN-debug] GET    /ptn/auth/me              --> gapi/controllers.(*PartnerAdmin).Me-fm (4 handlers)
```

###### gin framework documents
https://godoc.org/github.com/gin-gonic/gin

###### gorm documents 
http://doc.gorm.io/

### controller method name define rule:
Index, Show, Store, Update, Destroy

### model method name define rule:
Index, Create, Find, Update, Delete

### request & response process
router -> middleware -> controller -> model -> controller -> browser

### DDD for RESTful API patten define rule:
https://www.youtube.com/watch?v=aQVSzMV8DWc 

``` 
GET /hotels -> R:Index -> M:Index
GET /hotels/1 -> R:Show -> M:Find
POST /hotels -> R:Store -> M:Create
PUT /hotels -> R:Update -> M:Update
DELETE /hotels -> R:Destroy -> M:Delete
```
