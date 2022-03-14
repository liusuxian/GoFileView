FROM loads/alpine:3.8
MAINTAINER lsx "382185882@qq.com"
COPY fonts/* /usr/share/fonts/ChineseFonts/

###############################################################################
#                                INSTALLATION
###############################################################################
# 设置固定的项目路径
ENV WORKDIR /var/www/GoFileView
ENV LANG en_US.UTF-8
ENV LC_ALL en_US.UTF-8
ENV TZ=Asia/Shanghai

RUN  apk add --no-cache wget
RUN  apk add --no-cache java
RUN  apk add --no-cache deltarpm &&\
     apk add --no-cache libreoffice &&\
     apk add --no-cache libreoffice-headless &&\
     apk add --no-cache libreoffice-writer &&\
     apk add --no-cache ImageMagick &&\
     export DISPLAY=:0.0

# 安装wkhtmltopdf用于将eml（html）文件转pdf
RUN rpm --rebuilddb &&\
    apk add --no-cache openssl &&\
    apk add --no-cache xorg-x11-fonts-75dpi &&\
    apk add --no-cache wkhtmltopdf

# 添加应用可执行文件，并设置执行权限
ADD ./temp/linux_amd64/main $WORKDIR/main
RUN chmod +x $WORKDIR/main

# 添加I18N多语言文件、静态文件、配置文件、模板文件
ADD resource $WORKDIR/

# 创建 cache 目录以及子目录
RUN mkdir -p $WORKDIR/cache/convert/
RUN mkdir -p $WORKDIR/cache/download/
RUN mkdir -p $WORKDIR/cache/local/
RUN mkdir -p $WORKDIR/cache/pdf/

# jar包，用于将.msg文件转eml文件
COPY utility/emailconverter-2.5.3-all.jar /usr/local/emailconverter-2.5.3-all.jar

# pdf 添加水印
COPY utility/pdfcpu /usr/local/pdfcpu

###############################################################################
#                                   START
###############################################################################
WORKDIR $WORKDIR
CMD ./main

# ------------------------------------本地打包镜像---------------------
# gf docker main.go -t gofileview:v1.0
# docker run -d -p 8089:80 <imageID>