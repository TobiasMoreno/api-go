# Guía de Contribución

¡Gracias por tu interés en contribuir a este proyecto! Esta guía te ayudará a entender cómo contribuir de manera efectiva.

## Código de Conducta

- Sé respetuoso y profesional
- Acepta críticas constructivas
- Ayuda a otros cuando sea posible

## Proceso de Contribución

### 1. Fork y Clone

```bash
# Fork el repositorio en GitHub
# Luego clona tu fork
git clone https://github.com/tu-usuario/go-api.git
cd go-api
```

### 2. Crear una Rama

```bash
git checkout -b feature/nombre-de-tu-feature
# o
git checkout -b fix/descripcion-del-bug
```

### 3. Configurar el Entorno

```bash
# Instalar dependencias
make deps

# Instalar herramientas de desarrollo
make install-tools

# Copiar .env.example a .env
cp .env.example .env
```

### 4. Desarrollo

- Sigue los principios SOLID
- Escribe código limpio y legible
- Agrega comentarios cuando sea necesario
- Mantén las funciones pequeñas y enfocadas

### 5. Testing

```bash
# Ejecutar tests
make test

# Ver cobertura
make test-coverage
```

### 6. Linting

```bash
# Verificar código
make lint

# Corregir automáticamente
make lint-fix
```

### 7. Commit

Usa mensajes de commit descriptivos siguiendo [Conventional Commits](https://www.conventionalcommits.org/):

```
feat: agregar autenticación JWT
fix: corregir error en validación de email
docs: actualizar README
refactor: mejorar estructura del repositorio
test: agregar tests para servicio de usuarios
```

### 8. Push y Pull Request

```bash
git push origin feature/nombre-de-tu-feature
```

Luego crea un Pull Request en GitHub con:
- Descripción clara de los cambios
- Referencias a issues relacionados (si aplica)
- Screenshots (si es relevante)

## Estándares de Código

### Estilo

- Usa `gofmt` o `goimports` para formatear
- Sigue las convenciones de Go
- Usa nombres descriptivos

### Estructura

- Mantén funciones pequeñas (< 50 líneas idealmente)
- Una responsabilidad por función (SRP)
- Usa interfaces para desacoplar

### Testing

- Escribe tests para nueva funcionalidad
- Mantén cobertura > 80%
- Usa nombres descriptivos para tests

## Checklist antes de PR

- [ ] Código compila sin errores
- [ ] Tests pasan (`make test`)
- [ ] Linter pasa (`make lint`)
- [ ] Documentación actualizada
- [ ] .env.example actualizado si hay nuevas variables
- [ ] Commits siguen Conventional Commits

## Preguntas?

Si tienes preguntas, abre un issue o contacta a los mantenedores.

¡Gracias por contribuir! 🎉

