FROM golang:1.19-alpine

ENV PKGS "curl git"
ENV ROOT /app
ENV CGO_ENABLED 0

RUN apk update && apk add --no-cache ${PKGS} \
  && go install golang.org/x/tools/gopls@latest \
  && go install github.com/cweill/gotests/gotests@latest \
  && go install github.com/fatih/gomodifytags@latest

WORKDIR ${ROOT}

COPY ./ ./

RUN go mod download

EXPOSE 8080

CMD ["go", "run", "main.go"]
