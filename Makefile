.venv:
	python3 -m venv .venv

.PHONY: scan
scan: .venv
	. ./.venv/bin/activate \
	&& pip3 install -r requirements.txt \
	&& tartufo -v scan-local-repo .
