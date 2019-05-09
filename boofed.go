package main

// https://github.com/faiface/beep/wiki/Hello,-Beep!

import (
  "github.com/faiface/beep/mp3"
  "github.com/faiface/beep/speaker"
  "github.com/faiface/beep"

  "log"
  "os"
  "time"
  "fmt"
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

  // choose one of the songs in the directory
  f, err := os.Open(cold)
  if err != nil {
    log.Fatal(err)
  }

  // our streamer is a beep.StreamSeeker (Len() int, Position() int, Seek(p int) error)
  streamer, format, err := mp3.Decode(f)
  if err != nil {
    log.Fatal(err)
  }
  defer streamer.Close()

  sr := format.SampleRate
  speaker.Init(sr, sr.N(time.Second/10))

  // resampled := beep.Resample(4, format.SampleRate, sr, streamer)

  loop := beep.Loop(3, streamer)
  fast := beep.ResampleRatio(4, 5, loop)

  done := make(chan bool)
  // switch streamer out for resampled yo
  speaker.Play(beep.Seq(fast, beep.Callback(func() {
    done <- true
    })))

    for {
      select {
      case <- done:
        return
      case <- time.After(time.Second):
        speaker.Lock()
        fmt.Println(format.SampleRate.D(streamer.Position()).Round(time.Second))
        speaker.Unlock()
      }
    }

}
