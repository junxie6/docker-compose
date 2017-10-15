#!/bin/bash

longStr=$(printf 'asdf %d ' {1..1000})

for i in {1..1000}; do
	mysql -e "INSERT INTO exp.student (Name, Age, Intro) VALUES ('Jun ${i}', 15, '${longStr}');"
done


