package main

import "fmt"

func main() {
	a1 := []int{}
	a2 := []int{1}
	b1 := []int{1, 2}
	b2 := []int{3, 4}

	fmt.Printf("a1,a2:%v\n", findMedianSortedArrays(a1, a2))
	fmt.Printf("b1,b2:%v\n", findMedianSortedArrays(b1, b2))
}
// runtime test:38ms

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1 := len(nums1)
	len2 := len(nums2)
	array := make([]int, len1+len2)
	var i, j int
	var k int
	for i, j, k = 0, 0, 0; i < len1 && j < len2; {
		switch t := nums1[i] - nums2[j]; {
		case t < 0:
			array[k] = nums1[i]
			i++
			k++
		case t > 0:
			array[k] = nums2[j]
			j++
			k++
		case t == 0:
			array[k] = nums1[i]
			k++
			i++
			array[k] = nums2[j]
			k++
			j++
		}
	}
	if i != len1 {
		for ; i < len1; i++ {
			array[k] = nums1[i]
			k++
		}
	}
	if j != len2 {
		for ; j < len2; j++ {
			array[k] = nums2[j]
			k++
		}
	}
	fmt.Printf("array :%v\n", array)
	lenArray := len(array)
	if lenArray%2 == 0 {
		return float64((array[lenArray/2-1] + array[lenArray/2])) / 2
	} else {
		return float64(array[lenArray/2])
	}
}

//the simplest way runtime test :52ms
// import "sort"
//func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
//    merged := append(nums1, nums2...)
//    sort.Ints(merged)
//
//    totalLen := len(merged)
//    if totalLen % 2 == 0 {
//
//        return float64((merged[(totalLen/2)-1] + merged[totalLen/2])) / 2.0
//    } else {
//
//        return float64(merged[totalLen/2])
//    }
//}
