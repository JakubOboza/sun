package relay_control

import (
  "github.com/stianeikeland/go-rpio"
)

type Relay struct {
  pin rpio.Pin
}

func Open() error {
  return rpio.Open()
}

func Close() {
  rpio.Close()
}

func New(pin int) *Relay {
  rc :=  &Relay{pin: rpio.Pin(pin)}
  rc.pin.Output() // Set pin mode
  return rc
}

func (rc *Relay) TurnOn() {
  rc.pin.High()
}

func (rc *Relay) TurnOff() {
  rc.pin.Low()
}


