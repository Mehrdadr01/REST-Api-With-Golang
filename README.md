# RESTful api in golang

go mod init restapi
then create our main.go file inside main package

## create our web server

- write our web server
  https://pkg.go.dev/net/http#HandleFunc

http.HandleFunc(pattern for route, func handleRoute(writer http.ResponseWriter, req \* http.Request))
handleRoute(writer req){
costume error for diff URL route
}

- https://pkg.go.dev/net/http#ListenAndServe

## installing dependencies before making it RESTful

- go get github.com/asdine/storm/v3 ----> don't need to use GO111MODULE=on it is deprecated and also dont need to use gopath anymore for go 1.17 i think (https://github.com/asdine/storm)
  -- storm is toolkit for BoltBD
- go get gopkg.in/mgo.v2/bson

## users package

## handlers package
