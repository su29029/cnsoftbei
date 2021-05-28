package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"gorm.io/gorm"
	"gorm.io/driver/mysql"
)

type Data struct {
	Name string `json:"name"`
	Value []int	`json:"value"`
}

type EpidemicData struct {
	gorm.Model
	Date string
	County string
	State string
	Cases string
	Death string
}

var data = [52]Data{
	{"Alabama",[]int{ 688, 52717 } },
	{"Alaska",[]int{ 42, 85335 } },
	{"Arizona",[]int{690, 69927} },
	{"Arkansas",[]int{975, 24924} },
	{"California",[]int{529, 51127} },
	{"Colorado",[]int{817, 80437} },
	{"Connecticut",[]int{665, 57869} },
	{"Delaware",[]int{772, 84146} },
	{"District of Columbia",[]int{260, 36095} },
	{"Florida",[]int{629, 66176} },
	{"Georgia",[]int{582, 86849} },
	{"Hawaii",[]int{555, 43028} },
	{"Idaho",[]int{370, 11296} },
	{"Illinois",[]int{203, 67732} },
	{"Indiana",[]int{657, 60185} },
	{"Iowa",[]int{122, 77489} },
	{"Kansas",[]int{384, 3470} },
	{"Kentucky",[]int{292, 70978} },
	{"Louisiana",[]int{790, 14394} },
	{"Maine",[]int{334, 17926} },
	{"Maryland",[]int{354, 56374} },
	{"Massachusetts",[]int{638, 52548} },
	{"Michigan",[]int{896, 27238} },
	{"Minnesota",[]int{627, 74360} },
	{"Mississippi",[]int{701, 39155} },
	{"Missouri",[]int{225, 65674} },
	{"Montana",[]int{694, 57283} },
	{"Nebraska",[]int{302, 78257} },
	{"Nevada",[]int{176, 9676} },
	{"New Hampshire",[]int{679, 93855} },
	{"New Jersey",[]int{988, 37088} },
	{"New Mexico",[]int{354, 6478} },
	{"New York",[]int{987, 3257} },
	{"North Carolina",[]int{403, 94568} },
	{"North Dakota",[]int{680, 81917} },
	{"Ohio",[]int{143, 41841} },
	{"Oklahoma",[]int{569, 68609} },
	{"Oregon",[]int{509, 55236} },
	{"Pennsylvania",[]int{744, 92852} },
	{"Rhode Island",[]int{304, 3461} },
	{"South Carolina",[]int{434, 17271} },
	{"South Dakota",[]int{405, 90326} },
	{"Tennessee",[]int{137, 33063} },
	{"Texas",[]int{40, 60259} },
	{"Utah",[]int{159, 59944} },
	{"Vermont",[]int{391, 30527} },
	{"Virginia",[]int{681, 32020} },
	{"Washington",[]int{991, 19176} },
	{"West Virginia",[]int{695, 31184} },
	{"Wisconsin",[]int{100, 91178} },
	{"Wyoming",[]int{509, 72238} },
	{"Puerto Rico",[]int{794, 5621} },
}

func main() {
	r := gin.Default()
	r.GET("/api/states/all", getCountyHandler)
	r.GET("/api/search", handleSearch)
	_ = r.Run(":18111")
}

func getCountyHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

func handleSearch(c *gin.Context) {
	searchString := c.Query("string")

	dsn := "root:tF#262420228@tcp(127.0.0.1:3306)/epidemic?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
		return
	}

	_ = db.AutoMigrate(&EpidemicData{})


	for _, v := range data { // 最终会替换成数据库中读取数据
		if v.Name == searchString {
			c.JSON(http.StatusOK, gin.H{
				"data": v,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"err": "city not found",
	})
}