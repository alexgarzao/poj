package parser

//go:generate java -Xmx500M -cp "antlr-4.13.1-complete.jar:$(CLASSPATH)" org.antlr.v4.Tool -Dlanguage=Go -listener -package parsing Pascal.g4 -o ../parsing
