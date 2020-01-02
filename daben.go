package daben

import (
	"regexp"
)

// func main() {
// 	myText := ""
// 	pick := PickMonth(myText)
// 	fmt.Println(pick)
// }

//PickYear :  this function helps you to pick out the years in a sparse text
func PickYear(text string) []string {
	var yearArray []string
	if len(text) == 0 {
		errorMessage := "error : String cannot be empty"
		return []string{errorMessage}
	}

	re := regexp.MustCompile(`(?m) \d{4} `)

	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			yearArray = append(yearArray, match)
		}
	}

	return yearArray

}

//PickMonth :  this function helps to pick months in a sparse text
func PickMonth(text string) []string {
	if len(text) == 0 {
		return []string{"String cannot be empty"}
	}

	re := regexp.MustCompile(`(?m) Jan(?:uary)?|Feb(?:ruary)?|Mar(?:ch)?|Apr(?:il)?|May|June?|
	| July?|Aug(?:ust)?|Sep(?:tember)?|Oct(?:ober)?|Nov(?:ember)?|Dec(?:ember)?`)

	var monthArray []string
	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			monthArray = append(monthArray, match)
		}
	}

	return monthArray

}

//PickDate :  this function should be used to pick all date formats from a sparse text
func PickDate(text string) []string {
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
	if len(text) == 0 {
		return []string{"String cannot be empty"}
	}

	re := regexp.MustCompile(`[ADFJMNOS]\w* [\d]{1,2}, [\d]{4}| \d{4} |
	| [\d]{1,2}-[ADFJMNOS]\w*-[\d]{1,4}| [\d]{1,2}/[\d]{1,2}/[\d]{4} | 
	| [\d]{1,2}.[\d]{1,2}.[\d]{4} |
	| [\d{1,2}](?:st|th|nd|rd) [ADFJMNOS]\w*,[\d]{4}| 
	|[\d{1,2}](?:st|th|nd|rd) [ADFJMNOS]\w*, [\d]{4} | 
	| [\d]{1,2}/[ADFJMNOS]\w*/[\d]{4} `)

	var dateArray []string
	for _, match := range re.FindAllString(text, -1) {
		if match != "" {
			dateArray = append(dateArray, match)
		}
	}

	return dateArray

}
