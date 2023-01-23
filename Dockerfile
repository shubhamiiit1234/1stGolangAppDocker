#Alpine is base image and using golang version 1.19.5
FROM golang:1.19.5-alpine
#Current working directory inside the image
WORKDIR /app
#Before downloading and insatalling the modules and packages, we need to copy the go.mod and go.sum files into the current directory
COPY go.mod ./
COPY go.sum ./
#Downloading the modules required, mentioned in the go.mod file
RUN go mod download
#Copying all .go files into the current directory
COPY *.go ./
#Building the application, creating the executable (i.e. /fileName)
RUN go build -o /first-docker
#Port on which we want to host
EXPOSE 8080
#Run executable
CMD [ "/first-docker" ]