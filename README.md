# go-circuit-breaker
Implemetación de Patron Circuit Breaker en Go usando https://github.com/sony/gobreaker en un servidor que actua como middleware
redireccionando las peticiones a la API deseada y cortando el flujo al darse las condiciones para abrir el circuito.

# How to
Para probar el patrón se utilizará la api melisearch (https://github.com/mauryparra/melisearch) 
del cual se usará el endpoint http://localhost:8080/ping el cual se modifico para que devuelva un error 500 con una probabilidad del 50%

En otro terminal se ejecutara el proyecto actual el cual correra en http://localhost:8181 y llamando al endpoint /cbmiddle?req={URL}
reemplazando {URL} por http://localhost:8080/ping que seria la API final de consulta.

Por consola se podran ver los detalles de cantidad de request que se respondieron de forma correcta, incorrecta, el ratio de fallas
y el estado del circuit breaker cuando se encuentra en condicion 'abierto'


# Integrantes:
* Mauricio Parra Casado
* Marcos Lopez
* Rodrigo Vicente
