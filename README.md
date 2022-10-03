# TechnicalTest-Back-End-For-PT-Millenium-Pharmacon-International

Hi, Thank you for taking the time to take this home test.
All we need, you have to solve this case:
1. Make converter int number to string number in bahasa max 6 digit, example:
a. 12 -> dua belas,
b. 2012 -> dua ribu dua belas
c. 999999 -> sembilan ratus ribu sembilan puluh sembilan sembilan ratus
sembilan puluh sembilan
You can choose whatever language that you want
2. Using Go, write a simple HTTP server and handler that serves a HTML form
with 2 fields:
a. A hidden input field named "auth" that receives its input value passed
from theGo code. This token should be passed as an environment
variable to the application.
b. A file upload field named "data" (ie, <input type="file"/>) that uploads a
file that the user selects
(please do not waste time trying to make the HTML form pretty -- we don't
need HTML developers, we need Go developers)
- The form above should POST data to the /upload handler, which
should write the received file data to a temporary file.
- Before accepting any data, you should check that the content type of
the uploaded file is an image, and that the auth token matches. If the
submission is bad, please return a 403 HTTP error code. Images larger
than 8 megabytes should also be rejected.
- Write the image metadata (content type, size, etc) to a database of
your choice, including all relevant HTTP information.
- Do a short 1 or 2 paragraph write up with information on
building/running the application if relevant / necessary or non-standard
