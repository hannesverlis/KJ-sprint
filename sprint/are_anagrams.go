package sprint

import (
    "sort"
    "strings"
)

func AreAnagrams(str1, str2 string) bool {
    // 1. Normalize Strings:
    //    - Convert to lowercase for case-insensitivity
    //    - Remove whitespace to disregard spaces
    str1 = strings.ToLower(strings.ReplaceAll(str1, " ", ""))
    str2 = strings.ToLower(strings.ReplaceAll(str2, " ", ""))

    // 2. Check Length:
    //    If lengths are different, they can't be anagrams
    if len(str1) != len(str2) {
        return false
    }

    // 3. Convert to Runes and Sort:
    //    - Convert strings to runes for easier sorting
    //    - Sort the rune slices alphabetically
    r1 := []rune(str1)
    r2 := []rune(str2)
    sort.Slice(r1, func(i, j int) bool { return r1[i] < r1[j] })
    sort.Slice(r2, func(i, j int) bool { return r2[i] < r2[j] })

    // 4. Compare Sorted Runes:
    //    If the sorted rune slices are identical, the strings are anagrams
    return string(r1) == string(r2) 
}