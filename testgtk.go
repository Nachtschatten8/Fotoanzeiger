package main

import (
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"

	"github.com/gotk3/gotk3/gdk"

	"fmt"

	"github.com/gotk3/gotk3/gtk"
)

func main() {
	// Initialize GTK without parsing any command line arguments.
	gtk.Init(nil)

	// Create a new toplevel window, set its title, and connect it to the
	// "destroy" signal to exit the GTK main loop when it is destroyed.
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		log.Fatal("Unable to create window:", err)
	}
	win.SetTitle("Simple Example")
	win.Connect("destroy", func() {
		gtk.MainQuit()
	})

	mainBox, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 6)

	// Create a new label widget to show in the window.
	l, err := gtk.LabelNew("Hallo Adrian")
	if err != nil {
		log.Fatal("Unable to create label:", err)
	}

	// Add the label to the window.
	mainBox.Add(l)

	Dateinamen := make([]string, 0)
	err = filepath.Walk("/home/abarth/Pictures",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			//fmt.Println(path, info.Size())
			Dateinamen = append(Dateinamen, path)
			return nil
		})
	if err != nil {
		log.Println(err)
	}

	f, err := gtk.FrameNew("ll")
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator

	filename := Dateinamen[rand.Intn(len(Dateinamen))]
	//filename := "/home/abarth/Pictures/wallpaper/132_3222.JPG"
	//imageOK, err := gdk.PixbufNewFromFile(filename)
	img, err := gtk.ImageNewFromFile(filename)

	pb := img.GetPixbuf()
	fmt.Printf("%d\n", pb.GetWidth())
	fmt.Printf("%d\n", pb.GetHeight())
	width := pb.GetWidth()
	height := pb.GetHeight()
	var height2 int
	var width2 int

	if width > height {
		width2 = 500
		height2 = (height * 500) / width
	} else {
		height2 = 500
		width2 = (width * 500) / height
	}

	pb2, err := pb.ScaleSimple(width2, height2, gdk.INTERP_BILINEAR)
	fmt.Printf("%T\n", pb)

	img2, err := gtk.ImageNewFromPixbuf(pb2)

	f.Add(img2)
	mainBox.Add(f)
	win.Add(mainBox)
	//pixbuf = gtk.

	// Set the default window size.
	win.SetDefaultSize(800, 600)

	// Recursively show all widgets contained in this window.
	win.ShowAll()

	// Begin executing the GTK main loop.  This blocks until
	// gtk.MainQuit() is run.
	gtk.Main()
}
