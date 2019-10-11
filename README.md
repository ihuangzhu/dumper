# dumper
分片视频下载合并器

## 有嘛用？
用于下载被分片的视频，目前只实现了ATV视频。

## 示例内容
```
{"autoplay":true,"video":[{"video":[{"file":"http://119.84.134.35/0/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":30.32,"bytesTotal":23025181},{"file":"http://119.84.134.35/1/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/2/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/3/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/4/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/5/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/6/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/7/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/8/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/9/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/10/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/11/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/12/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/13/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/14/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/15/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/16/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/17/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/18/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/19/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/20/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/21/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/22/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/23/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/24/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/25/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0},{"file":"http://119.84.134.35/26/61f17eb4582e3449b983385ff3866a98.mp4?type=atv&k=db82ba3a4a909a5ef40c8853c8751a20-e14b-1570796984%26segment%3D152ba5fe_152ba5ff_1570782584%26bppcataid%3D38","duration":25,"bytesTotal":0}]}]}
```

## 环境依赖
- FFMPEG >= v4.1
- Golang >= 1.13
