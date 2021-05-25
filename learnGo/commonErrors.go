package main

import "fmt"

func main() {
	/**
	 * 在go语言中，循环（迭代）所使用的变量是同一个变量，只是在每次循环的时候被赋于不同的值
	 */
	var out []*int //切片
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("values:", *out[0], *out[1], *out[2])
	fmt.Println("address:", out[0], out[1], out[2])

	//解决办法
	for j := 0; j < 3; j++ {
		j := j
		out = append(out, &j)
	}
	fmt.Println("values:", *out[3], *out[4], *out[5])
	fmt.Println("address:", out[3], out[4], out[5])

	/**
	 * 在协程中使用循环变量 可能会发现输出的结果是一摸一样的！ 因为go的协程跑起来也是需要一点时间的，循环结束的时候，可能一个goroute都没有跑完
	 */
	var sliceVal = []int{1, 2, 3}
	for _, val := range sliceVal {
		go func() {
			fmt.Println(val)
		}()
	}

	fmt.Println(&sliceVal)
	//解决办法1
	for _, val := range out {
		go func(val interface{}) {
			fmt.Println(&val)
		}(val)
	}

	//解决办法2
	for _, val := range out {
		val := &val
		go func() {
			fmt.Println(val)
		}()
	}

}
