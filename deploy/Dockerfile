FROM alpine:3.8

WORKDIR /home/
COPY main.bin .

RUN chmod +x main.bin

EXPOSE 443

CMD ["./main.bin"]