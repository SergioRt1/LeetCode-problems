struct Solution;
struct Solution1;

use std::collections::HashMap;
use std::collections::BTreeMap;
use std::hash::{Hash, Hasher};

//Use a Map calculate the anagram the hash the map and use a second map<Hash, Index>, then build the response
impl Solution1 {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut maps = HashMap::with_capacity(strs.len());

        for (i, s) in strs.iter().enumerate() {
            let mut letters = BTreeMap::new();
            let mut hasher = std::collections::hash_map::DefaultHasher::new();
            for c in s.chars() {
                *letters.entry(c as u8 - b'a').or_insert(0 as i32) += 1;
            }
            letters.hash(&mut hasher);
            let list_of_anagrams = maps.entry(hasher.finish()).or_insert(Vec::new());
            list_of_anagrams.push(i);
        }
        // build the response
        let mut response = Vec::with_capacity(maps.len());
        for (_, v) in &maps {
            let mut group = Vec::with_capacity(v.len());
            for &i in v.iter() {
                group.push(strs[i].clone());
            }
            response.push(group);
        }
        return response;
    }
}

// Use a vec instead of a BTreeMap
impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut maps = HashMap::with_capacity(strs.len());

        for (i, s) in strs.iter().enumerate() {
            let mut letters = vec![0; 26];

            for c in s.chars() {
                letters[c as usize - b'a' as usize] += 1;
            }
            maps.entry(letters).or_insert(Vec::new()).push(i);
        }
        // build the response
        let mut response = Vec::with_capacity(maps.len());
        for (_, v) in &maps {
            let mut group = Vec::with_capacity(v.len());
            for &i in v.iter() {
                group.push(strs[i].clone());
            }
            response.push(group);
        }
        return response;
    }
}


fn main() {
    let strs = vec!["eat", "tea", "tan", "ate", "nat", "bat"].iter().map(|&s| s.to_owned()).collect();

    let solution = Solution::group_anagrams(strs);
    println!("{:?}", solution);
}