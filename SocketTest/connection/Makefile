main: build
	java Server

pre: ServerThread.java
	javac ServerThread.java

build: pre Server.java Client.java
	javac Server.java
	javac Client.java

client: Client.class
	java Client
