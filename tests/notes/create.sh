curl http://0.0.0.0:8080/notes \
     --include \
     --header "Content-Type: application/json" \
     --header "Authorization: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im5oZWZuZXIiLCJleHAiOjE3MTA3NTk0NDh9.k_Uec6d92roND7McpicrYS4hT3vKK_KQcwtpadJiXl4" \
     --request "POST" \
     --data '{"FileName": "ACOOLFILE.md", "Content":"#this is another test"}'
