#!/bin/bash
# chkconfig: 2345 55 25
# description: Vez Service

### BEGIN INIT INFO
# Provides:          bt
# Required-Start:    $all
# Required-Stop:     $all
# Default-Start:     2 3 4 5
# Default-Stop:      0 1 6
# Short-Description: starts Vez
# Description:       starts the Vez
### END INIT INFO

PATH=/bin:/sbin:/usr/bin:/usr/sbin:/usr/local/bin:/usr/local/sbin:~/bin

# Source function library.
if [ -f /etc/init.d/functions ];then
  . /etc/init.d/functions
fi

if [ -f /etc/rc.d/init.d/functions ];then
  . /etc/rc.d/init.d/functions
fi

app_path={APP_PATH}
SERVICENAME="vez-en"

app_start(){
    isStart=`ps -ef|grep "$SERVICENAME web" |grep -v grep|awk '{print $2}'`
    if [ "$isStart" == '' ];then
        echo -e "Starting $SERVICENAME... \c"
        cd $app_path && daemon "${app_path}/$SERVICENAME web &"
        isStart=""
        while [[ "$isStart" == "" ]];
        do
            echo -e ".\c"
            sleep 0.5
            isStart=$(lsof -n -P -i:11011|grep LISTEN|grep -v grep|awk '{print $2}'|xargs)
            let n+=1
            if [ $n -gt 15 ];then
                break;
            fi
        done
        if [ "$isStart" == '' ];then
                echo -e "\033[31mfailed\033[0m"
                echo '------------------------------------------------------'
                tail -n 20 ${app_path}/logs/run_away.log
                echo '------------------------------------------------------'
                echo -e "\033[31mError: ${SERVICENAME} service startup failed.\033[0m"
                return;
        fi
        echo -e "\033[32mdone\033[0m"
    else
        echo "Starting ${SERVICENAME}(pid $(echo $isStart)) already running"
    fi

  
}

app_stop(){
	pids=`ps -ef|grep "$SERVICENAME web" |grep -v grep|awk '{print $2}'`
    arr=($pids)
    echo -e "Stopping ${SERVICENAME}... \c"
    for p in ${arr[@]}
    do
            kill -9 $p
    done
    echo -e "\033[32mdone\033[0m"

    pids=`ps -ef|grep "$SERVICENAME robot" |grep -v grep|awk '{print $2}'`
    arr=($pids)
    echo -e "Stopping ${SERVICENAME} robot ... \c"
    for p in ${arr[@]}
    do
            kill -9 $p
    done
    echo -e "\033[32mdone\033[0m"
}

app_status(){
    isStart=`ps -ef|grep "${SERVICENAME} web" |grep -v grep|awk '{print $2}'`
    if [ "$isStart" == '' ];then
      echo -e "${SERVICENAME} not running"
    else
      echo -e "${SERVICENAME}(pid $(echo $isStart)) already running"
    fi

    isStart=`ps -ef|grep "${SERVICENAME} robot" |grep -v grep|awk '{print $2}'`
    if [ "$isStart" == '' ];then
      echo -e "${SERVICENAME} not running"
    else
      echo -e "${SERVICENAME}(pid $(echo $isStart)) already running"
    fi
}

app_robot(){
    isStart=$(ps aux |grep "$SERVICENAME robot" | grep -v grep | awk '{print $2}')
    if [ "$isStart" == '' ];then
            echo -e "Starting $SERVICENAME robot... \c"
            cd $app_path && daemon "${app_path}/$SERVICENAME robot &"
            sleep 0.3
            isStart=$(ps aux |grep "$SERVICENAME robot" |grep -v grep|awk '{print $2}')
            if [ "$isStart" == '' ];then
                    echo -e "\033[31mfailed\033[0m"
                    echo '------------------------------------------------------'
                    tail -n 20 $app_path/logs/task.log
                    echo '------------------------------------------------------'
                    echo -e "\033[31mError: ${SERVICENAME} service startup failed.\033[0m"
                    return;
            fi
            echo -e "\033[32mdone\033[0m"
    else
            echo "Starting ${SERVICENAME}... robot (pid $isStart) already running"
    fi
}

case "$1" in
    'start') app_start;;
    'stop') app_stop;;
    'robot') app_robot;;
    'status') app_status;;
    'reload') 
        app_stop
        app_start;;
    'restart') 
        app_stop
        app_start;;
esac