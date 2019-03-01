FROM golang:latest
WORKDIR /app 
COPY ./ .
RUN CGO_ENABLED=0 && export CGO_ENABLED && make all 
CMD ["/bin/sh"]

FROM alpine:latest  
WORKDIR /app
COPY --from=0 /app/build /app/
CMD ["./bootstrapper serve"]  
