love:
	@echo "What is love?"

gen:
	@for f in generators/*.go; do go build $$f; done
	@for f in transactions/*; do ./generateHandler "$$f"; done


