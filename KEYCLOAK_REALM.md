# 🔑 Configuración del Realm en Keycloak

## 🏰 Información del Realm

- **Nombre del Realm**: `intercors-realm`
- **URL del Servidor Keycloak**: `http://localhost:8081`
- **Cliente**: `intercors-client`
- **Protocolo**: `openid-connect`

## 🔄 Rutas JWT del Realm

### 1. Endpoint de Configuración
```
GET /realms/intercors-realm/.well-known/openid-configuration
```

### 2. Endpoint de Autenticación
```
POST /realms/intercors-realm/protocol/openid-connect/token
```

### 3. Endpoint de UserInfo
```
GET /realms/intercors-realm/protocol/openid-connect/userinfo
```

### 4. Endpoint de Cierre de Sesión
```
POST /realms/intercors-realm/protocol/openid-connect/logout
```

### 5. Endpoint de Validación de Token
```
POST /realms/intercors-realm/protocol/openid-connect/token/introspect
```

## 🔧 Configuración del Cliente

1. **Client ID**: `intercors-client`
2. **Client Protocol**: `openid-connect`
3. **Access Type**: `confidential`
4. **Valid Redirect URIs**: `*` (en producción, reemplazar con los dominios correctos)
5. **Web Origins**: `*` (en producción, limitar a los dominios permitidos)

## 🔐 Configuración de Mappers (Claims JWT)

1. **Mappers por defecto**:
   - `aud` (audience)
   - `sub` (subject)
   - `iss` (issuer)
   - `preferred_username`
   - `email`
   - `email_verified`
   - `name`
   - `given_name`
   - `family_name`

## 🔄 Flujo de Autenticación

1. **Obtener Token**:
   ```bash
   curl -X POST 'http://localhost:8081/realms/intercors-realm/protocol/openid-connect/token' \
     -H 'Content-Type: application/x-www-form-urlencoded' \
     -d 'client_id=intercors-client' \
     -d 'username=admin' \
     -d 'password=your-password' \
     -d 'grant_type=password' \
     -d 'client_secret=your-client-secret'
   ```

2. **Usar Token**:
   ```bash
   curl -X GET 'http://localhost:8080/api/protected' \
     -H 'Authorization: Bearer your-jwt-token'
   ```

## 🔒 Configuración de Seguridad

- **Soporte para HTTPS**: Habilitado
- **Access Token Lifespan**: 5 minutos (configurable)
- **Refresh Token Lifespan**: 60 minutos (configurable)
- **SSO Session Idle**: 30 minutos
- **SSO Session Max**: 10 horas

## 🛠️ Configuración Avanzada

### 1. Configuración de Realm
```json
{
  "realm": "intercors-realm",
  "enabled": true,
  "sslRequired": "external",
  "registrationAllowed": false,
  "loginWithEmailAllowed": true,
  "duplicateEmailsAllowed": false,
  "resetPasswordAllowed": true,
  "editUsernameAllowed": false,
  "bruteForceProtected": true
}
```

### 2. Configuración del Cliente
```json
{
  "clientId": "intercors-client",
  "enabled": true,
  "clientAuthenticatorType": "client-secret",
  "redirectUris": ["*"],
  "webOrigins": ["*"],
  "bearerOnly": false,
  "consentRequired": false,
  "standardFlowEnabled": true,
  "implicitFlowEnabled": false,
  "directAccessGrantsEnabled": true,
  "serviceAccountsEnabled": true,
  "publicClient": false,
  "frontchannelLogout": false,
  "protocol": "openid-connect"
}
```

## 🔍 Solución de Problemas

### Error: Invalid token
- Verifica que el token no haya expirado
- Verifica la firma del token
- Verifica el issuer (iss) y audience (aud)

### Error: Unauthorized
- Verifica las credenciales del cliente
- Verifica los permisos del usuario
- Verifica los roles asignados

## 📝 Notas de Implementación

- El token JWT incluye los claims estándar de OpenID Connect
- Se recomienda implementar refresh tokens para una mejor experiencia de usuario
- En producción, siempre usar HTTPS
- Rotar los client secrets regularmente
- Monitorear los logs de Keycloak para detectar intentos de acceso no autorizados
