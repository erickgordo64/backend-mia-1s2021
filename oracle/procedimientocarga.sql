create or replace PROCEDURE CARGAMASIVA AS 
BEGIN
insert into usuario(codigo_usuario, username, password, nombre, apellido, correo)
select distinct id, username, password, nombre_cliente, apellido_cliente, username from masiva;
-------------------------membresia----------------------------
insert into tipo_membresia(tipo_membresia)
select distinct tier from masiva;
-----------------DEPORTE
insert into deporte(nombre_deporte)
select distinct deporte from masiva;
--------TERMPORADA
insert into temporada(nombre_temporada, fecha_inicio, fecha_fin,estado_temporada) 
select distinct Temporada,TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(1,'/'),SUBSTR(Temporada,7,7)),'/'),SUBSTR(Temporada,0,4)),'DD/MM/YYYY'), 
TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(28,'/'),SUBSTR(Temporada,7,7)),'/'),SUBSTR(Temporada,0,4)),'DD/MM/YYYY'),'Finalizada' FROM masiva;        
---- EQUIPOS
insert into equipo_local(nombre_equipo)
select distinct local_nombre from masiva;
insert into equipo_visita(nombre_equipo)
select distinct visitante_nombre from masiva;
----------------------------jorandas ------------------------------------
insert into jornada(nombre_jornada, fecha_inicio, fecha_fin, estado, temporada_idtemporada)
SELECT DISTINCT JORNADA, TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(1,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY'),
TO_DATE(TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(1,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY')+7),'FINALIZADA',temporada.idtemporada
FROM masiva
inner join Temporada on masiva.TEMPORADA = Temporada.NOMBRE_TEMPORADA
where JORNADA ='J1'
union 
SELECT DISTINCT JORNADA, TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(8,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY'),
TO_DATE(TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(8,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY')+7),'FINALIZADA',temporada.idtemporada
FROM masiva
inner join Temporada on masiva.TEMPORADA = Temporada.NOMBRE_TEMPORADA
where JORNADA ='J2'
union
SELECT DISTINCT JORNADA, TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(15,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY'),
TO_DATE(TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(15,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY')+7),'FINALIZADA',temporada.idtemporada
FROM masiva
inner join Temporada on masiva.TEMPORADA = Temporada.NOMBRE_TEMPORADA
where JORNADA ='J3'
union
SELECT DISTINCT JORNADA, TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(22,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY'),
TO_DATE(TO_DATE(CONCAT(CONCAT(CONCAT(CONCAT(22,'/'),SUBSTR(nombre_temporada,7,7)),'/'),SUBSTR(nombre_temporada,0,4)),'DD/MM/YYYY')+7),'FINALIZADA',temporada.idtemporada
FROM masiva
inner join Temporada on masiva.TEMPORADA = Temporada.NOMBRE_TEMPORADA
where JORNADA ='J4';
----------------------------eventos-----------------------
insert into evento(fecha_inicio_evento, resultado_local, resultado_visitante, jornada_idjornada, equipo_local_idequipo, equipo_visita_idequipo, deporte_iddeporte)
select distinct masiva.fecha, masiva.local_resultado, masiva.visitante_resultado, jornada.idjornada,
equipo_local.idequipo, equipo_visita.idequipo, deporte.iddeporte from jornada, masiva, temporada, equipo_visita, equipo_local, deporte
where temporada.nombre_temporada = masiva.temporada and jornada.temporada_idtemporada=temporada.idtemporada
and masiva.fecha>jornada.fecha_inicio and masiva.fecha<jornada.fecha_fin
and masiva.visitante_nombre=equipo_visita.nombre_equipo and masiva.local_nombre=equipo_local.nombre_equipo
and deporte.nombre_deporte=masiva.deporte;
------------------------------predicciones---------------------------
insert into prediccion(local,visitante,usuario_idusuario,evento_idevento)
select masiva.local_prediccion, visitante_prediccion, usuario.idusuario, evento.idevento from masiva
inner join evento on evento.fecha_inicio_evento=masiva.fecha
inner join usuario on usuario.codigo_usuario=masiva.id;
----------------------------membresia---------------------------------------
insert into membresia(fecha_inicio,fecha_fin,membresia_idtipo_membresia,temporada_idtemporada,usuario_idusuario)
select distinct temporada.fecha_inicio, temporada.fecha_fin, tipo_membresia.idtipo_membresia, temporada.idtemporada, usuario.idusuario from masiva
inner join temporada on masiva.temporada=temporada.nombre_temporada
inner join usuario on masiva.id=usuario.codigo_usuario
inner join tipo_membresia on masiva.tier=tipo_membresia.tipo_membresia;

END CARGAMASIVA;