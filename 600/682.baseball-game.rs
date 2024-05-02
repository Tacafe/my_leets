/*
 * @lc app=leetcode id=682 lang=rust
 *
 * [682] Baseball Game
 */

// @lc code=start
impl Solution {
    pub fn cal_points(operations: Vec<String>) -> i32 {
        let mut scores = Vec::new();
        let mut prev_score = 0;
        for op in operations.iter() {
            match op.as_str() {
                "C" => {
                    scores.pop();
                    if scores.len() > 0 {
                        prev_score = scores[scores.len() - 1];
                    }
                },
                "D" => {
                    let doubled = prev_score * 2;
                    scores.push(doubled);
                    prev_score = doubled;
                },
                "+" => {
                    scores.push(scores[scores.len() - 1] + scores[scores.len() - 2]);
                    prev_score = scores[scores.len() - 1];
                }
                _ => {
                    let op_score = op.parse::<i32>().unwrap();
                    scores.push(op_score);
                    prev_score = op_score;
                }
            }
            // println!("{:?}", scores)
        }
        scores.iter().sum()
    }
}
// @lc code=end

