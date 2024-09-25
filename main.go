package main

import (
	"machine"
	"time"
)

const (
	measureInterval   = time.Second * 60
	moistureThreshold = 23000 // Higher moisture level means the soil is drier
)

var (
	moistureSensor1 = machine.ADC{Pin: machine.ADC0}
	greenLed        = machine.D13
	redLed          = machine.D12
)

func main() {
	machine.InitADC() //! Better don't forget to initialize ADC
	configureDigitalOutputPins(greenLed, redLed)

	moistureSensor1.Configure(machine.ADCConfig{
		Reference:  5000, // 5 Volts
		Resolution: 10,   // 10 bits
	})

	// Check the moisture level in a certain interval
	for {
		moisture := readMoisture(moistureSensor1)
		println("Moisture level: ", moisture)
		controlLights(moisture, greenLed, redLed)

		time.Sleep(measureInterval)
	}
}

// configureDigitalOutputPins configures the given pins as digital output pins
func configureDigitalOutputPins(pins ...machine.Pin) {
	for _, p := range pins {
		p.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

// readMoisture reads the moisture level from the given sensor and returns it
func readMoisture(sensor machine.ADC) uint16 {
	return sensor.Get() // Analog value of moisture level
}

// controlLights controls the LEDs based on the moisture level
func controlLights(moisture uint16, ok machine.Pin, warn machine.Pin) {
	const ledBlinkDuration = time.Millisecond * 500

	// If moisture level is dried, turn on the warning light
	if moisture > moistureThreshold {
		blink(warn, ok, ledBlinkDuration)
		return
	}

	// If moisture level is wet, turn on the OK light
	blink(ok, warn, ledBlinkDuration)
}

// blink blinks the given LED for the given duration and turns off the counter LED
func blink(led machine.Pin, counterLed machine.Pin, duration time.Duration) {
	counterLed.Low()

	led.High()
	time.Sleep(duration)
	led.Low()
}
