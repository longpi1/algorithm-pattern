<?php

//链表是一种线性数据结构，由一组节点组成，每个节点包含数据和指向下一个节点的指针。
//在PHP中，可以使用对象来实现链表
class ListNode {
    public $data;
    public $next;

    public function __construct($data = null, $next = null) {
        $this->data = $data;
        $this->next = $next;
    }
}

class SingleLinkList {
    public function headInsert($n) {
        $head = new ListNode();
        for ($i=$n; $i > 0; $i--) {
            $newNode = new ListNode($i, $head->next);
            $head->next = $newNode;
        }
        return $head;
    }
}
