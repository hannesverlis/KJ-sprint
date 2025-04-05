package sprint
//tehtud 5 juuni kl 15:28

func IsNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {  // ekomtrollivada numbrid tuleb panna märkidesse, kuna tegemist on stringiga. INt võrdluses läheb ilma ülakomata. 
			return false   //juhl kui r on väiksem kui 0 ja suurem kui 9, on tegemist mingi märgiga, mis ei ole number. 
		}
	}
	return true
}