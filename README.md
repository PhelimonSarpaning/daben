## daben
A Go package for picking dates from text.

[![GoDoc](https://godoc.org/github.com/PhelimonSarpaning/daben?status.svg)](https://godoc.org/github.com/PhelimonSarpaning/daben)

Daben is an Ghanaian Akan language which literally means "when".

## Install

```bash
go get github.com/PhelimonSarpaning/daben
```
## Usage
```
$ PickYear(text) : picks all the years in a text and outputs it as a slice
$ PickMonth(text) : picks all the months in a text and outputs it as a slice
$ PickDate(text) : picks all the date formats in a text and outputs it as a slice
```
## Example

```go
package main

import (
	"fmt"
	"log"

	"github.com/PhelimonSarpaning/daben"
)

func main() {
    text := "I know of a man who was born on July 17, 1954. Trained as a theologian but entered politics after the year 1959. In 2nd June, 2015 he rose to the position as chairmain of the peace council of the world, Ganel became Badain's first preacher, following the 2005 national elections 20 - May- 2019."
    
    //to get date of any format from a text
    allDates := daben.PickDate(text)
    fmt.Println(allDates) // to print all of them as an array
    
    //printing individual dates
    for _, date := range allDates {
        fmt.Println(date)
    }
    
    
    //to get months from a text
    allMonths := daben.PickMonth(text)
    fmt.Println(allMonths) // to print all of them as an array
    
    //printing individual months
    for _, month := range allMonths {
        fmt.Println(month)
    }

    //to get years from a text
    allYears := daben.PickYear(text)
    fmt.Println(allYears) // to print all of them as an array
    
    //printing individual years
    for _, year := range allYears {
        fmt.Println(year)
    }

    


}

```