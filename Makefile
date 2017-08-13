love:
	@echo "What is love?"

gen:
	@cd generators/; go build
	@for f in transactions/*; do ./generators/generators "Handler" "$$f"; done
	@./generators/generators "Main" transactions/
