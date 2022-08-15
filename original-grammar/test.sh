#!/bin/bash

unameOut="$(uname -s)"
case "${unameOut}" in
    Linux*)     machine=Linux;;
    Darwin*)    machine=Mac;;
    CYGWIN*)    machine=Cygwin;;
    MINGW*)     machine=MinGw;;
    *)          machine="UNKNOWN:${unameOut}"
esac
if [ "$machine" == "Linux" ]
then
	makefiles="makefile.linux"
else
	makefiles="makefile.win"
fi

for i in csharp java go
do
	pushd $i
	make -f "$makefiles" clean
	make -f "$makefiles" build
	make -f "$makefiles" run RUNARGS="-file ../../p15.txt" 2>&1
	popd
done	
