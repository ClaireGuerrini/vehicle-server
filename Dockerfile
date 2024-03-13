FROM gcr.io/distroless/static-debian12
COPY ./dist /app
ENTRYPOINT ["/app/server"]
