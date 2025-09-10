// Copyright 2025 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This application is Go port of `cut` command.
// This is just a sample application so it's not the perfect clone.
// This is used for the exercise part of https://connpass.com/event/144347/
//
// For testing, prepare large file where fields are splitted by same characters.
//
// Sample data: https://excelbianalytics.com/wp/downloads-21-sample-csv-files-data-sets-for-testing-till-5-million-records-hr-analytics-for-attrition/
// * 10k records: https://excelbianalytics.com/wp/wp-content/uploads/2021/09/10000-HRA-Records.zip
// * 1.5M records: https://excelbianalytics.com/wp/wp-content/uploads/2021/09/1500000-HRA-Records.zip
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

var (
	delimiter byte
	field     int
)

func init() {
	var delimiterStr string
	flag.StringVar(&delimiterStr, "d", ",", "delimiter")
	delimiter = delimiterStr[0]
	flag.IntVar(&field, "f", 1, "field position")
}

// 使用方法: ./cut1 -f 3 -d ',' foo.csv
// フィールドに関しては範囲を表すハイフンは利用しない。
func main() {
	flag.Parse()
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Could not open file %q: %v", flag.Arg(0), err)
	}
	defer f.Close()

	// create pprof report to record
	report, _ := os.Create("cpu.prof.2")
	defer report.Close()
	_ = pprof.StartCPUProfile(report)
	defer pprof.StopCPUProfile()

	r := bufio.NewReader(f)

	for {
		// syscall.read が多すぎるので、その回数を減らすために1行ごと読み込むことにした
		line, err := r.ReadBytes('\n')
		// ファイルからの読み込みが出来ない場合やファイル末尾の場合は終了する
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not read file %q properly: %v", flag.Arg(0), err)
		}
		// bytes.SplitNは標準パッケージで十分テストされているので速いと考える。
		// ただし []byte(",") は毎回変換されているのでコストがかかっていないか
		// よく確認した方が良い
		fields := bytes.SplitN(line, []byte{delimiter}, field+1)
		// ここもまだ効率が悪いけど、残しておく
		// - 1行ごとにflushしてるので効率が悪い
		// - stringに変換してるのも効率が悪い
		fmt.Println(string(fields[field-1]))
	}
}
