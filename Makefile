deploy:
	go build -tags=netgo -ldflags="-extldflags=-static" -o app.run .
	docker build -t gcr.io/***/csv2xls:latest .
	docker push gcr.io/***/csv2xls:latest
	gcloud run deploy csv2xls \
	--image=gcr.io/***/csv2xls:latest \
	--service-account=health-diary@***.iam.gserviceaccount.com \
	--region=us-central1 \
	--concurrency=100 \
	--cpu=1 \
	--memory=1Gi \
	--execution-environment=gen2 \
	--max-instances=1 \
	--timeout=15 \
	--project=*** \
	&& gcloud run services update-traffic csv2xls --to-latest --region=us-central1