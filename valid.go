package noppe

import (
	// "fmt"
	"strings"
	"regexp"
)


var zipnreg *regexp.Regexp
var ziphreg *regexp.Regexp
var telnreg *regexp.Regexp
var telhreg *regexp.Regexp
var mailreg *regexp.Regexp
var nbrnreg *regexp.Regexp
var nbrfreg *regexp.Regexp


func init() {

	zipnreg = regexp.MustCompile(`^[0-9]{7}$`)
	ziphreg = regexp.MustCompile(`^[0-9]{3}-{1}[0-9]{4}$`)

	telnreg = regexp.MustCompile(`^[0-9]{10,11}$`)
	telhreg = regexp.MustCompile(`^[0-9]{3,4}-{1}[0-9]{2,4}-{1}[0-9]{4}$`)

	mailreg = regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@((?:[-a-z0-9]+\.)+[a-z]{2,})$`)

	nbrnreg = regexp.MustCompile(`^[1-9]+[0-9]?$`)
	nbrfreg = regexp.MustCompile(`^(0|[1-9]+)\.[0-9]+$`)
}



func IsNotEmpty(input *interface{}) bool {

	if input == nil {
		return false
	}

	res := true

	in := *input

	switch in.(type) {
		case string :

			val := in.(string)
			if len(val)==0 {
				res = false
			}

	}




	return res

}



func IsAllowedChar(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	xx := []rune(input)

	b := true 

	for _,x := range xx {

		if 0x00 < x && x <= 0x1F {

			if x != 0x0A && x != 0x0D {
				// fmt.Printf("%d %d %c %x \n",i,x,x,x)
				b=false
				// return false
			}

		} else if 0x20 <= x && x <= 0x7F {
			
			if x==0x27 || x==0x2C || x==0x3B || x==0x3C || x==0x3E {
				// fmt.Printf("%d %d %c %x \n",i,x,x,x)
				b=false
				// return false
			}

		} else if 0xFF06 < x && x < 0xFFA0 {
			// fmt.Printf("%d %d %c %x \n",i,x,x,x)
			b=false
			// return false;
		}

	} 

	return b
}




func IsFullKana(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	xx := []rune(input)

	b := true 

	for _,x := range xx {
		if x < 0x30A1 || 0x30F6 < x {
			// fmt.Printf("%d %d %c %x \n",i,x,x,x)
			b=false
			// return false
		}
	} 

	return b
}


func IsTel(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	res := telnreg.MatchString(input)
	
	if res {
		return true
	}

	return telhreg.MatchString(input)
}



func IsZipcode(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	res := zipnreg.MatchString(input)

	if res {
		return true
	}

	return ziphreg.MatchString(input)
}



func IsMailadr(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	return mailreg.MatchString(input)
}



func IsNumeric(input string,require bool) bool {

	if !require && len(input)==0 {
		return true
	}

	if input == "0"  {
		return true
	}

	p := strings.Split(input,".")

	if len(p) == 1 {
		return nbrnreg.MatchString(input)
	}

	if len(p) == 2 {
		return nbrfreg.MatchString(input)
	}

	return false

}
