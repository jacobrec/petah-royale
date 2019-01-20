FROM golang

RUN go get github.com/gorilla/websocket
RUN go get github.com/gin-gonic/gin
RUN go get github.com/SolarLune/resolv/resolv
RUN go get github.com/gin-contrib/cors

COPY server      /go/src/github.com/jacobrec/petah-royale/server
COPY games/alpha /go/src/github.com/jacobrec/petah-royale/games/alpha

RUN go install github.com/jacobrec/petah-royale/games/alpha

CMD alpha
