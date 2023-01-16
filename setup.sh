echo -e "\033[33mStart checking whether the Go development environment is installed locally ...... \033[0m"
	go version
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo "\033[32m[Success] The Go development environment has been detected!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the GO development environment, please refer to the manual for installation.\033[0m"
    	exit
    fi

echo "\033[33mStart checking whether the docker application is installed locally ...... \033[0m"
	docker -v
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo "\033[32m[Success] The docker application has been detected!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the docker application, please refer to the manual for installation.\033[0m"
      exit
    fi

echo "\033[33mStart checking whether the docker-compose application is installed locally ...... \033[0m"
	docker-compose version
    # shellcheck disable=SC2181
    if [ $? -eq  0 ]; then
        echo "\033[32m[Success] The docker-compose application has been detected!\033[0m"
    else
    	echo "\033[31m[Error] The current device does not detect the docker-compose application, please refer to the manual for installation.\033[0m"
      exit
    fi

echo "\033[33mStart build test...... \033[0m"
  go build -o young_engine
    if [ $? -eq  0 ]; then
        echo "\033[32m[Success] build success!\033[0m"
        rm ./young_engine
    else
    	echo "\033[31m[Error] build failed, Please resolve the issue based on the error message\033[0m"
      exit
    fi

echo ""
echo "\033[32mCongratulations on having installed all your environment dependencies and start your project journey!\033[0m"