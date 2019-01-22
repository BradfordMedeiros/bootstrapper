# bootstrapper-
Simple bootstrapper to find and connect to other automate clients 
This allows you to discovery new data sources, emphasizing mqtt (at least at beginning)

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
				


bootstrapper use some-server 
bootstrapper download some-server
bootstrapper use some-server
bootstrapper get room1/humidity (or -s some-server) --tag=hello
bootstrapper set myresourcename 'data in here' --tag=sometag (no tag is default)

~~~~
