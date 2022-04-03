# gin-vue

##  $\delta \ 1$ goproject
###  $\delta \ 1.1$  Package Install


    go get -u github.com/gin-gonic/gin           #gin框架
    go get -u github.com/jinzhu/gorm             #gorm
    go get -u github.com/go-sql-driver/mysql     #mysql驱动
    go get -u github.com/dgrijalva/jwt-go        #jwt鉴权
    go get -u golang.org/x/crypto/bcrypt         #bcrypt提供解析token

###  $\delta \ 1.2$  Project Framework

    ├── common
    │   ├── database.go                          #数据库连接
    │   └── jwt.go                               #生成以及解析token
    ├── controller
    │   └── UserContorller.go                    #用户handle
    ├── dto
    │   └── user_dto.go                          #用户可见返回数据
    ├── gin-vue
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── middleware
    │   └── AuthMiddleware.go                    #中间件 配合jwt用户鉴权
    ├── model
    │   └── user.go                              #用户类
    ├── response
    │   └── response.go                          #封装统一返回格式
    ├── routes.go                                #路由
    └── util
    └── util.go                                  #工具包









