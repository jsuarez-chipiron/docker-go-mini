FROM scratch
#RUN "CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ."
ADD main /
EXPOSE 8080
CMD ["/main"]