# olcdsample
以下のように接続

|OLEDディスプレイ|Raspberry Pi|
|----------------|------------|
|GND             |GND         |
|VCC             |3.3V        |
|SCL             |SCL（GPIO3）|
|SDA             | SDA（GPIO2）|

![回路図](./kairo.png "回路図")

# build && install

``` bash

$ make pi
$ ./copypi.sh [raspberrypi]

```
