<?php

//49. 字母异位词分组
//给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
//
//字母异位词 是由重新排列源单词的所有字母得到的一个新单词。
//示例 1:
//
//输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
//输出: [["bat"],["nat","tan"],["ate","eat","tea"]]
//示例 2:
//输入: strs = [""]
//输出: [[""]]
//示例 3:
//
//输入: strs = ["a"]
//输出: [["a"]]

class Solution {

    /**
     * @param String[] $strs
     * @return String[][]
     */
    function groupAnagrams($strs) {
        $map = [];  // 创建一个空数组，用于存储分组后的字母异位词
        for ($i = 0; $i < count($strs); $i++) {
            $arr = str_split($strs[$i]);  // 将当前字符串拆分为字符数组
            sort($arr);  // 对字符数组进行排序
            $sorted_str = implode('', $arr);  // 将排序后的字符数组合并为一个字符串
            $map[$sorted_str][] = $strs[$i];  // 将当前字符串添加到以排序后的字符串为键的分组中
        }
        return array_values($map);  // 返回分组后的结果，array_values() 函数返回 map 数组中所有的值
    }
}