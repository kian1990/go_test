package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"

    _ "github.com/go-sql-driver/mysql"
    "github.com/tealeg/xlsx"
)

// 定义一个结构体来表示数据库中的一行数据
type User struct {
    Name  string `json:"名称"`
    Count string `json:"数量"`
}

// 连接数据库并查询数据
func getUsers(w http.ResponseWriter, r *http.Request) {
    // 设置数据库连接参数
    dsn := "root:root@tcp(127.0.0.1:3306)/test"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 查询数据库
    rows, err := db.Query("SELECT SUBSTR(GKLX,1,4),COUNT(SUBSTR(GKLX,1,4)) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY SUBSTR(GKLX,1,4) ORDER BY item_count DESC")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // 遍历查询结果
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.Name, &user.Count)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    // 检查是否有错误
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    // 将数据转换为 JSON 并写入响应
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

// 将 JSON 数据转换为 Excel 文件
func jsonToExcel(w http.ResponseWriter, r *http.Request) {
    // 设置数据库连接参数
    dsn := "root:root@tcp(127.0.0.1:3306)/test"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 查询数据库
    rows, err := db.Query("SELECT SUBSTR(GKLX,1,4),COUNT(SUBSTR(GKLX,1,4)) AS item_count FROM `xnsdsjj_12345cbs` GROUP BY SUBSTR(GKLX,1,4) ORDER BY item_count DESC")
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // 遍历查询结果
    var users []User
    for rows.Next() {
        var user User
        err := rows.Scan(&user.Name, &user.Count)
        if err != nil {
            log.Fatal(err)
        }
        users = append(users, user)
    }

    // 检查是否有错误
    err = rows.Err()
    if err != nil {
        log.Fatal(err)
    }

    // 创建一个新的 Excel 文件
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("Users")
    if err != nil {
        log.Fatal(err)
    }

    // 添加表头
    row := sheet.AddRow()
    row.AddCell().Value = "名称"
    row.AddCell().Value = "数量"

    // 添加数据
    for _, user := range users {
        row := sheet.AddRow()
        row.AddCell().Value = user.Name
        row.AddCell().Value = string(user.Count)
    }

    // 设置响应头
    w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
    w.Header().Set("Content-Disposition", "attachment; filename=users.xlsx")

    // 写入文件
    err = file.Write(w)
    if err != nil {
        log.Fatal(err)
    }
}

func main() {
    // 设置 HTTP 路由
    http.HandleFunc("/", getUsers)
    http.HandleFunc("/excel", jsonToExcel)

    // 启动 HTTP 服务器
    log.Println("浏览器访问：http://127.0.0.1 查看")
    log.Println("浏览器访问：http://127.0.0.1/excel 下载")
    log.Fatal(http.ListenAndServe(":80", nil))
}
