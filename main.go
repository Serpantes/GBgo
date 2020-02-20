package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type vehilce struct {
	model                     string
	volume, volumeOccup, year int
	isRunning, isWindowOpen   bool
}
type car struct {
	model        string
	volume       int
	volumeOccup  int
	year         int
	isRunning    bool
	isWindowOpen bool
}

type queue struct {
	nodes []*node
	size  int
	front int
	back  int
	count int
}

func newQueue(size int) *queue {
	return &queue{
		nodes: make([]*node, size),
		size:  size,
	}
}

func (q *queue) push(n *node) {
	if q.front == q.back && q.count > 0 {
		nodes := make([]*node, len(q.nodes)+q.size)
		copy(nodes, q.nodes[q.front:])
		copy(nodes[len(q.nodes)-q.front:], q.nodes[:q.front])
		q.front = 0
		q.back = len(q.nodes)
		q.nodes = nodes
	}
	q.nodes[q.back] = n
	q.back = (q.back + 1) % len(q.nodes)
	q.count++
}
func (q *queue) pop() *node {
	if q.count == 0 {
		return nil
	}
	node := q.nodes[q.front]
	q.front = (q.front + 1) % len(q.nodes)
	q.count--
	return node
}

type phones []int
type node struct {
	Value int
}

func (p phones) ViewList() {
	for i, phone := range p {
		fmt.Printf("\t %v) %v \n", i+1, phone)
	}
}
func viewAddressBook(addressBook map[string]phones) {
	for name, p := range addressBook {
		fmt.Println(name)
		p.ViewList()
	}
}
func saveAddressBook(addressBook map[string]phones) {
	file, _ := json.MarshalIndent(addressBook, "", "")

	_ = ioutil.WriteFile("testoutput.json", file, 0644)
}

// PercentOf : [part] is X precent of [total]. return X
func PercentOf(part int, total int) float64 {
	return (float64(part) * float64(100)) / float64(total)
}

func vehInfo(obj vehilce) string {
	x := PercentOf(obj.volumeOccup, obj.volume)
	if x == 0 {
		return "пустой"
	}
	if x < 20 {
		return "Почти пустой"
	}
	if x >= 20 && x <= 40 {
		return "Небольная заполненость"
	}
	if x > 40 && x < 60 {
		return "Половина заполнена"
	}
	if x >= 60 && x < 100 {
		return "Почти полный"
	}
	if x == 100 {
		return "Полный"
	}
	return "obj.volumeOccup, obj.volume error"
}
func carInfo(obj car) string {
	if obj.isRunning {
		return "В пути"
	}
	return "Не в пути"

}
func printVehInfo(x vehilce) {
	fmt.Println(x.model, vehInfo(x))
}
func printCarInfo(x car) {
	fmt.Println(x.model, carInfo(x))
}
func main() {
	machina := vehilce{
		model:        "gruzovik",
		volume:       100,
		volumeOccup:  82,
		year:         2018,
		isRunning:    true,
		isWindowOpen: false,
	}
	machinka := vehilce{
		model:        "gruzovichok",
		volume:       50,
		volumeOccup:  11,
		year:         1999,
		isRunning:    false,
		isWindowOpen: true,
	}
	tachka := car{
		model:        "cabrio",
		volume:       15,
		volumeOccup:  14,
		year:         2020,
		isRunning:    true,
		isWindowOpen: true,
	}
	avtomobil := car{
		model:        "lada",
		volume:       10,
		volumeOccup:  2,
		year:         2007,
		isRunning:    false,
		isWindowOpen: true,
	}

	fmt.Println()

	printVehInfo(machina)
	printVehInfo(machinka)
	printCarInfo(tachka)
	printCarInfo(avtomobil)

	fmt.Println()

	queue := newQueue(1)
	queue.push(&node{1})
	queue.push(&node{22})
	queue.push(&node{333})
	fmt.Println(queue.pop(), queue.pop(), queue.pop())

	fmt.Println()

	addressBook := make(map[string]phones)

	addressBook["Alex"] = phones{89996543210}
	addressBook["Alex"] = append(addressBook["Alex"], 89993157365)
	addressBook["Alex"] = append(addressBook["Alex"], 89999983121)
	addressBook["Bob"] = phones{89167243812}
	addressBook["Bob"] = append(addressBook["Bob"], 89161437244)
	addressBook["Bob"] = append(addressBook["Bob"], 89163255424)
	addressBook["Rob"] = phones{89264323453}
	addressBook["Rob"] = append(addressBook["Rob"], 89260889352)
	addressBook["Rob"] = append(addressBook["Rob"], 89265841488)

	viewAddressBook(addressBook)

	saveAddressBook(addressBook)
}
