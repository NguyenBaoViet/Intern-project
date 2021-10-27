FROM alpine:3.8

WORKDIR /home/
COPY cr-adapter-service.bin .

RUN chmod +x main.bin

EXPOSE 443

CMD [ "./main.bin" ]