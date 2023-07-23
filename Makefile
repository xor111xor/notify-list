BIN 					:= notify-list
VERSION 			:= $(shell git describe --tags)
PREFIX 				:= $(shell go env GOPATH)
UNIT_PREFIX 	:= $${HOME}/.config/systemd/user
UNIT_TEMPLATE := ./systemd/${BIN}.service.template

.PHONY: check
check:
	@notify-send Test "This is a check" || (echo "Notify-send not installed $$?"; exit 1)

.PHONY: build
build: 
	@go build -ldflags="-s -w -X main.AppVersion=${VERSION}" ./cmd/${BIN}

.PHONY: install
install: 
	@install  -m744 ${BIN} $(PREFIX)/bin/${BIN}

	@cp ${UNIT_TEMPLATE} ${BIN}.service
	@sed -i "s|ExecStart=|ExecStart=${PREFIX}\/bin\/${BIN}|g" ${BIN}.service
	@install  -m644 ${BIN}.service ${UNIT_PREFIX}

	@systemctl --user daemon-reload
	@systemctl --user enable --now ${BIN}.service

.PHONY: clean
clean: 
	@rm -f ${BIN}
	@rm -f ${BIN}.service

.PHONY: uninstall
uninstall: 
	@systemctl --user disable --now ${BIN}.service
	@rm -f $(PREFIX)/bin/${BIN}
	@rm -f  ${UNIT_PREFIX}/${BIN}.service
	@systemctl --user daemon-reload
