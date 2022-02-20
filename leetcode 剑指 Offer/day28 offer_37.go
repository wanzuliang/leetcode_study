/*
题目描述：  剑指 Offer 37. 序列化二叉树
难度：  困难
编写日期：   2022 年 2 月 20 日 晚 20：00
重写日期：  
*/
/*题目：  

	请实现两个函数，分别用来序列化和反序列化二叉树。

    你需要设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

    提示：输入输出格式与 LeetCode 目前使用的方式一致，详情请参阅 LeetCode 序列化二叉树的格式(https://support.leetcode-cn.com/hc/kb/article/1194353/)。你并非必须采取这种方式，你也可以采用其他的方法解决这个问题。

    示例：
           1

       2      3

            4    5

    输入：root = [1,2,3,null,null,4,5]
    输出：[1,2,3,null,null,4,5]

*/
/*  
        暂时跳过
*/
package main

import (
    "fmt"
)

func main() {
    fmt.Println( "T")
}

// 默认代码模版
/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {
    private TreeNode root;
    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        this.root = root;
        return null;
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        return root;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec codec = new Codec();
// codec.deserialize(codec.serialize(root));