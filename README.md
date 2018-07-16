# gowebapp

Top Needs dependencies

- Testing tool [goconvey](http://goconvey.co)
- Task Runner [realize](https://gorealize.io)
- Http framework/router [chi](https://github.com/go-chi/chi)
- [Mongo driver](https://github.com/mongodb/mongo-go-driver)

Teach materials

- mongodb-driver-[official](https://github.com/mongodb/mongo-go-driver/blob/master/mongo/doc.go)
- [Awesome go](https://github.com/avelino/awesome-go)
- Russian site with useful articles [4gophers](https://4gophers.ru)
- Books and screencast [stackoverflow](https://ru.stackoverflow.com/questions/436505/Книги-документация-статьи-и-курсы-по-go)

Ready packages

- Static generator sites [hugo](https://gohugo.io)

Simplistic Go Web App

Read about it [here](https://grisha.org/blog/2017/04/27/simplistic-go-web-app/).

To install:

1. You need PostgreSQL, install it somehow, make an empty database
   called `gowebapp`. (Usually `createdb gowebapp` on the command line is
   all you need).

1. You need Go, preferably the latest version (1.8).

1. Now do this:

```
$ export GOPATH=~/go # optional, adjust as necessary
$ go get github.com/grisha/gowebapp
```

1. That's it. You should now be able to run:
```
$ $GOPATH/bin/gowebapp
```
