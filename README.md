# Overview

This is a small home arduino project to automatically water my plants. The project is based on an Arduino Uno R3, a soil moisture sensor, a water pump and a electronic switch (MOSFET). The soil moisture sensor is used to measure the moisture level of the soil and the water pump is used to water the plants. The MOSFET is used to control the water pump because the pump runs on another voltage than the board itself.

## Getting Started

### Prerequisites

- [TinyGo](https://tinygo.org/getting-started/install/)
- [avrdude](https://github.com/avrdudes/avrdude) to flash the project to the Arduino board
- _Optional:_ [go-task](https://taskfile.dev/installation/) to use handy `makefile` like commands

### Flashing the project

To flash the project to the Arduino board, run the following command:

```bash
tinygo flash -target arduino .
# Or if you have go-task installed
task flash
```
