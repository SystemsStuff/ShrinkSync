FROM python:3-alpine
COPY . /DataGrid/
WORKDIR /DataGrid/
# Input generation
RUN python3 input_gen.py /DataGrid/Input wordCount 3 5
# Starting FTP server
RUN pip3 install -r requirements.txt
CMD python3 ftp_server.py
EXPOSE 21
