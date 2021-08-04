package main

import (
	"hash/crc32"
	"sort"
	"strconv"
)

// refer: https://mp.weixin.qq.com/s?__biz=MzAxMTA4Njc0OQ==&mid=2651440746&idx=2&sn=1491963e2132b4de2f90d0d21161dd1f&chksm=80bb1898b7cc918e959b43ffb355193a2d46791745303393b4490176253924638d8fb1c01836&xtrack=1&scene=0&subscene=91&sessionid=1594956676&clicktime=1594956696&enterid=1594956696&ascene=7&devicetype=android-26&version=2700103b&nettype=cmnet&abtest_cookie=AAACAA%3D%3D&lang=zh_CN&exportkey=AXACZuIiTiLVXLL2HklUbJM%3D&pass_ticket=60PxHVDAKpkkCeV4OT3ENwdVFvqjvQX08MRyI%2FzalCL3m3EfFWG1S2PhZywRlaA8&wx_header=1
// 一致性hash

type Hash func(data []byte) uint32

type Map struct {
	hash     Hash           // 计算 hash 的函数
	replicas int            // 这个是副本数，这里影响到虚拟节点的个数
	keys     []int          // 有序的列表，从大到小排序的，这个很重要
	hashMap  map[int]string // 可以理解成用来记录虚拟节点和物理节点元数据关系的
}

func New(replicas int, fn Hash) *Map {
	m := &Map{
		replicas: replicas,
		hash:     fn,
		hashMap:  make(map[int]string),
	}
	if m.hash == nil {
		// 默认可以用 crc32 来计算hash值
		m.hash = crc32.ChecksumIEEE
	}
	return m
}

func (m *Map) Add(keys ...string) {
	// keys => [ A, B, C ]
	for _, key := range keys {
		for i := 0; i < m.replicas; i++ {
			// hash 值 = hash (i+key)
			hash := int(m.hash([]byte(strconv.Itoa(i) + key)))
			m.keys = append(m.keys, hash)
			// 虚拟节点 <-> 实际节点
			m.hashMap[hash] = key
		}
	}
	sort.Ints(m.keys)
}

func (m *Map) Get(key string) string {
	if m.IsEmpty() {
		return ""
	}
	// 根据用户输入key值，计算出一个hash值
	hash := int(m.hash([]byte(key)))
	// 查看值落到哪个值域范围？选择到虚节点
	idx := sort.Search(len(m.keys), func(i int) bool { return m.keys[i] >= hash })
	if idx == len(m.keys) {
		idx = 0
	}
	// 选择到对应物理节点
	return m.hashMap[m.keys[idx]]
}

func (m *Map) IsEmpty() bool {
	return m == &Map{}
}
