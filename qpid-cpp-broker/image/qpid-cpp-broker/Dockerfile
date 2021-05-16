ARG CENTOS_TAG
FROM ${CENTOS_TAG}

RUN yum -y update

RUN yum -y install cmake boost-devel libuuid-devel pkgconfig gcc-c++ make ruby help2man doxygen graphviz

RUN yum -y install cyrus-sasl-devel cyrus-sasl-md5 cyrus-sasl-plain nss-devel nspr-devel

RUN yum -y git