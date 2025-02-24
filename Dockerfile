FROM golang:1.24 as euphoria

RUN apt-get update && apt install libluajit-5.1-dev -y
COPY . /euphoria
WORKDIR /euphoria

RUN go mod download && go mod verify

RUN CGO_LDFLAGS="-lm -lluajit-5.1" go build -tags luajit -ldflags '-linkmode external -extldflags "-fno-PIC -static"' -v -o /usr/bin

FROM alpine:3.19
ARG AWGTOOLS_RELEASE="0.1.0"
RUN apk --no-cache add iproute2 iptables bash && \
    cd /usr/bin/ && \
    wget https://github.com/amnezia-vpn/euphoria-tools/releases/download/v${AWGTOOLS_RELEASE}/alpine-3.19-amneziawg-tools.zip && \
    unzip -j alpine-3.19-amneziawg-tools.zip && \
    chmod +x /usr/bin/awg /usr/bin/awg-quick && \
    ln -s /usr/bin/awg /usr/bin/wg && \
    ln -s /usr/bin/awg-quick /usr/bin/wg-quick
COPY --from=euphoria /usr/bin/euphoria /usr/bin/amneziawg-go
