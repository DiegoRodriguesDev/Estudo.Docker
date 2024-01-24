FROM mcr.microsoft.com/dotnet/sdk:6.0

WORKDIR /var/www

RUN apt-get update

RUN dotnet new mvc -au None --name docker_estudo

#run dotnet run docker_estudo --urls=http://localhost:8000/