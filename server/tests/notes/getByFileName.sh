curl http://0.0.0.0:8080/notes \
     --include \
     --header "Content-Type: application/json" \
     --request "GET" \
     --data '{"FileName": "ACOOLFILE.md"}'
