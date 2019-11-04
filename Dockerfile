FROM 301301/ubuntulatest
WORKDIR /app
RUN apt install sysstat -y
ADD /client /app/client
ADD project-1 /app/p1
EXPOSE 8080
CMD ["/app/p1"]