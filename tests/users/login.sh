curl http://0.0.0.0:8080/login \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"Username":"nhefner", "Password": "pass"}'
