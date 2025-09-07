APP_DIR := /Users/charith/Desktop/Golang-Learn/blogger
LOG_FILE := /tmp/blogger_server.log

.PHONY: build run stop logs smoke clean

build:
	cd $(APP_DIR) && go build ./...

run:
	@if lsof -tiTCP:3000 -sTCP:LISTEN >/dev/null; then echo "server already running"; \
	else \
		cd $(APP_DIR) && nohup go run ./cmd/server > $(LOG_FILE) 2>&1 & echo $$! > /tmp/blogger_server.pid && echo "started: $$(cat /tmp/blogger_server.pid)"; \
		sleep 1; \
	fi

stop:
	-@if lsof -tiTCP:3000 -sTCP:LISTEN >/dev/null; then kill $$(lsof -tiTCP:3000 -sTCP:LISTEN); echo stopped; else echo not running; fi
	-@rm -f /tmp/blogger_server.pid

logs:
	@tail -n 100 $(LOG_FILE) | cat

smoke: run
	cd $(APP_DIR) && ./smoke.sh

clean:
	rm -f $(LOG_FILE) /tmp/blogger_cookies.txt /tmp/blogger_server.pid

