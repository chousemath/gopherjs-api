#!/bin/bash

echo "Type your commit comment, followed by [ENTER]:"
# the read command stores the user input as a string
read comment
/Users/jo/go/bin/godep save .
git add .
eval "git commit -am \"${comment}\""
git push
git push heroku
