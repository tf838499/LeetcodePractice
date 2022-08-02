package main

// "strconv"

/*
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

func canConstruct(ransomNote string, magazine string) bool {
	t := make(map[int32]int)
	ans := true
	for _, i := range magazine {
		t[i]++
	}
	for _, i := range ransomNote {
		t[i]--
		if t[i] < 0 {
			ans = false
		}
	}

	return ans
}

func main() {
	ransomNote := "aacc"
	magazine := "aabb"
	// words := ["bella","label","roller"]
	canConstruct(ransomNote, magazine)

}
