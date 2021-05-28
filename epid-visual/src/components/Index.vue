<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          src="https://cdn.jsdelivr.net/gh/apache/echarts-website@asf-site/zh/images/index-home-pie.png?_v_=20200710_1"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4" cols="12"> <!-- title -->
        <h2 class="display-1 font-weight-bold mb-3">
          欢迎使用疫情信息查询及趋势预测系统
        </h2>
      </v-col>

      <v-row justify="center">
        <v-col
          class="mb-5"
          cols="7"
        > <!-- search -->
          <v-text-field
            v-model="searchString"
            label="输入想要查询的城市/省份/州"
            dense
            outlined
            clearable
            @keyup.enter="searchString"
          ></v-text-field>
        </v-col>
        <v-col
          cols="1"
        >
          <v-btn
            color="primary"
            @click="search"
            :loading="loading"
            :disabled="loading"
          >
            搜索
          </v-btn>
        </v-col>
        <v-col
          cols="12"
        > <!-- map -->
          <div ref="echart" style="margin: 0 auto; left: 0; right: 0;height: 400px; width: 60%;"></div>
        </v-col>
      </v-row>
    </v-row>
  </v-container>
</template>

<script>
  import * as echarts from 'echarts';
  export default {
    name: 'Index',

    data: () => ({
      searchString: '',
      loading: false,
      echart: {}
    }),
    mounted: function(){
      this.echart = echarts.init(this.$refs.echart);
      this.echart.showLoading();
      fetch("https://cdn.jsdelivr.net/gh/apache/echarts-website@asf-site/examples/data/asset/geo/USA.json").then( res => {
        return res.json();
      }).then( usaJSON => {
        this.echart.hideLoading();

        echarts.registerMap('USA', usaJSON, {
          Alaska: {              // 把阿拉斯加移到美国主大陆左下方
            left: -131,
            top: 25,
            width: 15
          },
          Hawaii: {
            left: -110,        // 夏威夷
            top: 28,
            width: 5
          },
          'Puerto Rico': {       // 波多黎各
            left: -76,
            top: 26,
            width: 2
          }
        });

        var option = {
          
          title: {
            text: 'USA Covid-19 Estimates',
            left: 'right'
          },
          tooltip: {
            trigger: 'item',
            showDelay: 0,
            transitionDuration: 0.2,
            formatter: function (params) {
              var cases = (params.data.value[1] + '').split('.');
              cases = cases[0].replace(/(\d{1,3})(?=(?:\d{3})+(?!\d))/g, '$1,');
              var death = (params.data.value[0] + '').split('.');
              death = death[0].replace(/(\d{1,3})(?=(?:\d{3})+(?!\d))/g, '$1,');
              return params.data.name + '<br/>' + 'cases: ' + cases + '<br/>' + 'death: ' + death;
            }
          },
          visualMap: {
            left: 'right',
            min: 0,
            max: 100000,
            seriesIndex: 0,
            inRange: {
              color: ['#313695', '#4575b4', '#74add1', '#abd9e9', '#e0f3f8', '#ffffbf', '#fee090', '#fdae61', '#f46d43', '#d73027', '#a50026']
            },
            text: ['High', 'Low'],           // 文本，默认为数值文本
            calculable: true
          },
          series: [{
            name: 'USA PopEstimates',
            type: 'map',
            roam: false,
            map: 'USA',
            emphasis: {
              label: {
                show: true
              }
            },
            // 文本位置修正
            textFixed: {
              Alaska: [20, -20]
            },
            data: [] // e.g. {"name": "test", "value": [100, 10000]}
          }]
        };

        fetch(`http://34.96.239.6:8111/api/states/all`).then( res => {
          return res.json();
        }).then( res => {
          option.series[0].data = res.data;
          this.echart.setOption(option);
        }).catch( err => {
          throw new Error(err);
        })

      });
    },
    methods: {
      search: function(){
        var dataValue = this.dealWithData();
        var option = {
          geo: {
            roam: false,
            center: [31.38, -96.08],
            label: {
              show: true,
            },
          },      
          series: [
            {
              type: "map",
              map: 'USA',
              roam: false,
              label: {
                normal: {
                  show: false
                },
                emphasis: {
                  show: false
                }
              },
            },
            // {
            //   name: "",
            //   type: "scatter",
            //   coordinateSystem: "geo",
            //   data: dataValue,
            //   symbol: "circle",
            //   symbolSize: 5,
            //   hoverSymbolSize: 10,
            //   tooltip: {
            //     formatter(value) {
            //       return value.data.name + "<br/>" + "确诊人数：" + value.data.case.confirmed + "<br/>" + "死亡人数：" + value.data.case.death;
            //     },
            //     show: true
            //   },
            //   encode: {
            //     value: 2
            //   },
            //   label: {
            //     formatter: "{b}",
            //     position: "right",
            //     show: false
            //   },
            //   itemStyle: {
            //     color: "#0efacc"
            //   },
            //   emphasis: {
            //     label: {
            //       show: false
            //     }
            //   }
            // },
            {
              name: "Selected",
              type: "effectScatter",
              coordinateSystem: "geo",
              data: dataValue,
              symbolSize: 9,
              tooltip: {
                formatter(value) {
                  return value.data.name + "<br/>" + "确诊人数：" + value.data.case.confirmed + "<br/>" + "死亡人数：" + value.data.case.death;
                },
                show: true
              },
              encode: {
                value: 2
              },
              showEffectOn: "render",
              rippleEffect: {
                brushType: "stroke",
                color: "#0efacc",
                period: 9,
                scale: 5
              },
              hoverAnimation: true,
              label: {
                formatter: "{b}",
                position: "right",
                show: true
              },
              itemStyle: {
                color: "#0efacc",
                shadowBlur: 2,
                shadowColor: "#333"
              },
              zlevel: 1
            }
          ]
        };
        this.echart.setOption(option);
        
        this.loading = true;
        fetch(`http://34.96.239.6:8111/api/search?string=${this.searchString}`).then( res => {
          return res.json();
        }).then( res => {
          console.log(res);
        }).catch( err => {
          throw new Error(err);
        })
        this.loading = false;
      },
      dealWithData: function () {
        // var geoCoordMap = {
        //   北京: { geo: [116.46, 39.92], case: {confirmed: 100, death: 1} },
        //   上海: { geo: [121.48, 31.22], case: {confirmed: 745, death: 15} },
        //   广州: { geo: [113.23, 23.16], case: {confirmed: 280, death: 7} },
        //   深圳: { geo: [114.07, 22.62], case: {confirmed: 113, death: 1} },
        //   成都: { geo: [104.06, 30.67], case: {confirmed: 113, death: 1} },
        //   西安: { geo: [108.95, 34.27], case: {confirmed: 57, death: 2} },
        //   拉萨: { geo: [91.11, 29.97], case: {confirmed: 1, death: 0} },
        //   南宁: { geo: [108.33, 22.84], case: {confirmed: 4,death: 0} },
        //   太原: { geo: [112.53, 37.87], case: {confirmed: 10, death: 1} },
        //   昆明: { geo: [102.73, 25.04], case: {confirmed: 19, death: 3} },
        //   海口: { geo: [110.35, 20.02], case: {confirmed: 9, death: 0} },
        //   沈阳: { geo: [123.38, 41.8], case: {confirmed: 23, death: 2} },
        //   银川: { geo: [106.27, 38.47], case: {confirmed: 3, death: 0} },
        //   南昌: { geo: [115.89, 28.68], case: {confirmed: 9, death: 0} },
        //   吉林: { geo: [126.57, 43.87], case: {confirmed: 27, death: 2} },
        //   西宁: { geo: [101.74, 36.56], case: {confirmed: 7, death: 0} },
        //   福州: { geo: [119.3, 26.08], case: {confirmed: 33, death: 1} },
        //   呼和浩特: { geo: [111.65, 40.82], case: {confirmed: 29, death: 1} },
        //   重庆: { geo: [106.54, 29.59], case: {confirmed: 89, death: 5} }, 
        //   南京: { geo: [118.78, 32.04], case: {confirmed: 102, death: 9} },
        //   贵阳: { geo: [106.71, 26.57], case: {confirmed: 15, death: 1} },
        //   乌鲁木齐: { geo: [87.68, 43.77], case: {confirmed: 13, death: 0} },
        //   杭州: { geo: [120.19, 30.26], case: {confirmed: 89, death: 7} },
        //   济南: { geo: [117, 36.65], case: {confirmed: 46, death: 2} },
        //   兰州: { geo: [103.73, 36.03], case: {confirmed: 18, death: 0} },
        //   天津: { geo: [117.2, 39.13], case: {confirmed: 55, death: 1} },
        //   郑州: { geo: [113.65, 34.76], case: {confirmed: 78, death: 6} },
        //   哈尔滨: { geo: [126.63, 45.75], case: {confirmed: 24, death: 3} },
        //   石家庄: { geo: [114.48, 38.03], case: {confirmed: 33, death: 4} },
        //   长沙: { geo: [113, 28.21], case: {confirmed: 44, death: 3} },
        //   合肥: { geo: [117.27, 31.86], case: {confirmed: 80, death: 3} },
        //   武汉: { geo: [114.31, 30.52], case: {confirmed: 10486, death: 208} },
        // };
        var data = [];
        if (this.searchString && geoCoordMap[this.searchString])
          data.push({ name: this.searchString, value: geoCoordMap[this.searchString].geo, case: geoCoordMap[this.searchString].case });
        
        return data;
      }
    },
  }
</script>
