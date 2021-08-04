package dirtywords

import (
	"fmt"
	"strings"
)

type Trie struct {
	Word  rune //字符数而非字节数
	IsEnd bool
	Child map[rune]*Trie
}

func (t *Trie) Insert(word string) {
	word = strings.TrimSpace(word)
	ptr := t
	for _, u := range word {
		_, ok := ptr.Child[u]
		if !ok {
			node := make(map[rune]*Trie)
			ptr.Child[u] = &Trie{Word: u, Child: node}
		}
		ptr = ptr.Child[u]
	}
	ptr.IsEnd = true
}

func (t *Trie) Walk() {
	var walk func(string, *Trie)
	walk = func(pfx string, node *Trie) {
		if node == nil {
			return
		}
		if node.Word != 0 {
			pfx += string(node.Word)
		}
		if node.IsEnd {
			fmt.Println(pfx)
		}
		for _, v := range node.Child {
			walk(pfx, v)
		}
	}
	walk("", t)
}
func (t *Trie) Search(segment string) []string {
	segment = strings.TrimSpace(segment)
	segmentRune := []rune(segment)
	var matched []string

	ptr := t
	item := ""
	index := 0
	for i := 0; i < len(segmentRune); i++ {
		c, ok := ptr.Child[segmentRune[i]]
		if !ok {
			i = index
			index++
			item = ""
			ptr = t
			continue
		}
		item += string(c.Word)
		if c.IsEnd {
			matched = append(matched, item)

			// 重置检索起点
			if len(c.Child) == 0 {
				// 重置检索起点
				i = index
				index++
				item = ""
				ptr = t
				continue
			}
		}
		ptr = c
	}
	return matched
}

func (t *Trie) Delete(word string) {
	word = strings.TrimSpace(strings.ToLower(word))
	var branch []*Trie
	ptr := t
	for _, u := range word {
		branch = append(branch, ptr)
		c, ok := ptr.Child[u]
		if !ok {
			return
		}
		ptr = c
	}

	// 只命中字典中部分词
	if !ptr.IsEnd {
		return
	}

	// 如bitch和bitches
	// 删除bitch时，只需要将bitch最后一个节点的IsEnd改为false即可
	if len(ptr.Child) != 0 {
		ptr.IsEnd = false
		return
	}

	for len(branch) > 0 {
		p := branch[len(branch)-1]
		branch = branch[:len(branch)-1]

		delete(p.Child, ptr.Word)
		// IsEnd == true 如bitch和bitches，删除bitches时，只需要删除后面的"es"即可
		// len(Child) != 0 整个敏感词全删除
		if p.IsEnd || len(p.Child) != 0 {
			break
		}
		ptr = p
	}
}

// refer : https://juejin.cn/post/6870077377863647239
