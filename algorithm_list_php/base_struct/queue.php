<?php

// 队列是一种遵循先进先出（FIFO）原则的数据结构。在PHP中，可以使用数组来实现队列
class Queue {
    private $queue = array();

    // 入队操作
    public function enqueue($item) {
        array_push($this->queue, $item);
    }

    // 出队操作
    public function dequeue() {
        return array_shift($this->queue);
    }

    // 获取队首元素
    public function front() {
        return current($this->queue);
    }

    // 判断队列是否为空
    public function isEmpty() {
        return empty($this->queue);
    }

    // 获取队列的大小
    public function size() {
        return count($this->queue);
    }
}
