curl http://localhost:8080/notes/ \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"FileName": "coolnote2.md", "Content":"#this is another test"}'
