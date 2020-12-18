package passivebuzzer

import (
	"math"

	"github.com/stianeikeland/go-rpio/v4"
)

const (
	pwmClockSpeed int = 1000000 // 1Mhz
)

// PassiveBuzzer represents a passive buzzer
type PassiveBuzzer struct {
	pin rpio.Pin
}

// NewPassiveBuzzer creates a new PassiveBuzzer instance
func NewPassiveBuzzer(pinID int) *PassiveBuzzer {
	buzzer := PassiveBuzzer{
		pin: rpio.Pin(uint8(pinID)),
	}
	buzzer.pin.Mode(rpio.Pwm)
	buzzer.pin.Freq(pwmClockSpeed)
	return &buzzer
}

// Tone activate the buzzer tone ate the specified sound frequency in Hz
func (buzzer *PassiveBuzzer) Tone(soundFrequency float64) {
	cycle := math.RoundToEven(float64(pwmClockSpeed) / soundFrequency)
	duty := math.RoundToEven((float64(pwmClockSpeed) / soundFrequency) / 2.0)
	buzzer.pin.DutyCycle(uint32(duty), uint32(cycle))
}

// StopTone stops the buzzer tone
func (buzzer *PassiveBuzzer) StopTone() {
	buzzer.pin.DutyCycle(0, 1)
}
