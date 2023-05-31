#!/bin/bash

# user api
osascript -e 'tell app "Terminal" to do script "cd /Users/liuxian/GoProjects/project/MouHu/service/app/user/rpc && go run user.go -f etc/user.yaml"'

# user rpc
osascript -e 'tell app "Terminal" to do script "cd /Users/liuxian/GoProjects/project/MouHu/service/app/user/api && go run user.go -f etc/user-api.yaml"'

# qa api
osascript -e 'tell app "Terminal" to do script "cd /Users/liuxian/GoProjects/project/MouHu/service/app/qa/rpc && go run qa.go -f etc/qa.yaml"'

# qa rpc
osascript -e 'tell app "Terminal" to do script "cd /Users/liuxian/GoProjects/project/MouHu/service/app/qa/api && go run qa.go -f etc/qa-api.yaml"'
