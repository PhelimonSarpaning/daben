package extract

import (
	"regexp"
)

// func main() {
// 	myText := "Angela Jan dec Dorothea Kasner, 1st January,2014 better known as Angela Merkel, was born in Hamburg, West Germany, on July 17, 1954. Trained as a physicist, Merkel entered politics after the 1989 fall of the Berlin Wall.2nd June, 2015 Rising to the position of chairwoman of the Christian Democratic Union party, Merkel became Germany's first female chancellor and one of the leading figures of the European Union, following the 2005 national elections 20-May-2019."
// 	pick := pickMonth(myText)
// 	fmt.Println(pick)
// }

func pickYear(text string) []string {
	re := regexp.MustCompile(`(?m) \d{4} `)

	var yearArray []string
	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			yearArray = append(yearArray, match, ", ")
		}
	}

	return yearArray
}

func pickMonth(text string) []string {
	re := regexp.MustCompile(`(?m) Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|June?|
	| July?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?`)

	var monthArray []string
	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			monthArray = append(monthArray, match, ", ")
		}
	}

	return monthArray
}

func pickDate(text string) []string {
	/*
				-- current date formats supported --
				 month dd, yyyy
				 yyyy
				 dd-month-yyyy
				 dd/mm/yyyy
				 dd.mm.yyyy
				 mm^st month,yyyy
				 mm^st month, yyyy
				 dd/month/yyyy

				 -- you can add the following for more date format support --
				 [\d]{1,2}-[ADFJMNOS]\w*-[\d]{4} # for format  xx-may-xxxx
		         [\d]{1,2}-[ADFJMNOS]\w*-[\d]{1,2} # for format  xx-may-xx
		         [\d]{1,2}-[\d]{1,2}-[\d]{4}   # xx-xx-xxxx
		         [\d]{1,2}/[\d]{1,2}/[\d]{1,2}   #xx/xx/xx
		         [\d]{1,2}-[\d]{1,2}-[\d]{1,2} #xx-xx-xx
		         [\d]{1,2}/[\d]{1,2}/[\d]{4}   #xx/xx/xxxx
		         [\d]{1,2}.[\d]{1,2}.[\d]{4} # xx.xx.xxxx
		         [\d]{1}/[\d]{1,2}/[\d]{1,2} #x/xx/xx
		         [\d]{1,2}-[ADFJMNOS]\w*-[\d]{1,2}  # for format  xx-may-xx
				 [\d]{1,2} [ADFJMNOS]\w* [\d]{4} # xx full month name xxxx
		         [\d]{1}/[\d]{1,2}/[\d]{1,2}  #x/xx/xx
		         [\d]{4}-[\d]{1,2}-[\d]{1,2} # xxxx-xx-xx
		         [ADFJMNOS]\w* [\d]{1,2}, [\d]{4} # for format  may xx, xxxx
	*/

	re := regexp.MustCompile(`[ADFJMNOS]\w* [\d]{1,2}, [\d]{4}| \d{4} |
	| [\d]{1,2}-[ADFJMNOS]\w*-[\d]{1,4}| [\d]{1,2}/[\d]{1,2}/[\d]{4} | 
	| [\d]{1,2}.[\d]{1,2}.[\d]{4} |
	| [\d{1,2}](?:st|th|nd|rd) [ADFJMNOS]\w*,[\d]{4}| 
	|[\d{1,2}](?:st|th|nd|rd) [ADFJMNOS]\w*, [\d]{4} | 
	| [\d]{1,2}/[ADFJMNOS]\w*/[\d]{4} `)

	var dateArray []string
	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			dateArray = append(dateArray, match, ", ")
		}
	}

	return dateArray
}
