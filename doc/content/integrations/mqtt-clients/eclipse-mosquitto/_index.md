---
title: "Eclipse Mosquitto"
description: ""
weight: 
---

[Eclipse Mosquitto](https://mosquitto.org/) is a project which provides an open source MQTT broker, a C and C++ library for MQTT client implementations and the popular command line MQTT clients. Its lightweight MQTT protocol implementation makes it suitable for full power machines, as well as for the low power and embedded ones. 

<!--more-->

This guide shows how to receive upstream messages and send downlink messages with the Eclipse Mosquitto command line clients and {{% tts %}} [MQTT Server]({{< ref "/integrations/mqtt" >}}).

>Note: Eclipse Mosquitto MQTT server supports 3.1, 3.1.1 and 5.0 MQTT protocol versions.

## Requirements

1. [Eclipse Mosquitto MQTT server](https://github.com/eclipse/mosquitto) installed on your system.
