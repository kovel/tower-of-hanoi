package main

import (
	"fmt"
)

const DISKS = 3

func doTowers(topN int, from rune, inter rune, to rune, padding string) {
	if topN == 1 {
		fmt.Printf("Disk 1 from %c to %c\n", from, to)
	} else {
		doTowers(topN-1, from, to, inter, padding+"\t")
		fmt.Printf("%sDisk %d from %c to %c\n", padding, topN, from, to)
		doTowers(topN-1, inter, from, to, padding+"\t")
	}
}

func main() {
	doTowers(DISKS, 'A', 'B', 'C', "")
}
