package main

import (
	"log"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

//pwm pins: 13,19,12,18

//18 and 12 on one channel
//13 and 19 on the other channel

const escPinID = 12
const servoPinID = 13

// const panPinID = 14
// const tiltPinID = 15

const frequency = 100000
const maxvalue = uint32(1024)
const midvalue = uint32(512)
const minvalue = uint32(0)

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal("failed to open rpio")
	}

	servo := rpio.Pin(servoPinID)
	servo.Mode(rpio.Pwm)
	//c.pins.servo.Pwm()
	servo.Freq(frequency)

	i := uint32(0)
	for {
		if i > maxvalue {
			i = 0
		}
		servo.DutyCycle(i, maxvalue)
		time.Sleep(100 * time.Millisecond)
	}
}
