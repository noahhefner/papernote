curl http://localhost:8080/notes/ \
     --include \
     --header "Content-Type: application/json" \
     --request "PATCH" \
     --data '{"FileName": "coolnote.md", "Content": "this content will replace the file content"}'
