package video

import (
	"regexp"
	"strings"
)

func parseOrderString(orderSlice []string) (col string, direction string) {
	orderString := ""
	direction = "desc"
	if len(orderSlice) == 1 {
		orderString = orderSlice[0]
		orderString = strings.ToLower(deleteExtraSpace(strings.TrimSpace(orderString)))
	}

	if len(orderSlice) > 1 {
		direction = orderSlice[len(orderSlice)-1]
		orderString = strings.Join(orderSlice[:len(orderSlice)-2], ",")
	}

	if len(orderString) == 0 {
		col = GetVideoQueryCols().GetPrimaryColsStr()
		return
	}

	if strings.Contains(orderString, " ") {
		splitStrings := strings.Split(orderString, " ")
		if len(splitStrings) != 2 {
			panic("order str error")
		}
		direction = splitStrings[1]
		col = splitStrings[0]
	}

	if strings.Contains(col, ",") {
		for _, colName := range strings.Split(col, ",") {
			if !GetVideoQueryCols().Has(colName) {
				panic("unknown col name")
			}
		}
	} else {
		if !GetVideoQueryCols().Has(col) {
			panic("unknown col name")
		}
	}

	if direction != "asc" && direction != "desc" {
		panic("unknown direction")
	}

	return
}

func deleteExtraSpace(s string) string {
	//删除字符串中的多余空格，有多个空格时，仅保留一个空格
	s1 := strings.Replace(s, "	", " ", -1)      //替换tab为空格
	register := "\\s{2,}"                       //两个及两个以上空格的正则表达式
	reg, _ := regexp.Compile(register)          //编译正则表达式
	s2 := make([]byte, len(s1))                 //定义字符数组切片
	copy(s2, s1)                                //将字符串复制到切片
	spcIndex := reg.FindStringIndex(string(s2)) //在字符串中搜索
	for len(spcIndex) > 0 {                     //找到适配项
		s2 = append(s2[:spcIndex[0]+1], s2[spcIndex[1]:]...) //删除多余空格
		spcIndex = reg.FindStringIndex(string(s2))           //继续在字符串中搜索
	}
	return string(s2)
}
