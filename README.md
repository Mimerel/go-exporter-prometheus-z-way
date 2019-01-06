# go-exporter-prometheus-z-way
Collect Zway metrics to store them in prometheus

* Prerequisits

The possibility to run a go webservice on your zway device
I use this exporter on a raspberry pi running the z-wave.me software

* Modification to be done to you Z-wave devices

You must name your devices in your z-wave.me website as follows :

<name of the device> | <Room it is in> | <type of device>

the type of device could be :

```
plug
Temperature
Co2
sensor
Lampe
...
```

* RUN : to run the application either : 
```
go run main.go
```
or
```
go build  // to build the application

then

./go-exporter-prometheus-z-way // to run the build
```

You will probably be missing dependencies
```
github.com/apsdehal/go-logger"
github.com/prometheus/client_golang/prometheus"
github.com/prometheus/client_golang/prometheus/promhttp"
gopkg.in/yaml.v2"
gopkg.in/alecthomas/kingpin.v2"
github.com/sirupsen/logrus"
golang.org/x/text/transform"
golang.org/x/text/unicode/norm"
```

to add a dependency run 

```
go get <name of dependency>
```

for example : 
```
go get github.com/Mimerel/go-logger-client
```

* Configuration file

```

host: zwave // I have several zwave servers therefore each one has its own name garage, kitchen, ...

port: 2112  // port on which you want the exporter to run.. 

zway_Server: http://<ip of you zwave raspberry>:8083

logger: http://<ip of your logger>:9999 // if you want to centralize logs in an elasticSearch

// As well as uploading the zwave metrics, the exporter can also upload details on specific running threads
followed_Services:
  go-exporter-prometheus-z-way: exporter_service   // real name of the process : name of the metric used for prometheus
  z-way-server: zway_server

// name of the modules you want the exporter to run
// systemData creates metrics for the processes defined in followed_service
// zway creates the z-wave.me metrics
activated_Modules:
- systemData
- zway

// device types you which to create metrics for
// the types are those used in the z-wave.me interface to name the devices
device_Types:
- plug
- Temperature
- Co2
- sensor
- Lampe

// additionnal configuration to override default settings in z-wave.me
// 54_0 means device Id 54, instance 0
// this is necessary as some devices have several fonctionnalities and you
// may not want all of them.
// in my case, for example, 54_1, 54_2 :
// two entries exist for may alarm system.
// as I only use the first one, I have renammed it to Alarm Sirène => 54_1 -> name : Alarm Sirène
// and decided to ignore the other one 54_2 => ignore: true
device_configuration:
  54_0:
    ignore: true
  54_1:
    name: Alarm Sirène
  54_2:
    ignore: true
  53_0:
    ignore: true
  53_1:
    name: Alarm Totale
  53_2:
    name: Alarm Partielle
```