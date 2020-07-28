# Configuration examples for EVCC

[![Build Status](https://travis-ci.org/andig/evcc-config.svg?branch=master)](https://travis-ci.org/andig/evcc-config)

Configuration examples for the [EVCC EV Charge Controller](https://github.com/andig/evcc).

## Meters

- [E3DC (Battery Meter)](#meter-0)
- [E3DC (PV Meter)](#meter-1)
- [Generisch (MQTT)](#meter-2)
- [Generisch (Script)](#meter-3)
- [Kostal Inverter (Grid Meter)](#meter-4)
- [Kostal Inverter (PV Meter)](#meter-5)
- [Kostal Smart Energy Meter (Grid Meter)](#meter-6)
- [Modbus](#meter-7)
- [Modbus (RTU)](#meter-8)
- [Multiple Grid Inverters combined (PV Meter)](#meter-9)
- [SMA Home Manager 2.0 / SMA Energy Meter 30](#meter-10)
- [Solarlog (Grid Meter)](#meter-11)
- [Solarlog (PV Meter)](#meter-12)
- [vzlogger (HTTP)](#meter-13)
- [vzlogger (Push Server/ Websocket)](#meter-14)
- [vzlogger (split import/export channels)](#meter-15)

## Chargers

- [EVSE Wifi](#charger-0)
- [Generisch](#charger-1)
- [Generisch (MQTT)](#charger-2)
- [go-eCharger (Cloud)](#charger-3)
- [go-eCharger (Lokal)](#charger-4)
- [KEBA Connect](#charger-5)
- [Mobile Charger Connect](#charger-6)
- [NRGKick Bluetooth](#charger-7)
- [NRGKick Connect](#charger-8)
- [Phoenix EMCP Controller (Ethernet/Modbus TCP)](#charger-9)
- [Phoenix EVCC Controller (Modbus)](#charger-10)
- [Simple EVSE (Ethernet/Modbus TCP)](#charger-11)
- [Simple EVSE (USB)](#charger-12)
- [Wallbe (Eco, Pro)](#charger-13)
- [Wallbe (pre 2019)](#charger-14)

## Vehicles

- [Generisch](#vehicle-0)
- [Generisch (Script)](#vehicle-1)

## Details

### Meters


<a id="meter-0"></a>
#### E3DC (Battery Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40070
      type: holding
      decode: int32
    scale: -1 # reverse direction
```

<a id="meter-1"></a>
#### E3DC (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: e3dc.fritz.box:502
    id: 1 # ModBus slave id
    register: # manual register configuration
      address: 40067 # (40068 - 1) "Photovoltaikleistung in Watt"
      type: holding
      decode: int32s
    scale: -1 # reverse sign
```

<a id="meter-2"></a>
#### Generisch (MQTT)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-3"></a>
#### Generisch (Script)

```yaml
- type: default
  power:
    type: script # use script
    cmd: /bin/sh -c "echo 0" # actual command
    timeout: 3s # kill script after 3 seconds
```

<a id="meter-4"></a>
#### Kostal Inverter (Grid Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.168.178.52:1502 
  id: 71 # Configured Modbus Device ID 
  register: # manual register configuration
    address: 252 # (see https://www.kostal-solar-electric.com/de-de/download/-/media/document-library-folder---kse/2018/08/30/08/53/ba_kostal_interface_modbus-tcp_sunspec.pdf)
    type: holding
    decode: float32s #swapped float encoding
```

<a id="meter-5"></a>
#### Kostal Inverter (PV Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.168.0.1:1502
  id: 71
  power: Power
```

<a id="meter-6"></a>
#### Kostal Smart Energy Meter (Grid Meter)

```yaml
- type: modbus
  model: kostal
  uri: 192.168.0.1:502
  id: 71
  power: Power
  energy: Energy
```

<a id="meter-7"></a>
#### Modbus

```yaml
- type: modbus
  model: sdm
  uri: rs485.fritz.box:23
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

<a id="meter-8"></a>
#### Modbus (RTU)

```yaml
- type: modbus
  model: sdm
  uri: rs485.fritz.box:23
  rtu: true # rs485 device connected using ethernet adapter
  id: 2
  power: Power # default values, optionally override
  energy: Sum # default values, optionally override
```

<a id="meter-9"></a>
#### Multiple Grid Inverters combined (PV Meter)

```yaml
- type: default
  power:
    type: calc # use the calc plugin
    add: # The add function sums up both string values
    - type: modbus
      model: sunspec
      value: 160:1:DCW # string 1
      uri: 192.168.178.52:1502 
      id: 71 # Configured Modbus Device ID 
    - type: modbus  
      value: 160:2:DCW # string 2
      uri: 192.168.178.52:1502 
      id: 71 # Configured Modbus Device ID 
```

<a id="meter-10"></a>
#### SMA Home Manager 2.0 / SMA Energy Meter 30

```yaml
- type: sma
  serial: 1234567890 # Serial number of the device
```

<a id="meter-11"></a>
#### Solarlog (Grid Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.168.0.32:502 # IP address of the SolarLog device and ModBus port address
    id: 1
    register:
      address: 3518
      type: input
      decode: uint32s
```

<a id="meter-12"></a>
#### Solarlog (PV Meter)

```yaml
- type: default
  power:
    type: modbus
    uri: 192.168.0.32:502 # IP address of the SolarLog  device and ModBus port address
    id: 1
    register:
      address: 3502
      type: input
      decode: uint32s
```

<a id="meter-13"></a>
#### vzlogger (HTTP)

```yaml
- type: default
  power: # power reading
    type: mqtt # use mqtt
    topic: mbmd/sdm1-1/Power # mqtt topic
    timeout: 10s # don't use older values
```

<a id="meter-14"></a>
#### vzlogger (Push Server/ Websocket)

```yaml
- type: default
  type: default
  power:
    type: ws # use websocket plugin
    uri: ws://volkszaehler:8082/socket
    jq: .data | select(.uuid=="<uuid>") .tuples[0][1] # parse response json
    timeout: 30s
    scale: 1
```

<a id="meter-15"></a>
#### vzlogger (split import/export channels)

```yaml
- type: default
  power:
    type: calc # use calc plugin
    add:
      - type: http # import channel
        uri: http://volkszaehler/api/data/<import-uuid>.json?from=now
        jq: .data.tuples[0][1] # parse response json
      - type: http # export channel
        uri: http://volkszaehler/api/data/<export-uuid>.json?from=now
        jq: .data.tuples[0][1] # parse response json
        scale: -1 # export must result in negative values
```


### Chargers


<a id="charger-0"></a>
#### EVSE Wifi

```yaml
- type: evsewifi
  uri: http://192.168.1.4 # SimpleEVSE-Wifi address
```

<a id="charger-1"></a>
#### Generisch

```yaml
- type: default
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

<a id="charger-2"></a>
#### Generisch (MQTT)

```yaml
- type: default
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

<a id="charger-3"></a>
#### go-eCharger (Cloud)

```yaml
- type: go-e
  token: 4711c # or go-e cloud API token
  cache: 10s # go-e cloud API cache duration
```

<a id="charger-4"></a>
#### go-eCharger (Lokal)

```yaml
- type: go-e
  uri: http://192.168.1.4 # either go-e local address
```

<a id="charger-5"></a>
#### KEBA Connect

```yaml
- type: keba
  uri: 192.168.1.4:7090 # KEBA address
  rfid:
    tag: 765765348 # RFID tag, see `evcc charger` to show tag
```

<a id="charger-6"></a>
#### Mobile Charger Connect

```yaml
- type: mcc
  uri: https://192.168.1.4 # Mobile Charger Connect address
  password: # home user password
```

<a id="charger-7"></a>
#### NRGKick Bluetooth

```yaml
- type: nrgkick-bluetooth
  macaddress: 00:99:22 # MAC address
  pin: # pin
```

<a id="charger-8"></a>
#### NRGKick Connect

```yaml
- type: nrgkick-connect
  ip: 192.168.1.4 # IP
  macaddress: 00:99:22 # MAC address
  password: # password
```

<a id="charger-9"></a>
#### Phoenix EMCP Controller (Ethernet/Modbus TCP)

```yaml
- type: phoenix-emcp
  uri: 192.168.0.8:502 # TCP ModBus address
  id: 1
```

<a id="charger-10"></a>
#### Phoenix EVCC Controller (Modbus)

```yaml
- type: phoenix-evcc
  device: /dev/ttyUSB0
  baudrate: 9600
  comset: 8N1
  id: 1 # slave id
```

<a id="charger-11"></a>
#### Simple EVSE (Ethernet/Modbus TCP)

```yaml
- type: simpleevse
  uri: 192.168.0.8:502 # TCP ModBus address
```

<a id="charger-12"></a>
#### Simple EVSE (USB)

```yaml
- type: simpleevse
  device: /dev/usb1 # RS485 ModBus device
```

<a id="charger-13"></a>
#### Wallbe (Eco, Pro)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # TCP ModBus address
```

<a id="charger-14"></a>
#### Wallbe (pre 2019)

```yaml
- type: wallbe
  uri: 192.168.0.8:502 # TCP ModBus address
  legacy: true # enable for older Wallbes with Phoenix EV-CC-AC1-M3-CBC-RCM controller
```


### Vehicles


<a id="vehicle-0"></a>
#### Generisch

```yaml
- type: default
  title: Mein Auto # name
  capacity: 50 # kWh
  charge:
    type: ...
    # ...
  cache: 5m # cache duration
```

<a id="vehicle-1"></a>
#### Generisch (Script)

```yaml
- type: default
  title: Mein Auto # name
  capacity: 50 # kWh
  charge:
    type: script # use script
    cmd: /bin/sh -c "echo 50" # actual command
    timeout: 3s # kill script after 3 seconds
  cache: 5m # cache duration
```

