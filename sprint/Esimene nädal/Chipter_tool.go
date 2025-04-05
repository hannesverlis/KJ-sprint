package main

func main() {
    reader := bufio.NewReader(os.Stdin)

func rot13(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return (r-'a'+13)%26 + 'a'
	case r >= 'A' && r <= 'Z':
		return (r-'A'+13)%26 + 'A'
	default:
		return r
	}
}

func reverseAlphabet(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return 'z' - (r - 'a')
	case r >= 'A' && r <= 'Z':
		return 'Z' - (r - 'A')
	default:
		return r
	}
}

func atbash(r rune) rune {
	switch {
	case r >= 'a' && r <= 'z':
		return 'z' - (r - 'a')
	case r >= 'A' && r <= 'Z':
		return 'Z' - (r - 'A')
	default:
		return r
	}
}

// ... rest of the code (encryptMessage, decryptMessage, main)
