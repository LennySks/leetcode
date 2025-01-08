package main

import (
	"fmt"
)

func main() {
	fmt.Println(isAnagram("anagram", "nagaram"))
}

//	func isAnagram(s string, t string) bool {
//		hashS, hashT := make(map[string]int), make(map[string]int)
//		charsS := strings.Split(s, "")
//		charsT := strings.Split(t, "")
//		for _, letter := range charsS {
//			if _, hasLetter := hashS[letter]; hasLetter {
//				hashS[letter]++
//			} else {
//				hashS[letter] = 1
//			}
//		}
//		for _, letter := range charsT {
//			if _, hasLetter := hashT[letter]; hasLetter {
//				hashT[letter]++
//			} else {
//				hashT[letter] = 1
//			}
//		}
//		result := reflect.DeepEqual(hashS, hashT)
//		return result
//	}
func isAnagram(s string, t string) bool {
	chars := make([]int, 26)
	for _, v := range s {
		i := int(v - 'a')
		chars[i]++
	}

	for _, v := range t {
		i := int(v - 'a')
		chars[i]--
	}

	for _, v := range chars {
		if v != 0 {
			return false
		}
	}

	return true
}
