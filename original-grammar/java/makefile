# Template generated code from trgen 0.16.6
JAR = C:/Users/Kenne/Documents/GitHub/antlr4-upstream/tool/target/antlr4-4.10.2-SNAPSHOT-complete.jar
# JAR = C:/Users/Kenne/.m2/antlr4-4.10.1-complete.jar
CLASSPATH = $(JAR)\;.
.SUFFIXES: .g4 .java .class
.java.class:
	javac -cp $(CLASSPATH) $*.java
ANTLRGRAMMARS ?= $(wildcard *.g4)
GENERATED = SimpleBooleanParser.java SimpleBooleanLexer.java 
SOURCES = $(GENERATED) \
    Program.java \
    ErrorListener.java
default: classes
classes: $(GENERATED) $(SOURCES:.java=.class)
clean:
	rm -f *.class
	rm -f *.interp
	rm -f *.tokens
	rm -f $(GENERATED)
run:
	trwdog java -classpath $(CLASSPATH) Program $(RUNARGS)
SimpleBooleanParser.java : SimpleBoolean.g4
	java -jar $(JAR) -encoding utf-8  $<
SimpleBooleanLexer.java : SimpleBoolean.g4
	java -jar $(JAR) -encoding utf-8  $<

test:
	bash test.sh
