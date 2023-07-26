# go_pi_gpio_test

Example code utilizing github.com/stianeikeland/go-rpio/v4 to control an RC Servo and ESC.

*WARNING*- Throttle will be set to 75% (half speed forward) if everything works, so car might drive off table :).

- Servo will sweep from left to right (on a limited range to keep from breaking parts if car doesn't support full steering range)
- ESC will be armed to the full range of values, and then after arming will be set to half speed in the forward direction.
