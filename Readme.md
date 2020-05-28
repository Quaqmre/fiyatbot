Its work on epey.com 

Select your item in epey and add sender-receiver mail addres then if your item price under your triggerPrice will send mail
```console
fiyatbot \
        -item=https://www.epey.com/laptop/dell-inspiron-5490-s510f82n.html \
        -itemName=inspiron5490i7 \
        -smtp=smtp.gmail.com \
        -from=sender@mail \
        -to=receiver@mail \
        -pass=password \
        -price=triggerPrice \
        -interval=(minute) \
```

```console
docker run quaqmre/fiyatbot:v1 \
    -item=https://www.epey.com/laptop/dell-inspiron-5490-s510f82n.html \
    -itemName=inspiron5490i7 \
    -smtp=smtp.gmail.com \
    -from=sender@mail \
    -to=receiver@mail \
    -pass=smtpPasswprd \
    -price=triggerPrice(6000000=6bintl) \
    -interval=(minute)
```
