<?php
// 栈是一种遵循后进先出（LIFO）原则的数据结构。在PHP中，可以使用数组来实现栈，然后使用数组的内置函数来实现栈的各种操作
namespace base;
class Stack
{
    private $stack = array();

    // 入栈操作
    public function push($item)
    {
        array_push($this->stack, $item);
    }

    // 出栈操作
    public function pop()
    {
        return array_pop($this->stack);
    }

    // 获取栈顶元素
    public function peek()
    {
        return end($this->stack);
    }

    // 判断栈是否为空
    public function isEmpty()
    {
        return empty($this->stack);
    }

    // 获取栈的大小
    public function size()
    {
        return count($this->stack);
    }
}
