package main

import (
	"xixi_golang/pipeline"
	"fmt"
)

func main() {
	p := pipeline.ArraySource(3,2,1,5,8,9)

		for (v := range p)
		{
			fmt.Println(v)
		}

}
