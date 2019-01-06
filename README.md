# go-goole-home-requests
Execute z-wave commands and scenarios

This go software is to be used with Google Home, IFTTT, and z-wave.me razberry running on a raspberry
It enables you to send commands to on or several zwave domotique servers to execute commands

* Prerequisits

Google home must be linked to your google account
The same user account must be linked to you IFTT account

* RUN : to run the application either : 
```
go run main.go
```
or
```
go build  // to build the application
```
then 
```
./go-google-home-requests // to run the build
```

You will probably be missing dependencies
```
	github.com/op/go-logging
	github.com/Mimerel/go-logger-client
	go-goole-home-requests/configuration
	go-goole-home-requests/google_talk
	go-goole-home-requests/utils
```

to add a dependency run 

```
go get <name of dependency>
```

for example : 
```
go get github.com/Mimerel/go-logger-client
```



* Zwave - Razberry


On the Z-wave.me admin local website 
go to  Applications and find : Z-Wave Network Access
Enable : Allow public access to Z-Wave API

* Router // Public IP

IFTT requires a public ip to communicate with your network.
Make sur your router is open to farward requests to the server running this application.

* IFTTT setup

Create a new applet :

if -> Google Assitant => Say a phrase with a text ingredient
in my case, I use the "name" of the room as first command to enable this program
to make the difference between commands sent by my kitchen or living room

What do you want to say? -> maison $

then -> Webhook
in the url : http://your_public_ip:port_you_opened_on_you_network/maison/<< {{TextField}}>>

More examples:


What do you want to say? -> kitchen $
in the url : http://public_ip:port/kitchen/<< {{TextField}}>>

What do you want to say? -> Living room $
in the url : http://public_ip:port/room1/<< {{TextField}}>>

=> Kitchen & room1 will have to be set in the configuration.yaml file.

* Configuration file

```
// the elasticSearch params are used if you have also added my
// "github.com/Mimerel/go-logger-client" and have an elasticsearch running
// Otherwize omit this part.. logs will be displayed in console.
elasticSearch:
  url: http://elastic_search_ip:9200
host: go-google-log

// Google home seems to have difficulties in managing special chars
// As I am French, I use a lot of them..
// this part of the conf file is used to specify chars that have to be
// ignored when comparing command sent by google home to 
// the list of commands in the configuration.yaml
charsToRemove:
  - é
  - à
  - ç
  - è
  - ê
  - ö
  - à
  - ë
  - ô
  
// this part of the configuration file enables you
// to declare you google homes.
// the name used corresponds to the name used in the url on IFTTT
// I use the the fictive "home" to enable general commands and have messages
// Sent to both google homes.
googles:
  - name: kitchen
    ip:
    - 192.168.0.100
  - name: room1
    ip:
    - 192.168.0.101
  - name: home
    ip:
    - 192.168.0.100
    - 192.168.0.101
    
// this part enables you to declare you z-wave.me servers    
zwaves:
  - name: kitchen
    ip: 192.168.0.50
  - name: living_room
    ip: 192.168.0.51
  - name: garage
    ip: 192.168.0.52
  - name: room_parents
    ip: 192.168.0.53
    
// list of possible actions
// in the name array list all actions you could say that 
// end up executing the same thing
// only one word can be used at the moment
// you could use on / off / open / close
// the replacement value is the unique identifier you will use in the 
// configuration of your commands
// value is the default value you wish you zwave device(s) to be set to
// it can be overwritten in the commands 
// the type must always be set to domotiqueCommand as it is the only one 
// implemeneted at the moment.    
actions:
  - name:
      - on
    replacement: on
    type: domotiqueCommand
    value: 255
  - name:
      - off
      - cut
    replacement : off
    type: domotiqueCommand
    value: 0
  - name:
      - open
    replacement: open
    type: domotiqueCommand
    value: 255
  - name:
      - close
    replacement: close
    type: domotiqueCommand
    value: 0
    
// this part lists your zwave devices
// you have to give them a name that will be used in the configuration of your commands
// id, instance and commandClass can be found by hoving over the
// zwave switch in you z-wave me interface..
// you will see on the bottom of you navigator something like : 
// ....ZWayVDev_zway_8_0_37 => device id 8, instance 0, commandClass 37    
devices:
  - name: lampe_basse
    zwave: kitchen
    id: 2
    instance: 0
    commandClass: 37
  - name: lampe_halogene
    zwave: kitchen
    id: 4
    instance: 0
    commandClass: 37
  - name: lampe_leds
    zwave: kitchen
    id: 59
    instance: 0
    commandClass: 37
  - name: lampe_haute
    zwave: garage
    id: 8
    instance: 0
    commandClass: 37
  - name: lampe_etoile
    zwave: room_parents
    id: 2
    instance: 0
    commandClass: 37

// the main configuration of you software
// words => list of sentences that must trigger the actions
// rooms => list of google home that can trigger the actions
// actions => authorized actions (on, off, open, close)
// instructions: => device that should be activated / deactivated
// the first command below means:
// if I say "allume la lampe haute" or "eteins la lampe haute"
// from any google home
// it will run the device lampe_haute with the value defined in the action
// activated.. this can be overridden by completing the value: 
command:
  - words:
    - "la lampe haute"
    rooms:
    - home
    - kitchen
    - room_parents
    actions:
      - allume
      - éteins
    instructions:
      - name: lampe_haute
        value: 
  - words:
    - "la lampe étoile"
    - "les lampes étoile"
    - "la lampe étoiles"
    - "les lampes étoiles"
    - "les étoiles"
    rooms:
    - maison
    - salon
    - cuisine
    actions:
    - allume
    - éteins
    instructions:
      - name: lampe_etoile
        value: 
  - words:
    - "la lampe basse"
    rooms:
    - maison
    - salon
    - cuisine
    actions:
    - allume
    - éteins
    instructions:
      - name: lampe_basse
        value:
```