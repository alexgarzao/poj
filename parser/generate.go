package parser

//go:generate java -Xmx500M -cp "antlr-4.13.1-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool -Dlanguage=Go -no-visitor -package parsing JSON.g4 -o ../parsing
