PROJECT_NAME=template
IMAGE_NAME=${USER}_${PROJECT_NAME}
CONTAINER_NAME=${USER}_${PROJECT_NAME}
PORT=8885
SHM_SIZE=2g
FORCE_RM=true
build_python_linux:
	docker build \
		--build-arg USER_ID=$(shell id -u) \
		--build-arg GROUP_ID=$(shell id -g) \
		-f python_3_10/linux/Dockerfile \
		-t unilab_train \
		--force-rm=${FORCE_RM}\
		.
restart: stop run
run_python:
	docker run \
		-dit \
		-v $(PWD):/workspace \
		-p 8883:8883 \
		--name unilab_train \
		--rm \
		--shm-size $(SHM_SIZE) \
		unilab_train
exec_python:
	docker exec \
		-it \
		unilab_train bash
stop:
	docker stop $(IMAGE_NAME)
run_jupyter:
	jupyter nbextension enable --py widgetsnbextension
	jupyter notebook --ip=0.0.0.0 --port ${PORT} --allow-root