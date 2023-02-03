# Base Application
## run

### bootstrap
```inventory bootstrap```

### agent
```inventory agent run```

## local deploy
```
unset KAFKA_ADDR
cd foliage/
docker compose up -d

export KAFKA_ADDR=127.0.0.1:9094
inventory agent run
```