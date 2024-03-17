curl http://0.0.0.0:8080/users \
     --include \
     --header "Content-Type: application/json" \
     --request "POST" \
     --data '{"Username":"nhefner", "Password": "pass"}'
