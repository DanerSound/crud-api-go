package validators

import "regexp"

func IsEmailValid(email string) bool{
	var emailRegex = regexp.MustCompile(^([A-Za-z][0-9])+$) // toFIX 

	if len(email)<3 || len(email)> 254 || emailRegex.MatchString(email){
		return false
	} 
	return true
}