package main

import (
    "fmt"
    "github.com/360EntSecGroup-Skylar/excelize"
)

func main() {
    f, err := excelize.OpenFile("/Users/kowilson/Downloads/Market Assessment RFI - ELR.xlsx")
    if err != nil {
        fmt.Println(err)
        return
    }
    // Get value from cell by given worksheet name and axis.
    cell := f.GetCellValue("Sheet2", "B2")
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