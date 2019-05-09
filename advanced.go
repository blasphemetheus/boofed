package main

// official gui? https://golangr.com/gui/


// interesting gui option, complex? C:
//  https://github.com/golang-ui/nuklear
// resume reading pls:
//  https://github.com/faiface/beep/wiki/To-buffer,-or-not-to-buffer,-that-is-the-question

// options for gui:
//     https://github.com/therecipe/qt
//    yo https://awesome-go.com/#gui
//

// list of stuff:
//    https://github.com/avelino/awesome-go
//

// concise, clear: https://medium.com/from-the-couch/building-a-desktop-ui-in-go-1b6474afa736

// more reading: https://medium.com/jexia/go-1-11-scheduling-and-gui-with-golang-more-ce41ffdcca1d

// oo interesting medium post, tutorial using electron:
// https://medium.com/benchkram/tutorial-adding-a-gui-to-golang-6aca601e277d
import (
  "log"
  "os"
  "time"
  "fmt"

  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
  "github.com/faiface/beep"
  "github.com/faiface/beep/effects"

)

func main() {

  const (
    alex = "Alex.mp3"
    changes = "Mutemath - Changes.mp3"
    god = "33 GOD.mp3"
    creeks = "715 - CRKS.mp3"
    holocene = "Bon Iver - Holocene (Instrumental).mp3"
    cold = "Exitmusic - The Cold (Original Edit).mp3"
    limbo = "Jimmy Cliff - Sitting In Limbo.mp3"
    madworld = "Mad World - Gary Jules.mp3"
    anothernewworld = "Punch Brothers - Another New World.mp3"
    myohmy = "My Oh My (Punch Brothers).mp3"
    other_myohmy = "Punch Brothers - My Oh My.mp3"
    simizzoke = "Odell Brown - Simizzoke.mp3"
    rainbow = "Somewhere over the Rainbow.mp3"
    smack = "The Prodigy - Smack My Bitch Up.mp3"
    wine = "UB40 - Red Red Wine.mp3"
    blue = "Willow Beats - Blue.mp3"
    dances = "Wilson Pickett - Land of 1000 Dances.mp3"
    wind = "Wind Beneath My Wings - Bette Middler.mp3"
    come_closer = "WizKid - Come Closer.mp3"
    bed_rock = "Young Money - Bed Rock.mp3"
  )

    f, err := os.Open(come_closer)
    if err != nil {
      log.Fatal(err)
    }

    streamer, format, err := mp3.Decode(f)
    if err != nil {
      log.Fatal(err)
    }
    defer streamer.Close()

    speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

    ctrl := &beep.Ctrl{
      Streamer: beep.Loop(-1, streamer),
      Paused: false}
    volume := &effects.Volume{
      Streamer: ctrl,
      Base: 2,
      Volume: 0,
      Silent: false,
    }
    speedy := beep.ResampleRatio(4, 1, volume)
    speaker.Play(speedy)


    for {
      fmt.Print("Press [ENTER] to pause/resume. ")
      fmt.Scanln()

      speaker.Lock()
      ctrl.Paused = !ctrl.Paused
      volume.Volume += 0.5 //// change
      speedy.SetRatio(speedy.Ratio() + 0.1) /// change
      speaker.Unlock()
    }


    // Then we delve into Volume

    /*
    type Volume struct {
	   Streamer beep.Streamer
	    Base     float64
	     Volume   float64
	      Silent   bool
      }
    */

}
