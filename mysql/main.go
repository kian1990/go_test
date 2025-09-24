package main

import (
    "database/sql"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "log"
)

func main() {
    // 连接数据库
    dsn := "root:root@tcp(127.0.0.1:3306)/test"
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    // 确认数据库连接是否正常
    err = db.Ping()
    if err != nil {
        log.Fatal(err)
    }

    // 执行查询
    query := "SELECT HEX(AES_ENCRYPT('你好', 'abcdefg123456789'));"
    rows, err := db.Query(query)
    if err != nil {
        log.Fatal(err)
    }
    defer rows.Close()

    // 执行查询
    var decryptedValue []byte
    err = db.QueryRow(query).Scan(&decryptedValue)
    if err != nil {
        log.Fatal(err)
    }

    // 输出解密结果
    fmt.Printf("%s", string(decryptedValue))

    // 检查查询过程中是否有错误
    if err = rows.Err(); err != nil {
        log.Fatal(err)
    }
}
