clean:
	rm -Rf *.class

jasm-from-pas:
	javac $(program).java && jasm -d $(program).class

class-from-jasm:
	jasm $(program).jasm && java $(program)
