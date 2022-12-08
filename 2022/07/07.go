package aoc202207

import (
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const (
	RegexpFile = `(?P<Size>\d+)\s(?P<Name>.*)`
)

type Directory struct {
	Name           string
	Subdirectories []*Directory
	Files          []*File
	Parent         *Directory
}

func (d Directory) Size() (total int) {
	for _, file := range d.Files {
		total += file.Size
	}

	for _, dir := range d.Subdirectories {
		total += dir.Size()
	}

	return
}

func (d Directory) Flatten() (dirs []*Directory) {
	dirs = d.Subdirectories
	for _, dir := range d.Subdirectories {
		dirs = append(dirs, dir.Flatten()...)
	}

	return
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{
		Name:           name,
		Subdirectories: []*Directory{},
		Files:          []*File{},
		Parent:         parent,
	}
}

func (d *Directory) FindDirectory(name string) (dir *Directory) {
	if d.Name == "/" && name == "/" {
		return d
	}

	if name == ".." {
		return d.Parent
	}

	for _, dir = range d.Subdirectories {
		if dir.Name == name {
			return
		}
	}

	return nil
}

type File struct {
	Name string
	Size int
}

func Walk(commands []string) (d Directory) {
	fileExp := regexp.MustCompile(RegexpFile)

	d = *NewDirectory("/", nil)

	cd := &d

	for _, command := range commands {
		if strings.TrimSpace(command) == "$ ls" || command == "" {
			continue
		}

		if strings.HasPrefix(command, "$ cd ") {
			cd = cd.FindDirectory(command[5:])
			continue
		}

		if strings.HasPrefix(command, "dir ") {
			cd.Subdirectories = append(cd.Subdirectories, NewDirectory(command[4:], cd))
			continue
		}

		matches := fileExp.FindAllStringSubmatch(command, -1)

		size, _ := strconv.Atoi(matches[0][1])
		cd.Files = append(cd.Files, &File{Name: matches[0][2], Size: size})
	}

	return
}

func Part1(input string) (output string) {
	commands := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	d := Walk(commands)
	dirs := d.Flatten()

	total := 0
	for _, d := range dirs {
		if size := d.Size(); size <= 100000 {
			total += size
		}
	}

	return strconv.Itoa(total)
}

func Part2(input string) (output string) {
	commands := strings.Split(strings.ReplaceAll(input, "\r", ""), "\n")
	d := Walk(commands)
	dirs := d.Flatten()

	sort.SliceStable(dirs, func(i, j int) bool {
		return dirs[i].Size() < dirs[j].Size()
	})

	unused := 70000000 - d.Size()
	needed := 30000000 - unused

	for _, d := range dirs {
		if size := d.Size(); size >= needed {
			return strconv.Itoa(size)
		}
	}

	return
}
