# Template generated code from trgen 0.16.6
err=0
SAVEIFS=$IFS
IFS=$(echo -en "\n\b")
for g in `find ../examples -type f | grep -v '.errors$' | grep -v '.tree$'`
do
  file="$g"
  x1="${g##*.}"
  if [ "$x1" != "errors" ]
  then
    echo "$file"
    cat "$file" | trwdog ./bin/Debug/net6.0/Test.exe
    status="$?"
    if [ -f "$file".errors ]
    then
      if [ "$stat" = "0" ]
      then
        echo Expected parse fail.
        err=1
        break
      else
        echo Expected.
      fi
    else
      if [ "$status" != "0" ]
      then
        err=1
        break
      fi
    fi
  fi
done
exit $err
