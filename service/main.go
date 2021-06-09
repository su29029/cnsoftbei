package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type Data struct {
	Name string    `json:"name"`
	Geo  []float32 `json:"geo"`
	Case Case      `json:"case"`
}

type Case struct {
	Confirmed int `json:"confirmed"`
	Death     int `json:"death"`
}

type EpidemicData struct {
	gorm.Model
	Date   string
	County string
	State  string
	Cases  string
	Death  string
}

var data = [52]Data{
	{"Alabama", []float32{32.57, -87.80}, Case{Confirmed: 84520, Death: 407}},
	{"Alaska", []float32{60.15, -176.44}, Case{Confirmed: 85335, Death: 42}}, // 修正后的经纬数据
	{"Arizona", []float32{34.15, -114.17}, Case{Confirmed: 69927, Death: 690}},
	{"Arkansas", []float32{34.73, -94.37}, Case{Confirmed: 24924, Death: 975}},
	{"California", []float32{37.19, -123.79}, Case{Confirmed: 51127, Death: 529}},
	{"Colorado", []float32{38.98,-107.79}, Case{Confirmed: 80437, Death: 817}},
	{"Connecticut", []float32{38.98, -107.79}, Case{Confirmed: 57869, Death: 665}},
	{"Delaware", []float32{39.14, -75.95}, Case{Confirmed: 84146, Death: 772}},
	{"District of Columbia", []float32{38.89, -77.08}, Case{Confirmed: 36095, Death: 260}},
	{"Florida", []float32{27.53, -88.30}, Case{Confirmed: 66176, Death: 629}},
	{"Georgia", []float32{32.66, -85.42}, Case{Confirmed: 86849, Death: 582}},
	{"Hawaii", []float32{20.45, -159.75}, Case{Confirmed: 43028, Death: 555}},
	{"Idaho", []float32{45.48, -116.39}, Case{Confirmed: 11296, Death: 370}},
	{"Illinois", []float32{39.72, -91.51}, Case{Confirmed: 67732, Death: 203}},
	{"Indiana", []float32{39.76,-87.56}, Case{Confirmed: 60185, Death: 657}},
	{"Iowa", []float32{41.92, -95.63}, Case{Confirmed: 77489, Death: 122}},
	{"Kansas", []float32{38.48, -100.56}, Case{Confirmed: 3470, Death: 384}},
	{"Kentucky", []float32{37.81, -88.01}, Case{Confirmed: 70978, Death: 292}},
	{"Louisiana", []float32{30.92, -93.64}, Case{Confirmed: 14394, Death: 790}},
	{"Maine", []float32{45.18, -70.10}, Case{Confirmed: 17926, Death: 334}},
	{"Maryland", []float32{38.79, -79.48}, Case{Confirmed: 56374, Death: 354}},
	{"Massachusetts", []float32{42.03, -72.80}, Case{Confirmed: 52548, Death: 638}},
	{"Michigan", []float32{44.98, -88.51}, Case{Confirmed: 27238, Death: 896}},
	{"Minnesota", []float32{46.42, -95.60}, Case{Confirmed: 74360, Death: 627}},
	{"Mississippi", []float32{32.57,-91.00}, Case{Confirmed: 39155, Death: 701}},
	{"Missouri", []float32{38.29, -94.68}, Case{Confirmed: 65674, Death: 225}},
	{"Montana", []float32{46.60, -114.53}, Case{Confirmed: 57283, Death: 694}},
	{"Nebraska", []float32{41.42, -104.17}, Case{Confirmed: 78257, Death: 302}},
	{"Nevada", []float32{38.49, -119.27}, Case{Confirmed: 9676, Death: 176}},
	{"New Hampshire", []float32{44.00, -72.69}, Case{Confirmed: 93855, Death: 679}},
	{"New Jersey", []float32{40.07, -75.85}, Case{Confirmed: 37088, Death: 988}},
	{"New Mexico", []float32{34.15, -108.27}, Case{Confirmed: 6478, Death: 354}},
	{"New York", []float32{40.70, -74.26}, Case{Confirmed: 3257, Death: 987}},
	{"North Carolina", []float32{35.10, -84.35}, Case{Confirmed: 94568, Death: 403}},
	{"North Dakota", []float32{47.45, -102.54}, Case{Confirmed: 81917, Death: 680}},
	{"Ohio", []float32{40.35, -84.91}, Case{Confirmed: 41841, Death: 143}},
	{"Oklahoma", []float32{35.24, -103.21}, Case{Confirmed: 68609, Death: 569}},
	{"Oregon", []float32{44.13, -122.83}, Case{Confirmed: 55236, Death: 509}},
	{"Pennsylvania", []float32{41.10, -79.85}, Case{Confirmed: 92852, Death: 744}},
	{"Rhode Island", []float32{41.56, -71.78}, Case{Confirmed: 3461, Death: 304}},
	{"South Carolina", []float32{33.61, -83.17}, Case{Confirmed: 17271, Death: 434}},
	{"South Dakota", []float32{44.19, -102.49}, Case{Confirmed: 90326, Death: 405}},
	{"Tennessee", []float32{35.76, -90.47}, Case{Confirmed: 33063, Death: 137}},
	{"Texas", []float32{31.10, -104.57}, Case{Confirmed: 60259, Death: 40}},
	{"Utah", []float32{39.48, -113.79}, Case{Confirmed: 59944, Death: 159}},
	{"Vermont", []float32{43.87, -73.01}, Case{Confirmed: 30527, Death: 391}},
	{"Virginia", []float32{37.93, -83.91}, Case{Confirmed: 32020, Death: 681}},
	{"Washington", []float32{47.25, -123.12}, Case{Confirmed: 19176, Death: 991}},
	{"West Virginia", []float32{38.90, -82.42}, Case{Confirmed: 31184, Death: 695}},
	{"Wisconsin", []float32{44.88, -91.81}, Case{Confirmed: 91178, Death: 100}},
	{"Wyoming", []float32{42.98, -109.80}, Case{Confirmed: 72238, Death: 509}},
	{"Puerto Rico", []float32{18.20, -67.71}, Case{Confirmed: 5621, Death: 794}},
}

func main() {
	r := gin.Default()
	r.GET("/api/states/all", getCountyHandler)
	r.GET("/api/search", handleSearch)
	_ = r.Run(":18111")
}

func getCountyHandler(c *gin.Context) {
	var county []string
	for _, v := range data {
		county = append(county, v.Name)
	}
	c.JSON(http.StatusOK, gin.H{
		"data": county,
	})
}

func handleSearch(c *gin.Context) {
	searchString := c.Query("string")

	dsn := "cnsoftbei:tF#262420228@tcp(127.0.0.1:3306)/epidemic?charset=utf8"
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
