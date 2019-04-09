#!/bin/bash

for ((i = 1; i <= 21; i++)); do
  hugo new docs/${i}-xx.md
  sleep 1
done
