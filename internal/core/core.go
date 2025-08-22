package core

import (
	"os"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var initialized bool

func Timer(win *glfw.Window, done *bool) {
	time.Sleep(5*time.Second)
	win.Show()
	*done = true
	playMP3()
	time.Sleep(5*time.Second)
	playMP3()
	win.Hide()
}

func initSpeaker(format beep.Format) {
	if initialized {return}
	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	initialized = true
}

func playMP3() {
	f, err := os.Open("assets/beep.mp3")
	if err != nil {
		log.Printf("Open mp3 error: %v", err)
		return
	}
	streamer, format, err := mp3.Decode(f)
	if err != nil {
		log.Printf("Decode mp3 error: %v", err)
		return
	}
	defer streamer.Close()

	initSpeaker(format)
	done := make(chan bool)

	speaker.Play(beep.Seq(streamer, beep.Callback(func() {
		done <- true
	})))

	<-done
}
