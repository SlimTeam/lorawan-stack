---
title: "Receive Messages"
description: ""
weight: 2
---

This section shows how to subscribe and listen to the messages being published by {{% tts %}} MQTT Server using the Eclipse Mosquitto **mosquitto_sub** client.

<!--more-->

>Note: this section follows the example for subscribing to upstream traffic in the [MQTT Server]({{< ref "/integrations/mqtt" >}}) guide.

The command for connecting to a host and subscribing to a topic has got the following syntax:

```bash 
mosquitto_sub -h {hostname} -p {port} -u {username} -P {password} -t {topic}
```

Use the command above and modify it to subscribe to {{% tts %}} MQTT Server and listen to the messages being sent by your end device.

<details><summary>Example</summary>

```bash
mosquitto_sub -h "thethings.example.com" -p "1883" -u "app1" -P "NNSXS.VEEBURF3KR77ZR.." -t "v3/app1/devices/dev1/up"
```
</details>

In you want to use TLS, you need to change the port value to `8883` and add the `--cafile` option to the command. `--cafile` option is used to define a path to the file containing trusted CA certificates that are PEM encoded.

>Note: read more about the command line options in the [mosquitto_sub manual](https://mosquitto.org/man/mosquitto_sub-1.html).