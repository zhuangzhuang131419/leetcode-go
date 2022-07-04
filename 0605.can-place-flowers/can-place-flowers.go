package leetcode

func canPlaceFlowers(flowerbed []int, n int) bool {

	result := 0

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 || flowerbed[i+1] == 1 && i+1 < len(flowerbed) || flowerbed[i-1] == 1 && i-1 >= 0 {
			continue
		}
		result++
		flowerbed[i] = 1
	}

	return result >= n

}
