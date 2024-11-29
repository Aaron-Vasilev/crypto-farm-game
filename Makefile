run:
	@templ generate
	@go build -o ./tmp/main ./src/main.go
dev:
	@npx concurrently "air" "npx tailwindcss -o ./public/styles/out.css --watch"
format:
	@gofmt -w .
	@templ fmt .
	@./tailwind_class_order.sh
start:
	@supervisord -c ./supervisord.conf
launch:
	@sudo supervisorctl shutdown
	@go build -o ./tmp/bot ./main.go
	@npx tailwindcss -o ./public/styles/out.css
	@echo Build ends
	@sudo supervisord -c ./supervisord.conf
	@echo Started
stop:
	@supervisorctl shutdown

