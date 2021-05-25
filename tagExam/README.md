Tag 由反引号包含，由一对或几对的键值对组成，通过空格来分割键值。格式如下

```key01:"value01" key02:"value02" key03:"value03"```

定义完后，如何从结构体中，取出 Tag 呢？

答案就是，我们上一节学过的 "反射"。

获取 Tag 可以分为三个步骤：

获取字段 field

获取标签 tag

获取键值对 key:value

演示如下
```
// 三种获取 field
field := reflect.TypeOf(obj).FieldByName("Name")
field := reflect.ValueOf(obj).Type().Field(i)  // i 表示第几个字段
field := reflect.ValueOf(&obj).Elem().Type().Field(i)  // i 表示第几个字段

// 获取 Tag
tag := field.Tag 

// 获取键值对
labelValue := tag.Get("label")
labelValue,ok := tag.Lookup("label")
```
获取键值对，有Get 和 Lookup 两种方法，但其实 Get 只是对 Lookup 函数的简单封装而已，当没有获取到对应 tag 的内容，会返回空字符串。
```
func (tag StructTag) Get(key string) string {
    v, _ := tag.Lookup(key)
    return v
}
```
空 Tag 和不设置 Tag 效果是一样的

```
package main
import (
    "fmt"
    "reflect"
)
type Person struct {
    Name string ``
    Age string
}
func main() {
    p := reflect.TypeOf(Person{})
    name, _ := p.FieldByName("Name")
    fmt.Printf("%q\n", name.Tag) //输出 ""
    age, _ := p.FieldByName("Age")
    fmt.Printf("%q\n", age.Tag) // 输出 ""
}
```