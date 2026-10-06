FROM ubuntu:24.04

# Update the package list and install necessary packages
RUN apt-get update && apt-get install -y \
    apt-get install net-tools curl vim  -y

# Expose container port 80 bind to the host
EXPOSE 80

CMD ["/bin/bash"]