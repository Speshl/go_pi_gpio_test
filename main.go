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

/*DON'T MODIFY THIS*/
const frequency = 100000
const cycleLen = uint32(2000)

const maxvalue_limited = uint32(200)
const maxvalue = uint32(250)

const midvalue = uint32(150)

const minvalue_limited = uint32(100)
const minvalue = uint32(50)

/*DON'T MODIFY THIS*/

const enableServo = true
const enableESC = true

func main() {
	err := rpio.Open()
	if err != nil {
		log.Fatal("failed to open rpio")
	}
	defer rpio.Close()

	if enableServo {
		go func() {
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
				if i > maxvalue_limited {
					i = minvalue_limited
				}
				//log.Printf("Sending %d of %d\n", i, maxvalue)
				servo.DutyCycle(i, cycleLen)
				time.Sleep(100 * time.Millisecond)
				if i == maxvalue {
					log.Println("waiting on top")
					time.Sleep(5 * time.Second)
				}
				if i == minvalue {
					log.Println("Waiting at bottom")
					time.Sleep(5 * time.Second)
				}
				i++
			}
		}()
	}
	if enableESC {
		go func() {
			esc := rpio.Pin(escPinID)
			esc.Mode(rpio.Pwm)
			//c.pins.servo.Pwm()
			esc.Freq(frequency)

			log.Println("Starting ESC Calibration")
			log.Println("Setting FULL throttle for 10 seconds")
			esc.DutyCycle(maxvalue, cycleLen)
			time.Sleep(10 * time.Second)
			log.Println("Setting MIN throttle for 5 seconds")
			esc.DutyCycle(minvalue, cycleLen)
			time.Sleep(5 * time.Second)
			log.Println("Setting CENTER throttle for 5 seconds")
			esc.DutyCycle(midvalue, cycleLen)
			time.Sleep(5 * time.Second)
			for {
				esc.DutyCycle(200, cycleLen)
				time.Sleep(5 * time.Second)
			}
		}()
	}

	for {
	}

}
