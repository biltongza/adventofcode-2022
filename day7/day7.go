package day7

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Day7() {
	file, err := os.Open("day7/input.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)

	var currentCmd string
	var cwd *Directory
	var root *Directory
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		if line[0:1] == "$" {
			// command mode
			parts := strings.Split(line, " ")
			cmd := parts[1]
			currentCmd = cmd
			args := parts[2:]
			switch cmd {
			case "cd":
				dest := args[0]
				if cwd == nil {
					cwd = makeDir(dest, nil)
					cwd.parent = cwd
				}
				if dest == "/" && root == nil {
					root = cwd
					continue
				}
				if dest == ".." {
					cwd = cwd.parent
					continue
				} else {
					exists := false
					for _, dir := range cwd.dirs {
						if dir.name == dest {
							cwd = dir
							exists = true
							break
						}
					}
					if !exists {
						cwd = makeDir(dest, cwd)
					}
				}
			}
		} else {
			switch currentCmd {
			case "ls":
				parts := strings.Split(line, " ")
				if parts[0] == "dir" {
					dir := makeDir(parts[1], cwd)
					cwd.dirs = append(cwd.dirs, dir)
				} else {
					size, _ := strconv.Atoi(parts[0])
					name := parts[1]
					file := &File{
						name: name,
						size: size,
						dir:  cwd,
					}
					cwd.files = append(cwd.files, file)
				}
			}
		}
	}
	totalUsed := getSize(root)
	totalFs := 70000000
	totalFree := totalFs - totalUsed
	totalRequired := 30000000 - totalFree

	smallestLargeDir := findSmallestLargeDir(root, totalRequired)
	smallestLargeDirSize := getSize(smallestLargeDir)

	fmt.Fprintf(os.Stdout, "smallestLargeDirSize = : %d\n", smallestLargeDirSize)
}

func makeDir(name string, parent *Directory) *Directory {
	dir := &Directory{
		name:   name,
		parent: parent,
		dirs:   make([]*Directory, 0),
		files:  make([]*File, 0),
	}

	return dir
}

func findSmallestLargeDir(dir *Directory, requiredSize int) *Directory {
	var smallestLargeDir *Directory = dir
	smallestLargeDirSize := getSize(smallestLargeDir)
	for _, subdir := range dir.dirs {
		subdirSize := getSize(subdir)
		if subdirSize < requiredSize {
			continue
		}
		if subdirSize >= requiredSize && subdirSize < smallestLargeDirSize {
			smallestLargeDir = subdir
			smallestLargeDirSize = subdirSize
		}
		nextSmallest := findSmallestLargeDir(subdir, requiredSize)
		nextSmallestSize := getSize(nextSmallest)
		if nextSmallestSize >= requiredSize && nextSmallestSize <= smallestLargeDirSize {
			smallestLargeDir = nextSmallest
			smallestLargeDirSize = nextSmallestSize
		}
	}

	return smallestLargeDir
}

func getSize(dir *Directory) int {
	total := 0
	for _, file := range dir.files {
		total += file.size
	}
	for _, subdir := range dir.dirs {
		total += getSize(subdir)
	}
	return total
}

type File struct {
	name string
	size int
	dir  *Directory
}

type Directory struct {
	name   string
	files  []*File
	parent *Directory
	dirs   []*Directory
}
