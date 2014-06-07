#!/bin/sh
trap 'killall' INT TERM
 

killall() {
    trap '' INT TERM     # ignore INT and TERM while shutting down
    echo "**** Shutting down... ****"     # added double quotes
    kill -TERM 0         # fixed order, send TERM not INT
    wait
    echo DONE
}

./stinger -config="$STINGER_CONFIG" -listen="$STINGER_LISTEN" &
wait 