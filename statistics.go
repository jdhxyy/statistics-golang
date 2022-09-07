// Copyright 2022-2022 The jdh99 Authors. All rights reserved.
// 统计模块
// Authors: jdh99 <jdh821@163.com>

package statistics_golang

import (
	"fmt"
	"sync"
)

type tItem struct {
	name string
	num int
}

var items []tItem
var lock sync.RWMutex

// Register 注册统计.返回统计项的id
func Register(name string) int {
	lock.Lock()
	defer lock.Unlock()

	items = append(items, tItem{name, 0})
	return len(items) - 1
}

// Add 统计项自增
func Add(itemID int) {
	lock.Lock()
	defer lock.Unlock()

	if itemID >= len(items) {
		return
	}
	items[itemID].num++
}

// Clear 清除统计项
func Clear(itemID int) {
	lock.Lock()
	defer lock.Unlock()

	if itemID >= len(items) {
		return
	}
	items[itemID].num = 0
}

// Set 设置统计项
func Set(itemID int, value int) {
	lock.Lock()
	defer lock.Unlock()

	if itemID >= len(items) {
		return
	}
	items[itemID].num = value
}

// ClearAll 清除所有统计项
func ClearAll() {
	lock.Lock()
	defer lock.Unlock()

	for i := 0; i < len(items); i++ {
		items[i].num = 0
	}
}

// GetItem 读取某项统计
// 返回的如果是nil,代表读取失败
func GetItem(itemID int) *tItem {
	lock.RLock()
	defer lock.RUnlock()

	if itemID >= len(items) {
		return nil
	}

	item := items[itemID]
	return &item
}

// GetItemNum 读取统计项数
func GetItemNum() int {
	return len(items)
}

// Output 输出统计信息
func Output() string {
	lock.RLock()
	defer lock.RUnlock()

	var s string
	for i := 0; i < len(items); i++ {
		s += fmt.Sprintf("%s:%d\n", items[i].name, items[i].num)
	}
	return s
}

// Print 打印统计信息
func Print() {
	lock.RLock()
	defer lock.RUnlock()

	for i := 0; i < len(items); i++ {
		fmt.Printf("%s:%d\n", items[i].name, items[i].num)
	}
}
