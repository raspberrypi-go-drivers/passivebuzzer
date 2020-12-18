package passivebuzzer_test

import (
	"os"
	"time"

	"github.com/raspberrypi-go-drivers/passivebuzzer"
	"github.com/stianeikeland/go-rpio/v4"
)

func Example() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	buzzer := passivebuzzer.NewPassiveBuzzer(18)
	buzzer.Tone(440) // 440Hz
	time.Sleep(3 * time.Second)
	buzzer.StopTone()
}

func ExampleNewPassiveBuzzer() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	buzzer := passivebuzzer.NewPassiveBuzzer(18)
	buzzer.Tone(440) // 440Hz
	time.Sleep(3 * time.Second)
	buzzer.StopTone()
}

func ExamplePassiveBuzzer_Tone() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	buzzer := passivebuzzer.NewPassiveBuzzer(18)
	buzzer.Tone(440) // 440Hz
	time.Sleep(3 * time.Second)
	buzzer.StopTone()
}

func ExamplePassiveBuzzer_StopTone() {
	err := rpio.Open()
	if err != nil {
		os.Exit(1)
	}
	defer rpio.Close()
	buzzer := passivebuzzer.NewPassiveBuzzer(18)
	buzzer.Tone(440) // 440Hz
	time.Sleep(3 * time.Second)
	buzzer.StopTone()
}
