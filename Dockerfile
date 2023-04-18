FROM golang:1.19.5
RUN mkdir /app
ADD . /app
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN apt update

RUN apt install -y wget fontconfig fontconfig-config fonts-dejavu-core libbsd0 \
libfontconfig1 libfontenc1 libfreetype6 libjpeg62-turbo libmd0 libpng16-16 \
libx11-6 libx11-data libxau6 libxcb1 libxdmcp6 libxext6 libxrender1 sensible-utils \
ucf x11-common xfonts-75dpi xfonts-base xfonts-encodings xfonts-utils

RUN wget https://github.com/wkhtmltopdf/packaging/releases/download/0.12.6-1/wkhtmltox_0.12.6-1.buster_amd64.deb

RUN dpkg -i wkhtmltox_0.12.6-1.buster_amd64.deb

RUN CGO_ENABLED=0 GOOS=linux go build -o /attempt4

CMD ["/attempt4"]

