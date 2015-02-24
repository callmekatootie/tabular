# tabular
Display a map in a tabular fashion. Think console.table() in javascript (but this one's for golang and  
in the system console)

To import the library - 
```golang
import "github.com/callmekatootie/tabular"
```

# Usage

Consider two maps as follows:
```golang
row := map[string]string{
  "#": "1",
  "Name": "Hermione Granger",
  "Age": "21",
  "Gender": "M",
}

row2 := map[string]string{
  "#": "1",
  "Name": "Frodo Baggins",
  "Age": "64",
  "Gender": "M",
}
```

Now, consider a slice containing the above maps (to simulate a json response from the server)
```golang
table := []map[string]string{row, row2}

//expected width of the columns - consider the header / key width too
widths := []int{1, 20, 3, 6}
```

To display the slice of maps in a tabular fashion, call the library as:
```golang
tabular.Display(table, widths)
```

The output will be:  
![result](http://i.imgur.com/NXuMeFo.png)

# TODO
Assumptions (TODO later rather)  
1. widths will contain width of all columns. No checks carried out here  
2. keys of the map will be used as the columns. Each map will have all the keys  
No checks carried out here.  
3. All keys will be used. Need to provide parameter to specify which keys to be used  
as columns  
4. Need to accept a generic map. Currently accepts only `map[string]string`  

# License
MIT
