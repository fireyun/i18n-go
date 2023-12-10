.PHONY: init
init:
	@echo Download gotext
	go install golang.org/x/text/cmd/gotext@v0.14.0

.PHONY: i18n
i18n:
	@go generate ./pkg/i18n/translations/translations.go
	@cp ./pkg/i18n/translations/locales/zh/out.gotext.json ./pkg/i18n/translations/locales/zh/messages.gotext.json