#POST: 
curl -X POST -i -d '{"nome":"Valoir Moura", "email":"valoirmoura@gmail.com"}' -H 'Content-Type: application/json' http://localhost:5000/usuarios

#PUT:
curl -X PUT -i -d '{"nome":"Valoir Moura", "email":"valoirmoura@outlook.com"}' -H 'Content-Type: application/json' http://localhost:5000/usuarios/1

#DELETE:
curl -X DELETE -i http://localhost:5000/usuarios/1

#GET:
curl -X GET -i http://localhost:5000/usuarios
