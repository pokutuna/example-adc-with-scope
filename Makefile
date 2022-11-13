PROJECT := pokutuna-dev
GCLOUD := gcloud --project $(PROJECT)

SPREADSHEET_ID := 1kXYux-TGOwf99lLPBOFAnrcLrsuv3R0BcJbiwTb74go
SHEET_RANGE := Sheet1!A2:C2

.PHONY: application-default-credentials
application-default-credentials:
	$(GCLOUD) auth application-default login --scopes=https://www.googleapis.com/auth/cloud-platform,https://www.googleapis.com/auth/spreadsheets


.PHONY: deploy
deploy: deploy-js

.PHONY: deploy-js
deploy-js:
	$(GCLOUD) functions deploy example-adc-spreadsheet-js \
		--source=./js \
		--runtime=nodejs16 \
		--entry-point=app \
		--trigger-http \
		--allow-unauthenticated \
		--region=asia-northeast1 \
		--set-env-vars=SPREADSHEET_ID=$(SPREADSHEET_ID) \
		--set-env-vars=SHEET_RANGE=$(SHEET_RANGE)
