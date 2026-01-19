package main

import "testing"

func TestRunner(t *testing.T) {
	runner := LocalRunner{}
	res := runner.Run(`
package main
import "fmt"
func main() {
	fmt.Println("test")
}
`)
	if res.IsError() {
		t.Fatal("failed to run!", res.GetError().Error())
	}
	t.Log(res.GetOutput())
}
