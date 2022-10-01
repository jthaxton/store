# store

curl -X 'POST' \
  'http://0.0.0.0:8080/create?customId=https://www.youtube.com/watch?v=test' \
   -d '{}' 

curl -X 'GET' \ 
  'http://0.0.0.0:8080/find?customId=https://www.youtube.com/watch?v=test'