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
const cycleLen = uint32(1000)
const maxvalue = uint32(200)
const midvalue = uint32(150)
const minvalue = uint32(100)

const enableServo = true

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal("failed to open rpio")
	}
	defer rpio.Close()

	if enableServo {
		servo := rpio.Pin(servoPinID)
		servo.Mode(rpio.Pwm)
		//c.pins.servo.Pwm()
		servo.Freq(frequency)

		servo.DutyCycle(midvalue, cycleLen)
		log.Printf("Servo to middle")
		time.Sleep(5 * time.Second)

		log.Println("Start sweeping")
		i := uint32(100)
		for {
			if i > maxvalue {
				i = minvalue
			}
			//log.Printf("Sending %d of %d\n", i, maxvalue)
			servo.DutyCycle(i, cycleLen)
			time.Sleep(20 * time.Millisecond)
			i++
		}
	} else {
		esc := rpio.Pin(servoPinID)
		esc.Mode(rpio.Pwm)
		//c.pins.servo.Pwm()
		esc.Freq(frequency)
	}

}
