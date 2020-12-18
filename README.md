# Passive Buzzer

[![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/passivebuzzer.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/passivebuzzer)
![golangci-lint](https://github.com/raspberrypi-go-drivers/passivebuzzer/workflows/golangci-lint/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/raspberrypi-go-drivers/passivebuzzer)](https://goreportcard.com/report/github.com/raspberrypi-go-drivers/passivebuzzer)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Package passivebuzzer is a driver allowing to control a passive buzzer from a GPIO pin

## Documentation

For full documentation, please visit [![Go Reference](https://pkg.go.dev/badge/github.com/raspberrypi-go-drivers/passivebuzzer.svg)](https://pkg.go.dev/github.com/raspberrypi-go-drivers/passivebuzzer)

## Quick start

```go
import (
	"math"
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/passivebuzzer"
	"github.com/stianeikeland/go-rpio/v4"
)

const (
	c4  float64 = 261.63 // Hz
	d4  float64 = 293.66 // Hz
	e4  float64 = 329.63 // Hz
	f4  float64 = 349.23 // Hz
	fd4 float64 = 369.99 // Hz
	g4  float64 = 392.00 // Hz
	a4  float64 = 440.00 // Hz
	b4  float64 = 493.88 // Hz
	c5  float64 = 523.25 // Hz
	d5  float64 = 587.33 // Hz
	e5  float64 = 659.25 // Hz
	f5  float64 = 698.46 // Hz
	fd5 float64 = 739.99 // Hz
	g5  float64 = 783.99 // Hz
	// https://pages.mtu.edu/~suits/notefreqs.html
)

type note struct {
	note     float64
	duration float64
}

func main() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	buzzer := passivebuzzer.NewPassiveBuzzer(18)
	melody := []note{
		{e4, 1}, {g4, 1}, {b4, 1}, {e5, 1}, {b4, 1},
		{g4, 1}, {e4, 1}, {g4, 1}, {b4, 1}, {g5, 1},
		{g4, 0.5}, {e5, 0.5}, {b4, 0.5}, {g4, 0.5},
		{fd5, 1}, {d4, 0.5}, {g5, 0.5}, {d5, 0.5},
		{g4, 0.5}, {g5, 1}, {d5, 0.5}, {g4, 0.5},
		{g4, 0.5}, {fd4, 0.5}, {e4, 2},
	}
	for _, n := range melody {
		buzzer.Tone(n.note)
		multiplier := int(math.RoundToEven(400.0 * n.duration))
		time.Sleep(time.Duration(multiplier) * time.Millisecond)
		buzzer.StopTone()
		time.Sleep(33 * time.Millisecond)
	}
}
```

## Raspberry Pi compatibility

This driver has has only been tested on an Raspberry Pi Zero WH using integrated bluetooth but may work well on other Raspberry Pi having integrated Bluetooth

## License

MIT License

---

Special thanks to @stianeikeland

This driver is based on his work in [stianeikeland/go-rpio](https://github.com/stianeikeland/go-rpio/)
