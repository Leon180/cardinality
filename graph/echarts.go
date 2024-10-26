package graph

import (
	"net/http"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/samber/lo"
)

func NewLineChart[T interface{}]() *LineChart[T] {
	return &LineChart[T]{
		line: charts.NewLine(),
	}
}

type LineChart[T interface{}] struct {
	line *charts.Line
}

func (lc *LineChart[T]) SetGlobalOptions(
	titleName *string,
	xAxisName *string,
	yAxisName *string,
) {
	if lc == nil {
		return
	}
	funcList := []charts.GlobalOpts{}
	if titleName != nil {
		funcList = append(funcList, charts.WithTitleOpts(opts.Title{Title: *titleName}))
	}
	if xAxisName != nil {
		funcList = append(funcList, charts.WithXAxisOpts(opts.XAxis{Name: *xAxisName}))
	}
	if yAxisName != nil {
		funcList = append(funcList, charts.WithYAxisOpts(opts.YAxis{Name: *yAxisName}))
	}
	lc.line.SetGlobalOptions(funcList...)
}

type Serie[T interface{}] struct {
	seriesName string
	yValues    []T
	options    []charts.SeriesOpts
}

func NewSerie[T interface{}](
	seriesName string,
	yValues []T,
	options []charts.SeriesOpts,
) Serie[T] {
	return Serie[T]{
		seriesName: seriesName,
		yValues:    yValues,
		options:    options,
	}
}

func (lc *LineChart[T]) AddSeries(
	xValues []T,
	series []Serie[T],
) {
	if lc == nil {
		return
	}
	lc.line.SetXAxis(xValues)
	for _, s := range series {
		if s.options != nil {
			lc.line.AddSeries(s.seriesName, generateLineItems(s.yValues), s.options...)
		} else {
			lc.line.AddSeries(s.seriesName, generateLineItems(s.yValues))
		}
	}
}

func (lc *LineChart[T]) Render(w http.ResponseWriter) {
	lc.line.Render(w)
}

func generateLineItems[T interface{}](data []T) []opts.LineData {
	return lo.Map(data, func(item T, _ int) opts.LineData {
		return opts.LineData{Value: item}
	})
}
