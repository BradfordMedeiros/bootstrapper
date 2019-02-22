# bootstrapper-
Simple bootstrapper to find and connect to other automate clients 
This allows you to discovery new data sources, emphasizing mqtt (at least at beginning)

Server functions
~~~~
bootstrapper serve -f path-to-data-folder	# all content gets saved in path-to-folder 
bootstrapper use-admit admit.sh 		# all additions go through this, and this script can simply reject stuff

data-folder/
    mqtt-topics/   #naive implementation can just add tags under these folders.   Initial implementation doesn't matter
	/192.23.321.23/
		/resources:
			room1/humidity  { type: "nqtt-topic", value: "room1/humidity" }
			room2/humidity   # tags1, tags2, tags3
			room1/temperature
			room2/temperature
                /description:
			{ "description": "fun tech shop located in downtown" }		
		
	/192.23.12.12/some/path
		/topics	
	/bradmedeiros.com/path/topics
		/topics
     game-server/
	quake1:
		something/	
		
	quake2:
		something/

     default/
	192.2.3.1../
		resources/
				


Client functions:
bootstrapper use some-server 
bootstrapper download some-server
bootstrapper use some-server
bootstrapper info some-server  # gets info on current server
bootstrapper get room1/humidity (or -s some-server) --tag=hello
bootstrapper set myresourcename 'data in here' --tag=sometag 

~~~~


automate implementation (what this will enable for automate)
~~~~
hippo extensions install bootstrapper
hippo extensions install auto-joiner

# this basically just creates bootstrapper server with whitelisted ips
hippo auto-joiner serve --allow-server # then you have to init on the other device

# this connects to the other device
hippo auto-joiner join ip-of-other-device --topic-map-file / automap
~~~~


so for example we might do:
bootstrapper get all --tag=mqtt 
bootstrapper set brad/temperature --tag=mqtt 
bootstrapper delete brad/temperature
bootstrapper download 

Ideas:
# Server Manipulation
bootstrapper servers list             # get a list of servers
bootstrapper servers use someserver   # set active server
bootstrapper servers active 	      # get name of active server
bootstrapper servers info             # get info for active server
bootstrapper servers add someotherserver
bootstrapper servers rm  someotherserver

# Get information from multiple
bootstrapper servers use-multi myserver:something,anotherserver:thing
and then we can query 
bootstrapper get #/thing/wow/go  and then we can do client side joins here


bootstrapper use brads-server
bootstrapper set brad/data/piano_temp --tag mqtt '{ url: 'someurl }'d

bootstrapper set seattle/lunches