# gowebapp

- [ ] Курс первый [вводный](https://www.coursera.org/learn/golang-webservices-1/home/welcome)
- [ ] Курс второй [вебный](https://www.coursera.org/learn/golang-webservices-2/lecture/jLqQf/o-chiem-etot-kurs)
- [ ] Тот же курс, вид с [youtube](https://www.youtube.com/watch?v=9Ia16QOY8rk&index=2&list=PLrCZzMib1e9q-X5V9pTM6J0AemRWseM7I)

Top Needs dependencies

- Testing tool [goconvey](https://github.com/smartystreets/goconvey)
- Task Runner [realize](https://github.com/oxequa/realize)
- Http framework/router [chi](https://github.com/go-chi/chi)
- [Mongo driver](https://github.com/mongodb/mongo-go-driver)
- Go [graphql](https://github.com/graphql-go/graphql)

> Umputun
> я по библитекам минималист, но без фанатизма :) Иx на самом деле у меня совсем немного, вот список примерно всего, что используется (понабирал из разных проектов)
> 1. github.com/pressly/chi
> 2. github.com/dgrijalva/jwt-go
> 3. github.com/stretchr/testify
> 4. gopkg.in/mgo.v2
> 5. gopkg.in/yaml.v2
> 6. github.com/didip/tollbooth
> 7. github.com/juju/ratelimit
> 8. github.com/hashicorp/go-mul...
> 9. github.com/aws/aws-sdk-go/...
> 10. github.com/dustin/go-humanize
> 11. github.com/jessevdk/go-flags
> 12. golang.org/x/sync/errgroup


Teach materials

- [ ] mongodb-driver-[official](https://github.com/mongodb/mongo-go-driver/blob/master/mongo/doc.go)
- [ ] [Awesome go](https://github.com/avelino/awesome-go)
- [ ] Russian site with useful articles [4gophers](https://4gophers.ru)
- [ ] Books and screencast [stackoverflow](https://ru.stackoverflow.com/questions/436505/Книги-документация-статьи-и-курсы-по-go)
- [ ] Введение в [go](http://golang-book.ru)
- [ ] [Тур по го](https://go-tour-ru-ru.appspot.com/flowcontrol/2)

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


