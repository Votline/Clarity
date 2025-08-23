package core

import (
	"os"
	"fmt"
	"log"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
	"github.com/go-gl/glfw/v3.3/glfw"
)

var initialized bool

func renderView(timerChan chan string) {
	totalMs := 20 * 1000
	for ms := totalMs; ms > 0; ms-=100{
		sec := ms/1000
		mil := (ms%1000)/10
		timerChan <- fmt.Sprintf("%02d.%02d", sec, mil/10)
		time.Sleep(100*time.Millisecond)
	}
	timerChan <- "00.00"
}

func Timer(win *glfw.Window, timerChan chan string, done *bool) {
	for {
		time.Sleep(1*time.Second)
		win.Show()
		*done = true
		playMP3()

		renderView(timerChan)
		playMP3()
		win.Hide()
	}
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
