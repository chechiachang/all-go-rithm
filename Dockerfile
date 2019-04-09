FROM scratch
EXPOSE 8080
ENTRYPOINT ["/all-go-rithm"]
COPY ./bin/ /