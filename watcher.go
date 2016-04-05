package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/fsnotify/fsnotify"
)

type ChangeChan chan bool

type Watcher struct {
	watcher *fsnotify.Watcher
}

func (w *Watcher) startWatcher() ChangeChan {
	changeChan := make(ChangeChan)
	var err error
	w.watcher, err = fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event := <-w.watcher.Events:
				if (event.Op&fsnotify.Write == fsnotify.Write) || (event.Op&fsnotify.Rename == fsnotify.Rename) {
					changeChan <- true
				}
			case err := <-w.watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	return changeChan
}

func (w *Watcher) closeChan() {
	w.watcher.Close()
}

func (w *Watcher) addWatchable(path string, watcheable os.FileInfo) {
	if watcheable.IsDir() {
		visit := func(fpath string, f os.FileInfo, err error) error {
			if !f.IsDir() {
				err = w.watcher.Add(fpath)
				if err != nil {
					log.Fatal(err)
				}
			}

			return nil
		}

		filepath.Walk(path, visit)
	} else {
		w.watcher.Add(path)
	}
}
