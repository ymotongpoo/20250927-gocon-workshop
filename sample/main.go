// Copyright 2025 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"os"
	"runtime/pprof"
)

func Fib(n int) int {
	switch n {
	case 0:
		return 0
	case 1:
		return 1
	default:
		return Fib(n-1) + Fib(n-2)
	}
}

func main() {
	report, _ := os.Create("cpu.prof")
	defer report.Close()
	pprof.StartCPUProfile(report)
	defer pprof.StopCPUProfile()

	fmt.Println("running fibonatti")
	i := 40
	fmt.Printf("%dth fibonatti: %d\n", i, Fib(i))
	fmt.Println("done")
}
