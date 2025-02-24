package main

import (
    "fmt"
    "os"
    "github.com/xuri/excelize/v2"
)

func main() {
    // 打开xlsx文件
    file, err := excelize.OpenFile("data.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    // 选择第一个sheet
    sheetName := file.GetSheetName(0)
    rows, err := file.GetRows(sheetName)
    if err != nil {
        fmt.Println(err)
        return
    }

    if len(rows) < 2 {
        fmt.Println("文件至少需要两行：第一行为字段信息，第二行为字段注释")
        return
    }

    fields := rows[0]
    comments := rows[1]

    if len(fields) != len(comments) {
        fmt.Println("字段信息和字段注释行的列数不一致")
        return
    }

    // 创建SQL语句
    sql := "CREATE TABLE `test`.`table_name` (\n"
    for i := range fields {
        sql += fmt.Sprintf("`%s` VARCHAR(255) COMMENT '%s',\n", fields[i], comments[i])
    }
    sql = sql[:len(sql)-2] // 去掉最后一个逗号
    sql += "\n) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;"

    // 写入sql文件
    sqlFile, err := os.Create("data.sql")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer sqlFile.Close()

    _, err = sqlFile.WriteString(sql)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("SQL文件已生成")
}
