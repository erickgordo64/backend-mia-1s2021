CREATE SEQUENCE "admin_idadmin_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "admin_idadmin_TRG" BEFORE
    INSERT ON administrador
    FOR EACH ROW
    WHEN ( new.idadmin IS NULL )
BEGIN
    :new.idadmin := "admin_idadmin_SEQ".nextval;
END;
/

CREATE SEQUENCE "chat_idchat_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "chat_idchat_TRG" BEFORE
    INSERT ON chat
    FOR EACH ROW
    WHEN ( new.idchat IS NULL )
BEGIN
    :new.idchat := "chat_idchat_SEQ".nextval;
END;
/

CREATE SEQUENCE "deporte_iddeporte_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "deporte_iddeporte_TRG" BEFORE
    INSERT ON deporte
    FOR EACH ROW
    WHEN ( new.iddeporte IS NULL )
BEGIN
    :new.iddeporte := "deporte_iddeporte_SEQ".nextval;
END;
/

CREATE SEQUENCE "chat_iddetalle_chat_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "chat_iddetalle_chat_TRG" BEFORE
    INSERT ON detalle_chat
    FOR EACH ROW
    WHEN ( new.iddetalle_chat IS NULL )
BEGIN
    :new.iddetalle_chat := "chat_iddetalle_chat_SEQ".nextval;
END;
/

CREATE SEQUENCE "evento_idevento_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "evento_idevento_TRG" BEFORE
    INSERT ON evento
    FOR EACH ROW
    WHEN ( new.idevento IS NULL )
BEGIN
    :new.idevento := "evento_idevento_SEQ".nextval;
END;
/

CREATE SEQUENCE "membresia_idmembresia_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "membresia_idmembresia_TRG" BEFORE
    INSERT ON membresia
    FOR EACH ROW
    WHEN ( new.idmembresia IS NULL )
BEGIN
    :new.idmembresia := "membresia_idmembresia_SEQ".nextval;
END;
/

CREATE SEQUENCE "posicion_idposicion_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "posicion_idposicion_TRG" BEFORE
    INSERT ON posicion
    FOR EACH ROW
    WHEN ( new.idposicion IS NULL )
BEGIN
    :new.idposicion := "posicion_idposicion_SEQ".nextval;
END;
/

CREATE SEQUENCE "prediccion_idprediccion_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "prediccion_idprediccion_TRG" BEFORE
    INSERT ON prediccion
    FOR EACH ROW
    WHEN ( new.idprediccion IS NULL )
BEGIN
    :new.idprediccion := "prediccion_idprediccion_SEQ".nextval;
END;
/

CREATE SEQUENCE "recompensa_idrecompensa_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "recompensa_idrecompensa_TRG" BEFORE
    INSERT ON recompensa
    FOR EACH ROW
    WHEN ( new.idrecompensa IS NULL )
BEGIN
    :new.idrecompensa := "recompensa_idrecompensa_SEQ".nextval;
END;
/

CREATE SEQUENCE "temporada_idtemporada_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "temporada_idtemporada_TRG" BEFORE
    INSERT ON temporada
    FOR EACH ROW
    WHEN ( new.idtemporada IS NULL )
BEGIN
    :new.idtemporada := "temporada_idtemporada_SEQ".nextval;
END;
/

CREATE SEQUENCE "membresia_idtipomembresia_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "membresia_idtipomembresia_TRG" BEFORE
    INSERT ON tipo_membresia
    FOR EACH ROW
    WHEN ( new.idtipo_membresia IS NULL )
BEGIN
    :new.idtipo_membresia := "membresia_idtipomembresia_SEQ".nextval;
END;
/


CREATE SEQUENCE "equipol_idequipo_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "equipol_idequipo_TRG" BEFORE
    INSERT ON equipo_local
    FOR EACH ROW
    WHEN ( new.idequipo IS NULL )
BEGIN
    :new.idequipo := "equipol_idequipo_SEQ".nextval;
END;
/



CREATE SEQUENCE "equipov_idequipo_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "equipov_idequipo_TRG" BEFORE
    INSERT ON equipo_visita
    FOR EACH ROW
    WHEN ( new.idequipo IS NULL )
BEGIN
    :new.idequipo := "equipov_idequipo_SEQ".nextval;
END;
/

CREATE SEQUENCE "jornada_idjornada_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "jornada_idjornada_TRG" BEFORE
    INSERT ON jornada
    FOR EACH ROW
    WHEN ( new.idjornada IS NULL )
BEGIN
    :new.idjornada := "jornada_idjornada_SEQ".nextval;
END;
/

CREATE SEQUENCE "usuario_idusuario_SEQ" START WITH 1 NOCACHE ORDER;

CREATE OR REPLACE TRIGGER "usuario_idusuario_TRG" BEFORE
    INSERT ON usuario
    FOR EACH ROW
    WHEN ( new.idusuario IS NULL )
BEGIN
    :new.idusuario := "usuario_idusuario_SEQ".nextval;
END;
/