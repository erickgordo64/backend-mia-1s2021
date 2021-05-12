-- Generado por Oracle SQL Developer Data Modeler 21.1.0.092.1221
--   en:        2021-05-06 22:22:00 CST
--   sitio:      Oracle Database 11g
--   tipo:      Oracle Database 11g



-- predefined type, no DDL - MDSYS.SDO_GEOMETRY

-- predefined type, no DDL - XMLTYPE

CREATE TABLE administrador (
    idadmin        INTEGER NOT NULL,
    usuario_admin  VARCHAR2(150 BYTE),
    contrasena     VARCHAR2(150 BYTE)
);

ALTER TABLE administrador ADD CONSTRAINT administrador_pk PRIMARY KEY ( idadmin );

CREATE TABLE chat (
    idchat                 INTEGER NOT NULL,
    usuario_idusuario      NUMBER NOT NULL,
    administrador_idadmin  INTEGER NOT NULL
);

ALTER TABLE chat ADD CONSTRAINT chat_pk PRIMARY KEY ( idchat );

CREATE TABLE deporte (
    iddeporte       INTEGER NOT NULL,
    nombre_deporte  VARCHAR2(100 BYTE),
    imagen          VARCHAR2(250 BYTE),
    color           VARCHAR2(50 BYTE)
);

ALTER TABLE deporte ADD CONSTRAINT deporte_pk PRIMARY KEY ( iddeporte );

CREATE TABLE detalle_chat (
    iddetalle_chat         INTEGER NOT NULL,
    contenido_mensaje      VARCHAR2(260 BYTE),
    fecha_chat             DATE,
    usuario_idusuario      NUMBER NOT NULL,
    administrador_idadmin  INTEGER NOT NULL,
    chat_chat_id           NUMBER NOT NULL,
    idchat                 INTEGER NOT NULL,
    chat_idchat            INTEGER NOT NULL
);

ALTER TABLE detalle_chat ADD CONSTRAINT detalle_chat_pk PRIMARY KEY ( iddetalle_chat );

CREATE TABLE equipo_local (
    idequipo       INTEGER NOT NULL,
    nombre_equipo  VARCHAR2(150 BYTE)
);

ALTER TABLE equipo_local ADD CONSTRAINT equipo_local_pk PRIMARY KEY ( idequipo );

CREATE TABLE equipo_visita (
    idequipo       INTEGER NOT NULL,
    nombre_equipo  VARCHAR2(150 BYTE)
);

ALTER TABLE equipo_visita ADD CONSTRAINT equipo_visita_pk PRIMARY KEY ( idequipo );

CREATE TABLE evento (
    idevento                INTEGER NOT NULL,
    fecha_inicio_evento     DATE,
    fecha_fin_evento        DATE,
    color                   VARCHAR2(50 BYTE),
    resultado_local         INTEGER,
    resultado_visitante     INTEGER,
    jornada_idjornada       INTEGER NOT NULL,
    equipo_local_idequipo   INTEGER NOT NULL,
    equipo_visita_idequipo  INTEGER NOT NULL,
    deporte_iddeporte       INTEGER NOT NULL
);

ALTER TABLE evento ADD CONSTRAINT evento_pk PRIMARY KEY ( idevento );

CREATE TABLE jornada (
    idjornada              INTEGER NOT NULL,
    nombre_jornada         VARCHAR2(100 BYTE),
    fecha_inicio           DATE,
    fecha_fin              DATE,
    estado                 VARCHAR2(100 BYTE),
    temporada_idtemporada  INTEGER NOT NULL
);

ALTER TABLE jornada ADD CONSTRAINT jornada_pk PRIMARY KEY ( idjornada );

CREATE TABLE masiva (
    id                    VARCHAR2(100 BYTE),
    nombre_cliente        VARCHAR2(100 BYTE),
    apellido_cliente      VARCHAR2(100 BYTE),
    password              VARCHAR2(150 BYTE),
    username              VARCHAR2(150 BYTE),
    temporada             VARCHAR2(100 BYTE),
    tier                  VARCHAR2(100),
    jornada               VARCHAR2(100),
    deporte               VARCHAR2(100),
    fecha                 DATE,
    visitante_nombre      VARCHAR2(100 BYTE),
    local_nombre          VARCHAR2(250),
    visitante_prediccion  INTEGER,
    local_prediccion      INTEGER,
    visitante_resultado   INTEGER,
    local_resultado       INTEGER
);

CREATE TABLE membresia (
    idmembresia                      INTEGER NOT NULL,
    fecha_inicio                     DATE,
    fecha_fin                        DATE,
    estado_membresia                 VARCHAR2(100 BYTE), 
    membresia_idtipo_membresia       INTEGER NOT NULL,
    temporada_idtemporada            INTEGER NOT NULL,
    usuario_idusuario                NUMBER NOT NULL
);

ALTER TABLE membresia ADD CONSTRAINT membresia_pk PRIMARY KEY ( idmembresia );

CREATE TABLE posicion (
    idposicion             INTEGER NOT NULL,
    posicion               INTEGER,
    punteo                 INTEGER,
    p10                    INTEGER,
    p5                     INTEGER,
    p3                     INTEGER,
    p0                     INTEGER,
    incremento             INTEGER,
    temporada_idtemporada  INTEGER NOT NULL,
    usuario_idusuario      NUMBER NOT NULL
);

ALTER TABLE posicion ADD CONSTRAINT posicion_pk PRIMARY KEY ( idposicion );

CREATE TABLE prediccion (
    idprediccion       INTEGER NOT NULL,
    local              INTEGER,
    visitante          INTEGER,
    puntaje            INTEGER,
    usuario_idusuario  NUMBER NOT NULL,
    evento_idevento    INTEGER NOT NULL
);

ALTER TABLE prediccion ADD CONSTRAINT prediccion_pk PRIMARY KEY ( idprediccion );

CREATE TABLE recompensa (
    idrecompensa           INTEGER NOT NULL,
    recompensa             NUMBER,
    ultima_recompenza      NUMBER,
    incremento             NUMBER,
    temporada_idtemporada  INTEGER NOT NULL,
    usuario_idusuario      NUMBER NOT NULL
);

ALTER TABLE recompensa ADD CONSTRAINT recompensa_pk PRIMARY KEY ( idrecompensa );

CREATE TABLE temporada (
    idtemporada       INTEGER NOT NULL,
    nombre_temporada  VARCHAR2(150 BYTE),
    fecha_inicio      DATE,
    fecha_fin         DATE NOT NULL,
    estado_temporada  VARCHAR2(100 BYTE)
);

ALTER TABLE temporada ADD CONSTRAINT temporada_pk PRIMARY KEY ( idtemporada );

CREATE TABLE tipo_membresia (
    idtipo_membresia  INTEGER NOT NULL,
    tipo_membresia    VARCHAR2(100 BYTE),
    varlor_membresia  INTEGER
);

ALTER TABLE tipo_membresia ADD CONSTRAINT tipo_membresia_pk PRIMARY KEY ( idtipo_membresia );

CREATE TABLE usuario (
    idusuario         NUMBER NOT NULL,
    codigo_usuario    VARCHAR2(100 BYTE),
    username          VARCHAR2(100 BYTE),
    password          VARCHAR2(250 BYTE),
    nombre            VARCHAR2(100 BYTE),
    apellido          VARCHAR2(100 BYTE),
    fecha_nacimiento  DATE,
    fecha_registro    DATE,
    correo            VARCHAR2(150 BYTE)
);

ALTER TABLE usuario ADD CONSTRAINT usuario_pk PRIMARY KEY ( idusuario );

ALTER TABLE chat
    ADD CONSTRAINT chat_administrador_fk FOREIGN KEY ( administrador_idadmin )
        REFERENCES administrador ( idadmin );

ALTER TABLE chat
    ADD CONSTRAINT chat_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );

ALTER TABLE detalle_chat
    ADD CONSTRAINT detalle_chat_administrador_fk FOREIGN KEY ( administrador_idadmin )
        REFERENCES administrador ( idadmin );

ALTER TABLE detalle_chat
    ADD CONSTRAINT detalle_chat_chat_fk FOREIGN KEY ( chat_idchat )
        REFERENCES chat ( idchat );

ALTER TABLE detalle_chat
    ADD CONSTRAINT detalle_chat_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );

ALTER TABLE evento
    ADD CONSTRAINT evento_deporte_fk FOREIGN KEY ( deporte_iddeporte )
        REFERENCES deporte ( iddeporte );

ALTER TABLE evento
    ADD CONSTRAINT evento_equipo_local_fk FOREIGN KEY ( equipo_local_idequipo )
        REFERENCES equipo_local ( idequipo );

ALTER TABLE evento
    ADD CONSTRAINT evento_equipo_visita_fk FOREIGN KEY ( equipo_visita_idequipo )
        REFERENCES equipo_visita ( idequipo );

ALTER TABLE evento
    ADD CONSTRAINT evento_jornada_fk FOREIGN KEY ( jornada_idjornada )
        REFERENCES jornada ( idjornada );

ALTER TABLE jornada
    ADD CONSTRAINT jornada_temporada_fk FOREIGN KEY ( temporada_idtemporada )
        REFERENCES temporada ( idtemporada );

ALTER TABLE membresia
    ADD CONSTRAINT membresia_temporada_fk FOREIGN KEY ( temporada_idtemporada )
        REFERENCES temporada ( idtemporada );

ALTER TABLE membresia
    ADD CONSTRAINT membresia_tipo_membresia_fk FOREIGN KEY ( membresia_idtipo_membresia )
        REFERENCES tipo_membresia ( idtipo_membresia );

ALTER TABLE membresia
    ADD CONSTRAINT membresia_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );

ALTER TABLE posicion
    ADD CONSTRAINT posicion_temporada_fk FOREIGN KEY ( temporada_idtemporada )
        REFERENCES temporada ( idtemporada );

ALTER TABLE posicion
    ADD CONSTRAINT posicion_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );

ALTER TABLE prediccion
    ADD CONSTRAINT prediccion_evento_fk FOREIGN KEY ( evento_idevento )
        REFERENCES evento ( idevento );

ALTER TABLE prediccion
    ADD CONSTRAINT prediccion_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );

ALTER TABLE recompensa
    ADD CONSTRAINT recompensa_temporada_fk FOREIGN KEY ( temporada_idtemporada )
        REFERENCES temporada ( idtemporada );

ALTER TABLE recompensa
    ADD CONSTRAINT recompensa_usuario_fk FOREIGN KEY ( usuario_idusuario )
        REFERENCES usuario ( idusuario );



-- Informe de Resumen de Oracle SQL Developer Data Modeler: 
-- 
-- CREATE TABLE                            16
-- CREATE INDEX                             0
-- ALTER TABLE                             34
-- CREATE VIEW                              0
-- ALTER VIEW                               0
-- CREATE PACKAGE                           0
-- CREATE PACKAGE BODY                      0
-- CREATE PROCEDURE                         0
-- CREATE FUNCTION                          0
-- CREATE TRIGGER                           0
-- ALTER TRIGGER                            0
-- CREATE COLLECTION TYPE                   0
-- CREATE STRUCTURED TYPE                   0
-- CREATE STRUCTURED TYPE BODY              0
-- CREATE CLUSTER                           0
-- CREATE CONTEXT                           0
-- CREATE DATABASE                          0
-- CREATE DIMENSION                         0
-- CREATE DIRECTORY                         0
-- CREATE DISK GROUP                        0
-- CREATE ROLE                              0
-- CREATE ROLLBACK SEGMENT                  0
-- CREATE SEQUENCE                          0
-- CREATE MATERIALIZED VIEW                 0
-- CREATE MATERIALIZED VIEW LOG             0
-- CREATE SYNONYM                           0
-- CREATE TABLESPACE                        0
-- CREATE USER                              0
-- 
-- DROP TABLESPACE                          0
-- DROP DATABASE                            0
-- 
-- REDACTION POLICY                         0
-- 
-- ORDS DROP SCHEMA                         0
-- ORDS ENABLE SCHEMA                       0
-- ORDS ENABLE OBJECT                       0
-- 
-- ERRORS                                   1
-- WARNINGS                                 0
