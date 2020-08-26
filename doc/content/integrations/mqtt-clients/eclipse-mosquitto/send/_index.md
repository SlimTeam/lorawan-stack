---
title: "Send Messages"
description: ""
weight: 3
---

Learn how to send data to your end device by scheduling a downlink message with the Eclipse Mosquitto **mosquitto_pub** client

<!--more-->

>Note: this section follows the example for publishing downlink traffic in the [MQTT Server]({{< ref "/integrations/mqtt" >}}) guide.

For connecting to a host and publishing a message, **mosquitto_pub** client defines a command with the following syntax:

```bash 
mosquitto_pub -h {hostname} -p {port} -u {username} -P {password} -t {topic} -m {message}
```

Modify the command above to publish a message to {{% tts %}} MQTT Server and in that way schedule a downlink to be sent to your end device.

<details><summary>Example</summary>

```bash
mosquitto_pub -h "thethings.example.com" -p "1883" -u "app1" -P "NNSXS.VEEBURF3KR77ZR.." -t "v3/app1/devices/dev1/up" -m '{"downlinks":[{"f_port": 15,"frm_payload":"vu8=","priority": "NORMAL"}]}'
```
</details>

If TLS is being used, change the port value to `8883` and add the `--cafile` option to the command as described in the [Receive Messages]({{< ref "/integrations/mqtt-clients/eclipse-mosquitto/receive" >}}) section.

>Note: read more about the command line options in the [mosquitto_pub manual](https://mosquitto.org/man/mosquitto_pub-1.html).