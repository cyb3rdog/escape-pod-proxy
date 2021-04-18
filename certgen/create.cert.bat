
SET CERT_NAME=Cyb3rVector-Extension
SET CONF="escapepod-cert.conf"

copy escapepod-cert.conf.in %CONF%
echo DNS.1       = %CERT_NAME% >> %CONF%

openssl req -config %CONF% -subj "/C=CZ/OU=cyb3rdog/O=PHACT/CN=%CERT_NAME%" -new -x509 -days 36500 -nodes -newkey rsa:2048 -keyout %CERT_NAME%.pem -out %CERT_NAME%.crt

rem del %CONF% 