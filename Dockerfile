FROM alpine:latest

WORKDIR /app

COPY build/blog-svc .

RUN chmod +x blog-svc

EXPOSE 8001

CMD ["./blog-svc"]
