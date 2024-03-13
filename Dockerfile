# Utilise une image de base officielle de Go avec Debian
FROM golang:latest

# Installe Graphviz
RUN apt-get update && apt-get install -y graphviz

# Copie votre code source dans le conteneur (ajustez le chemin selon votre structure de projet)
COPY . /app

# Définit le répertoire de travail
WORKDIR /app
