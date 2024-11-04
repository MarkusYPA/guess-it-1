package main

import (
	"fmt"
	"log"
	"os/exec"
	"runtime"
	"strconv"
)

var avgScores [][]int = [][]int{{94720, 100640, 100320, 87360, 104160}, {640, 320, 640, 480, 1280}, {960, 640, 320, 500, 1280}}
var medScores [][]int = [][]int{{94608, 99499, 101835, 98404, 102273}, {96433, 96287, 100229, 99134, 101543}, {100959, 99791, 102930, 94973, 97966}}

func main() {

	for i := 1; i < 4; i++ {

		fmt.Println("Data", i)
		fmt.Println()

		os := runtime.GOOS
		extension := ""
		if os == "windows" {
			extension = ".exe"
		}

		for j := 1; j < 6; j++ {
			command := fmt.Sprintf("./numwriter%v %v %v | go run . -points=true | tail -1", extension, i, j)
			cmd := exec.Command("bash", "-c", command)

			points, err1 := cmd.CombinedOutput()
			if err1 != nil {
				log.Fatalf("Command execution failed: %s", err1)
			}
			pointStr := string(points[:len(points)-1]) // last char is line feed
			pointInt, err2 := strconv.Atoi(pointStr)
			if err2 != nil {
				log.Fatalln(err2.Error())
			}

			winLossA := ""
			winLossM := ""
			if pointInt < avgScores[i-1][j-1] {
				winLossA = "L"
			} else {
				winLossA = "W"
			}
			if pointInt < medScores[i-1][j-1] {
				winLossM = "L"
			} else {
				winLossM = "W"
			}

			fmt.Printf("%v %v %v\n", pointStr, winLossA, winLossM)
		}

		fmt.Println()
	}
}
