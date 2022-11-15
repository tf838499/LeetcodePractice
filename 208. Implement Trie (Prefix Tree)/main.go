package main

import (
// "fmt"
// "strconv"
)

/*
medium
trie
tree
doing

A trie (pronounced as "try") or prefix tree is a tree data structure used to efficiently store and
retrieve keys in a dataset of strings. There are various applications of this data structure,
such as autocomplete and spellchecker.
Trie() Initializes the trie object.

void insert(String word) Inserts the string word into the trie.
boolean search(String word) Returns true
if the string word is in the trie (i.e., was inserted before), and false otherwise.
boolean startsWith(String prefix) Returns true
if there is a previously inserted string word that has the prefix prefix, and false otherwise.

Input
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
Output
[null, null, true, false, true, null, true]
Explanation
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // return True
trie.search("app");     // return False
trie.startsWith("app"); // return True
trie.insert("app");
trie.search("app");     // return True

Implement the Trie class:
*/
// type Trie struct {
//     val strings{}
// }

// func Constructor() Trie {

// }

// func (this *Trie) Insert(word string)  {

// }

// func (this *Trie) Search(word string) bool {

// }

// func (this *Trie) StartsWith(prefix string) bool {

// }

// /*____________________ ____________________ map solution ____________________ ____________________*/
// /** Initialize your data structure here. */
type Trie struct {
	IsTerminated bool
	Children     map[rune]*Trie
}

func Constructor() Trie {
	m := make(map[rune]*Trie)
	return Trie{IsTerminated: false, Children: m}
}

/** Inserts a word into the trie. */
func (this *Trie) Insert(word string) {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
		} else {
			newChild := &Trie{IsTerminated: false, Children: make(map[rune]*Trie)}
			parent.Children[ch] = newChild
			parent = newChild
		}
	}
	parent.IsTerminated = true
}

/** Returns if the word is in the trie. */
func (this *Trie) Search(word string) bool {
	parent := this
	for _, ch := range word {
		if child, ok := parent.Children[ch]; ok {
			parent = child
			continue
		}
		return false
	}
	return parent.IsTerminated
}

/** Returns if there is any word in the trie that starts with the given prefix. */
func (this *Trie) StartsWith(prefix string) bool {

}

/*____________________ ____________________ normal solution ____________________ ____________________*/
// type Trie struct {
// 	children [26]*Trie
// 	isEnd    bool
// }

// /** Initialize your data structure here. */
// func Constructor() Trie {
// 	return Trie{}
// }

// /** Inserts a word into the trie. */
// func (this *Trie) Insert(word string) {
// 	curr := this
// 	for _, ch := range word {
// 		idx := ch - 'a'
// 		if curr.children[idx] == nil {
// 			curr.children[idx] = &Trie{}
// 		}
// 		curr = curr.children[idx]
// 	}
// 	curr.isEnd = true
// }

// /** Returns if the word is in the trie. */
// func (this *Trie) Search(word string) bool {
// 	curr := this
// 	for _, ch := range word {
// 		idx := ch - 'a'
// 		if curr.children[idx] == nil {
// 			return false
// 		}
// 		curr = curr.children[idx]
// 	}
// 	return curr.isEnd
// }

// /** Returns if there is any word in the trie that starts with the given prefix. */
// func (this *Trie) StartsWith(prefix string) bool {
// 	curr := this
// 	for _, ch := range prefix {
// 		idx := ch - 'a'
// 		if curr.children[idx] == nil {
// 			return false
// 		}
// 		curr = curr.children[idx]
// 	}
// 	return true
// }

func main() {
	// s := Constructor()
	trie := Constructor()
	trie.Insert("apple")
	trie.Search("apple")   // return True
	trie.Search("app")     // return False
	trie.StartsWith("app") // return True
	trie.Insert("app")
	trie.Search("app") // return True
}
