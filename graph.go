package accmeter

import (
	"fmt"
	"slices"

	ui "github.com/gizak/termui/v3"
	"github.com/gizak/termui/v3/widgets"
)

func Graph(label string, pos []int, data [][2]float64) (*widgets.Plot, error) {
	if len(pos) != 4 {
		return nil, fmt.Errorf("pos must have 4 values")
	}

	lc := widgets.NewPlot()
	dataY := func() []float64 {
		result := make([]float64, len(data))
		for i, d := range data {
			result[i] = d[1]
		}
		return result
	}()
	dataX := func() []float64 {
		result := make([]float64, len(data))
		for i, d := range data {
			result[i] = d[0]
		}
		return result
	}()

	Ymax := slices.Max(dataY)
	Xmax := dataX[len(dataX)-1] - dataX[0]
	lc.Title = fmt.Sprintf("tmax: %.2fms, %smax: %.2fg", Xmax, label, Ymax)
	lc.Data = [][]float64{dataY}

	lc.SetRect(pos[0], pos[1], pos[2], pos[3])
	lc.AxesColor = ui.ColorWhite
	lc.LineColors[0] = ui.ColorGreen
	// lc.PlotType = widgets.ScatterPlot

	return lc, nil

	// ui.Render(lc)

	// for e := range ui.PollEvents() {
	// 	if e.Type == ui.KeyboardEvent {
	// 		break
	// 	}
	// }
}
