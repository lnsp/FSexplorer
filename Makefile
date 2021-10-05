all:
	cd client && yarn run generate
	go build -o fsexplorer