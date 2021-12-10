package readFlagFromFile

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strings"
)

// refer:https://mp.weixin.qq.com/s/7bqZWoPwLbI1v9Xp4kn6QA

type FlagFile string

func (f FlagFile) String() string {
	return string(f)
}
func (f FlagFile) Set(str string) error {
	return nil
}

// FlagFile 存储命令行传过来的文件路径
var flagFile FlagFile

func init() {
	//注册命令行的flagFile参数
	flag.Var(&flagFile, "flagfile", "")
}

//在Parse函数中调用，将解析到的命令行参数打印出来
func visitFlag(f *flag.Flag) {
	fmt.Println(f.Name + "=" + f.Value.String())
}

func Parse() error {
	//先解析命令行中的-flagfile参数
	flag.Parse()

	var validFlagLines []string

	flagContents, _ := ioutil.ReadFile(string(flagFile))
	configContent := string(flagContents)
	configContent = strings.Replace(configContent, "\r\n", "\n", -1)
	flagLines := strings.Split(configContent, "\n")
	for _, line := range flagLines {
		//忽略掉以 # 开头的注释行
		if len(strings.Replace(line, " ", "", -1)) != 0 && string([]rune(line)[0]) != "#" {
			//将每一行作为一个有效的命令行参数
			validFlagLines = append(validFlagLines, line)
		}
	}

	fmt.Println(validFlagLines)
	//实际执行解析命令行参数的地方，这里就又和常规的flag调用一样了
	_ = flag.CommandLine.Parse(validFlagLines)
	//将解析到的命令行参数都按visitFlag函数的格式输出
	flag.VisitAll(visitFlag)
	return nil
}
