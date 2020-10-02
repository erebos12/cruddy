#!/usr/bin/env bash

#
#  If you build an image without tagging it, the image will appear on the list of dangling images 
#  because it has no association with a tagged image. 
#  See also https://www.digitalocean.com/community/tutorials/how-to-remove-docker-images-containers-and-volumes
#
# Dangling images are layers that have no relationship to any tagged images. 

# this will remove all dangling images   

CMD=$(docker volume ls -qf dangling=true)


if [[ -z "$CMD" ]]; then 
    echo "Nothing to purge!" 
    exit 0
fi

docker volume rm $(docker volume ls -qf dangling=true)
echo "Purged!" 
