package sprint


func StrCompare(a , b string) int {

	for i := 0; i < len(a) && i < len(b); i++ {
        if a[i] < b[i] {
            return -1
        } else if a[i] > b[i] {
            return 1
        }
    }

    // Handle cases where one string is a prefix of the other
    if len(a) < len(b) {
        return -1
    } else if len(a) > len(b) {
        return 1
    }

    return 0
}