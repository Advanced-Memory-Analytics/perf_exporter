FROM ubuntu:20.04


RUN apt -y update \
    && apt install -y linux-tools-common
RUN apt-get -y update 
# This will update the repositories list
RUN apt-get -y upgrade 
# This will update all the necessary packages on your system
RUN apt-get -y dist-upgrade 
# This will add/remove any needed packages
# RUN reboot 
# You may need this since sometimes after a upgrade/dist-upgrade, there are some left over entries that get fixed after a reboot
# RUN apt-get install linux-headers-$(uname -r) 
# This should work now
RUN apt-get install -y linux-tools-$(uname -r)