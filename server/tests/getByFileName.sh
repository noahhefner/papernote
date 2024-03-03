curl http://localhost:8080/notes/ \
     --include \
     --header "Content-Type: application/json" \
     --request "GET" \
     --data '{"FileName": "coolnote.md"}'
