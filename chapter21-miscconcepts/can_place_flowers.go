// Copyright (c) July 12, 2020 CareerMonk Publications and others.
// E-Mail           		: info@careermonk.com
// Creation Date    		: 2020-07-12 06:15:46
// Last modification		: 2020-07-12
//               by		: Narasimha Karumanchi
// Book Title			: Data Structures And Algorithmic Thinking With Go
// Warranty         		: This software is provided "as is" without any
// 				   warranty; without even the implied warranty of
// 				    merchantability or fitness for a particular purpose.

package main
import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
    lenF, count := len(flowerbed), 0
    for i := range flowerbed{
        if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == lenF - 1 || flowerbed[i+1] == 0){
            flowerbed[i] = 1
            count++
        } 
    }
    return count >= n
}

func main() {
	fmt.Println(countSetBits(5))
	fmt.Println(countSetBits(10))
	fmt.Println(countSetBits(2))
	fmt.Println(countSetBits(4))
}
