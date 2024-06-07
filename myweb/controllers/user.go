package controllers

import (
    
    "log"
    "net/http"
    "database/sql"
    "github.com/gin-gonic/gin"
    "golang.org/x/crypto/bcrypt"
    "github.com/gin-contrib/sessions"
    _ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// 初始化函数，用于连接到数据库
func init() {
    var err error

    // 连接到 MySQL 数据库
    dsn := "root:root@tcp(127.0.0.1:3306)/myweb"
    db, err = sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }

    // 确保数据库连接是有效的
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }
}

// AuthRequired 中间件，用于保护需要认证的路由
func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        session := sessions.Default(c)
        user := session.Get("user")
        if user == nil {
            c.Redirect(http.StatusFound, "/")
            c.Abort()
            return
        }
        c.Next()
    }
}

// 显示登录页面
func LoginGet(c *gin.Context) {
    c.HTML(http.StatusOK, "login.html", gin.H{
        "title": "登录",
    })
}

// 处理登录请求
func LoginPost(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    // 检查用户是否存在
    var storedPassword string
    err := db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "用户或密码错误"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
        }
        return
    }

    // 检查密码是否匹配
    err = bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(password))
    if err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "用户或密码错误"})
        return
    }

    // c.JSON(http.StatusOK, gin.H{"message": "登录成功"})

    // 登录成功，设置会话
    session := sessions.Default(c)
    session.Set("user", username)
    session.Save()
    c.Redirect(http.StatusFound, "/")
}

// 显示注册页面
func RegisterGet(c *gin.Context) {
    c.HTML(http.StatusOK, "register.html", gin.H{
        "title": "注册",
    })
}

// 处理注册请求
func RegisterPost(c *gin.Context) {
    username := c.PostForm("username")
    password := c.PostForm("password")

    // 检查用户名是否已存在
    var existingUser string
    err := db.QueryRow("SELECT username FROM users WHERE username = ?", username).Scan(&existingUser)
    if err != sql.ErrNoRows {
        c.JSON(http.StatusBadRequest, gin.H{"error": "用户已存在"})
        return
    }

    // 哈希用户密码
    hashedPassword, err := HashPassword(password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
        return
    }

    // 将用户保存到数据库
    _, err = db.Exec("INSERT INTO users (username, password) VALUES (?, ?)", username, hashedPassword)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "内部错误"})
        return
    }

    // c.JSON(http.StatusOK, gin.H{"message": "注册成功"})
    c.Redirect(http.StatusFound, "/login")
}

// 密码哈希函数
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    if err != nil {
        return "", err
    }
    return string(bytes), nil
}
