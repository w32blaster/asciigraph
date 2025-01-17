package asciigraph

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

const precision = 0

// Plot returns ascii graph for a series.
func Plot(series []float64, options ...Option) (string, int) {
	config := configure(config{
		Offset: 3,
	}, options)

	if config.Width > 0 {
		series = interpolateArray(series, config.Width)
	}

	minimum, maximum := minMaxFloat64Slice(series)
	interval := math.Abs(maximum - minimum)

	if config.Height <= 0 {
		if int(interval) <= 0 {
			config.Height = int(interval * math.Pow10(int(math.Ceil(-math.Log10(interval)))))
		} else {
			config.Height = int(interval)
		}
	}

	if config.Offset <= 0 {
		config.Offset = 3
	}

	var ratio float64
	if interval != 0 {
		ratio = float64(config.Height) / interval
	} else {
		ratio = 1
	}
	min2 := round(minimum * ratio)
	max2 := round(maximum * ratio)

	intmin2 := int(min2)
	intmax2 := int(max2)

	rows := int(math.Abs(float64(intmax2 - intmin2)))
	width := len(series) + config.Offset

	var plot [][]string

	// initialise empty 2D grid
	for i := 0; i < rows+1; i++ {
		var line []string
		for j := 0; j < width; j++ {
			line = append(line, " ")
		}
		plot = append(plot, line)
	}

	maxNumLength := len(fmt.Sprintf("%0.*f", precision, maximum))
	minNumLength := len(fmt.Sprintf("%0.*f", precision, minimum))
	maxWidth := int(math.Max(float64(maxNumLength), float64(minNumLength)))

	// axis and labels
	for y := intmin2; y < intmax2+1; y++ {
		var magnitude float64
		if rows > 0 {
			magnitude = maximum - (float64(y-intmin2) * interval / float64(rows))
		} else {
			magnitude = float64(y)
		}

		label := fmt.Sprintf("%*.*f", maxWidth+1, precision, magnitude)
		w := y - intmin2
		h := int(math.Max(float64(config.Offset)-float64(len(label)), 0))

		plot[w][h] = label
		if y == 0 {
			plot[w][config.Offset-1] = "┼"
		} else {
			plot[w][config.Offset-1] = "┤"
		}
	}

	y0 := int(round(series[0]*ratio) - min2)
	var y1 int

	plot[rows-y0][config.Offset-1] = "┼" // first value

	for x := 0; x < len(series)-1; x++ { // plot the line
		y0 = int(round(series[x+0]*ratio) - float64(intmin2))
		y1 = int(round(series[x+1]*ratio) - float64(intmin2))
		if y0 == y1 {
			plot[rows-y0][x+config.Offset] = "─"
		} else {
			if y0 > y1 {
				plot[rows-y1][x+config.Offset] = "╰"
				plot[rows-y0][x+config.Offset] = "╮"
			} else {
				plot[rows-y1][x+config.Offset] = "╭"
				plot[rows-y0][x+config.Offset] = "╯"
			}

			start := int(math.Min(float64(y0), float64(y1))) + 1
			end := int(math.Max(float64(y0), float64(y1)))
			for y := start; y < end; y++ {
				plot[rows-y][x+config.Offset] = "│"
			}
		}
	}

	// join columns
	var lines bytes.Buffer
	for h, horizontal := range plot {
		if h != 0 {
			lines.WriteRune('\n')
		}
		for _, v := range horizontal {
			lines.WriteString(v)
		}
	}

	leftLabelOffset := config.Offset + maxWidth

	// bottom line
	lines.WriteRune('\n')
	lines.WriteString(strings.Repeat(" ", leftLabelOffset-1))
	lines.WriteRune('╰')
	for i := 0; i < len(plot[0])-3; i++ {
		if i%3 == 0 {
			lines.WriteRune('┬')
		} else {
			lines.WriteRune('─')
		}
	}

	// add caption if not empty
	if config.Caption != "" {
		lines.WriteRune('\n')
		lines.WriteString(strings.Repeat(" ", leftLabelOffset+2))
		lines.WriteString(config.Caption)
	}

	return lines.String(), leftLabelOffset
}
