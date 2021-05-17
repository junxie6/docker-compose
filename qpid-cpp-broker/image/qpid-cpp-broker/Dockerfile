ARG CENTOS_TAG
FROM ${CENTOS_TAG}

RUN yum -y update

RUN yum -y install cmake boost-devel libuuid-devel pkgconfig gcc-c++ make ruby help2man doxygen graphviz

RUN yum -y install cyrus-sasl-devel cyrus-sasl-md5 cyrus-sasl-plain nss-devel nspr-devel

RUN yum -y git

RUN yum -y install libdb-cxx-devel libaio-devel

RUN yum -y install openssl-devel

RUN yum -y install git gcc gcc-c++ cmake make swig pkgconfig libuuid-devel openssl-devel python-devel python2-devel python3-devel cyrus-sasl-devel \
    boost-devel boost-filesystem boost-program-options cyrus-sasl-lib libaio-devel libdb4-cxx-devel libibverbs-devel librdmacm-devel nspr-devel nss-devel pkgconfig python-setuptools xerces-c-devel xqilla-devel ruby ruby-devel
