curl http://localhost:8080/notes/2 \
     --include \
     --header "Content-Type: application/json" \
     --request "PATCH" \
     --data '{"ID": "2", "FileName": "newfile", "Content": "some really epic content"}'
