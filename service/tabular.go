package service

import (
    "fmt"
    "errors"
)

func printDivider(width int) {
    for j := 0; j < width; j++ {
        fmt.Print("-")
    }
    fmt.Println()
}

/*
    Public function taht expects a slice of maps as argument when called along
    with the width of the columns of the generated table
    Assumptions (TODO later rather)
    1. widths will contain width of all columns. No checks carried out here
    2. keys of the map will be used as the columns. Each map will have all the keys
        No checks carried out here.
    3. All keys will be used. Need to provide parameter to specify which keys to be used
        as columns
*/
func Display(data []map[string]int, widths []int) error {
    //Calculate the width of the table
    //For each column, we will have an additional character
    //to the left of the column
    width := (len(widths) + 1) + (2 * len(widths))

    for _, w := range widths {
        if width + w > 255 {
            return errors.New("Width of the table exceeds 255. Cannot display table with such a large width")
        }
        width += w
    }

    //Start displaying the table
    //First display the header rows
    //Start with the display of the top of the header row
    printDivider(width)

    //Print the beginning of the first column
    fmt.Print("|")

    //Now print the column headers
    //Use the first map's key - assume that it contains all the key columns
    for key, _ := range data[0] {
        fmt.Print(" ", key, " |")
    }


    //Move to the next row
    //Here we being printing the contents of each row
    fmt.Println()

    printDivider(width)

    for _, row := range data {
        //Print the column beginning
        fmt.Print("|")
        for _, value := range row {
            fmt.Print(" ", value, " |")
        }
        //Print the column end
        fmt.Println()
        printDivider(width)
    }

    return nil
}
