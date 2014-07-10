CREATE TABLE SSDM_RAW(
   SSN varchar(9) PRIMARY KEY     NOT NULL,
   FIRSTNAME           CHAR(15)    NOT NULL,
   MIDDLENAME           CHAR(15)    NOT NULL,
   LASTNAME           CHAR(20)    NOT NULL,
   SUFFIX           varCHAR(4),
   BDAY            varchar(2),
   BMONTH            varchar(2),
   BYEAR            varchar(4),
   DDAY            varchar(2),
   DMONTH            varchar(2),
   DYEAR            varchar(4),
   RAW        varchar(82)
);

copy SSDM_RAW from '/home/ec2-user/output1.csv' DELIMITERS ',' CSV;