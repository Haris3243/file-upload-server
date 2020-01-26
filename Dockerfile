FROM alpine:latest
RUN mkdir /lib64 && ln -s /lib/libc.musl-x86_64.so.1 /lib64/ld-linux-x86-64.so.2
RUN mkdir /server
RUN mkdir /server/data
WORKDIR /server
COPY ./main /server/
EXPOSE 8083
CMD ["./main"]