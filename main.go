package main

import (
  // "encoding/json"
  "flag"
  "log"
  // "net/http"
  // "os"
  // "os/exec"
  // "runtime"
  "strconv"
  "time"
  "fmt"
  "./lib/xethru"

  // "github.com/gorilla/websocket"
  "github.com/jacobsa/go-serial/serial"
)

var resetflag bool

func main() {
  fmt.Println("X2M200 Web Demo")

  commPort := flag.String("com", "/dev/ttyACM0", "the comm port you wish to use")
  baudrate := flag.Uint("baud", 115200, "the baud rate for the comm port you wish to use")
  sensitivity := flag.Uint("sensitivity", 7, "the sensitivity")
  start := flag.Float64("start", 0.5, "start of detection zone")
  end := flag.Float64("end", 2.1, "end of detection zone")
  // listen := flag.String("listen", "127.0.0.1:23000", "host:port to start webserver")
  reset := flag.Bool("reset", false, "try to reset the sensor")
  // format := flag.String("format", "json", "format for the log files valid choices are csv and json")
  flag.Parse()
  resetflag = *reset

  time.Sleep(time.Second * 1)
  baseband := make(chan xethru.BaseBandAmpPhase)
  resp := make(chan xethru.Respiration)
  sleep := make(chan xethru.Sleep)
  // url := "http://" + *listen
  openXethru(*commPort, *baudrate, *sensitivity, *start, *end, baseband, resp, sleep)
}


func openXethru(comm string, baudrate uint, sensivity uint, start float64, end float64, baseband chan xethru.BaseBandAmpPhase, resp chan xethru.Respiration, sleep chan xethru.Sleep) {

  fmt.Println("baba")
  time.Sleep(time.Second * 1)

  if resetflag {
    err := resetSensor(comm, baudrate)
    if err != nil {
      log.Panic(err)
    }
  }


  count := 5
  for {
    select {
    case <-time.After(time.Second):
      count--
      log.Println("Waiting for sensor " + strconv.Itoa(count))
    }
    if count <= 0 {
      break
    }
  }

  options := serial.OpenOptions{
    PortName:        comm,
    BaudRate:        baudrate,
    DataBits:        8,
    StopBits:        1,
    MinimumReadSize: 4,
  }

  port, err := serial.Open(options)

  // c := &serial.Config{Name: comm, Baud: int(baudrate)}
  // port, err := serial.OpenPort(c)
  if err != nil {
    log.Fatalf("serial.Open: %v", err)
  }

  x2 := xethru.Open("x2m200", port)
  defer x2.Close()

  m := xethru.NewModule(x2, "sleep")

  log.Printf("%#+v\n", m)
  err = m.Load()
  if err != nil {
    log.Panicln(err)
  }

  log.Println("Setting LED MODE")
  m.LEDMode = xethru.LEDInhalation
  err = m.SetLEDMode()
  if err != nil {
    log.Panicln(err)
  }

  log.Println("SetDetectionZone")
  err = m.SetDetectionZone(start, end)
  if err != nil {
    log.Panicln(err)
  }

  log.Println("SetSensitivity")
  err = m.SetSensitivity(int(sensivity))
  if err != nil {
    log.Panicln(err)
  }

  err = m.Enable("phase")
  if err != nil {
    log.Panicln(err)
  }

  stream := make(chan interface{})

  // log.Println("Opening browser to: ", url)
  // open(url)

  go m.Run(stream)

  for {
    select {
    case s := <-stream:
      switch s.(type) {
      case xethru.Respiration:
        resp <- s.(xethru.Respiration)
      case xethru.BaseBandAmpPhase:
        baseband <- s.(xethru.BaseBandAmpPhase)
      case xethru.Sleep:
        sleep <- s.(xethru.Sleep)
      default:
        log.Printf("%#v", s)
      }

    }
  }
}

func resetSensor(comm string, baudrate uint) error {
  // c := &serial.Config{Name: comm, Baud: int(baudrate)}
  options := serial.OpenOptions{
    PortName:        comm,
    BaudRate:        baudrate,
    DataBits:        8,
    StopBits:        1,
    MinimumReadSize: 4,
  }

  port, err := serial.Open(options)
  if err != nil {
    log.Printf("serial.Open: %v\n", err)
  }
  // port.Flush()

  x2 := xethru.Open("x2m200", port)
  // defer port.Close()
  defer x2.Close()

  reset, err := x2.Reset()
  if err != nil {
    log.Printf("serial.Reset: %v\n", err)
    return err
  }
  if !reset {
    log.Fatal("Could not reset")
  }
  return nil
}