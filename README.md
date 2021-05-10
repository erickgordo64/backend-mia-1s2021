# MANUAL TECNICO/ QUINIELA APP
---------------------------------------------------------------------------
## UNIVERSIDAD SAN CARLOS DE GUATEMALA
## FACULTAD DE INGENIERIA 
## MANEJO E IMPLEMENTACION DE ARCHIVOS

### ERICK VALENZUELA
---------------------------------------------------------------------------

### Lenguajes y herramientas utlizadas proyecto
- Para el backend se utilizo el lenguaje de programación Go (Golang).
- Para el frontend se utilizo el framework web React.
    - HTML5 y CSS para que el sitio sea funcional, atractivo y refleja la imagen de cada entidad.
- Base de datos oracle
---------------------------------------------------------------------------
# REACT
React es una librería Javascript focalizada en el desarrollo de interfaces de usuario. Así se define la propia librería y evidentemente, esa es su principal área de trabajo. Sin embargo, lo cierto es que en React encontramos un excelente aliado para hacer todo tipo de aplicaciones web, SPA (Single Page Application) o incluso aplicaciones para móviles. Para ello, alrededor de React existe un completo ecosistema de módulos, herramientas y componentes capaces de ayudar al desarrollador a cubrir objetivos avanzados con relativamente poco esfuerzo. 
![myimage-alt-tag](https://www.programacion.com.py/wp-content/uploads/2016/11/react-logo-1024x576.png)
## DESCRIPCION DEL PROBLEMA
- se nos pidio realizar un solucion para el app quiniela, dicha app consta de diferentes modulos y roles de usuario que se tendran que encontrar soluciones para que no se mezclen las diferentes paginas realizadas para cada uno de estos roles.\
La distribucion de la de los ficheros quedo de la siguiente forma
```
    \Administrador
        \inicio
            -Capital de temporada: Muestra el total de capital para la temporada actual que será puesto en juego.
            -Bronze: Número total de clientes con membresı́a bronze activa en la temporada actual.
            -Silver: Número total de clientes con membresı́a silver activa en la temporada actual.
            -Gold: Número total de clientes con membresı́a gold activa en la temporada actual.
    \Cliente
        \Registro cliente
        \login
        \recuperar contraseña
        \perfil de usuario
        \pagar membresia
        \ingresar predicciones
        \navegar en eventos deportivos
            - vista mensual
            - vista semanal
             -- codigo de colores
                --- gris: evento ya ocurrio
                --- morado: evento no a ocurrido pero ya se realizo la prediccion
                --- otros: colores propios por cada deporte
            - restriccion de intento
        \resultados
        \tabla de posiciones
        \recompensas

```

### librerias utilizadas en react
- axios
```
    npm install axios
```
- universal cookies 
```
    install univeral-cookies
```
- react router dom 
```
    intall reat-router-dom
```
- big calendar 
```
    npm install big-calender
```
### librerias utilizadas golang
- [github.com/godror/godror](https://github.com/godror/godror)
    - libreria que nos permite realizar la conexion entre nuestro api en go con orale 
 ```
    go get -u github.com/godror/godror
```   

- [github.com/gorilla/mux](https://github.com/gorilla/mux)
    - libreria que no permite volver nuestro app en go en un api rest
 ```
    go get -u github.com/gorilla/mux
```   
