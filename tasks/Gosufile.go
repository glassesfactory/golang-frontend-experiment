package main

import (
	"log"
	"os"

	. "github.com/mgutz/gosu"
	"github.com/yosssi/gcss"
)

// Tasks タスク一覧
func Tasks(p *Project) {

	// CSS のコンパイル
	p.Task("css", "compile css...", W{"**/*.gcss"}, func(c *Context) {

		f, err := os.Open(c.Task.WatchFiles[0].Path)

		if err != nil {
			panic(err)
		}

		defer func() {
			if err := f.Close(); err != nil {
				panic(err)
			}
		}()
		fout, err := os.Create("src/assets/stylesheets/main.css")
		if err != nil {
			panic(err)
		}
		defer fout.Close()

		n, err := gcss.Compile(fout, f)
		if err != nil {
			log.Print("some err.....", err)
		}
		log.Print(n)

	})

	// サーバーの起動
	p.Task("server", D{"css"}, W{"**/*.go"}, W{"**/*.ace"}, func() {
		Start("main.go", M{"Dir": "src"})
	})
}

func main() {
	Gosu(Tasks)
}
