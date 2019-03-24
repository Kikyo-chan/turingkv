#!/bin/bash

function randStr()
{
	j=0;
	for i in {a..z};do array[$j]=$i;j=$(($j+1));done
	for i in {A..Z};do array[$j]=$i;j=$(($j+1));done
	for ((i=0;i<512;i++));do strs="$strs${array[$(($RANDOM%$j))]}"; done;
		echo $strs
}


for i in {1..9}
do
str=$(randStr)
echo "RUN SET TEST"
ab -n 100 -c 10 -T application/json -p "data/test."$i".json"  "http://127.0.0.1:9988/keys/"$str"/"
echo "RUN SET TEST OVER"
echo "RUN GET TEST"
ab -n 100 -c 10 "http://127.0.0.1:9988/keys/"$str"/"
echo "RUN GET TEST OVER"
done



