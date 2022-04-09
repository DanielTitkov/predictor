FROM golang:latest AS dependencies
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM dependencies AS build
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY . .
RUN ["make", "build"]

FROM alpine
ARG ENV
ARG VERSION
ENV NAME "app"
WORKDIR /opt/${NAME}
COPY --from=build /opt/${NAME}/bin/${NAME} ./${NAME}
COPY --from=build /opt/${NAME}/data ./data
COPY --from=build /opt/${NAME}/templates ./templates
COPY --from=build /opt/${NAME}/web/media ./web/media
COPY --from=build /opt/${NAME}/web/dist ./web/dist
CMD ./${NAME} ${CONFIG_PATH}