FROM loads/alpine:3.8

LABEL maintainer="rygo"

###############################################################################
#                                INSTALLATION
###############################################################################

# 设置固定的项目路径
ENV      WORKDIR   /root/apps
WORKDIR  $WORKDIR
# 添加应用可执行文件，并设置执行权限
ADD main      $WORKDIR/main
RUN chmod +x  $WORKDIR/main

# 添加I18N多语言文件、配置文件、静态文件、模板文件
#ADD i18n     $WORKDIR/i18n
ADD config.toml   $WORKDIR/config.toml
ADD public        $WORKDIR/public
ADD template      $WORKDIR/template

###############################################################################
#                                   START
###############################################################################
EXPOSE 80 8080
CMD ./main
