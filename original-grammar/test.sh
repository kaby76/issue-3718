#!/bin/bash

pushd csharp; make clean; make; ./bin/Debug/net6.0/Test.exe -file ../../p15.txt 2>&1 > out.txt; dos2unix out.txt; popd
pushd java; make clean; make; make run RUNARGS="-file ../../p15.txt" 2>&1 > out.txt ; dos2unix out.txt; popd
pushd go; rm -f Test.exe; make clean; make; trwdog -t 4 ./Test.exe -file ../../p15.txt 2>&1 > out.txt ; dos2unix out.txt; popd
