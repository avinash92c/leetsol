# Read from the file file.txt and output the tenth line to stdout.
input="file.txt"
count=0

while read line;
do
((count=count+1))
if [ $count -eq 10 ];then
echo $line;
exit;
fi
done < $input;
echo $line
