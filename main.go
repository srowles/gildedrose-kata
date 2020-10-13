package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/srowles/gildedrose-kata/shop"
)

func main() {
	fmt.Println("OMGHAI!")

	var items = []*shop.Item{
		&shop.Item{"+5 Dexterity Vest", 10, 20},
		&shop.Item{"Aged Brie", 2, 0},
		&shop.Item{"Elixir of the Mongoose", 5, 7},
		&shop.Item{"Sulfuras, Hand of Ragnaros", 0, 80},
		&shop.Item{"Sulfuras, Hand of Ragnaros", -1, 80},
		&shop.Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		&shop.Item{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		&shop.Item{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		&shop.Item{"Conjured Mana Cake", 3, 6}, // <-- :O
	}

	days := 2
	var err error
	if len(os.Args) > 1 {
		days, err = strconv.Atoi(os.Args[1])
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		days++
	}

	for day := 0; day < days; day++ {
		fmt.Printf("-------- day %d --------\n", day)
		fmt.Println("name, sellIn, quality")
		for i := 0; i < len(items); i++ {
			fmt.Printf("%+v\n", items[i])
		}
		fmt.Println("")
		shop.UpdateQuality(items)
	}
}
