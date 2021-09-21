.SHELLFLAGS += -e
ui-setup:
	cd ui && yarn
ui:
	cd ui
	yarn serve
server-setup:
	cd server && go mod vendor
serve:
	cd server && go run ./
test:
	cd server && go test ./...
setup :: ui-setup server-setup
