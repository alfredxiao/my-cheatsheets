package main

import (
    "fmt"
    "os"
    "strconv"
    "time"
    "os/exec"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func toInt(s string) int {
  n, e := strconv.Atoi(s)
  check(e)
  return n
}

func main() {
  switch os.Args[1] {
	case "parent":
		parent()
	case "child":
		child()
	default:
		panic("help")
	}
}

func parent() {
    count := toInt(os.Args[2])
    sleep := toInt(os.Args[3])

    cmd := exec.Command("/proc/self/exe", append([]string{"child"}, os.Args[2:]...)...)
  	cmd.Stdin = os.Stdin
  	cmd.Stdout = os.Stdout
  	cmd.Stderr = os.Stderr
    cmd.Start()

    openAndCloseFiles("parent", 1, count, sleep)
}

func child() {
  count := toInt(os.Args[2])
  sleep := toInt(os.Args[3])
  openAndCloseFiles("child", count + 1, count * 2, sleep)
}

func openAndCloseFiles(who string, start, end, sleep int) {
  var files []*os.File

  for i := start; i <= end; i++ {
      f := read(i)
      files = append(files, f)
  }

  fmt.Printf("%s - %d files opened\n", who, (end - start + 1))
  fmt.Printf("%s - sleep for %d seconds\n", who, sleep)
  time.Sleep(time.Duration(sleep) * time.Second)

  for _, file := range(files) {
    err := file.Close()
    check(err)
  }

  fmt.Printf("%s - %d files closed\n", who, len(files))
}

func read(n int) *os.File {
    f, err := os.Open(fmt.Sprintf("./data/%d", n))
    check(err)
    // fmt.Printf("file %d opened, ", n)

    return f
}
