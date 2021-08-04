package dirtywords

import (
	"testing"
)

var trieTree = Trie{Word: 0, Child: make(map[rune]*Trie)}

func TestTrie(t *testing.T) {
	trieTree.Insert("你大爷")
	trieTree.Insert("大姨妈")
	trieTree.Insert("姨妈jin")
	trieTree.Insert("jin子")
	trieTree.Insert("大姨父")
	trieTree.Insert("妈了个吧")
	trieTree.Insert("狗日的")
	trieTree.Insert("去你吗的")
	trieTree.Insert("bitch")
	trieTree.Insert("bitches")

	matched := trieTree.Search("你大爷")
	t.Log(matched)
	if len(matched) != 1 || matched[0] != "你大爷" {
		t.Errorf("search: %s, expected: %v, actual: %v", "你大爷", []string{"你大爷"}, matched)
	}

	matched = trieTree.Search("英文单词bitches意思是母狗")
	t.Log(matched)
	if len(matched) != 2 || matched[0] != "bitch" || matched[1] != "bitches" {
		t.Errorf("search: %s, expected: %v, actual: %v", "英文单词bitches意思是母狗", []string{"bitch", "bitches"}, matched)
	}

	matched = trieTree.Search("狗日的大姨妈啊")
	t.Log(matched)
	if len(matched) != 2 || matched[0] != "狗日的" || matched[1] != "大姨妈" {
		t.Errorf("search: %s, expected: %v, actual: %v", "狗日的大姨妈啊", []string{"狗日的", "大姨妈"}, matched)
	}

	matched = trieTree.Search("我去你大爷的")
	t.Log(matched)
	if len(matched) != 1 {
		t.Errorf("search: %s, expected: %v, actual: %v", "我去你大爷的", []string{"你大爷"}, matched)
	}

	matched = trieTree.Search("大姨妈jin子")
	t.Log(matched)
	if len(matched) != 3 {
		t.Errorf("search: %s, expected: %v, actual: %v", "大姨妈jin子", []string{"大姨妈", "姨妈jin", "jin子"}, matched)
	}

	matched = trieTree.Search("我很正常")
	t.Log(matched)
	if len(matched) > 0 {
		t.Errorf("search: %s, expected: %v, actual: %v", "我很正常", "", matched)
	}

	trieTree.Delete("你大爷")
	matched = trieTree.Search("你大爷")
	t.Log(matched)
	if len(matched) != 0 {
		t.Errorf("search: %s, expected: %v, actual: %v", "你大爷", []string{"你大爷"}, matched)
	}
}

// go test -v .
