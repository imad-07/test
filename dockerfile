FROM golang:1.23.4

LABEL Team="Orca"
LABEL Project_Name="Forum"
LABEL Description="Simple social media web where you can share your posts with friends"

WORKDIR /forum
COPY . .
EXPOSE 8081
WORKDIR /forum/server
RUN go build

CMD ["./server"]