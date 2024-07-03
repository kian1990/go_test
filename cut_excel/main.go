package main

import (
    "fmt"
    "github.com/xuri/excelize/v2"
)

func main() {
    // 打开原始Excel文件
    f, err := excelize.OpenFile("in.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }

    // 获取所有的Sheet名称
    sheets := f.GetSheetList()
    if len(sheets) == 0 {
        fmt.Println("No sheets found")
        return
    }

    // 获取第一个Sheet的名称
    sheet := sheets[0]

    // 获取所有行
    rows, err := f.GetRows(sheet)
    if err != nil {
        fmt.Println(err)
        return
    }

    // 设定每个子文件的行数
    rowsPerFile := 500
    totalRows := len(rows)
    fileIndex := 1

    for i := 0; i < totalRows; i += rowsPerFile {
        // 创建一个新的Excel文件
        newFile := excelize.NewFile()
        newSheetIndex, err := newFile.NewSheet("Sheet1")
        if err != nil {
            fmt.Println(err)
            return
        }

        // 复制行到新的Excel文件中
        for j := i; j < i+rowsPerFile && j < totalRows; j++ {
            newFile.SetSheetRow("Sheet1", fmt.Sprintf("A%d", j-i+1), &rows[j])
        }

        // 设置为活跃Sheet
        newFile.SetActiveSheet(newSheetIndex)

        // 保存新的Excel文件
        outputFileName := fmt.Sprintf("out_%d.xlsx", fileIndex)
        if err := newFile.SaveAs(outputFileName); err != nil {
            fmt.Println(err)
            return
        }

        fileIndex++
    }

    fmt.Println("拆分完成")
}
