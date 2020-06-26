#!/bin/sh

GOOS=linux go build -o main && zip main.zip main && rm main
