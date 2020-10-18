package pq

import (
    "testing"
)

func TestMaxPQ(t *testing.T) {
    t.Logf("=====>>>测试开始<<<=====")
    N := 15 // 设置容量
    data := NewMaxPQ(N)
    if !data.IsEmpty() {
        t.Errorf("此时优先队列应该是空的！")
    }
    t.Logf("54 76 11 100 32 1 1024 78 1 33 33 5 6 12 13 依次加入..")
    data.Insert(54)
    data.Out()
    data.Insert(76)
    data.Out()
    data.Insert(11)
    data.Out()
    data.Insert(100)
    data.Out()
    data.Insert(32)
    data.Out()
    data.Insert(1)
    data.Out()
    data.Insert(1024)
    data.Out()
    data.Insert(78)
    data.Out()
    data.Insert(1)
    data.Out()
    data.Insert(33)
    data.Out()
    data.Insert(33)
    data.Out()
    data.Insert(5)
    data.Out()
    data.Insert(6)
    data.Out()
    data.Insert(12)
    data.Out()
    data.Insert(13)
    data.Out()
    data.Pretty()
    max := data.DelMax()
    t.Logf("取出最大元素：%v，观察..", max)
    data.Out()
    data.Pretty()
    max = data.DelMax()
    t.Logf("取出最大元素：%v，观察..", max)
    data.Out()
    data.Pretty()
    max = data.DelMax()
    t.Logf("取出最大元素：%v，观察..", max)
    data.Out()
    data.Pretty()
    t.Logf("=====>>>测试结束<<<=====")
}
