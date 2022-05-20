# golang_word_search_v2

Given an `m x n` `board` of characters and a list of strings `words`, return *all words on the board*.

Each word must be constructed from letters of sequentially adjacent cells, where **adjacent cells** are horizontally or vertically neighboring. The same letter cell may not be used more than once in a word.

## Examples

**Example 1:**

![https://assets.leetcode.com/uploads/2020/11/07/search1.jpg](https://assets.leetcode.com/uploads/2020/11/07/search1.jpg)

```
Input: board = [["o","a","a","n"],["e","t","a","e"],["i","h","k","r"],["i","f","l","v"]], words = ["oath","pea","eat","rain"]
Output: ["eat","oath"]

```

**Example 2:**

![https://assets.leetcode.com/uploads/2020/11/07/search2.jpg](https://assets.leetcode.com/uploads/2020/11/07/search2.jpg)

```
Input: board = [["a","b"],["c","d"]], words = ["abcb"]
Output: []

```

**Constraints:**

- `m == board.length`
- `n == board[i].length`
- `1 <= m, n <= 12`
- `board[i][j]` is a lowercase English letter.
- `1 <= words.length <= $3*10^4$`
- `1 <= words[i].length <= 10`
- `words[i]` consists of lowercase English letters.
- All the strings of `words` are unique.

## 解析

題目給定一個 m by n 字元矩陣 board 還有一個字串 array words

要求實作一個演算法找出字串 array words 有哪些字串存在於 m by n 矩陣

在 m by n 矩陣搜詢一個字串做的法是從每個字元當作起點針對上下左右四個方向做 DFS 找尋有可能的字串，所以最遭的狀況就是  $(4^m)^n$

Trie 的結構能夠有效讓要搜尋的字串可以用一個很有效率的方式做比對

如下圖

![](https://i.imgur.com/rrf2Mmj.png)

## 程式碼
```go
package sol

type Node struct {
	Children map[byte]*Node
	isWord   bool
	word     string
}

type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			Children: make(map[byte]*Node),
			isWord:   false,
		},
	}
}
func (this *Trie) AddWord(word string) {
	cur := this.root
	wLen := len(word)
	for idx := 0; idx < wLen; idx++ {
		ch := word[idx]
		if _, exists := cur.Children[ch]; !exists {
			cur.Children[ch] = &Node{
				Children: make(map[byte]*Node),
				isWord:   false,
			}
		}
		cur = cur.Children[ch]
	}
	cur.isWord = true
	cur.word = word
}

type Pair struct {
	row, col int
}

func findWords(board [][]byte, words []string) []string {
	trie := Constructor()
	for _, word := range words {
		trie.AddWord(word)
	}
	visit := make(map[Pair]struct{})
	res := make(map[string]struct{})
	col, row := len(board), len(board[0])
	var dfs func(r int, c int, node *Node)
	dfs = func(r int, c int, node *Node) {
		cur := node
		if r < 0 || c < 0 || r >= row || c >= col {
			return
		}
		if _, visited := visit[Pair{row: r, col: c}]; visited {
			return
		}
		ch := board[r][c]
		if _, match := cur.Children[ch]; !match {
			return
		}
		visit[Pair{row: r, col: c}] = struct{}{}
		cur = cur.Children[ch]
		if cur.isWord {
			res[cur.word] = struct{}{}
		}
		if len(res) == len(words) {
			return
		}
		// 往上下左右找詢
		dfs(r-1, c, cur)
		dfs(r+1, c, cur)
		dfs(r, c-1, cur)
		dfs(r, c+1, cur)
		delete(visit, Pair{row: r, col: c})
	}
	for r := 0; r < row; r++ {
		for c := 0; c < col; c++ {
			dfs(r, c, trie.root)
		}
	}
	result := make([]string, len(res))
	count := 0
	for key := range res {
		result[count] = key
		count++
	}
	return result
}

```
## 困難點

1. 理解使用 DFS 針對 4個方向去找可能的值
2. 理解透過 Tries 來減少搜尋的次數，透過 Trie 結構可以一次找到多個 prefix 一樣的字串

## Solve Point

- [x]  Understand what problem need to solve
- [x]  Analysis complexity