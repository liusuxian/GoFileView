FROM centos:7.9.2009
MAINTAINER liusuxian "382185882@qq.com"
COPY utility/external/fonts/* /usr/share/fonts/

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /GoFileView
RUN yum install kde-l10n-Chinese -y
RUN yum install glibc-common -y
RUN localedef -c -f UTF-8 -i zh_CN zh_CN.utf8
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV TZ=Asia/Shanghai

RUN yum install wget -y
RUN yum install vim -y
RUN yum install java -y
RUN yum install deltarpm -y &&\
    yum install ImageMagick -y &&\
    export DISPLAY=:0.0

# 添加应用可执行文件，并设置执行权限
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD resource $WORKDIR/resource
ADD manifest/config $WORKDIR/config

# 创建 cache 目录以及子目录
RUN mkdir -p $WORKDIR/cache/convert/
RUN mkdir -p $WORKDIR/cache/download/
RUN mkdir -p $WORKDIR/cache/local/
RUN mkdir -p $WORKDIR/cache/pdf/

# jar包，用于将.msg文件转eml文件
COPY utility/external/lib/emailconverter/linux/emailconverter-2.5.3-all.jar /usr/local/emailconverter-2.5.3-all.jar
RUN chmod +x /usr/local/emailconverter-2.5.3-all.jar

# pdf 添加水印
# pdfcpu 下载地址 https://github.com/pdfcpu/pdfcpu/releases
COPY utility/external/lib/pdfcpu/linux/pdfcpu /usr/local/pdfcpu
RUN chmod +x /usr/local/pdfcpu

# 安装 libreoffice
RUN wget https://mirrors.cloud.tencent.com/libreoffice/libreoffice/stable/7.3.3/rpm/x86_64/LibreOffice_7.3.3_Linux_x86-64_rpm.tar.gz
RUN tar -zxvf LibreOffice_7.3.3_Linux_x86-64_rpm.tar.gz
RUN rpm -Uivh LibreOffice_7.3.3.2_Linux_x86-64_rpm/RPMS/*.rpm --nodeps
RUN rm -rf LibreOffice_7.3.3*

# 安装wkhtmltopdf 用于将eml（html）文件转pdf
RUN wget https://github.com/wkhtmltopdf/wkhtmltopdf/releases/download/0.12.5/wkhtmltox-0.12.5-1.centos7.x86_64.rpm &&\
    rpm --rebuilddb && yum install -y openssl && yum install -y xorg-x11-fonts-75dpi &&\
    rpm -ivh wkhtmltox-0.12.5-1.centos7.x86_64.rpm

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

# ------------------------------------本地打包镜像---------------------
# gf docker main.go -t gofileview:v1.0
# docker run -d -p 8089:80 <imageID>