# LeetcodePractice

🎈   easy       1
🔥   medium     4
💀   hard       0

## topic : LinkList

⭕🔥 2. Add Two Numbers  
⭕🔥 3. Longest Substring Without Repeating Characters
⭕🔥 19. Remove Nth Node From End of List  
⭕🔥 24. Swap Nodes in Pairs  
⭕🔥 142. Linked List Cycle II  
⭕🔥 143. Reorder List 
⭕🎈 206. Reverse Linked List  
⭕🔥 707. Design Linked List  

## topic : array

⭕🎈 724. Find Pivot Index 
⭕🎈 1480. Running Sum of 1d Array 

## topic : hash map

⭕ 🎈 1. Two Sum  
🚩 🔥 15. 3Sum  
🚩 🔥 18. 4Sum  
⭕ 🎈 202. Happy Number  
⭕ 🎈 242. Valid Anagram  
⭕ 🎈 349. Intersection of Two Arrays  
❌ 🔥 454. 4Sum II  
⭕ 🎈 1002. Find Common Characters  
⭕ 🎈 383. Ransom Note  

## topic : binary search   

⭕ 🎈 35. Search Insert Position 
⭕ 🎈 278. First Bad Version 
⭕ 🎈 374. Guess Number Higher or Lower
⭕ 🎈 704. Binary Search

## topic : Dynamic Programming  

🚩 🔥 38. Count and Say
🚩 🔥 53. Maximum Subarray 

## topic : string  

⭕ 🎈 20. Valid Parentheses  
⭕ 🎈 26. Remove Duplicates from Sorted Array 
⭕ 🎈 27. Remove Element  
⭕ 🔥 28. Implement strStr() 
⭕ 🎈 151. Reverse Words in a String  
⭕ 🎈 217. Contains Duplicate   
⭕ 🎈 344. Reverse String  
⭕ 🎈 345. Reverse Vowels of a String   
⭕ 🎈 541. Reverse String II  
🚩 🎈 459. Repeated Substring Pattern  
⭕ 🎈 1323. Maximum 69 Number 
🚩 🔥 2131. Longest Palindrome by Concatenating Two Letter Words   

## topic : queue and stack

🚩 🔥 901. Online Stock Span copy 
⭕ 🎈 1047. Remove All Adjacent Duplicates In String

## topic : tree
⭕ 🎈 94. Binary Tree Inorder Traversal  
⭕ 🎈 101. Symmetric Tree  
⭕ 🔥 102. Binary Tree Level Order Traversal  
⭕ 🎈 104. Maximum Depth of Binary Tree
⭕ 🔥 107. Binary Tree Level Order Traversal II
⭕ 🎈 110. Balanced Binary Tree 
⭕ 🎈 112. Path Sum 
⭕ 🎈 144. Binary Tree Preorder Traversal  
⭕ 🎈 145. Binary Tree Postorder Traversal  
⭕ 🎈 199. Binary Tree Right Side View   
🚩 🔥 208. Implement Trie (Prefix Tree) 
⭕ 🎈 222. Count Complete Tree Nodes  
⭕ 🎈 226. Invert Binary Tree  
⭕ 🎈 257. Binary Tree Paths  
⭕ 🔥 429. N-ary Tree Level Order Traversal  
⭕ 🔥 515. Find Largest Value in Each Tree Row  
🚩 🔥 947. Most Stones Removed with Same Row or Column

2131
724
947


leetcode 筆記
34 在數組中找尋目標數，數組已排序，且目標數會重複出現，找尋頭尾的目標數index
想法 :
使用暴力解直接歷遍找尋出現位置，時間複雜度會有O(n)，此題可以使用binary Search有效率的找尋可以達到O(log n)，由於有兩個index要找，想法可以用兩組low(l) high(r)去找尋頭個出現以及結尾出現
心得 : 
1. 要注意! 寫題時可以先寫，判斷Input是否正常，以此題當給空矩陣時可直接返回，或是只接看數組index 0
是否大於targe"[大於時，已經沒有解可以找"，或是數組尾部是否小於target也是代表沒有解可以找
2. 可以使用分治法，直接用兩次binary search另外搜尋頭尾的index
3. 寫binary search要注意邊界以及如何做出Mid
12/1
leetcode 88 未寫

12/2
leetcode 455
greedy貪婪法 每次操作都是局部最優，從而得到結果為全局最優

12/5
leetcode 88 
two pointer，
想法:使用two pointer，跟搭配一個pointer指向第二個矩陣
作法:先用Fast找到0位置(注意0不算，但判斷式會算)，然後交換Slow跟fast交換0位置，在跟矩陣2交換值，此時有問題與0交換值可能比前面還小，所以在思考要怎麼將前面再重新排序，但會導致不斷地重新排位置
別人作法:直接重後面開始排序，就不必考慮要重新排序前面大小!
心得:這次解了快兩天，卡在死胡同，原先一定要重前面排到後面，但轉向思考一下，從後面變寫到前面快速簡單
ex:在多個條件式下 switch case比if else還快，""因為golang compiler 會把 switch 編譯成 binary search，所以不管各個 case 順序如何，根據 input variable 找到對應的 condition 的時間大致相同""注意 如果條件式通常發生在前面的話if else會比較快。