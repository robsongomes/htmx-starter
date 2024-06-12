.PHONY: start stop

start:
	@echo "Starting process"
	@templ generate -watch -proxy="http://localhost:3000" & echo $$! > p1.pid & air & echo $$! > p2.pid &

stop:
	@echo "Stopping process"
	@kill `cat p1.pid` || true
	@kill `cat p2.pid` || true
	@rm -f p1.pid p2.pid