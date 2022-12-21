# Gist clone

## Table of contents

* [Description](#description)
* [Tech stack](#tech-stack)
* [How to run](#how-to-run)

## Description

This project is the backend part of Github Gist clone.\
The purpose of creating this project is to learn backend development in the Go language.

## Tech stack

* Go 1.19.4
* Gin 1.8.1
* MongoDB 6.0.3
* Viper 1.14.0

## How to run

1. Install Docker, Docker Compose, Go (≥ 1.19.4), GNU Make
2. Run following commands to get the ability to use Docker as non-root user:
    ```
    groupadd docker
    usermod -aG docker $(id -un)
    systemctl enable docker
    ```
3. Reboot your machine to apply changes performed in the previous step
4. Run following commands:
    ```
    docker-compose up -d
    make
    ```
