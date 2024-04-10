#!/bin/bash

# Array contenente i nomi degli dei greci
dei=("Zeus" "Era" "Efesto" "Atena" "Apollo" "Artemide" "Ares" "Afrodite" "Estia" "Ermes" "Demetra" "Poseidone" "Ade")

# Ciclo attraverso gli elementi dell'array e invio della richiesta curl per ciascun dio
for dio in "${dei[@]}"
do
    echo "Invio richiesta per il dio: $dio"
    curl -v \
        -X POST \
        -H 'Content-Type: application/json' \
        -d "{\"username\":\"$dio\"}" \
        localhost:3000/session/
    echo -e "\n"
done

###random follow
for dio in "${dei[@]}"
do
    echo "Faccio seguire altri dei a: $dio"
    
    # Fai seguire ad ogni dio altri 4 dei
    for ((i=0; i<4; i++))
    do
        # Scegli casualmente un dio diverso da quello attuale
        # (assicurati che non stia seguendo già questo dio)
        while true
        do
            index=$(( RANDOM % ${#dei[@]} ))
            altro_dio="${dei[$index]}"
            if [ "$altro_dio" != "$dio" ]; then
                break
            fi
        done

        # Invia la richiesta per far seguire il dio attuale all'altro dio
        curl -v \
            -X PUT \
            -H "Authorization: $dio" \
            "localhost:3000/following/$altro_dio/"

        echo "Il dio $dio ora segue $altro_dio"
    done

    echo -e "\n"
done

###random ban
for dio in "${dei[@]}"
do
    echo "Faccio seguire altri dei a: $dio"
    
    # Fai seguire ad ogni dio altri 4 dei
    for ((i=0; i<2; i++))
    do
        # Scegli casualmente un dio diverso da quello attuale
        # (assicurati che non stia seguendo già questo dio)
        while true
        do
            index=$(( RANDOM % ${#dei[@]} ))
            altro_dio="${dei[$index]}"
            if [ "$altro_dio" != "$dio" ]; then
                break
            fi
        done

        # Invia la richiesta per far seguire il dio attuale all'altro dio
        curl -v \
            -X PUT \
            -H "Authorization: $dio" \
            "localhost:3000/ban/$altro_dio/"

        echo "Il dio $dio ora ha bannato $altro_dio"
    done

    echo -e "\n"
done

# Upload photos
for ((i=1; i<=3; i++))
do
    curl -v \
        -X POST \
        -H 'Authorization: Zeus' \
        -H 'Content-Type: image/jpeg' \
        --data-binary "@./photo-samples/atene$i.jpg" \
        localhost:3000/photos/
done

for ((i=1; i<=3; i++))
do
    curl -v \
        -X POST \
        -H 'Authorization: Ares' \
        -H 'Content-Type: image/png' \
        --data-binary "@./photo-samples/sparta$i.png" \
        localhost:3000/photos/
done

curl -v \
    -X POST \
    -H 'Content-Type: application/json' \
    -d '{"username":"Urano"}' \
    localhost:3000/session/


curl -v \
    -X PUT \
    -H 'Authorization: Urano' \
    localhost:3000/following/{Ares}/

curl -v \
    -X PUT \
    -H 'Authorization: Urano' \
    localhost:3000/following/{Zeus}/
