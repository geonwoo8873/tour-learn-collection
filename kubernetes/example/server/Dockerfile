FROM golang:1.27.4-alpine3.23 AS build
WORKDIR /src

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# `CGO_ENABLED=0`는 C 라이브러리의 의존성을 사용하지 않도록 설정하여, 빌드의 시간을 단축시키고 정적 바이너리를 생성한다.
# `-trimpath`는 빌드 경로 정보를 제거하여 빌드 결과물의 경로 의존성을 없앤다.
# `-ldflags="-s -w"`는 디버그 정보를 제거하여 바이너리 크기를 줄인다.
# 이러한 옵션들을 통해 최적화된 정적 바이너리를 생성해 런타임 성능과 보안성을 향상 시킨다.
# 하지만, 향후 C 라이브러리를 사용해야 하는 경우에는 빌드가 깨질 수 있다.
RUN CGO_ENABLED=0 go build -trimpath -ldflags="-s -w" -o /out/server .

#-------------------------------[Multi-stage Build]-----------------------------------------

FROM alpine:3.23
RUN apk add --no-cache ca-certificates
WORKDIR /app

COPY --from=build /out/server ./server

# `USER`는 컨테이너가 실행될 때 사용할 프로세스를 지정을하고 기본 GID 65532 권한을 실행할 사용자로 설정하라는 명령이다.
# 기본으로 root 사용자 권한으로 실행되지만 USER 명령으로 비루트 사용자로 변경한다.
USER 65532:65532

EXPOSE 8080
ENTRYPOINT ["/app/server"]

#### 이미지 태그와 버전간 빌드 및 소요 시간 비교
# 1. golang:alpine3.23 + final image = 14.9sec / 978.67MB
# 2. golang:1.27.4-alpine3.23 + alpine:3.23 final image = 29.1sec / 34.92MB