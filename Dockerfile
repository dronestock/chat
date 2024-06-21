FROM dockerhub.storezhang.tech/dockerat/alpine:3.20.0

LABEL author="storezhang<华寅>" \
    email="storezhang@gmail.com" \
    qq="160290688" \
    wechat="storezhang" \
    description="聊天相关功能，比如：1、发送消息；2、上传文件等 "

# 复制文件
COPY from=builder /docker /

RUN set -ex \
    \
    \
    \
    # 增加执行权限
    && chmod +x /bin/plugin \
    \
    \
    \
    && rm -rf /var/cache/apk/*

# 执行命令
ENTRYPOINT /usr/local/bin/plugin
