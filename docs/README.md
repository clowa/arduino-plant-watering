# Wiring

![Wiring](./wiring.png)

## Components

- 1x Arduino Uno R3
- 3x Water pump (12V)
- 3x MOSFET (TODO: Add type)
- 3x Resistor (10k Ohm)
- 3x Flyback diode

## Explanations

- Flyback diodes are needed, because the water pumps are inductive loads and can generate voltage spikes when turned off. The diodes are used to protect the MOSFETs from these spikes.
- The resistors are used to pull the gate of the MOSFET to ground when the Arduino is not driving the gate. This is to prevent the MOSFET from turning `on` when it should be `off`.

## Open questions

- Can the pumps actually be powered by the same 12 V power supply?
- What can be used to control the speed of the pumps? _PWM?_
- What can be used to measure the water level in the water tank?

## Hardware

### Moisture Sensor

`APKLVSR Capacitive Soil Moisture Sensor v1.2` [Amazon](https://www.amazon.de/dp/B0CQNF7S7L?ref=ppx_yo2ov_dt_b_fed_asin_title)

#### Test measures

| Type     | Moisture Value 1 | Moisture Value 2 |
|----------|------------------|------------------|
| In Water | 11776            | _ToDo_           |
| At Air   | 27840            | _ToDo_           |

Coconut fibre:

| Type     | Moisture Value 1 | Moisture Value 2 |
|----------|------------------|------------------|
| Wet soil | 15168            | _ToDo_           |
| Dry soil | 22656            | _ToDo_           |

Plant soil:

| Type     | Moisture Value 1 | Moisture Value 2 |
|----------|------------------|------------------|
| Wet soil | 16064            | _ToDo_           |
| Dry soil | _ToDo_           | _ToDo_           |

### Ultrasonic Sensors

`Fdit 5V water proof ultrasonic sensor (Fdith9obv7uqge)` [Amazon](https://www.amazon.de/dp/B07N5GHZVX?ref=ppx_yo2ov_dt_b_fed_asin_title)
