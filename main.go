package main

// see:
// https://code.visualstudio.com/docs/languages/go

import (
	"math"
	"fmt"
	"math/rand"
)

func abs(x int) int {
	if x<0 {
		return -x
	}
	return x
}

func dot(v1, v2 []int) (res int) {
	for i, v := range v1 {
		res += v*v2[i]
	}
	return
}

func is_close(target, suggestion, tolerance int) (res bool) {
	res = abs(target-suggestion)<tolerance
	return
}

func find_W(V []int, T, max_weight, max_iter, TOL int, SEED int64) (W []int, Error int) {
	
	rand.Seed(SEED)
	
	/*
	let n be the cardinality of V, i.e. n=len(V), then there are max_weight^n possible available configurations of W. 
	*/
	var configurations int = int(math.Pow(float64(max_weight), float64(len(V))))
	max_iter = int(math.Min(float64(max_iter), float64(4*configurations)))

	fmt.Println("Randomly propose", max_iter, "out of totally", configurations ,"configurations of the weights, each bounded by", max_weight)
	
	var res int
	var P int = max_iter/100 // logging...
	
	for i:=0; i<max_iter; i++ { 

		// logging...
		if i%P==0 { 
			fmt.Print(".")
		}
		
		// array of weights
		W = []int{}
		WS := 0 // space, i.e. sum of W
		tmp := 0

		// randomize weigt array
		for j:=0; j<len(V); j++ {
			tmp = rand.Intn(max_weight + 1)
			WS += tmp
			W = append(W, tmp)
		}
		
		// solution?
		res = dot(V, W)
		if is_close(res, T*WS, TOL) {
			Error := abs(res/WS-T)
			fmt.Println("\nfound weights after", i+1, "iterations. Error: ", Error)
			return W, Error
		}
	}
	fmt.Println("\nCould not find weights after", max_iter, "iterations...")
	W = []int{}
	return W, -1
}

func main() {
	// input: values
	V1 := []int{1200, 1100, 1150, 1300, 1600, 2500}
	T1 := 1580
	
	V2 := []int{1500, 1800, 2000, 2100, 2400}
	T2 := 2000
	
	max_weight := 4
	max_iter := 50000000
	TOL := 10
	var SEED int64 = 20210213

	fmt.Println("\n\nrunning first vals...\n")
	W, Error := find_W(V1, T1, max_weight, max_iter, TOL, SEED)
	fmt.Println(W, "\n", Error)

	fmt.Println("\n\nrunning second vals...\n")
	W, Error = find_W(V2, T2, max_weight, max_iter, TOL, SEED)
	fmt.Println(W, Error)
	
}