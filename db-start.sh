###################################################
## JUST FOR ISOLATED DEV                          #
###################################################

CONTAINER_NAME='valhallablog-db-isolated'
# IMAGE_NAME='valhallablog-db'
IMAGE_TAG='12-alpine'

cont_exist=`docker ps -aqf name=${CONTAINER_NAME}`

if [ -z $cont_exist ]; then
  echo "Creating container"
  docker run -d --name ${CONTAINER_NAME} \
    --env-file "$PWD/.env" \
    -p "5432:5432" \
    -v ${PWD}/db:/var/lib/postgresql/data \
    postgres:${IMAGE_TAG}
else
  echo "Running ${CONTAINER_NAME} ${IMAGE_TAG}"
  docker start ${CONTAINER_NAME}
fi
