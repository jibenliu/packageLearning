#base-refer:[https://mp.weixin.qq.com/s/0TJRcbaWlfwEjbKMKAxNAQ](https://mp.weixin.qq.com/s/0TJRcbaWlfwEjbKMKAxNAQ)


#upgrade refer:[https://colobu.com/2021/05/11/unbounded-channel-in-go/](https://colobu.com/2021/05/11/unbounded-channel-in-go/)
```bigquery
缓存无限的channel拥有下面的特性：

不会阻塞write。 它总是能处理write的数据，或者放入到待读取的channel中，或者放入到缓存中
无数据时read会被阻塞。当没有可读的数据时，从channel中读取的goroutine会被阻塞
读写都是通过channel操作。 内部的缓存不会暴露出来
能够查询当前待读取的数据数量。因为缓存中可能也有待处理的数据，所以需要返回len(buffer)+len(chan)
关闭channel后，还未读取的channel还是能够被读取，读取完之后才能发现channel已经完毕。这和正常的channel的逻辑是一样的，这种情况叫"drain"未读的数据
```