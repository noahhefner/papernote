curl http://0.0.0.0:8080/nhefner/notes/coolnote.md \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"Content":"#this is another test"}'
