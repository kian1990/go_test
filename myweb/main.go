package main

import (
    "fmt"
    "net/http"
    "io/ioutil"
    "path/filepath"
    "myweb/controllers"
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/sessions"
    "github.com/russross/blackfriday/v2"
    "github.com/gin-contrib/sessions/cookie"
)

func main() {
    router := gin.Default()
    router.LoadHTMLGlob("views/*")

    // 设置session存储
    store := cookie.NewStore([]byte("secret"))
    router.Use(sessions.Sessions("mysession", store))

    // 主页
    router.GET("/", controllers.AuthRequired(), func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H {"title": "主页"})
    })

    // 退出
    router.GET("/logout", controllers.AuthRequired(), func(c *gin.Context) {
        session := sessions.Default(c)
        session.Clear()
        session.Save()
        c.Redirect(http.StatusFound, "/login")
    })

    // 文档页面
    router.GET("/doc", controllers.AuthRequired(), func(c *gin.Context) {
        c.HTML(http.StatusOK, "doc.html", gin.H {"title": "文档"})
    })

    // 关于页面
    router.GET("/about", controllers.AuthRequired(), func(c *gin.Context) {
        c.HTML(http.StatusOK, "about.html", gin.H {"title": "关于"})
    })

    // 联系页面
    router.GET("/contact", controllers.AuthRequired(), func(c *gin.Context) {
        c.HTML(http.StatusOK, "contact.html", gin.H {"title": "联系"})
    })

    // 文件页面
    router.GET("/file", controllers.AuthRequired(), func(c *gin.Context) {
        folderPath := "./file"
        file, err := controllers.GetFile(folderPath)
        if err != nil {
            c.String(http.StatusInternalServerError, "读取文件错误: %v", err)
            return
        }
        c.HTML(http.StatusOK, "file.html", gin.H {"file": file})
    })
    // 处理文件上传的路由
    router.POST("/file", controllers.AuthRequired(), func(c *gin.Context) {
        file, err := c.FormFile("file")
        if err != nil {
            c.String(http.StatusBadRequest, fmt.Sprintf("上传错误: %s", err.Error()))
            return
        }
        // 保存文件到指定目录
        dst := filepath.Join("./file", file.Filename)
        if err := c.SaveUploadedFile(file, dst); err != nil {
            c.String(http.StatusInternalServerError, fmt.Sprintf("上传错误: %s", err.Error()))
            return
        }
        // c.String(http.StatusOK, fmt.Sprintf("文件: %s 上传成功", file.Filename))
        c.Redirect(http.StatusFound, "/file")
    })

    // Markdown文件渲染
    router.GET("/file/:filename", controllers.AuthRequired(), func(c *gin.Context) {
        filename := c.Param("filename")
        // 设置Markdown文件所在目录
        dir := "./file"
        // 组合文件路径
        filepath := filepath.Join(dir, filename)
        // 读取文件内容
        mdContent, err := ioutil.ReadFile(filepath)
        if err != nil {
            c.String(http.StatusNotFound, "文件不存在")
            return
        }
        // 渲染Markdown
        html := blackfriday.Run(mdContent)
        // 返回HTML内容
        c.Data(http.StatusOK, "text/html; charset=utf-8", html)
    })

    // 显示登录页面
    router.GET("/login", controllers.LoginGet)
    // 处理登录请求
    router.POST("/login", controllers.LoginPost)

    // 显示注册页面
    router.GET("/register", controllers.AuthRequired(), controllers.RegisterGet)
    // 处理注册请求
    router.POST("/register", controllers.AuthRequired(), controllers.RegisterPost)

    router.Static("/static", "./static")
    router.StaticFile("/favicon.ico", "./static/img/kian.ico")
    router.Run("0.0.0.0:80")
}