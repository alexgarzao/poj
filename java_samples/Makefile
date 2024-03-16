clean:
	rm -Rf *.class

jasm-from-java:
	javac $(program).java && jasm -d $(program).class

class-from-jasm:
	jasm $(program).jasm && java $(program)
