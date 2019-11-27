package dailyrotate

import (
	"time"
)

type option func(*File)

// WithOnOpen is an optional function that will be called every time a new file
// is opened
// If onOpen() takes a long time, you should do it in a background goroutine
// (it blocks all other operations, including writes)
func WithOnOpen(onOpen func(path string)) option {
	return func(f *File) {
		f.onOpen = onOpen
	}
}

// WithOnClose is an optional function that will be called every time existing file
// is closed, either as a result calling Close or due to being rotated.
// didRotate will be true if it was closed due to rotation.
// If onClose() takes a long time, you should do it in a background goroutine
// (it blocks all other operations, including writes)
func WithOnClose(onClose func(path string, didRotate bool)) option {
	return func(f *File) {
		f.onClose = onClose
	}
}

// WithBeforeClose is an optional function that will be called before every time
// existing file is closed, either as a result calling Close or due to being rotated.
// willRotate will be true if it was closed due to rotation.
// If beforeClose() takes a long time, you should do it in a background goroutine
// (it blocks all other operations, including writes)
func WithBeforeClose(beforeClose func(path string, willRotate bool)) option {
	return func(f *File) {
		f.beforeClose = beforeClose
	}
}

// WithPathGenerator is a function that will return a path for a daily log file.
// It should be unique in a given day e.g. time.Format of "2006-01-02.txt"
// creates a string unique for the day.
// Warning: time.Format might format more than you expect e.g.
// time.Now().Format(`/logs/dir-2/2006-01-02.txt`) will change "-2" in "dir-2" to
// current day. For better control over path generation, use NewFileWithPathGenerator
func WithPathGenerator(pathGenerator func(time.Time) string) option {
	return func(f *File) {
		f.pathGenerator = pathGenerator
	}
}

// WithPathFormat is file format accepted by time.Format that will be used to generate
// a name of the file. It should be unique in a given day e.g. 2006-01-02.txt.
// If you need more flexibility, use NewFileWithPathGenerator which accepts a
// function that generates a file path.
// Warning: time.Format might format more than you expect e.g.
// time.Now().Format(`/logs/dir-2/2006-01-02.txt`) will change "-2" in "dir-2" to
// current day. For better control over path generation, use NewFileWithPathGenerator
func WithPathFormat(pathFormat string) option {
	return func(f *File) {
		f.pathFormat = pathFormat
	}
}

// WithLocation lets you pass a location in to allow rotating in your favourite
// timezone.
func WithLocation(location *time.Location) option {
	return func(f *File) {
		f.Location = location
	}
}
