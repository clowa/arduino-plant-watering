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
)

func main() {
	machine.InitADC() //! Better don't forget to initialize ADC
	configureDigitalOutputPins(greenLed)

	moistureSensor1.Configure(machine.ADCConfig{
		Reference:  5000, // 5 Volts
		Resolution: 10,   // 10 bits
	})

	// Check the moisture level in a certain interval
	for {
		moisture := readMoisture(moistureSensor1)
		println("Moisture level: ", moisture)
		controlPump(moisture, greenLed)

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

// controlPump controls the LEDs based on the moisture level
func controlPump(moisture uint16, pump machine.Pin) {
	const wateringTime = time.Second * 1

	// If moisture level is dried, turn on the warning light
	if moisture > moistureThreshold {
		executePumpAction(pump, wateringTime)
		return
	}
}

// executePumpAction starts a pump for a certain duration
func executePumpAction(pump machine.Pin, duration time.Duration) {
	pump.High()
	time.Sleep(duration)
	pump.Low()
}
