#!/bin/bash

root=$(pwd -P $0)
PROGNAME=$(basename $0)

usage () {
    echo "usage: ${PROGNAME} [-f file | -i]"
    return
}

filename=
interactive=

if [ $# -eq 0 ];then
  echo "${PROGNAME}:  need params."
  usage
  exit 1
fi

while [[ -n $1 ]];do
    case $1 in
    -f | --file)
        shift
        filename=$1
        ;;
    -h | --help)
        usage
        exit 
        ;;
    -i | --intercative)
        shift
        interactive=$1
        ;;
    *)
        usage >&2
        exit 1
        ;;
    esac
    shift
done

echo filename=$filename,interactive=$interactive