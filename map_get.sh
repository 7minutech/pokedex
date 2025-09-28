
> map_log.json
echo "[" >> map_log.json
curl -s https://pokeapi.co/api/v2/location-area | jq '.' >> map_log.json
echo "," >> map_log.json
curl -s "https://pokeapi.co/api/v2/location-area?offset=20&limit=20" | jq '.' >> map_log.json
 echo "]" >> map_log.json
