package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/xuri/excelize/v2"
)

func main() {
    // 打开Excel文件
    f, err := excelize.OpenFile("data.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }

    // 获取所有行和列
    rows, err := f.GetRows("Sheet1")
    if err != nil {
        fmt.Println(err)
        return
    }

    // 遍历行，按A和B分类
    dataByAB := make(map[string]map[string][][]string)
    for _, row := range rows[1:] {
        A := row[3]  // A字段
        B := row[5]  // B字段

        if _, ok := dataByAB[A]; !ok {
            dataByAB[A] = make(map[string][][]string)
        }

        if _, ok := dataByAB[A][B]; !ok {
            dataByAB[A][B] = [][]string{rows[0]} // 加入标题行
        }

        dataByAB[A][B] = append(dataByAB[A][B], row)
    }

    // 创建文件夹并保存新表格
    for A, BS := range dataByAB {
        APath := filepath.Join(".", A)
        os.MkdirAll(APath, os.ModePerm)

        for B, data := range BS {
            BFile := filepath.Join(APath, fmt.Sprintf("%s.xlsx", B))
            newFile := excelize.NewFile()
            sheetIndex, err := newFile.NewSheet("Sheet1")
            if err != nil {
                fmt.Println(err)
                continue
            }

            // 创建黄色背景样式
            yellowFillStyleID, err := newFile.NewStyle(&excelize.Style{
                Fill: excelize.Fill{
                    Type:    "pattern",
                    Pattern: 1,
                    Color:   []string{"FFFF00"}, // 黄色背景
                },
            })
            if err != nil {
                fmt.Println(err)
                continue
            }

            // 创建宋体样式
            fontStyleID, err := newFile.NewStyle(&excelize.Style{
                Font: &excelize.Font{
                    Family: "SimSun", // 设置字体为宋体
                },
            })
            if err != nil {
                fmt.Println(err)
                continue
            }

            // 复制样式和数据
            for i, row := range data {
                for j, colCell := range row {
                    cell, _ := excelize.CoordinatesToCellName(j+1, i+1)
                    newFile.SetCellValue("Sheet1", cell, colCell)

                    // 应用宋体样式
                    newFile.SetCellStyle("Sheet1", cell, cell, fontStyleID)
                }
            }

            // 应用黄色背景从第一行11开始到最后
            for col := 11; col <= len(rows[0]); col++ {
                cell, _ := excelize.CoordinatesToCellName(col, 1)
                newFile.SetCellStyle("Sheet1", cell, cell, yellowFillStyleID)
            }

            newFile.SetActiveSheet(sheetIndex)
            if err := newFile.SaveAs(BFile); err != nil {
                fmt.Println(err)
            }
        }
    }

    fmt.Println("处理完成！")
}
