#!/bin/bash

pushd csharp; make; ./bin/Debug/net6.0/Test.exe -file ../p5.txt > out.txt 2>&1; dos2unix out.txt; popd
pushd go; rm Test.exe; make; ./Test.exe -file ../p5.txt > out.txt 2>&1; dos2unix out.txt; popd
pushd java; make clean; make; make run RUNARGS="-file ../p5.txt" > out.txt 2>&1; dos2unix out.txt; popd
