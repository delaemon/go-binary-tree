package main

import "fmt"

type Node struct {
	value  int
	parent *Node
	left   *Node
	right  *Node
}

func build(numbers []int) Node {
	pl := []Node{}
	l := len(numbers) - 3
	i := 0
	for i < l {
		p := Node{}
		l := Node{}
		r := Node{}
		l.parent = &p
		r.parent = &p
		p.left = &l
		p.right = &r
		l.value = numbers[i]
		p.value = numbers[i+1]
		r.value = numbers[i+2]
		pl = append(pl, p)
		i += 3
	}

	r := len(numbers) - i
	if r == 2 {
		p := Node{}
		l := Node{}
		l.parent = &p
		p.left = &l
		l.value = numbers[i]
		p.value = numbers[i+1]
		pl = append(pl, p)
	} else if r == 1 {
		p := Node{}
		p.value = numbers[i]
		pl = append(pl, p)
	}

	tree := []Node{}
	root := Node{}
	tmp := pl
	for {
		i := 0
		for i < len(tmp)-1 {
			l := tmp[i]
			r := tmp[i+1]
			last := &l
			p := l.right
			for {
				if p.right != nil {
					last = p
					p = p.right
				} else {
					last.right = nil
					p.parent = nil
					break
				}
			}
			l.parent = p
			r.parent = p
			p.left = &l
			p.right = &r
			tree = append(tree, *p)
			i += 2
		}
		if len(tmp) % 2 == 1 {
			l := tree[len(tree)-1]
			r := tmp[i]
			p := l.right
			last := &l
			tree = tree[:len(tree)-1]
			for {
				if p.right != nil {
					last = p
					p = p.right
				} else {
					last.right = nil
					p.parent = nil
					break
				}
			}
			p.left = &l
			p.right = &r
			l.parent = p
			r.parent = p
			tree = append(tree, *p)
		}
		if len(tree) == 1 {
			root = tree[0]
			break
		} else {
			tmp = tree
			tree = []Node{}
		}
	}

	return root
}

func test(root Node, length int) {
	for i := 0; i < length; i++ {
		target := i
		var hit Node
		n := root
		count := 0
		for {
			if target == n.value {
				hit = n
				break
			}
			if target < n.value && n.left != nil {
				n = *n.left
			} else if target > n.value && n.right != nil {
				n = *n.right
			} else {
				fmt.Printf("undefined. target: %d, node: %d", target, n.value)
				return
			}
			count++
		}
		fmt.Printf("hit count:%d, node: %v\n", count, hit)
	}
}

func dump(n Node) {
	fmt.Println(n)
	if n.left != nil {
		dump(*n.left)
	}
	if n.right != nil {
		dump(*n.right)
	}
	return
}

func main() {
	length := 1000000
	numbers := []int{}
	for i := 0; i < length; i++ {
		numbers = append(numbers, i)
	}
	root := build(numbers)
	test(root, len(numbers))
}
