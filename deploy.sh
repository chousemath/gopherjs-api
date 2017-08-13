#!/bin/bash

GREEN='\033[0;32m'
LIGHTCYAN='\033[1;36m'
NC='\033[0m'

echo -e "${GREEN}Type your commit comment, followed by [ENTER]:${NC}"
# the read command stores the user input as a string
read comment
/Users/jo/go/bin/godep save .
# the -e command enables the backslash necessary for colored text
echo -e "${GREEN}commit comment has been loaded${NC}"
# just a spacer
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}git add . started${NC}"
git add .
echo -e "${GREEN}git add . completed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}git commit -am started${NC}"
eval "git commit -am \"${comment}\""
echo -e "${GREEN}git commit -am completed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}git push started${NC}"
git push
echo -e "${GREEN}git push completed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}git push heroku was started${NC}"
git push heroku
echo -e "${GREEN}git push heroku was completed${NC}"
echo -e "${LIGHTCYAN}............${NC}"
echo -e "${GREEN}deployment complete!${NC}"
