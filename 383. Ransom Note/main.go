package main

// "strconv"

/*
easy
hashmap
done

Given two strings ransomNote and magazine,
return true if ransomNote can be constructed by using the letters from magazine and false otherwise.
Each letter in magazine can only be used once in ransomNote.

Input: ransomNote = "a", magazine = "b"
Output: false
Input: ransomNote = "aa", magazine = "ab"
Output: false
Input: ransomNote = "aa", magazine = "aab"
Output: true
*/

// func canConstruct(ransomNote string, magazine string) bool {
// 	t := make(map[int32]int)
// 	ans := true
// 	for _, i := range magazine {
// 		t[i]++
// 	}
// 	for _, i := range ransomNote {
// 		t[i]--
// 		if t[i] < 0 {
// 			ans = false
// 		}
// 	}

//		return ans
//	}
func canConstruct(ransomNote string, magazine string) bool {
	t := make(map[int32]int)

	for _, i := range magazine {
		t[i]++
	}
	for _, i := range ransomNote {
		_, ok := t[i]
		if !ok || t[i] <= 0 {
			return false
		} else {
			t[i]--
		}
	}

	return true
}

func main() {
	ransomNote := "aa"
	magazine := "ab"
	// words := ["bella","label","roller"]
	canConstruct(ransomNote, magazine)

}
