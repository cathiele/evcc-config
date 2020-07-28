# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters

- E3DC Batterie
- Generisch (MQTT)
- Generisch (Script)
- Modbus
- Modbus (RTU)
- SMA Home Manager 2.0 / SMA Energy Meter 30

## Chargers

- EVSE Wifi
- Generisch
- Generisch (MQTT)
- go-eCharger (Cloud)
- go-eCharger (Lokal)
- KEBA Connect
- Mobile Charger Connect
- NRGKick Bluetooth
- NRGKick Connect
- Phoenix EMCP Controller (Ethernet/Modbus TCP)
- Phoenix EVCC Controller (Modbus)
- Simple EVSE (Ethernet/Modbus TCP)
- Simple EVSE (USB)
- Wallbe (Eco, Pro)
- Wallbe (pre 2019)

## Details

### Meters


#### E3DC Batterie

```yaml
- type: meter
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1
    register: # manual register configuration
      address: 40070
      type: holding
      decode: int32
    scale: -1 # reverse direction
```

#### Generisch (MQTT)

```yaml
- type: meter
  power: # power reading
    type: mqtt # use mqtt
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

#### Generisch (Script)

```yaml
- type: meter
  power:
    type: script # use script
    cmd: /bin/sh -c "echo 0" # actual command
    timeout: 3s # kill script after 3 seconds
```

#### Modbus

```yaml
- type: meter
  model: sdm
  uri: rs485.fritz.box:23
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

#### Modbus (RTU)

```yaml
- type: meter
  model: sdm
  uri: rs485.fritz.box:23
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

#### SMA Home Manager 2.0 / SMA Energy Meter 30

```yaml
- type: meter
  serial: 1234567890 # Serial number of the device
```


### Chargers


#### EVSE Wifi

```yaml
- type: charger
  uri: http://192.168.1.4 # SimpleEVSE-Wifi address
```

#### Generisch

```yaml
- type: charger
  status: # charger status A..F
    type: ...
    # ...
  enabled: # charger enabled state (true/false or 0/1)
    type: ...
    # ...
  enable: # set charger enabled state
    type: ...
    # ...
  maxcurrent: # set charger max current
    type: ...
    # ...
```

#### Generisch (MQTT)

```yaml
- type: charger
  status: # charger status A..F
    type: mqtt
    topic: some/topic1
  enabled: # charger enabled state (true/false or 0/1)
    type: mqtt
    topic: some/topic2
  enable: # set charger enabled state
    type: script
    cmd: /bin/sh -c "echo ${enable}"
  maxcurrent: # set charger max current
    type: script
    cmd: /bin/sh -c "echo ${maxcurrent}"
```

#### go-eCharger (Cloud)

```yaml
- type: charger
  token: 4711c # or go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

#### go-eCharger (Lokal)

```yaml
- type: charger
  uri: http://192.168.1.4 # either go-e local address
```

#### KEBA Connect

```yaml
- type: charger
  uri: 192.168.1.4:7090 # KEBA address
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

#### Mobile Charger Connect

```yaml
- type: charger
  uri: https://192.168.1.4 # Mobile Charger Connect address
  password: # home user password
```

#### NRGKick Bluetooth

```yaml
- type: charger
  macaddress: 00:99:22 # MAC address
  pin: # pin
```

#### NRGKick Connect

```yaml
- type: charger
  ip: 192.168.1.4 # IP
  macaddress: 00:99:22 # MAC address
  password: # password
```

#### Phoenix EMCP Controller (Ethernet/Modbus TCP)

```yaml
- type: charger
  uri: 192.168.0.8:502 # ModBus address
  id: 1
```

#### Phoenix EVCC Controller (Modbus)

```yaml
- type: charger
  device: /dev/ttyUSB0
  baudrate: 9600
  comset: 8N1
  id: 1 # slave id
```

#### Simple EVSE (Ethernet/Modbus TCP)

```yaml
- type: charger
  uri: 192.168.0.8:502 # TCP ModBus address
```

#### Simple EVSE (USB)

```yaml
- type: charger
  device: /dev/usb1 # RS485 ModBus device
```

#### Wallbe (Eco, Pro)

```yaml
- type: charger
  uri: 192.168.0.8:502 # ModBus address
```

#### Wallbe (pre 2019)

```yaml
- type: charger
  uri: 192.168.0.8:502 # ModBus address
  legacy: true # enable for older Wallbes with Phoenix EV-CC-AC1-M3-CBC-RCM controller
```

