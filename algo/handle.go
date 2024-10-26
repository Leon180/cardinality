package algo

import (
	"cardinality/algo/morris"
	"cardinality/graph"
	"net/http"
	"strconv"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/samber/lo"
)

const (
	defaultTestTime          int    = 10
	defaultCount             int    = 1024
	chartTitle               string = "Morris Counter"
	chartXAxisName           string = "event"
	chartYAxisName           string = "estimate count"
	defaultExpectValueColor  string = "#808080"
	defaultExpectValueType   string = "dotted"
	defaultTestAvgValueColor string = "black"
	defaultTestAvgValueType  string = "dashed"
)

func NewMorrisCounterHandle() *MorrisCounterHandle {
	return &MorrisCounterHandle{}
}

type MorrisCounterHandle struct {
}

func (h *MorrisCounterHandle) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		testTime, countPerTest                                                 int
		expectValueColor, expectValueType, testAvgValueColor, testAvgValueType string
		err                                                                    error
	)

	// get params
	if testTime, err = strconv.Atoi(r.FormValue("testTime")); err != nil {
		testTime = defaultTestTime
	}
	if countPerTest, err = strconv.Atoi(r.FormValue("countPerTest")); err != nil {
		countPerTest = defaultCount
	}
	if expectValueColor = r.FormValue("expectValueColor"); expectValueColor == "" {
		expectValueColor = defaultExpectValueColor
	}
	if expectValueType = r.FormValue("expectValueType"); expectValueType == "" {
		expectValueType = defaultExpectValueType
	}
	if testAvgValueColor = r.FormValue("testAvgValueColor"); testAvgValueColor == "" {
		testAvgValueColor = defaultTestAvgValueColor
	}
	if testAvgValueType = r.FormValue("testAvgValueType"); testAvgValueType == "" {
		testAvgValueType = defaultTestAvgValueType
	}

	// run test
	testResults := make([][]int, testTime)
	for i := range testResults {
		mc := morris.NewMorrisCounterTest(countPerTest)
		mc.Run()
		testResults[i] = mc.GetEstimateCounts()
	}

	// draw chart
	intLineChart := graph.NewLineChart[int]()
	title, xAxisName, yAxisName := chartTitle, chartXAxisName, chartYAxisName
	intLineChart.SetGlobalOptions(
		&title,
		&xAxisName,
		&yAxisName,
	)
	series := []graph.Serie[int]{}
	// handle expect value
	expectYValues := make([]int, countPerTest)
	for i := range expectYValues {
		expectYValues[i] = i + 1
	}
	series = append(series, graph.NewSerie("expect value", expectYValues, []charts.SeriesOpts{charts.WithLineStyleOpts(opts.LineStyle{Color: expectValueColor, Type: expectValueType})}))
	// handle test value
	avgTestYValues := make([]int, countPerTest)
	for i := range avgTestYValues {
		avgTestYValues[i] = lo.Sum(lo.Map(testResults, func(item []int, _ int) int {
			return item[i]
		})) / len(testResults)
	}
	series = append(series, graph.NewSerie("avg test value", avgTestYValues, []charts.SeriesOpts{charts.WithLineStyleOpts(opts.LineStyle{Color: testAvgValueColor, Type: testAvgValueType})}))
	for i := range testResults {
		series = append(series, graph.NewSerie("test "+strconv.Itoa(i+1), testResults[i], nil))
	}
	xValues := make([]int, countPerTest)
	for i := range xValues {
		xValues[i] = i + 1
	}
	intLineChart.AddSeries(xValues, series)
	intLineChart.Render(w)
}
