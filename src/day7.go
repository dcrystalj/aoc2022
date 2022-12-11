package main

import (
	"fmt"
	"sort"
)

var total uint64

var dirSizes []uint64

type file struct {
	name string
	size uint64
}

type dir struct {
	dirs   map[string]*dir
	files  []file
	parent *dir
}

func (d dir) size() uint64 {
	s := uint64(0)
	for _, f := range d.files {
		s += f.size
	}

	for _, val := range d.dirs {
		s += val.size()
	}

	if s <= 100000 {
		total += s
	}
	dirSizes = append(dirSizes, s)
	return s
}

func NewDir(parent *dir) *dir {
	return &dir{
		dirs:   make(map[string]*dir),
		files:  []file{},
		parent: parent,
	}
}

func day7() {
	lines := ReadInput("day7.in")

	root := NewDir(nil)
	curr := root
	for _, line := range lines {
		curr = parseLine(line, curr)
	}
	rootSize := root.size()
	println(total)

	sort.Slice(dirSizes, func(i, j int) bool {
		return i < j
	})
	unused := 70000000 - rootSize
	for _, used := range dirSizes {
		if used+unused >= 30000000 {
			println(used)
			break
		}
	}
}

func parseLine(l string, c *dir) *dir {
	if l[0] == '$' {
		l = l[2:]
		if l[:2] == "ls" {
			return c
		} else if l[:3] == "cd " {
			if len(l[3:]) == 2 && l[3:5] == ".." {
				return c.parent
			}
			n := NewDir(c)
			c.dirs[l[3:]] = n
			return n
		} else { // ls
			panic("undefined")
		}
	} else {
		if l[:4] == "dir " {
			c.dirs[l[4:]] = NewDir(c)
		} else {
			f := file{}
			fmt.Sscanf(l, "%d %s", &f.size, &f.name)
			c.files = append(c.files, f)
		}
	}
	return c
}

func day7_2() {

}
