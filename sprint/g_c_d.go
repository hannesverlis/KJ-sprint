package sprint

func GCD(a, b int) int {
if a < 0 || b < 0 {
	return 0
} 


if a == 0 {
	return b
}
if b == 0 {
	return a
}

for b != 0 {
	a, b = b, a%b 
}
return a
}
