package main

func lonelyinteger(a []int32) int32 {
    // Write your code here
    b := int32(0)
    for _, i := range a {
        b ^= i
    }
    return b
}

func main(){
	lonelyinteger([]int32 {1,2,3,4,5,6,7,8,9,10,11})
}