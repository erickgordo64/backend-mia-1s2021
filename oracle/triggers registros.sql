CREATE OR REPLACE TRIGGER Registro
BEFORE
INSERT OR UPDATE ON usuario
FOR EACH ROW
declare
USUARIOT varchar2(50);
BEGIN
BEGIN
select COUNT(username) INTO USUARIOT FROM usuario WHERE username = :new.username;
END;
 IF INSERTING THEN
if (USUARIOT>0) then
RAISE_APPLICATION_ERROR(-20000, 'El usuario ya existe');
elsif not (regexp_like(:new.correo, '[A-Z0-9._]+@[A-Z0-9.-]+\.[A-Z]','i')) then
RAISE_APPLICATION_ERROR (-20000, 'El correo no es valido ');
end if;
end if;
end;

insert into usuario(username, password, nombre, apellido, fecha_nacimiento, correo) values('fer','1234','fernando','ambrosio',to_date('17/9/1996', 'dd/mm/yyyy'),'fernado18@gmail.com');

select * from usuario;
