struct Solution;

struct SolutionBasic;

//Simple double for checks all combinations O(n^2)
impl SolutionBasic {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let len = nums.len();
        for i in 0..len {
            for j in 1..len {
                if i != j && nums[i] + nums[j] == target {
                    return vec![i as i32, j as i32];
                }
            }
        }
        return Vec::new();
    }
}

use std::collections::HashMap;

//Use map and complement O(n)
impl Solution {
    pub fn two_sum(nums: Vec<i32>, target: i32) -> Vec<i32> {
        let mut map = HashMap::with_capacity(nums.len());

        for (i, &n) in nums.iter().enumerate() {
            let complement = target - n;
            if let Some(&idx) = map.get(&complement) {
                    return vec![i as i32, idx as i32]
            }
            map.insert(n, i);
        }
        return Vec::new();
    }
}

fn main() {
    let nums = vec![3,2,4];
    let target = 6;

    let solution = Solution::two_sum(nums, target);
    println!("{:?}", solution);
}