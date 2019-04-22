package main

// so this stuff is using
//   https://medium.com/@valentijnnieman_79984/how-to-build-an-audio-streaming-
//   server-in-go-part-1-1676eed93021
// as a resource

// we use https://github.com/gordonklaus/portaudio :)

import (
  "bytes"
  "encoding/binary"
  "fmt"
  "github.com/gordonklaus/portaudio"
  "io/ioutil"
  "net/http"
  "time"
)

const sampleRate = 44100
const seconds = 2

func main() {
  fmt.Println("Hello, guy but gender neutral")
  portaudio.Initialize()
  defer portaudio.Terminate()
  buffer := make([]float32, sampleRate * seconds)

  stream, err := portaudio.OpenDefaultStream(
    0, 1, sampleRate, len(buffer), func(out []float32) {
      resp, err := http.Get("http://localhost:8080/audio")
      chk(err)
      body, _ := ioutil.ReadAll(resp.Body)
      responseReader := bytes.NewReader(body)
      binary.Read(responseReader, binary.BigEndian, &buffer)
      for i:= range out {
        out[i] = buffer[i]
      }
    }) // end the OpenDefaultStream thing
    chk(err)
    chk(stream.Start())
    time.Sleep(time.Second * 40)
    chk(stream.Stop())
    defer stream.Close()
}

func chk(err error) {
  if err != nil {
    panic(err)
  }
}
