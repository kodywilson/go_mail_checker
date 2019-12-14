package main

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    f, err := excelize.OpenFile("/Users/kowilson/code/golang/go_mail_checker/Book1.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
    cell := f.GetCellValue("Sheet1", "B2")
    //if err != nil {
    //    fmt.Println(err)
    //    return
    //}
    fmt.Println(cell)
    // Get all the rows in the Sheet1.
    rows := f.GetRows("Sheet2")
    for _, row := range rows {
        for _, colCell := range row {
            fmt.Print(colCell, "\t")
        }
        fmt.Println()
    }
}