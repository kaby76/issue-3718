# Generated code from trgen 0.16.6
GENERATED = SimpleBooleanParser.cs SimpleBooleanLexer.cs 
default:
	dotnet restore
	dotnet build
run:
	trwdog dotnet run $(RUNARGS)
clean:
	dotnet clean
	rm -rf bin obj
test:
	bash test.sh
