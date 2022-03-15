FROM loads/alpine:3.8
# LABLE 给镜像添加元数据
# MAINTAINER 维护者信息
LABEL maintainer="382185882@qq.com"
COPY fonts/* /usr/share/fonts/ChineseFonts/

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /GoFileView
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

# Install cURL
#RUN echo -e "https://dl-cdn.alpinelinux.org/alpine/v3.9/main/\n" \
#            "https://dl-cdn.alpinelinux.org/alpine/v3.8/main/\n" \
#            "https://dl-cdn.alpinelinux.org/alpine/v3.7/main/\n" \
#            "https://dl-cdn.alpinelinux.org/alpine/v3.6/main/\n" \
#            "https://dl-cdn.alpinelinux.org/alpine/v3.5/main/\n" \
#            "https://mirror1.hs-esslingen.de/pub/Mirrors/alpine/edge/community/\n" \
#            "https://nl.alpinelinux.org/alpine/edge/main/\n" \
#            "https://nl.alpinelinux.org/alpine/edge/community/\n" \
#            > /etc/apk/repositories
RUN echo -e "https://mirrors.ustc.edu.cn/alpine/latest-stable/main\n https://mirrors.ustc.edu.cn/alpine/latest-stable/community\n" > /etc/apk/repositories
# RUN apk --update add curl bash openjdk8 && rm -rf /var/cache/apk/*
RUN apk update && apk add --no-cache curl bash openjdk8
# Set environment
ENV JAVA_HOME /usr/lib/jvm/default-jvm
ENV PATH ${PATH}:${JAVA_HOME}/bin

RUN apk add --no-cache wget vim libreoffice libreoffice-writer ImageMagick && export DISPLAY=:0.0

#RUN yum install java -y
#
#RUN yum install deltarpm -y &&\
#    yum install libreoffice -y &&\
#    yum install libreoffice-headless -y &&\
#    yum install libreoffice-writer -y &&\
#    yum install ImageMagick -y &&\
#    export DISPLAY=:0.0

# 添加应用可执行文件，并设置执行权限
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD resource $WORKDIR/resource
ADD manifest/config $WORKDIR/config
ADD manifest/test.pptx $WORKDIR/

# 创建 cache 目录以及子目录
RUN mkdir -p $WORKDIR/cache/convert/
RUN mkdir -p $WORKDIR/cache/download/
RUN mkdir -p $WORKDIR/cache/local/
RUN mkdir -p $WORKDIR/cache/pdf/

# jar包，用于将.msg文件转eml文件
COPY utility/emailconverter-2.5.3-all.jar /usr/local/emailconverter-2.5.3-all.jar

# pdf 添加水印
COPY utility/pdfcpu /usr/local/pdfcpu

# 安装wkhtmltopdf 用于将eml（html）文件转pdf
#RUN yum -y install wget &&\
#    wget http://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.5/wkhtmltox-0.12.5-1.centos7.x86_64.rpm &&\
#    rpm --rebuilddb && yum install -y openssl && yum install -y xorg-x11-fonts-75dpi &&\
#    rpm -ivh wkhtmltox-0.12.5-1.centos7.x86_64.rpm

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

# ------------------------------------本地打包镜像---------------------
# gf docker main.go -t gofileview:v1.0
# docker run -d -p 8089:80 <imageID>