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

    // 创建HQL语句
    hql := "CREATE TABLE table_name (\n"
    for i := range fields {
        hql += fmt.Sprintf("`%s` STRING COMMENT '%s',\n", fields[i], comments[i])
    }
    hql = hql[:len(hql)-2] // 去掉最后一个逗号
    hql += "\n) ROW FORMAT SERDE 'org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe'"
    hql += "\nWITH SERDEPROPERTIES ("
    hql += "\n\"field.delim\"='\\001',"
    hql += "\n\"escape.delim\"='\\n',"
    hql += "\n\"colelction.delim\"=',',"
    hql += "\n\"mapkey.delim\"=':'"
    hql += "\n)"
    hql += "\nstored as TEXTFILE;"

    // 写入hql文件
    hqlFile, err := os.Create("data.hql")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer hqlFile.Close()

    _, err = hqlFile.WriteString(hql)
    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println("HQL文件已生成")
}
