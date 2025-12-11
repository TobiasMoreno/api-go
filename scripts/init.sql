-- Script de inicialización de la base de datos
-- Este script se ejecuta automáticamente cuando MySQL se inicia por primera vez

-- Crear la base de datos si no existe (aunque ya se crea con MYSQL_DATABASE)
CREATE DATABASE IF NOT EXISTS usersdb CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

-- Usar la base de datos
USE usersdb;

-- La tabla se crea automáticamente por el repositorio MySQL
-- pero podemos agregar índices adicionales si es necesario

-- Crear índice adicional para búsquedas por email (si no existe)
-- CREATE INDEX IF NOT EXISTS idx_email ON users(email);

