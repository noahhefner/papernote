curl http://0.0.0.0:8080/notes \
     --include \
     --header "Content-Type: application/json" \
     --request "PATCH" \
     --data '{"FileName": "ACOOLNOTE.md", "Content": "boopdiddy scoop"}'
