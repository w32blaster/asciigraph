package main

import (
	"fmt"
	"github.com/w32blaster/asciigraph"
)

func main() {
	var data []float64

	// sine curve
	for i := 0; i < 105; i++ {
		//data = append(data, 15*math.Sin(float64(i)*((math.Pi*4)/120.0)))
		data = append(data, 4.0)
	}
	graph, offset := asciigraph.Plot(data, asciigraph.Height(10))

	fmt.Println(graph)
	fmt.Println(offset)
	// Output:
	//   15.00 ┤          ╭────────╮                                                  ╭────────╮
	//   12.00 ┤       ╭──╯        ╰──╮                                            ╭──╯        ╰──╮
	//    9.00 ┤    ╭──╯              ╰─╮                                       ╭──╯              ╰─╮
	//    6.00 ┤  ╭─╯                   ╰──╮                                  ╭─╯                   ╰──╮
	//    3.00 ┤╭─╯                        ╰─╮                              ╭─╯                        ╰─╮
	//    0.00 ┼╯                            ╰╮                            ╭╯                            ╰╮
	//   -3.00 ┤                              ╰─╮                        ╭─╯                              ╰─╮
	//   -6.00 ┤                                ╰─╮                   ╭──╯                                  ╰─╮
	//   -9.00 ┤                                  ╰──╮              ╭─╯                                       ╰──╮
	//  -12.00 ┤                                     ╰──╮        ╭──╯                                            ╰──╮
	//  -15.00 ┤                                        ╰────────╯                                                  ╰───
}
