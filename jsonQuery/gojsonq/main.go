package main

import (
	"encoding/json"
	"fmt"

	"github.com/thedevsaddam/gojsonq"
)

func main() {
	//	content := `{
	//  "user": {
	//    "name": "dj",
	//    "age": 18,
	//    "address": {
	//      "provice": "shanghai",
	//      "district": "xuhui"
	//    },
	//    "hobbies": [
	//      "chess",
	//      "programming",
	//      "game"
	//    ],
	//    "consume": [
	//      {
	//        "itemId": 1,
	//        "cnt": 20
	//      },
	//      {
	//        "itemId": 2,
	//        "cnt": 100
	//      }
	//    ]
	//  }
	//}`
	//
	//	gq := gojsonq.New().FromString(content)
	//	district := gq.Find("user.address.provice")
	//	fmt.Println(district)
	//
	//	gq.Reset()
	//
	//	hobby := gq.Find("user.consume.[1]")
	//	fmt.Println(hobby)
	//
	//	gq.Reset()

	arraryData := `{
  "identifier": 2001,
  "operation": "1709699486525182749",
  "pushData": {
    "event": "topic.pvp.sync",
    "topic": "Topic_Pvp",
    "seqId": 2203,
    "data": {
      "sid": 7,
      "score": 1000,
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
          "cnt": 20
        },
        {
          "itemId": 2,
          "cnt": 100
        },
        {
          "itemId": 3,
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

	jq := gojsonq.New().FromString(arraryData)

	//r := jq.From("consume").Select("itemId", "cnt").
	//	Where("itemId", ">", 0).Where("itemId", "<", 3).Sum("cnt" +
	//	"")
	//data, _ := json.MarshalIndent(r, "", "  ")
	//fmt.Println(string(data))

	//jq.Reset()

	//find := jq.Find("pushData.data.activityList")
	//fmt.Println("find-->", find)

	//find := jq.Find("pushData.data.score")
	//fmt.Println("find-->", find)

	//jq.Reset()
	//
	//reward1 := jq.Find("pushData.data.stageRwdStatus.reward1")
	//fmt.Println("reward1-->", reward1)

	//jq.Reset()
	//
	//get := jq.From("pushData.data.stageRwdStatus.reward2").Select("cnt").Get()
	//fmt.Println("get-->", get)
	//data, _ := json.MarshalIndent(get, "", "  ")
	//fmt.Println(string(data))

	// 聚合
	//jq.Reset()
	//get := jq.From("pushData.data.pick").
	//	OrWhere("itemId", "=", 3).
	//	OrWhere("itemId", "=", 2).
	//	Sum("cnt")
	//fmt.Println("get-->", get)
	//data, _ := json.MarshalIndent(get, "", "  ")
	//fmt.Println(string(data))

	jq.Reset()
	get := jq.From("pushData.data.pick").Select("cnt").WhereEqual("itemId", 1).First()
	fmt.Println("get-->", get)
	data, _ := json.MarshalIndent(get, "", "  ")
	fmt.Println(string(data))

}
