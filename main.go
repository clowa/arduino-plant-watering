package main

import (
	"machine"
	"time"
)

const (
	measureFrequency  = time.Second * 60
	moistureThreshold = 21000 // Higher moisture level means the soil is drier
)

var (
	// Plant pair 1
	moistureSensor1 = machine.ADC{Pin: machine.ADC0}
	pump1           = machine.D11
)

func main() {
	machine.InitADC() //! Better don't forget to initialize ADC
	configureDigitalOutputPins(pump1)
	configureAnalogInputPins(machine.ADCConfig{
		Reference:  5000, // 5 Volts
		Resolution: 10,   // 10 bits
	}, moistureSensor1)

	// Check the moisture level in a certain interval
	for {
		controlPump(moistureSensor1, pump1)
		time.Sleep(measureFrequency)
	}
}

// configureDigitalOutputPins configures the given pins as digital output pins
func configureDigitalOutputPins(pins ...machine.Pin) {
	for _, p := range pins {
		p.Configure(machine.PinConfig{Mode: machine.PinOutput})
	}
}

// configureAnalogInputPins configures the given pins as analog input pins with the given configuration
func configureAnalogInputPins(conf machine.ADCConfig, pins ...machine.ADC) {
	for _, p := range pins {
		p.Configure(conf)
	}
}

// controlPump controls the pump based on the moisture level
func controlPump(sensor machine.ADC, pump machine.Pin) {
	const wateringDuration = time.Millisecond * 1500

	moisture := readMoisture(sensor)
	println("Moisture level: ", moisture)

	// If moisture level is dried, turn on the warning light
	if moisture > moistureThreshold {
		runPump(pump, wateringDuration)
		return
	}
}

// readMoisture reads the moisture level from the given sensor and returns it
func readMoisture(sensor machine.ADC) uint16 {
	return sensor.Get() // Analog value of moisture level
}

// runPump starts a pump for a certain duration
func runPump(pump machine.Pin, duration time.Duration) {
	pump.High()
	time.Sleep(duration)
	pump.Low()
}
