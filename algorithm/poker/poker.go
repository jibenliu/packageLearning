package poker

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

// refer: https://www.jianshu.com/p/2686b68e3687

func CreateNew() []string {
	numbers := make([]string, 54)
	start := 0 //造牌游标
	for i := 3; i <= 16; i++ {
		if i == 16 { //i为16说明已经到大小王
			numbers[start] = "Q88"
			numbers[start+1] = "K99" //直接构造大小王
		} else {
			numbers[start] = "A" + strconv.Itoa(i)
			numbers[start+1] = "B" + strconv.Itoa(i)
			numbers[start+2] = "C" + strconv.Itoa(i)
			numbers[start+3] = "D" + strconv.Itoa(i)
			start += 4 //每造一套单值牌，游标移4位
		}
	}
	return numbers
}

func Shuffle(values []string) {
	r := rand.New(rand.NewSource(time.Now().Unix())) //根据系统时间戳初始化Random
	for len(values) > 0 {                            //根据牌面数组长度遍历
		n := len(values)                                                //数组长度
		randIndex := r.Intn(n)                                          //得到随机index
		values[n-1], values[randIndex] = values[randIndex], values[n-1] //最后一张牌和第randIndex张牌互换
		values = values[:n-1]
	}
}

// Dispatcher 发牌
// order int order==0 玩家1次序
// order int order==1 玩家2次序
// order int order==2 玩家3次序
// order int order==3 底牌次序
func Dispatcher(order int, values []string) []string {
	var playCards []string
	if order < 0 || order > 3 { //判断玩家次序是否正确
		return []string{}
	} else {
		size := 17 //默认总长度为17
		if order == 3 {
			size = 3 //次序为3(底牌次序)时，总长度为3
		}
		for i := 0; i < size; i++ {
			playCards = append(playCards, values[order*17+i]) //根据次序发牌
		}
	}
	return playCards
}

type CardTypeStatus int

const (
	_            CardTypeStatus = iota
	SINGLE      //单根
	DOUBLE      //对子
	THREE       //三不带
	ThreeAndOne //三带一
	ThreeAndTwo //三带二
	BOMB        //炸弹
	FourTwo     //四带二
	PLANE       //飞机
	PlaneEmpty  //三不带飞机
	DoubleAlone //连对
	SingleAlone //顺子
	KingBomb    //王炸
	ErrorType   //非法类型
)

func (c CardTypeStatus)String() string {
	switch c {
	case SINGLE:
		return "single"
	case DOUBLE:
		return "double"
	case THREE:
		return "three"
	case ThreeAndOne:
		return "three_and_one"
	case ThreeAndTwo:
		return "three_and_two"
	case BOMB:
		return "bomb"
	case FourTwo:
		return "four_and_two"
	case PLANE:
		return "plane"
	case PlaneEmpty:
		return "plane_empty"
	case DoubleAlone:
		return "double_alone"
	case SingleAlone:
		return "single_alone"
	case KingBomb:
		return "king_bomb"
	case ErrorType:
		return "error_type"
	default:
		return "Unknown"
	}
}

type CardShow struct {
	ShowValue      []string            //牌面数组
	CardMap        map[int]int         //牌面计算结果
	MaxCount       int                 //同值牌出现的最大次数
	MaxValues      []int               //同值牌出现的次数列表
	CompareValue   int                 //用于比较大小的值
	CardTypeStatus CardTypeStatus //牌面类型
	ShowTime time.Time
}

// ParseCardsInSize
// 根据牌面数量判断牌面类型
func ParseCardsInSize(plays []string) CardShow {
	cardShow := CardShow{
		ShowValue: plays,
		ShowTime:  time.Now(),
	}
	switch len(plays) {
	case 1:
		cardShow.CardTypeStatus = SINGLE
		cardShow.CompareValue = GetCardValue(plays[0])
		cardShow.MaxCount = 1
		cardShow.MaxValues = []int{cardShow.CompareValue}
		fmt.Printf("根%d", GetCardValue(plays[0]))
		break
	case 2:
		if plays[0] == "Q88" && plays[1] == "K99" {
			cardShow.CardTypeStatus = KingBomb
			cardShow.CompareValue = GetCardValue(plays[0])
			cardShow.MaxCount = 2
			cardShow.MaxValues = []int{cardShow.CompareValue}
			fmt.Println("王炸")
		} else {
			ParseCardsType(plays, &cardShow)
		}
		break
	}
	if len(plays) > 2 {
		ParseCardsType(plays, &cardShow)
	} else {
		cardShow.CardTypeStatus = ErrorType
	}
	return cardShow
}

// ParseCardsType
// 获取牌面类型
func ParseCardsType(cards []string, cardShow *CardShow) {
	mapCard, maxCount, maxValues := ComputerValueTimes(cards)
	cardShow.MaxCount = maxCount
	cardShow.MaxValues = maxValues
	cardShow.CardMap = mapCard
	cardShow.CompareValue = maxValues[len(maxValues)-1]
	switch maxCount {
	case 4:
		if maxCount == len(cards) {
			cardShow.CardTypeStatus = KingBomb
			fmt.Println("炸弹")
		} else if len(cards) == 6 {
			cardShow.CardTypeStatus = FourTwo
			fmt.Println("四带二")
		} else {
			cardShow.CardTypeStatus = ErrorType
			fmt.Println("不合法出牌")
		}
		break
	case 3:
		alive := len(cards) - len(maxValues)*maxCount
		if len(maxValues) == alive {
			if len(maxValues) == 1 {
				cardShow.CardTypeStatus = ThreeAndOne
				fmt.Println("三带一")
			} else if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = PLANE
					fmt.Printf("飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = ErrorType
					fmt.Println("非法飞机")
				}
			}
		} else if alive == 0 {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) {
					cardShow.CardTypeStatus = PlaneEmpty
					fmt.Printf("三不带飞机%d", len(maxValues))
				} else {
					cardShow.CardTypeStatus = ErrorType
					fmt.Println("非法三不带飞机")
				}

			} else {
				cardShow.CardTypeStatus = THREE
				fmt.Println("三不带")
			}
		} else {
			cardShow.CardTypeStatus = ErrorType
			fmt.Println("不合法飞机或三带一")
		}
		break
	case 2:
		if len(maxValues) == (len(cards) / 2) {
			if len(maxValues) > 1 {
				if IsContinuity(mapCard, false) && len(maxValues) > 2 {
					cardShow.CardTypeStatus = DoubleAlone
					fmt.Printf("%d连队", len(maxValues))
				} else {
					cardShow.CardTypeStatus = ErrorType
					fmt.Println("非法连对")
				}
			} else if len(maxValues) == 1 {
				cardShow.CardTypeStatus = DOUBLE
				fmt.Printf("对%d", GetCardValue(cards[0]))
			}
		} else {
			cardShow.CardTypeStatus = ErrorType
			fmt.Println("不合法出牌")
		}
		break
	case 1:
		if IsContinuity(mapCard, true) && len(cards) >= 5 {
			cardShow.CardTypeStatus = SingleAlone
			fmt.Printf("%d顺子", len(mapCard))
		} else {
			fmt.Println("非法顺子")
		}
		break
	}
}

// GetOrderKeys
// 获取顺序的key值数组
func GetOrderKeys(cardMap map[int]int, isSingle bool) []int {
	var keys []int
	for key, value := range cardMap {
		if (!isSingle && value > 1) || isSingle {
			keys = append(keys, key)
		}
	}
	sort.Ints(keys)
	return keys
}

// IsContinuity
// 计算牌面值是否连续
func IsContinuity(cardMap map[int]int, isSingle bool) bool {
	keys := GetOrderKeys(cardMap, isSingle)
	lastKey := 0
	for i := 0; i < len(keys); i++ {
		if (lastKey > 0 && (keys[i]-lastKey) != 1) || keys[i] == 15 {
			return false
		}
		lastKey = keys[i]
	}
	if lastKey > 0 {
		return true
	} else {
		return false
	}
}

// ComputerValueTimes
// 计算每张牌面出现的次数
// mapCard 标记结果
// MaxCount 出现最多的次数
// MaxValues 出现次数最多的所有值
func ComputerValueTimes(cards []string) (mapCard map[int]int, MaxCount int, MaxValues []int) {
	newMap := make(map[int]int)
	if len(cards) == 0 {
		return newMap, 0, nil
	}
	for _, value := range cards {
		cardValue := GetCardValue(value)
		if newMap[cardValue] != 0 {
			newMap[cardValue]++
		} else {
			newMap[cardValue] = 1
		}
	}
	var allCount []int //所有的次数
	var maxCount int   //出现最多的次数
	for _, value := range newMap {
		allCount = append(allCount, value)
	}
	maxCount = allCount[0]
	for i := 0; i < len(allCount); i++ {
		if maxCount < allCount[i] {
			maxCount = allCount[i]
		}
	}
	var maxValue []int
	for key, value := range newMap {
		if value == maxCount {
			maxValue = append(maxValue, key)
		}
	}
	sort.Ints(maxValue)
	return newMap, maxCount, maxValue
}

// GetCardValue
// 获取牌面值
func GetCardValue(card string) int {
	stringValue :=  card[1 : len(card)-1]
	value, err := strconv.Atoi(stringValue)
	if err == nil {
		return value
	}
	return -1
}
