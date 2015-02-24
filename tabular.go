package tabular

import (
	"errors"
	"fmt"
)

/*
   Prints the row divider (--------)
*/
func printDivider(width int) {
	for j := 0; j < width; j++ {
		fmt.Print("-")
	}
	fmt.Println()
}

/*
   Given a string and a width, this function prints spaces for
   the remaining width not occupied by the string
*/
func printRemainingWidth(width int, value string) {
	//Determine the remaining width
	rem := width - len(value)
	for j := 0; j < rem; j++ {
		fmt.Print(" ")
	}
}

/*
   Public function taht expects a slice of maps as argument when called along
   with the width of the columns of the generated table
*/
func Display(data []map[string]string, widths []int) error {
	//Calculate the width of the table
	//Each cell will have a space to its left and right (within the cell)
	//Also, each cell will have a wall to its left and right
	width := (len(widths) + 1) + (2 * len(widths))

	//Add the width of the column headers itself
	for _, w := range widths {
		if width+w > 255 {
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
	//Use the first map's keys - assume that it contains all the key columns

	//Keep track of the number of keys
	i := 0
	for key, _ := range data[0] {
		//Print the key seperated by a space to the left
		fmt.Print(" ", key)
		//Fill the remaining width (for the width of the column provided) with spaces
		printRemainingWidth(widths[i], key)
		//Print the column seperator with a space
		fmt.Print(" |")

		i++
	}

	//Move to the table data
	//Here we being printing the contents of each row
	fmt.Println()

	printDivider(width)

	for _, row := range data {
		//Print the column beginning
		fmt.Print("|")
		//Track the column number
		i = 0
		for _, value := range row {
			fmt.Print(" ", value)
			printRemainingWidth(widths[i], value)
			fmt.Print(" |")
			i++
		}
		//Move to the next row
		fmt.Println()
		//Print a divider
		printDivider(width)
	}

	return nil
}
