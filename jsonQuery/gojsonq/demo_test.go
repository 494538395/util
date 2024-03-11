package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/thedevsaddam/gojsonq"
)

var jq *gojsonq.JSONQ

func init() {
	jsonStr := `{
  "identifier": 2001,
  "operation": "1709699486525182749",
  "pushData": {
    "event": "topic.pvp.sync",
    "topic": "Topic_Pvp",
    "seqId": 2203,
    "data": {
      "sid": 7,
      "score": 1000,
      "jerry.name": "li",
      "seasonMaxScore": 1000,
      "stageRwdStatus": {
        "reward1": [
          {
            "itemId": 1,
            "cnt": 20
          },
          {
            "itemId": 2,
            "cnt": 50
          }
        ],
        "reward2": [
          {
            "itemId": 3,
            "cnt": 100
          }
        ]
      },
      "pick": [
        {
          "itemId": 1,
          "address": "beijing",
          "cnt": 20
        },
        {
          "itemId": 2,
          "address": "shenzhen",
          "cnt": 100
        },
        {
          "itemId": 4,
          "address": "shanghai",
          "cnt": 400
        },
        {
          "itemId": 5,
          "address": "nanjing",
          "cnt": 400
        }
      ],
      "stage": 1,
      "vectoryStar": {
        "maxProg": 5
      }
    },
    "errCode": 1,
    "errMsg": "ok"
  }
}


`

	jq = gojsonq.New().FromString(jsonStr)
}

func TestFiledWihPoint(t *testing.T) {
	jq.Reset()
	//find := jq.Find("pushData.data.jerry.name")
	find := jq.From("pushData.data").Find("score")
	fmt.Println("find-->", find)

}

func Test02(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").Select("itemId", "cnt").Where("itemId", ">=", 2.1).Get()
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))

}

func TestLen(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").Where("itemId", ">", 1).Count()
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))
}

func TestField(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").Select("itemId").Where("itemId", ">", 1).Get()
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))
}

func TestSum(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").Where("itemId", ">", 0).Sum("cnt")
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))
}

func Test01(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").Select("cnt").WhereEqual("itemId", 1).First()
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))

}

func TestMultiCond(t *testing.T) {
	jq.Reset()
	get := jq.From("pushData.data.pick").
		Where("itemId", ">=", 0).
		Where("itemId", "<=", 3).
		OrWhere("address", "=", "shanghai").
		Get()

	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))

}

func TestMultiCond2(t *testing.T) {

}

func TestParseFormulaStr(t *testing.T) {
	//formula1 := "pushData.data.pick[](itemId=1).sum(cnt)"
	formula2 := "pushData.data.score"

	query1, _ := parseFormula(formula2)
	fmt.Println("Query 1:", query1)
}

// Formula represents a parsed formula
type Formula struct {
	From     string   // From部分
	Filter   []cond   // 条件部分
	Selector selector // 最终部分
}

// parseFormula 解析公式字符串并返回From、Filter和Selector
func parseFormula(formula string) (*Formula, error) {
	parts := strings.Split(formula, "[]")
	if len(parts) == 1 {
		return &Formula{From: parts[0]}, nil
	}
	if len(parts) != 2 {
		return nil, fmt.Errorf("invalid formula: missing '[]' separator")
	}

	from := strings.TrimSpace(parts[0])

	// 解析条件部分
	second := strings.Trim(parts[1], "[]")

	idx := strings.IndexRune(second, ')')

	filter := second[1:idx]
	fmt.Println("filter-->", filter)

	cond, err := parseCondition(filter)
	if err != nil {
		panic(err)
	}

	return &Formula{
		From:     from,
		Filter:   cond,
		Selector: parseSelector(second[idx+1:]),
	}, nil
}

//func parseFormulaOrigin(formula string) (*Formula, error) {
//	parts := strings.Split(formula, "[]")
//	if len(parts) != 2 {
//		return nil, fmt.Errorf("invalid formula: missing '[]' separator")
//	}
//
//	from := strings.TrimSpace(parts[0])
//
//	// 解析条件部分
//	filter := strings.Trim(parts[1], "[]")
//
//	closeIndex := strings.IndexRune(filter, ')')
//	fmt.Println("closeIdx-->", closeIndex)
//
//	newFilter := filter[0 : closeIndex+1]
//	fmt.Println("newFilter-->", newFilter)
//
//	var filters []string
//	filterParts := strings.FieldsFunc(filter, func(r rune) bool {
//		return r == '&' || r == '|'
//	})
//	for _, cond := range filterParts {
//		filters = append(filters, strings.TrimSpace(cond))
//	}
//
//	// 解析最终部分
//	selector := ""
//	if strings.Contains(from, ".sum(") {
//		selector = "sum(cnt)"
//	} else if strings.Contains(from, ".count(") {
//		selector = "count()"
//	} else {
//		selector = strings.Split(parts[1], ".")[1]
//	}
//
//	return &Formula{
//		From:     from,
//		Filter:   nil,
//		Selector: selector,
//	}, nil
//}

//func parseFormula(formula string) (*Formula, error) {
//	// 正则表达式来解析
//	re := regexp.MustCompile(`(\w+)\[\](.*)\(([^)]+)\)`)
//	match := re.FindStringSubmatch(formula)
//	if len(match) != 4 {
//		return nil, fmt.Errorf("invalid formula format")
//	}
//
//	from := match[1]                        // 数据来源
//	conditions := match[2]                  // 条件部分
//	selector := strings.TrimSpace(match[3]) // 最终部分
//
//	// 解析条件部分，使用正则表达式分割条件
//	filterParts := strings.FieldsFunc(conditions, func(r rune) bool {
//		return r == '&' || r == '|'
//	})
//	var filters []string
//	for _, part := range filterParts {
//		filters = append(filters, strings.TrimSpace(part))
//	}
//
//	return &Formula{
//		From:     from,
//		Filter:   filters,
//		Selector: selector,
//	}, nil
//}

func TestParseCond(t *testing.T) {
	res, err := parseCondition("itemId>=1&& itemId <=5.2 && address = beijing")
	if err != nil {
		panic(err)
	}
	fmt.Println("res-->", res)
}

type cond struct {
	key     string
	keyType string
	expr    string
	val     string
}

func parseCondition(condition string) ([]cond, error) {
	var conds []cond
	parts := strings.FieldsFunc(condition, func(r rune) bool {
		return r == '&'
	})
	for _, part := range parts {
		c := cond{}
		expr := strings.ReplaceAll(part, " ", "")
		switch {
		case strings.Contains(expr, ">="):
			parts := strings.Split(expr, ">=")
			c.expr = ">="
			c.key = parts[0]
			c.val = parts[1]
			c.keyType = "int"
		case strings.Contains(expr, "<="):
			parts := strings.Split(expr, "<=")
			c.expr = "<="
			c.key = parts[0]
			c.val = parts[1]
			c.keyType = "int"
		case strings.Contains(expr, ">"):
			parts := strings.Split(expr, ">")
			c.expr = ">"
			c.key = parts[0]
			c.val = parts[1]
			c.keyType = "int"
		case strings.Contains(expr, "<"):
			parts := strings.Split(expr, "<")
			c.expr = "<"
			c.key = parts[0]
			c.val = parts[1]
			c.keyType = "int"
		case strings.Contains(expr, "="):
			parts := strings.Split(expr, "=")
			c.expr = "="
			c.key = parts[0]
			c.val = parts[1]
		case strings.Contains(expr, "!="):
			parts := strings.Split(expr, "!=")
			c.expr = "!="
			c.key = parts[0]
			c.val = parts[1]
		}
		c.keyType = getKeyType(c.val)
		conds = append(conds, c)
	}
	return conds, nil
}

func getKeyType(val string) string {
	_, err := strconv.Atoi(val)
	if err == nil {
		return "int"
	}

	_, err = strconv.ParseFloat(val, 64)
	if err == nil {
		return "float"
	}

	return "string"
}

type selector struct {
	sum   bool
	get   bool
	count bool
	field string
}

func parseSelector(suffix string) selector {
	s := selector{}
	if strings.Contains(suffix, ".sum(") {
		s.sum = true
	} else if strings.Contains(suffix, ".len") {
		s.count = true
	} else if strings.Contains(suffix, ".") {
		s.get = true
	}
	if s.sum {
		start := strings.Index(suffix, "(")
		end := strings.LastIndex(suffix, ")")
		if start != -1 && end != -1 {
			s.field = suffix[start+2 : end-1] // Exclude quotes and parentheses
		}
	}
	if s.get {
		s.field = suffix[1:]
	}

	return s
}
