package main

import (
	"fmt"
)

func main() {
	myText := "Angela Dorothea Kasner, 1st January,2014 better known as Angela Merkel, was born in Hamburg, West Germany, on July 17, 1954. Trained as a physicist, Merkel entered politics after the 1989 fall of the Berlin Wall.2nd June, 2015 Rising to the position of chairwoman of the Christian Democratic Union party, Merkel became Germany's first female chancellor and one of the leading figures of the European Union, following the 2005 national elections 20-May-2019."
	fmt.Println(extract.pickDate(myText))
}
