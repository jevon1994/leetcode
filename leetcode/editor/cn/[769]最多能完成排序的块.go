//leetcode submit region begin(Prohibit modification and deletion)
func maxChunksToSorted(arr []int) (ans int) {
	mx := 0
	for i, x := range arr {
		if x > mx {
			mx = x
		}
		if mx == i {
			ans++
		}
		fmt.Printf("i:%d, x:%d , mx:%d \n", i, x, mx)
	}
	return
}

//leetcode submit region end(Prohibit modification and deletion)
