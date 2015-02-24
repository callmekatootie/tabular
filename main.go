package main

import "github.com/callmekatootie/tabular/service"

func main() {
    row := map[string]int{
        "key1": 5,
        "key2": 6,
        "key3": 7,
        "key4": 8,
    }

    widths := []int{4, 4, 4, 4}

    table := []map[string]int{row}

    service.Display(table, widths)
}
