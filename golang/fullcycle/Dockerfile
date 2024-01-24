FROM nginx:1.15.0-alpine

WORKDIR /usr/src/go

COPY . ./usr/src/go

ENTRYPOINT [ "go", "run", "." ]

#CMD ["go","run", "."]