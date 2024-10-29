package main

import (
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"strings"
)

type average struct {
	input  float64
	result int
}

func main() {

	//inputs := []float64{0.1, 0.2, 0.3, 0.4, 0.5, 0.6, 0.7, 0.8, 0.9, 1.0, 1.1, 1.2, 1.3, 1.4, 1.5, 1.6, 1.7, 1.8, 1.9}
	//inputs := []float64{0.61, 0.62, 0.63, 0.64, 0.65, 0.66, 0.67, 0.68, 0.69, 0.70, 0.71, 0.72, 0.73, 0.74, 0.75, 0.76, 0.77, 0.78, 0.79}
	//inputs := []float64{0.66, 0.67, 0.68, 0.69, 0.70, 0.71, 0.72, 0.73, 0.74, 0.75, 0.76, 0.77, 0.78, 0.79, 0.80, 0.81, 0.82, 0.83, 0.84, 0.85}
	inputs := []float64{1.03, 1.04, 1.05, 1.06, 1.07, 1.08, 1.09, 1.10, 1.11, 1.12, 1.13, 1.14, 1.15, 1.16, 1.17, 1.18, 1.19, 1.20, 1.21, 1.22}

	avgs := []average{}

	for _, in := range inputs {

		outputs := []int{}
		for i := 1; i < 4; i += 2 {
			for j := 1; j < 6; j++ {
				line := fmt.Sprintf("./numwriter %v %v | go run . -points=true %f | tail -1", i, j, in)
				cmd := exec.Command("bash", "-c", line)

				output, err1 := cmd.CombinedOutput()
				if err1 != nil {
					log.Fatalf("Command execution failed: %s", err1)
				}
				resu, err2 := strconv.Atoi(strings.Fields(string(output))[0])
				if err2 != nil {
					log.Fatalln(err2.Error())
				}
				outputs = append(outputs, resu)
			}
		}
		sum := 0
		for _, n := range outputs {
			sum += n
		}
		avgs = append(avgs, average{in, sum / len(outputs)})
	}

	fmt.Println(avgs)
}
