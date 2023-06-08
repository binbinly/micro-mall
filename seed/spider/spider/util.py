# 解析url地址
def parse_img_url(img_url):
    index = img_url.find('images')
    if index > -1:
        return img_url[index:]
    else:
        return img_url


# 去除url ? 后的字符
def split_img_url(img_url):
    if img_url is None:
        return ""
    index = img_url.find('?')
    if index > -1:
        return img_url[0:index]
    return img_url
