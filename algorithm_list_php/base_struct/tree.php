<?php
// 树是一种非线性数据结构，由节点和连接节点的边组成。在PHP中，可以使用对象来实现树
 class TreeNode {
     public $val = null;
     public $left = null;
     public $right = null;
     function __construct($val = 0, $left = null, $right = null) {
         $this->val = $val;
         $this->left = $left;
         $this->right = $right;
     }
 }
