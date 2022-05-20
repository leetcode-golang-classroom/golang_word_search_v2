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
