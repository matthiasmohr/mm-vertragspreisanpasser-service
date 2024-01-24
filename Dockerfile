FROM alpine:3

RUN echo "test Dockerfile"
RUN apk update && \
    apk add --no-cache libxml2 \
                       tzdata \
                       libc6-compat \
                       git

WORKDIR /service

# Create an unprivileged user
RUN adduser \
    --disabled-password \
    --gecos "" \
    --home "$(pwd)" \
    --no-create-home \
    "service"


COPY dist/sample_service /service

RUN chown -R service /service

## Don't run as root if you don't need to
USER service

# Command
ENTRYPOINT [ "./sample_service" ]

EXPOSE 5000
