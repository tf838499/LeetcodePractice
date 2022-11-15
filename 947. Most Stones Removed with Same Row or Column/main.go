package main

import (
// "fmt"5
// // "strconv"
// "sort"
)

/*
Union-find Algorithm
tree
doing
并查集

On a 2D plane, we place n stones at some integer coordinate points. Each coordinate point may have at most one stone.

A stone can be removed if it shares either the same row or the same column as another stone that has not been removed.

Given an array stones of length n where stones[i] = [xi, yi] represents the location of the ith stone,
return the largest possible number of stones that can be removed.

Input: stones = [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]]
Output: 5
Explanation: One way to remove 5 stones is as follows:
1. Remove stone [2,2] because it shares the same row as [2,1].
2. Remove stone [2,1] because it shares the same column as [0,1].
3. Remove stone [1,2] because it shares the same row as [1,0].
4. Remove stone [1,0] because it shares the same column as [0,0].
5. Remove stone [0,1] because it shares the same row as [0,0].
Stone [0,0] cannot be removed since it does not share a row/column with another stone still on the plane.

Input: stones = [[0,0],[0,2],[1,1],[2,0],[2,2]]
Output: 3
Explanation: One way to make 3 moves is as follows:
1. Remove stone [2,2] because it shares the same row as [2,0].
2. Remove stone [2,0] because it shares the same column as [0,0].
3. Remove stone [0,2] because it shares the same row as [0,0].
Stones [0,0] and [1,1] cannot be removed since they do not share a row/column with another stone still on the plane.

Input: stones = [[0,0]]
Output: 0
Explanation: [0,0] is the only stone on the plane, so you cannot remove it.

*/

func removeStones(stones [][]int) int {
	ans := 0
	return ans
}

// [[0,1],[1,0],[1,1]] 2
// [[0,0],[0,1],[1,0],[1,2],[2,1],[2,2]] 5
// [[0,0],[0,2],[1,1],[2,0],[2,2]] 3
// [[0,0]] 1
// [[0,1],[1,0]] 0
func main() {
	// input := [][]int{[]int{0, 0}, []int{0, 2}, []int{1, 1}, []int{2, 0}, []int{2, 2}}
	input := [][]int{[]int{0, 1}, []int{1, 0}, []int{1, 1}}
	// input := [][]int{[]int{0, 0}, []int{0, 1}, []int{1, 0}, []int{1, 2}, []int{2, 1}, []int{2, 2}}
	// input := [][]int{[]int{0, 1}, []int{1, 0}}
	removeStones(input)
}

/* ____________________________________ DFS ____________________________________

func removeStones(stones [][]int) int {
    rows := make(map[int][]int)
    cols := make(map[int][]int)

    for _, stone := range stones {
        rows[stone[0]] = append(rows[stone[0]], stone[1])
        cols[stone[1]] = append(cols[stone[1]], stone[0])
    }

    count := 0
    visited := make(map[[2]int]bool)

    for _, stone := range stones {
        if _, ok := visited[[2]int{stone[0], stone[1]}]; !ok {
            count ++
            dfs(stone[0], stone[1], rows, cols, visited)
        }
    }
    return len(stones) - count
}

func dfs(x, y int, rows, cols map[int][]int, visited map[[2]int]bool) {

    visited[[2]int{x, y}] = true

    for _, col := range rows[x] {
        if !visited[[2]int{x, col}] {
            dfs(x, col, rows, cols, visited)
        }
    }
    for _, row := range cols[y] {
        if !visited[[2]int{row, y}] {
            dfs(row, y, rows, cols, visited)
        }
    }

}
____________________________________ DFS ____________________________________ */

/* ____________________________________ Union ____________________________________
type Node struct {
    rank int
    parent *Node
}

func find(n *Node) *Node {
    if n.parent == nil || n.parent == n {
        return n
    }
    // path compression
    n.parent = find(n.parent)
    return n.parent
}

// union by rank
func union(src, dst *Node) {
    psrc, pdst := find(src), find(dst)
    if psrc.rank > pdst.rank {
        pdst.parent = psrc
    } else if pdst.rank > psrc.rank {
        psrc.parent = pdst
    } else {
        psrc.parent = pdst
        pdst.rank++
    }
    return
}

func removeStones(stones [][]int) int {
    m := len(stones)
    if m == 0 {
        return 0
    }
    nodes :=[]*Node{}
    rowStone := map[int]*Node{}
    colStone := map[int]*Node{}
    for _, stone := range stones {
        x, y := stone[0], stone[1]
        node := &Node{}
        node.parent = node
        if _, ok := rowStone[x]; !ok {
            rowStone[x] = node
        }
        if _, ok := colStone[y]; !ok {
            colStone[y] = node
        }

        union(rowStone[x], node) // connect this node to the first node in the given row
        union(colStone[y], node) // connect this node to the first node in this col
        nodes = append(nodes, node)
    }

    res := 0
    for _, node := range nodes {
	    // count all nodes that have different parent than themselves since
		// those will be removed
        if find(node) != node {
            res++
        }
    }
    return res
}
____________________________________ Union ____________________________________ */
