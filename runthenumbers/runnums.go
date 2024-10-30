package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	for i := 1; i < 4; i++ {

		fmt.Println("Data", i)
		fmt.Println()

		for j := 1; j < 6; j++ {
			command := fmt.Sprintf("./numwriter %v %v | go run . -points=true | tail -1", i, j)
			cmd := exec.Command("bash", "-c", command)

			points, err1 := cmd.CombinedOutput()
			if err1 != nil {
				log.Fatalf("Command execution failed: %s", err1)
			}

			fmt.Print(string(points))
		}

		fmt.Println()
	}
}
