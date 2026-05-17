package code

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func GetPathSize(path string, recursive, human, all bool) (string, error) {
	size, err := getSize(path, recursive, all)
	if err != nil {
		return "", err
	}

	formatted := formatSize(size, human)
	return fmt.Sprintf("%s\t%s\n", formatted, path), nil
}

func getSize(path string, recursive, showAll bool) (int64, error) {
	fi, err := os.Lstat(path)
	if err != nil {
		return 0, err
	}

	if !fi.IsDir() {
		if fi.Mode()&os.ModeSymlink != 0 {
			target, err := os.Readlink(path)
			if err != nil {
				return 0, err
			}
			if !filepath.IsAbs(target) {
				target = filepath.Join(filepath.Dir(path), target)
			}
			return getSize(target, recursive, showAll)
		}
		return fi.Size(), nil
	}

	files, err := os.ReadDir(path)
	if err != nil {
		return 0, err
	}

	var total int64

	for _, file := range files {
		if !showAll && strings.HasPrefix(file.Name(), ".") {
			continue
		}

		fullPath := filepath.Join(path, file.Name())

		if !file.IsDir() {
			newfi, err := os.Lstat(fullPath)
			if err != nil {
				return 0, err
			}

			if newfi.Mode()&os.ModeSymlink != 0 {
				target, err := os.Readlink(fullPath)
				if err != nil {
					return 0, err
				}
				if !filepath.IsAbs(target) {
					target = filepath.Join(filepath.Dir(fullPath), target)
				}

				targetFi, err := os.Stat(target)
				if err != nil {
					return 0, err
				}

				if !targetFi.IsDir() {
					total += targetFi.Size()
				}
			} else {
				total += newfi.Size()
			}
		} else if recursive {
			subSize, err := getSize(fullPath, recursive, showAll)
			if err != nil {
				return 0, err
			}
			total += subSize
		}
	}

	return total, nil
}

func formatSize(size int64, human bool) string {
	if !human {
		return strconv.FormatInt(size, 10) + "B"
	}

	const (
		KB = 1024
		MB = KB * 1024
		GB = MB * 1024
		TB = GB * 1024
	)

	switch {
	case size >= TB:
		return fmt.Sprintf("%.1fTB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.1fGB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.1fMB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.1fKB", float64(size)/KB)
	default:
		return strconv.FormatInt(size, 10) + "B"
	}
}
