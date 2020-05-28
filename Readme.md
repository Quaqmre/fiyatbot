Its work on epey.com 

fiyatbot \
    -item=https://www.epey.com/laptop/dell-inspiron-5490-s510f82n.html \
    -itemName=inspiron5490i7
    -smtp=smtp.gmail.com
    -from=sender@mail
    -to=receiver@mail
    -pass=password 
    -price=triggerPrice
    -interval=(minute)

docker run quaqmre/fiyatbot:v1 \
    -item=https://www.epey.com/laptop/dell-inspiron-5490-s510f82n.html \
    -itemName=inspiron5490i7
    -smtp=smtp.gmail.com
    -from=sender@mail
    -to=receiver@mail
    -pass=smtpPasswprd 
    -price=triggerPrice
    -interval=(minute)