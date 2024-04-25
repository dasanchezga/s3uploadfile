MS for upload a song & delete a song in a AWS S3 bucket using GO

docker build -t upload_ms_image .
docker run -dp 8080:8080 upload_ms_image

Requests: 

Upload: http://localhost:8080/upload
Delete: http://localhost:8080/delete 