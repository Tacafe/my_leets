/*
 * @lc app=leetcode id=836 lang=golang
 *
 * [836] Rectangle Overlap
 */

// @lc code=start
func hasIntersection(range1 []int, range2 []int) bool {
	from1 := range1[0]
	to1 := range1[1]
	from2 := range2[0]
	to2 := range2[1]

	return from1 < to2 && to1 > from2 || from2 < to1 && to2 > to1
}

func isRectangleOverlap(rec1 []int, rec2 []int) bool {
	hasXintersection := hasIntersection([]int{rec1[0], rec1[2]}, []int{rec2[0], rec2[2]})
	hasYintersection := hasIntersection([]int{rec1[1], rec1[3]}, []int{rec2[1], rec2[3]})
	return hasXintersection && hasYintersection
}
// @lc code=end

