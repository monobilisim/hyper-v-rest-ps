#!/usr/bin/env bash

relpaths() {
    for arg in "$@"; do printf -- "%s\n" "${arg#$PWD/}"; done
}

bindpaths() {
    for arg in $(relpaths "$@"); do
        if [ -e "$arg" ] && [[ "$arg" == /* ]]; then
            printf -- "-v %s:%s" "$arg" "$arg"
        fi
    done
}

case $1 in
	32)
		exec docker run --rm -i --security-opt label=disable -v "$PWD":/work $(bindpaths "$@") -e WINEDEBUG=-all inno:32bit $(relpaths "${@:2}")
	;;
	64)
		exec docker run --rm -i --security-opt label=disable -v "$PWD":/work $(bindpaths "$@") -e WINEDEBUG=-all inno:64bit $(relpaths "${@:2}")
	;;
	*)
		exit 1
	;;
esac

chown -R $USER: $PWD/
