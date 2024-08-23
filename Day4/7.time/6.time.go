/*
Type	Placeholder
Day	2 or 02 or _2
Day of Week	Monday or Mon
Month	01 or 1 or Jan or January
Year	2006 or 06
Hour	03 or 3 or 15
Minutes	04 or 4
Seconds	05 or 5
Milli Seconds  (ms)	.000        //Trailing zero will be includedor .999   //Trailing zero will be omitted
Micro Seconds (μs)	.000000             //Trailing zero will be includedor .999999        //Trailing zero will be omitted
Nano Seconds (ns)	.000000000        //Trailing zero will be includedor .999999999 //Trailing zero will be omitted
am/pm	PM or pm
Timezone	MST
Timezone offset	Z0700 or Z070000 or Z07 or Z07:00 or Z07:00:00  or -0700 or  -070000 or -07 or -07:00 or -07:00:00
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Demo time.Parse")

	timeString := "2024/Aug/23 **** **** 14:05:14"
	timeFormat := "2006/Jan/02 **** **** 15:04:05"

	tm, err := time.Parse(timeFormat, timeString)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(tm)

}
