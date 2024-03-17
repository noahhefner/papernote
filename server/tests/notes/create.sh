curl http://0.0.0.0:8080/notes \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"FileName": "ACOOLFILE.md", "Content":"#this is another test"}'
